//go:build !viva_custom_backend && (windows || linux)

package app

// platformDrawsWindowControls is true where the system draws no minimize, maximize or
// close buttons for a window without a titlebar, so Viva draws them.
func platformDrawsWindowControls() bool { return true }

func platformTitlebarInsets() (leading, trailing float32) { return 0, 3 * WindowControlWidth }
