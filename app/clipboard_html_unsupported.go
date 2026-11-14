//go:build !viva_custom_backend && !darwin && !windows && !linux && !(js && wasm)

package app

func defaultNativeSetClipboardHTML(html, alt string) error { return ErrUnsupported }
