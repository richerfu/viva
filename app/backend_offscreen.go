package app

import (
	"fmt"
	"image"
	"math"
	"strings"
	"unicode/utf8"

	"github.com/vivaui/viva/internal/platform"
	"github.com/vivaui/viva/internal/raster"
	"github.com/vivaui/viva/internal/render"
	"github.com/vivaui/viva/internal/scene"
	"github.com/vivaui/viva/ui"
)

// BackendOffscreenConfig describes an independent drawing target. An offscreen
// backend owns its device and synchronization; it need not create a window.
type BackendOffscreenConfig struct {
	Width, Height int
	Scale         float64
	ClearColor    BackendColor
}

type BackendOffscreen interface {
	Set(BackendFrame) error
	Image() (*image.RGBA, error)
	Close()
}

// OffscreenBackend optionally implements headless snapshots with the same
// frames and text input representation used by native windows.
type OffscreenBackend interface {
	NewOffscreen(BackendOffscreenConfig) (BackendOffscreen, error)
}

func NewBackendOffscreen(config BackendOffscreenConfig) (BackendOffscreen, error) {
	if b, ok := selectedBackend().(OffscreenBackend); ok {
		return b.NewOffscreen(config)
	}
	return nil, ErrUnsupported
}

// BackendLabelRasterizer draws a label into width by height device pixels at
// scale. It returns tightly packed, premultiplied BGRA pixels, top row first.
// The returned buffer belongs to the offscreen renderer and must not be reused.
type BackendLabelRasterizer func(label BackendTextLabel, scale float32, width, height int) (bgra []byte, err error)

// NewRasterOffscreen creates a pure-Go scene renderer. Dimensions are logical
// pixels; scale must be in [1, 64], with at most 16384 device pixels per axis.
// A nil rasterize omits text but still draws shapes, bitmaps and native-view
// holes. The renderer is not safe for concurrent use.
func NewRasterOffscreen(config BackendOffscreenConfig, rasterize BackendLabelRasterizer) (BackendOffscreen, error) {
	const limit = 16384
	if config.Width < 1 || config.Height < 1 || config.Width > limit || config.Height > limit ||
		math.IsNaN(config.Scale) || config.Scale < 1 || config.Scale > 64 {
		return nil, fmt.Errorf("app: invalid offscreen size %dx%d at %gx", config.Width, config.Height, config.Scale)
	}
	width := int(math.Ceil(float64(config.Width) * config.Scale))
	height := int(math.Ceil(float64(config.Height) * config.Scale))
	if width > limit || height > limit {
		return nil, fmt.Errorf("app: offscreen size %dx%d at %gx is beyond %d device pixels", config.Width, config.Height, config.Scale, limit)
	}
	return &rasterOffscreen{
		canvas:    raster.NewCanvas(width, height),
		scale:     float32(config.Scale),
		clear:     scene.Color{Red: float32(config.ClearColor.Red), Green: float32(config.ClearColor.Green), Blue: float32(config.ClearColor.Blue), Alpha: float32(config.ClearColor.Alpha)},
		opaque:    config.ClearColor.Alpha >= 1,
		rasterize: rasterize,
	}, nil
}

type rasterOffscreen struct {
	canvas    *raster.Canvas
	raster    raster.Rasterizer
	scale     float32
	clear     scene.Color
	opaque    bool
	rasterize BackendLabelRasterizer
	rects     []render.Rectangle
	labels    []raster.Label
	editor    raster.Label
	holes     []raster.Hole
}

func (offscreen *rasterOffscreen) Set(frame BackendFrame) error {
	if offscreen.canvas == nil {
		return fmt.Errorf("app: offscreen renderer is closed")
	}
	offscreen.rects = append(offscreen.rects[:0], frame.Rectangles...)
	clear(offscreen.labels)
	offscreen.labels = offscreen.labels[:0]
	for index, label := range frame.TextLabels {
		prepared, ok, err := offscreen.prepareLabel(label)
		if err != nil {
			return fmt.Errorf("app: label %d: %w", index, err)
		}
		if ok {
			offscreen.labels = append(offscreen.labels, prepared)
		}
	}
	if frame.TextInput.Changed {
		offscreen.editor = raster.Label{}
		input := frame.TextInput.Input
		if input != nil && !input.HideNativeOverlays && input.State.Value != "" && input.Width > 0 && input.Height > 0 && offscreen.rasterize != nil {
			value := input.State.Value
			if input.Secure {
				value = strings.Repeat("•", utf8.RuneCountInString(value))
			}
			label := BackendTextLabel{
				Value: value, X: input.X, Y: input.Y, Width: input.Width, Height: input.Height,
				FontSize: input.FontSize, Weight: 400, Layer: input.Layer,
				Mono: input.Terminal || input.MonoFace != 0, MonoFace: input.MonoFace, Color: input.Color,
				Clipped: true, Clip: ui.Rect{X: input.X, Y: input.Y, Width: input.Width, Height: input.Height},
			}
			prepared, ok, err := offscreen.prepareLabel(label)
			if err != nil {
				return fmt.Errorf("app: text input: %w", err)
			}
			if ok {
				offscreen.editor = prepared
			}
		}
	}
	if offscreen.editor.Width > 0 {
		offscreen.labels = append(offscreen.labels, offscreen.editor)
	}
	if frame.NativeViews != nil {
		offscreen.holes = offscreen.holes[:0]
		for _, view := range frame.NativeViews {
			if view.Bounds.Width <= 0 || view.Bounds.Height <= 0 {
				continue
			}
			if box, visible := render.ScissorOf(scene.Rect(view.Bounds), offscreen.scale, offscreen.canvas.Width, offscreen.canvas.Height); visible {
				offscreen.holes = append(offscreen.holes, raster.Hole{Box: box, Layer: view.Layer})
			}
		}
	}
	return nil
}

