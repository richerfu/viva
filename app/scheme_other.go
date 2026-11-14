//go:build !viva_custom_backend && (!darwin && !windows && !linux && !(js && wasm))

package app

func defaultNativeSchemeHandlerAvailable(string) (bool, error) { return false, ErrUnsupported }

func driverSchemeHandler(string) (available, handled bool) { return false, false }
