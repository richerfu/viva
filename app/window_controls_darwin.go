//go:build !viva_custom_backend && darwin

package app

// platformDrawsWindowControls is false on macOS: the system draws the traffic lights.
func platformDrawsWindowControls() bool { return false }

func platformTitlebarInsets() (leading, trailing float32) { return 70, 0 }
