//go:build !viva_custom_backend && darwin

package app

import darwinplatform "github.com/vivaui/viva/internal/platform/darwin"

func defaultNativeSetClipboardHTML(html, alt string) error {
	return darwinplatform.SetClipboardHTML(html, alt)
}