func (offscreen *rasterOffscreen) prepareLabel(label BackendTextLabel) (raster.Label, bool, error) {
	if label.Width <= 0 || label.Height <= 0 || (label.Value == "" && len(label.Bitmap) == 0) {
		return raster.Label{}, false, nil
	}
	scale := offscreen.scale
	x, y := platform.DevicePosition(label.X, label.OriginX, scale), platform.DevicePosition(label.Y, label.OriginY, scale)
	var pixels []byte
	var width, height int
	if len(label.Bitmap) > 0 {
		width, height = label.BitmapWidth, label.BitmapHeight
		stride := label.BitmapStride
		// Validate without overflowing on malformed dimensions or strides.
		if width <= 0 || height <= 0 || width > len(label.Bitmap)/4 || stride < width*4 ||
			height-1 > (len(label.Bitmap)-width*4)/stride {
			return raster.Label{}, false, fmt.Errorf("invalid bitmap dimensions or stride")
		}
		shownWidth, shownHeight := platform.DeviceExtent(label.X, label.Width, scale), platform.DeviceExtent(label.Y, label.Height, scale)
		if shownWidth <= 0 || shownHeight <= 0 {
			return raster.Label{}, false, nil
		}
		if shownWidth <= width && shownHeight <= height && (shownWidth < width || shownHeight < height) {
			pixels = platform.ResampleBGRA(label.Bitmap, stride, width, height, shownWidth, shownHeight)
			width, height = shownWidth, shownHeight
		} else {
			// Copy even tightly packed bitmaps: frames are borrowed until Set returns.
			pixels = make([]byte, width*height*4)
			for row := range height {
				copy(pixels[row*width*4:(row+1)*width*4], label.Bitmap[row*stride:row*stride+width*4])
			}
		}
		x, y = platform.DevicePosition(label.X, 0, scale), platform.DevicePosition(label.Y, 0, scale)
	} else {
		if offscreen.rasterize == nil {
			return raster.Label{}, false, nil
		}
		width = int(math.Ceil(float64(label.Width * scale)))
		height = int(math.Ceil(float64(label.Height * scale)))
		var err error
		pixels, err = offscreen.rasterize(label, scale, width, height)
		if err != nil {
			return raster.Label{}, false, err
		}
	}
	prepared := raster.Label{Pix: pixels, Stride: width * 4, Width: width, Height: height, X: x, Y: y, Layer: label.Layer}
	if label.Clipped {
		left := max(int(math.Floor(float64(label.Clip.X*scale))), 0)
		top := max(int(math.Floor(float64(label.Clip.Y*scale))), 0)
		right := min(int(math.Ceil(float64((label.Clip.X+label.Clip.Width)*scale))), offscreen.canvas.Width)
		bottom := min(int(math.Ceil(float64((label.Clip.Y+label.Clip.Height)*scale))), offscreen.canvas.Height)
		if right <= left || bottom <= top {
			return raster.Label{}, false, nil
		}
		prepared.Clipped, prepared.Clip = true, render.Scissor{Left: left, Top: top, Right: right, Bottom: bottom}
	}
	return prepared, true, nil
}

func (offscreen *rasterOffscreen) Image() (*image.RGBA, error) {
	if offscreen.canvas == nil {
		return nil, fmt.Errorf("app: offscreen renderer is closed")
	}
	offscreen.raster.Draw(offscreen.canvas, raster.Params{Scale: offscreen.scale, Clear: offscreen.clear, Holes: offscreen.holes}, offscreen.rects, offscreen.labels)
	canvas := offscreen.canvas
	picture := image.NewRGBA(image.Rect(0, 0, canvas.Width, canvas.Height))
	for y := range canvas.Height {
		source := canvas.Pix[y*canvas.Stride : y*canvas.Stride+canvas.Width*4]
		target := picture.Pix[y*picture.Stride : y*picture.Stride+canvas.Width*4]
		for i := 0; i < len(source); i += 4 {
			target[i], target[i+1], target[i+2], target[i+3] = source[i+2], source[i+1], source[i], source[i+3]
			if offscreen.opaque {
				target[i+3] = 255
			}
		}
	}
	return picture, nil
}

func (offscreen *rasterOffscreen) Close() {
	*offscreen = rasterOffscreen{}
}
