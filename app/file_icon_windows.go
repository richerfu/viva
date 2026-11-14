//go:build !viva_custom_backend && windows

package app

import windowsplatform "github.com/vivaui/viva/internal/platform/windows"

func defaultNativeFileIcon(path string, pixels int) ([]byte, error) {
	return windowsplatform.FileIconBGRA(path, pixels)
}
