//go:build js && wasm && !viva_custom_backend

package app

import webplatform "github.com/vivaui/viva/internal/platform/web"

func defaultNativeRequestClipboardImage(done func(png []byte, err error)) {
	webplatform.ReadClipboardImagePNG(func(png []byte, err error) {
		done(png, webSystemError("read clipboard image", err))
	})
}
