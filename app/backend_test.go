//go:build darwin || windows || linux || viva_custom_backend

package app_test

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"math"
	"reflect"
	"testing"

	"github.com/vivaui/viva/app"
	"github.com/vivaui/viva/internal/scene"
	"github.com/vivaui/viva/ui"
)

type externalBackend struct {
	services app.BackendServices
	run      func(app.BackendConfig) error
}

func (b *externalBackend) Run(c app.BackendConfig) error { return b.run(c) }
func (*externalBackend) RunApplication(start app.BackendApplicationStart) error {
	return app.ErrUnsupported
}
func (*externalBackend) TextEngine() ui.TextEngine       { return ui.ReferenceTextEngine() }
func (b *externalBackend) Services() app.BackendServices { return b.services }

func TestExternalBackendReceivesProductionFrames(t *testing.T) {
	frames := 0
	b := &externalBackend{run: func(c app.BackendConfig) error {
		if c.Width != 320 || len(c.Rectangles) == 0 || len(c.TextLabels) == 0 {
			t.Fatalf("incomplete initial frame: %+v", c)
		}
		if c.OnResize == nil || c.OnRedraw == nil {
			t.Fatal("missing retained painter callbacks")
		}
		f, err := c.OnResize(app.BackendSize{Width: 640, Height: 360, Scale: 2})
		if err != nil || len(f.Rectangles) == 0 || f.Update.Empty() {
			t.Fatalf("resize frame: %+v, %v", f, err)
		}
		frames++
		c.OnClosed()
		return nil
	}}
	if err := app.UseBackend(b); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = app.UseBackend(nil) })
	err := app.RunCanvas(app.Config{Width: 320, Height: 240, Paint: func(c *ui.Canvas, s ui.Size) {
		c.Fill(ui.Rect{Width: s.Width, Height: s.Height}, ui.RGB(20, 30, 40))
		c.Label("custom backend", ui.Rect{Width: 200, Height: 30}, ui.TextStyle{Size: 16, Color: ui.RGB(255, 255, 255)})
	}}, nil)
	if err != nil || frames != 1 {
		t.Fatalf("custom RunCanvas: frames=%d err=%v", frames, err)
	}
}

func TestSelectedBackendOwnsServices(t *testing.T) {
	clipboard := "external"
	b := &externalBackend{services: app.BackendServices{
		ClipboardText:      func() (string, error) { return clipboard, nil },
		SetClipboardText:   func(s string) error { clipboard = s; return nil },
		PreferredLanguages: func() []string { return []string{"zh-CN", "en"} },
	}}
	if err := app.UseBackend(b); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = app.UseBackend(nil) })
	if err := app.SetClipboardText("custom"); err != nil {
		t.Fatal(err)
	}
	if got, err := app.ClipboardText(); got != "custom" || err != nil {
		t.Fatal(got, err)
	}
	if languages := app.PreferredLanguages(); len(languages) != 2 || languages[0] != "zh-CN" {
		t.Fatal(languages)
	}
	if err := app.OpenURL("https://example.com"); !errors.Is(err, app.ErrUnsupported) {
		t.Fatalf("missing external service fell through to the built-in platform: %v", err)
	}
}

func newRasterOffscreen(t *testing.T, config app.BackendOffscreenConfig, rasterize app.BackendLabelRasterizer) app.BackendOffscreen {
	t.Helper()
	target, err := app.NewRasterOffscreen(config, rasterize)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(target.Close)
	return target
}

func rasterImage(t *testing.T, target app.BackendOffscreen) *image.RGBA {
	t.Helper()
	picture, err := target.Image()
	if err != nil {
		t.Fatal(err)
	}
	return picture
}

func checkRasterPixel(t *testing.T, picture *image.RGBA, x, y int, want color.RGBA) {
	t.Helper()
	if got := picture.RGBAAt(x, y); got != want {
		t.Errorf("pixel (%d, %d) = %v, want %v", x, y, got, want)
	}
}

