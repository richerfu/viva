//go:build !viva_custom_backend && !darwin

package app

func defaultNativeFileIcon(string, int) ([]byte, error) { return nil, ErrUnsupported }
