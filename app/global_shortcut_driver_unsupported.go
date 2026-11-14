//go:build !viva_custom_backend && !darwin && !windows && !linux

package app

func driverRegisterGlobalShortcut(GlobalShortcut, func()) (func(), error, bool) {
	return nil, nil, false
}