func TestRasterOffscreenValidatesDimensions(t *testing.T) {
	for _, config := range []app.BackendOffscreenConfig{
		{Width: 0, Height: 1, Scale: 1},
		{Width: 1, Height: -1, Scale: 1},
		{Width: 16385, Height: 1, Scale: 1},
		{Width: 1, Height: 16385, Scale: 1},
		{Width: 1, Height: 1, Scale: 0},
		{Width: 1, Height: 1, Scale: 0.99},
		{Width: 1, Height: 1, Scale: 64.01},
		{Width: 1, Height: 1, Scale: math.NaN()},
		{Width: 1, Height: 1, Scale: math.Inf(1)},
		{Width: 1, Height: 1, Scale: math.Inf(-1)},
		{Width: 8193, Height: 1, Scale: 2},
		{Width: 1, Height: 8193, Scale: 2},
	} {
		target, err := app.NewRasterOffscreen(config, nil)
		if target != nil {
			target.Close()
		}
		if err == nil || target != nil {
			t.Errorf("NewRasterOffscreen(%+v) = %v, %v; want nil and error", config, target, err)
		}
	}
	for _, config := range []app.BackendOffscreenConfig{
		{Width: 1, Height: 1, Scale: 64},
		{Width: 16384, Height: 1, Scale: 1},
		{Width: 1, Height: 8192, Scale: 2},
	} {
		target := newRasterOffscreen(t, config, nil)
		want := image.Rect(0, 0, int(float64(config.Width)*config.Scale), int(float64(config.Height)*config.Scale))
		if got := rasterImage(t, target).Bounds(); got != want {
			t.Errorf("bounds = %v, want %v", got, want)
		}
	}
}

func TestRasterOffscreenRectanglesAndLifecycle(t *testing.T) {
	target := newRasterOffscreen(t, app.BackendOffscreenConfig{
		Width: 4, Height: 3, Scale: 1.5, ClearColor: app.BackendColor{Blue: 1, Alpha: 1},
	}, nil)
	frame := app.BackendFrame{Rectangles: []app.BackendRectangle{{
		Bounds: scene.Rect{Width: 2, Height: 2}, Color: scene.Color{Red: 1, Alpha: 1},
	}}}
	if err := target.Set(frame); err != nil {
		t.Fatal(err)
	}
	frame.Rectangles[0].Color = scene.Color{Green: 1, Alpha: 1}
	picture := rasterImage(t, target)
	if picture.Bounds() != image.Rect(0, 0, 6, 5) {
		t.Fatalf("fractional-scale bounds = %v", picture.Bounds())
	}
	checkRasterPixel(t, picture, 1, 1, color.RGBA{R: 255, A: 255})
	checkRasterPixel(t, picture, 4, 4, color.RGBA{B: 255, A: 255})
	picture.SetRGBA(1, 1, color.RGBA{})
	checkRasterPixel(t, rasterImage(t, target), 1, 1, color.RGBA{R: 255, A: 255})
	if err := target.Set(app.BackendFrame{}); err != nil {
		t.Fatal(err)
	}
	checkRasterPixel(t, rasterImage(t, target), 1, 1, color.RGBA{B: 255, A: 255})
	target.Close()
	target.Close()
	if err := target.Set(frame); err == nil {
		t.Error("Set succeeded after Close")
	}
	if picture, err := target.Image(); picture != nil || err == nil {
		t.Errorf("Image after Close = %v, %v", picture, err)
	}
}

func TestRasterOffscreenTextPlacementClippingAndLayers(t *testing.T) {
	calls := 0
	label := app.BackendTextLabel{
		Value: "text", X: 0.8, Y: 1, OriginX: 0.2, OriginY: 0.3, Width: 2.2, Height: 1.1,
		FontSize: 12, Weight: 600, Layer: 1, Color: app.BackendColor{Green: 1, Alpha: 1},
		Clipped: true, Clip: ui.Rect{X: 1.8, Y: 1, Width: 1, Height: 2},
	}
	target := newRasterOffscreen(t, app.BackendOffscreenConfig{Width: 6, Height: 4, Scale: 1.5},
		func(got app.BackendTextLabel, scale float32, width, height int) ([]byte, error) {
			calls++
			if !reflect.DeepEqual(got, label) || scale != 1.5 || width != 4 || height != 2 {
				t.Fatalf("rasterize(%+v, %g, %d, %d)", got, scale, width, height)
			}
			return bytes.Repeat([]byte{0, 255, 0, 255}, width*height), nil
		})
	if err := target.Set(app.BackendFrame{
		Rectangles: []app.BackendRectangle{{Bounds: scene.Rect{X: 1, Y: 1, Width: 1, Height: 1}, Color: scene.Color{Blue: 1, Alpha: 1}, Layer: 2}},
		TextLabels: []app.BackendTextLabel{label, {Width: 1, Height: 1}, {Value: "empty bounds"}},
	}); err != nil {
		t.Fatal(err)
	}
	if calls != 1 {
		t.Fatalf("rasterize called %d times", calls)
	}
	picture := rasterImage(t, target)
	checkRasterPixel(t, picture, 2, 2, color.RGBA{B: 255, A: 255})
	checkRasterPixel(t, picture, 3, 2, color.RGBA{G: 255, A: 255})
	checkRasterPixel(t, picture, 4, 3, color.RGBA{G: 255, A: 255})
	checkRasterPixel(t, picture, 5, 2, color.RGBA{})
	checkRasterPixel(t, picture, 1, 3, color.RGBA{})
}

