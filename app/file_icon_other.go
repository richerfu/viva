//go:build !viva_custom_backend && (!darwin && !windows)

package app

func defaultNativeFileIcon(string, int) ([]byte, error) { return nil, ErrUnsupported }
