//go:build !(js && wasm) || viva_custom_backend

package app

func defaultNativeRequestClipboardImage(done func(png []byte, err error)) {
	png, err := ClipboardImagePNG()
	done(png, err)
}