func TestRasterOffscreenBitmapPackingAndResampling(t *testing.T) {
	bitmap := []byte{
		0, 0, 255, 255, 0, 255, 0, 255, 99, 99, 99, 99,
		255, 0, 0, 255, 255, 255, 255, 255,
	}
	target := newRasterOffscreen(t, app.BackendOffscreenConfig{Width: 5, Height: 4, Scale: 1},
		func(app.BackendTextLabel, float32, int, int) ([]byte, error) {
			t.Fatal("bitmap invoked the text rasterizer")
			return nil, nil
		})
	if err := target.Set(app.BackendFrame{TextLabels: []app.BackendTextLabel{{
		Value: "bitmap takes precedence", X: 1, Y: 1, OriginX: 2, OriginY: 2, Width: 2, Height: 2,
		Bitmap: bitmap, BitmapWidth: 2, BitmapHeight: 2, BitmapStride: 12,
	}}}); err != nil {
		t.Fatal(err)
	}
	clear(bitmap)
	picture := rasterImage(t, target)
	checkRasterPixel(t, picture, 1, 1, color.RGBA{R: 255, A: 255})
	checkRasterPixel(t, picture, 2, 1, color.RGBA{G: 255, A: 255})
	checkRasterPixel(t, picture, 1, 2, color.RGBA{B: 255, A: 255})
	checkRasterPixel(t, picture, 2, 2, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	checkRasterPixel(t, picture, 3, 3, color.RGBA{})

	// At this fractional position both edges round to a one-pixel span,
	// whereas ceil(logical width * scale) would incorrectly give two pixels.
	target = newRasterOffscreen(t, app.BackendOffscreenConfig{Width: 4, Height: 4, Scale: 1.5}, nil)
	if err := target.Set(app.BackendFrame{TextLabels: []app.BackendTextLabel{{
		X: 0.4, Y: 0.4, Width: 1, Height: 1, BitmapWidth: 4, BitmapHeight: 2, BitmapStride: 16,
		Bitmap: bytes.Repeat([]byte{0, 0, 255, 255, 0, 255, 0, 255, 255, 0, 0, 255, 255, 255, 255, 255}, 2),
	}}}); err != nil {
		t.Fatal(err)
	}
	picture = rasterImage(t, target)
	checkRasterPixel(t, picture, 1, 1, color.RGBA{R: 128, G: 128, B: 128, A: 255})
	checkRasterPixel(t, picture, 2, 1, color.RGBA{})
}

func TestRasterOffscreenPreservesPremultipliedAlpha(t *testing.T) {
	target := newRasterOffscreen(t, app.BackendOffscreenConfig{Width: 3, Height: 1, Scale: 1}, nil)
	if err := target.Set(app.BackendFrame{
		Rectangles: []app.BackendRectangle{{Bounds: scene.Rect{Width: 1, Height: 1}, Color: scene.Color{Red: 1, Alpha: 0.5}}},
		TextLabels: []app.BackendTextLabel{{X: 1, Width: 1, Height: 1, Bitmap: []byte{32, 64, 96, 128}, BitmapWidth: 1, BitmapHeight: 1, BitmapStride: 4}},
	}); err != nil {
		t.Fatal(err)
	}
	picture := rasterImage(t, target)
	checkRasterPixel(t, picture, 0, 0, color.RGBA{R: 128, A: 128})
	checkRasterPixel(t, picture, 1, 0, color.RGBA{R: 96, G: 64, B: 32, A: 128})
	checkRasterPixel(t, picture, 2, 0, color.RGBA{})
}

func TestRasterOffscreenActiveTextInput(t *testing.T) {
	for _, test := range []struct {
		name   string
		secure bool
		mono   ui.MonoFace
		term   bool
		value  string
	}{
		{name: "plain", value: "a界🙂"},
		{name: "secure", secure: true, value: "•••"},
		{name: "mono", mono: ui.MonoClassic, value: "a界🙂"},
		{name: "terminal", term: true, value: "a界🙂"},
	} {
		t.Run(test.name, func(t *testing.T) {
			calls := 0
			input := app.BackendTextInput{
				X: 1, Y: 1, Width: 2, Height: 1, FontSize: 12, Layer: 3,
				Secure: test.secure, MonoFace: test.mono, Terminal: test.term,
				Color: app.BackendColor{Red: 1, Alpha: 1}, State: app.BackendTextState{Value: "a界🙂"},
			}
			target := newRasterOffscreen(t, app.BackendOffscreenConfig{Width: 5, Height: 3, Scale: 2},
				func(label app.BackendTextLabel, scale float32, width, height int) ([]byte, error) {
					calls++
					want := app.BackendTextLabel{
						Value: test.value, X: 1, Y: 1, Width: 2, Height: 1, FontSize: 12, Weight: 400, Layer: 3,
						Mono: test.mono != 0 || test.term, MonoFace: test.mono, Color: input.Color,
						Clipped: true, Clip: ui.Rect{X: 1, Y: 1, Width: 2, Height: 1},
					}
					if !reflect.DeepEqual(label, want) || scale != 2 || width != 4 || height != 2 {
						t.Fatalf("editor label = %+v at %gx, %dx%d; want %+v", label, scale, width, height, want)
					}
					return bytes.Repeat([]byte{0, 0, 255, 255}, width*height), nil
				})
			if err := target.Set(app.BackendFrame{TextInput: app.BackendTextInputUpdate{Changed: true, Input: &input}}); err != nil {
				t.Fatal(err)
			}
			checkRasterPixel(t, rasterImage(t, target), 2, 2, color.RGBA{R: 255, A: 255})
			// Unchanged input retains its prepared pixels, independent of the borrowed frame.
			input.State.Value = "changed by caller"
			if err := target.Set(app.BackendFrame{}); err != nil {
				t.Fatal(err)
			}
			checkRasterPixel(t, rasterImage(t, target), 2, 2, color.RGBA{R: 255, A: 255})
			if calls != 1 {
				t.Errorf("unchanged editor rasterized %d times", calls)
			}
			if err := target.Set(app.BackendFrame{TextInput: app.BackendTextInputUpdate{Changed: true}}); err != nil {
				t.Fatal(err)
			}
			checkRasterPixel(t, rasterImage(t, target), 2, 2, color.RGBA{})
		})
	}
}

func TestRasterOffscreenSkipsHiddenOrEmptyTextInput(t *testing.T) {
	target := newRasterOffscreen(t, app.BackendOffscreenConfig{Width: 2, Height: 2, Scale: 1},
		func(app.BackendTextLabel, float32, int, int) ([]byte, error) {
			t.Fatal("invisible editor invoked the rasterizer")
			return nil, nil
		})
	for _, input := range []*app.BackendTextInput{
		nil,
		{Width: 1, Height: 1},
		{Width: 1, Height: 1, State: app.BackendTextState{Value: "hidden"}, HideNativeOverlays: true},
		{Height: 1, State: app.BackendTextState{Value: "zero width"}},
		{Width: 1, State: app.BackendTextState{Value: "zero height"}},
	} {
		if err := target.Set(app.BackendFrame{TextInput: app.BackendTextInputUpdate{Changed: true, Input: input}}); err != nil {
			t.Fatal(err)
		}
		checkRasterPixel(t, rasterImage(t, target), 0, 0, color.RGBA{})
	}
	// Without a text callback, text and editor content are simply omitted.
	target = newRasterOffscreen(t, app.BackendOffscreenConfig{Width: 2, Height: 2, Scale: 1}, nil)
	if err := target.Set(app.BackendFrame{
		TextLabels: []app.BackendTextLabel{{Value: "label", Width: 1, Height: 1}},
		TextInput: app.BackendTextInputUpdate{Changed: true, Input: &app.BackendTextInput{
			Width: 1, Height: 1, State: app.BackendTextState{Value: "input"},
		}},
	}); err != nil {
		t.Fatal(err)
	}
	checkRasterPixel(t, rasterImage(t, target), 0, 0, color.RGBA{})
}

func TestRasterOffscreenNativeViewHoles(t *testing.T) {
	for _, alpha := range []float64{0, 1} {
		target := newRasterOffscreen(t, app.BackendOffscreenConfig{Width: 4, Height: 4, Scale: 1.5, ClearColor: app.BackendColor{Alpha: alpha}}, nil)
		frame := app.BackendFrame{
			Rectangles: []app.BackendRectangle{
				{Bounds: scene.Rect{Width: 4, Height: 4}, Color: scene.Color{Red: 1, Alpha: 1}, Layer: 1},
				{Bounds: scene.Rect{X: 2, Y: 2, Width: 1, Height: 1}, Color: scene.Color{Blue: 1, Alpha: 1}, Layer: 3},
			},
			NativeViews: []app.BackendNativeView{
				{Bounds: ui.Rect{X: 0.8, Y: 0.8, Width: 2.2, Height: 2.2}, Layer: 1},
				{Bounds: ui.Rect{Width: -1, Height: 4}, Layer: 4},
				{Bounds: ui.Rect{Width: 4}, Layer: 4},
				{Bounds: ui.Rect{X: 10, Y: 10, Width: 1, Height: 1}, Layer: 4},
			},
		}
		if err := target.Set(frame); err != nil {
			t.Fatal(err)
		}
		picture := rasterImage(t, target)
		checkRasterPixel(t, picture, 0, 0, color.RGBA{R: 255, A: 255})
		checkRasterPixel(t, picture, 1, 1, color.RGBA{A: byte(alpha * 255)})
		checkRasterPixel(t, picture, 4, 1, color.RGBA{A: byte(alpha * 255)})
		checkRasterPixel(t, picture, 5, 1, color.RGBA{R: 255, A: 255})
		checkRasterPixel(t, picture, 3, 3, color.RGBA{B: 255, A: 255})
		frame.NativeViews = nil // nil preserves the previously supplied holes.
		if err := target.Set(frame); err != nil {
			t.Fatal(err)
		}
		checkRasterPixel(t, rasterImage(t, target), 1, 1, color.RGBA{A: byte(alpha * 255)})
		frame.NativeViews = []app.BackendNativeView{}
		if err := target.Set(frame); err != nil {
			t.Fatal(err)
		}
		checkRasterPixel(t, rasterImage(t, target), 1, 1, color.RGBA{R: 255, A: 255})
	}
}

func TestRasterOffscreenLabelErrors(t *testing.T) {
	want := errors.New("rasterization failed")
	target := newRasterOffscreen(t, app.BackendOffscreenConfig{Width: 2, Height: 2, Scale: 1},
		func(app.BackendTextLabel, float32, int, int) ([]byte, error) { return nil, want })
	for _, frame := range []app.BackendFrame{
		{TextLabels: []app.BackendTextLabel{{Value: "label", Width: 1, Height: 1}}},
		{TextInput: app.BackendTextInputUpdate{Changed: true, Input: &app.BackendTextInput{Width: 1, Height: 1, State: app.BackendTextState{Value: "input"}}}},
	} {
		if err := target.Set(frame); !errors.Is(err, want) {
			t.Errorf("Set error = %v, want %v", err, want)
		}
	}
	if err := target.Set(app.BackendFrame{TextLabels: []app.BackendTextLabel{{
		Width: 1, Height: 1, Bitmap: []byte{0, 0, 0, 255}, BitmapWidth: 2, BitmapHeight: 2, BitmapStride: 8,
	}}}); err == nil {
		t.Error("malformed bitmap was accepted")
	}
}

func TestBackendAutostartServices(t *testing.T) {
	entry := app.Autostart{ID: "dev.viva.backend-test", Name: "Backend test", Args: []string{"--hidden", "two words"}}
	on, sets, queries := false, 0, 0
	var serviceErr error
	b := &externalBackend{services: app.BackendServices{
		SetAutostart: func(got app.Autostart, enabled bool) error {
			sets++
			if !reflect.DeepEqual(got, entry) {
				t.Fatalf("SetAutostart entry = %+v, want %+v", got, entry)
			}
			on = enabled
			return serviceErr
		},
		AutostartEnabled: func(id string) (bool, error) {
			queries++
			if id != entry.ID {
				t.Fatalf("AutostartEnabled id = %q", id)
			}
			return on, serviceErr
		},
	}}
	if err := app.UseBackend(b); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = app.UseBackend(nil) })
	for _, enabled := range []bool{true, false} {
		if err := app.SetAutostart(entry, enabled); err != nil {
			t.Fatal(err)
		}
		if got, err := app.AutostartEnabled(entry.ID); got != enabled || err != nil {
			t.Fatalf("AutostartEnabled = %v, %v; want %v, nil", got, err, enabled)
		}
	}
	for _, id := range []string{"", "bad/id", "bad\x00id", "bad id"} {
		if err := app.SetAutostart(app.Autostart{ID: id}, true); err == nil {
			t.Errorf("accepted autostart id %q", id)
		}
		if _, err := app.AutostartEnabled(id); err == nil {
			t.Errorf("queried autostart id %q", id)
		}
	}
	if err := app.SetAutostart(app.Autostart{ID: entry.ID, Args: []string{"bad\narg"}}, true); err == nil {
		t.Error("accepted invalid autostart argument")
	}
	if sets != 2 || queries != 2 {
		t.Fatalf("invalid requests reached backend: sets=%d queries=%d", sets, queries)
	}
	serviceErr = errors.New("backend autostart error")
	if err := app.SetAutostart(entry, true); !errors.Is(err, serviceErr) {
		t.Errorf("SetAutostart error = %v", err)
	}
	if _, err := app.AutostartEnabled(entry.ID); !errors.Is(err, serviceErr) {
		t.Errorf("AutostartEnabled error = %v", err)
	}
	b.services = app.BackendServices{}
	if err := app.SetAutostart(entry, true); !errors.Is(err, app.ErrUnsupported) {
		t.Errorf("missing SetAutostart service = %v", err)
	}
	if enabled, err := app.AutostartEnabled(entry.ID); enabled || !errors.Is(err, app.ErrUnsupported) {
		t.Errorf("missing AutostartEnabled service = %v, %v", enabled, err)
	}
}

