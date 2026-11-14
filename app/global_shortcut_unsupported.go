//go:build !viva_custom_backend && !windows && !linux && !darwin

package app

func defaultNativeRegisterGlobalShortcut(GlobalShortcut, func()) (func(), error) {
	return nil, ErrUnsupported
}
