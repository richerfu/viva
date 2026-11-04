//go:build darwin || windows || linux || viva_custom_backend

package app

func (painter *painter) fullscreenHandler(callback func(bool)) func(bool) {
	return func(fullscreen bool) {
		painter.fullscreen = fullscreen
		if callback != nil {
			callback(fullscreen)
		}
	}
}
