//go:build !viva_custom_backend && darwin

package app

import darwinplatform "github.com/vivaui/viva/internal/platform/darwin"

func defaultNativeFileIcon(path string, pixels int) ([]byte, error) {
	return darwinplatform.FileIconBGRA(path, pixels)
}