func TestBackendWindowControlsAndTitlebarInsets(t *testing.T) {
	b := &externalBackend{}
	if err := app.UseBackend(b); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = app.UseBackend(nil) })
	for _, test := range []struct {
		name     string
		services app.BackendServices
		buttons  bool
		leading  float32
		trailing float32
	}{
		{name: "missing"},
		{name: "host controls", services: app.BackendServices{
			DrawsWindowControls: func() bool { return false },
			TitlebarInsets:      func() (float32, float32) { return 42, 7 },
		}, leading: 42, trailing: 7},
		{name: "Viva controls", services: app.BackendServices{
			DrawsWindowControls: func() bool { return true },
			TitlebarInsets:      func() (float32, float32) { return 0, 3 * app.WindowControlWidth },
		}, buttons: true, trailing: 3 * app.WindowControlWidth},
	} {
		t.Run(test.name, func(t *testing.T) {
			b.services = test.services
			if leading, trailing := app.TitlebarInsets(); leading != test.leading || trailing != test.trailing {
				t.Errorf("TitlebarInsets = %g, %g; want %g, %g", leading, trailing, test.leading, test.trailing)
			}
			for _, mode := range []app.WindowMode{app.WindowCustomTitlebar, app.WindowTitled} {
				driver, err := app.NewDriver(app.Config{
					Title: "controls", Width: 300, Height: 200, WindowMode: mode,
					Paint: func(*ui.Canvas, ui.Size) {},
				}, nil)
				if err != nil {
					t.Fatal(err)
				}
				_, hit := driver.HitTest(280, 16)
				driver.Close()
				if want := test.buttons && mode == app.WindowCustomTitlebar; hit != want {
					t.Errorf("mode %v: window control hit = %v, want %v", mode, hit, want)
				}
			}
		})
	}
}
