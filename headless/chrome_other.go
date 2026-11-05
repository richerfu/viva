//go:build viva_custom_backend || !darwin

package headless

import (
	"image"

	"github.com/vivaui/viva/app"
)

// drawWindowChrome adds nothing: Viva's own buttons are already in the scene,
// and a custom backend supplies any additional chrome in its offscreen target.
func drawWindowChrome(*image.RGBA, app.Config, float64) {}
