//go:build !viva_custom_backend && js && wasm

package proxy

import "net/url"

func nativeSystemProxy(*url.URL) (*url.URL, error) { return nil, ErrUnavailable }
