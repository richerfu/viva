//go:build viva_custom_backend || (!windows && !linux && !darwin)

package app

// platformDrawsWindowControls is false where there is no native window to give controls.
func platformDrawsWindowControls() bool { return false }

func platformTitlebarInsets() (leading, trailing float32) { return 0, 0 }
