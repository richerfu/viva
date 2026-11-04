//go:build darwin || windows || linux || viva_custom_backend

package app_test

import (
	"errors"
	"testing"

	"github.com/vivaui/viva/app"
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
