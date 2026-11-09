//go:build viva_custom_backend

package headless

import (
	"math"

	"github.com/vivaui/viva/app"
	"github.com/vivaui/viva/ui"
)

// Exercise custom headless rendering without native windows or fonts. Colored
// blocks stand in for glyphs; the tests cover frame routing and compositing,
// while a host adapter is responsible for testing its own font rasterizer.
func init() {
	if err := app.UseBackend(rasterTestBackend{}); err != nil {
		panic(err)
	}
}

type rasterTestBackend struct{}

func (rasterTestBackend) Run(app.BackendConfig) error { return app.ErrUnsupported }
func (rasterTestBackend) RunApplication(app.BackendApplicationStart) error {
	return app.ErrUnsupported
}
func (rasterTestBackend) Services() app.BackendServices {
	return app.BackendServices{
		RequestClipboardImage: func(done func([]byte, error)) {
			// Use the active headless driver's simulated clipboard.
			pixels, err := app.ClipboardImagePNG()
			done(pixels, err)
		},
	}
}
func (rasterTestBackend) TextEngine() ui.TextEngine {
	return rasterTestTextEngine{ui.ReferenceTextEngine()}
}
func (rasterTestBackend) NewOffscreen(config app.BackendOffscreenConfig) (app.BackendOffscreen, error) {
	return app.NewRasterOffscreen(config, rasterTestLabel)
}

// Distinct metrics let the existing engine-routing test distinguish the
// backend's measurements from an explicit ReferenceText option.
type rasterTestTextEngine struct{ ui.TextEngine }

func (e rasterTestTextEngine) Width(value string, size, weight float32) float32 {
	return e.TextEngine.Width(value, size*1.1, weight)
}
func (e rasterTestTextEngine) Caret(value string, size, weight float32, offset int) float32 {
	return e.TextEngine.Caret(value, size*1.1, weight, offset)
}
func (e rasterTestTextEngine) Selection(dst []ui.TextSpan, value string, size, weight float32, start, end int) []ui.TextSpan {
	return e.TextEngine.Selection(dst, value, size*1.1, weight, start, end)
}
func (e rasterTestTextEngine) OffsetAt(value string, size, weight, x float32) int {
	return e.TextEngine.OffsetAt(value, size*1.1, weight, x)
}

func rasterTestLabel(label app.BackendTextLabel, scale float32, width, height int) ([]byte, error) {
	pixels := make([]byte, width*height*4)
	fill := func(left, right float32, color ui.Color) {
		x0 := max(0, int(math.Round(float64(left*scale))))
		x1 := min(width, int(math.Round(float64(right*scale))))
		for y := height / 4; y < height-height/4; y++ {
			for x := x0; x < x1; x++ {
				i := (y*width + x) * 4
				pixels[i] = byte(color.Blue*color.Alpha*255 + 0.5)
				pixels[i+1] = byte(color.Green*color.Alpha*255 + 0.5)
				pixels[i+2] = byte(color.Red*color.Alpha*255 + 0.5)
				pixels[i+3] = byte(color.Alpha*255 + 0.5)
			}
		}
	}
	if len(label.Runs) > 0 {
		for _, run := range label.Runs {
			fill(float32(run.Column)*label.CellWidth, float32(run.Column+run.Cells)*label.CellWidth, run.Color)
		}
	} else {
		width := (rasterTestBackend{}).TextEngine().Width(label.Value, label.FontSize, label.Weight)
		fill(0, width, ui.Color{Red: float32(label.Color.Red), Green: float32(label.Color.Green), Blue: float32(label.Color.Blue), Alpha: float32(label.Color.Alpha)})
	}
	return pixels, nil
}
