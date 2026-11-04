//go:build !viva_custom_backend && !linux

package app

func defaultNativeSetPrimarySelection(string) error  { return ErrUnsupported }
func defaultNativePrimarySelection() (string, error) { return "", ErrUnsupported }
