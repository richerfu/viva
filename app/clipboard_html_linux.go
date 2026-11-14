//go:build !viva_custom_backend && linux

package app

import linuxplatform "github.com/vivaui/viva/internal/platform/linux"

func defaultNativeSetClipboardHTML(html, alt string) error {
	return linuxSystemError(linuxplatform.SetClipboardHTML(html, alt))
}
