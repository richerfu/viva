//go:build viva_custom_backend

package headless

import (
	"image"

	"github.com/vivaui/viva/app"
)

const (
	rendererAvailable = true
	defaultScale      = 2
	rendersTextEditor = true
)

type renderer struct {
	target app.BackendOffscreen
	config app.BackendOffscreenConfig
}

func newRenderer() *renderer { return &renderer{} }

func (r *renderer) close() {
	if r.target != nil {
		r.target.Close()
		r.target = nil
	}
}

func (r *renderer) render(s app.Snapshot, scale float64) (*image.RGBA, error) {
	c := app.BackendOffscreenConfig{Width: s.Width, Height: s.Height, Scale: scale,
		ClearColor: app.BackendColor{Red: s.Clear.Red, Green: s.Clear.Green, Blue: s.Clear.Blue, Alpha: s.Clear.Alpha}}
	if r.target == nil || r.config != c {
		r.close()
		var err error
		r.target, err = app.NewBackendOffscreen(c)
		if err != nil {
			return nil, err
		}
		r.config = c
	}
	rectangles, labels, input := scenery(s)
	if err := r.target.Set(app.BackendFrame{Rectangles: rectangles, TextLabels: labels,
		TextInput: app.BackendTextInputUpdate{Changed: true, Input: input}}); err != nil {
		return nil, err
	}
	return r.target.Image()
}
