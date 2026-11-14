//go:build !viva_custom_backend && js && wasm

package app

import webplatform "github.com/vivaui/viva/internal/platform/web"

func defaultNativeSetClipboardHTML(html, alt string) error {
	return webSystemError("write clipboard html", webplatform.SetClipboardHTML(html, alt))
}
