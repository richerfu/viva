//go:build !viva_custom_backend && windows

package app

import windowsplatform "github.com/vivaui/viva/internal/platform/windows"

func defaultNativeSetClipboardHTML(html, alt string) error {
	return windowsplatform.SetClipboardHTML(html, alt)
}
