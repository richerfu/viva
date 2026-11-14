//go:build !viva_custom_backend

package app

import darwinplatform "github.com/vivaui/viva/internal/platform/darwin"

func defaultNativeRegisterGlobalShortcut(s GlobalShortcut, pressed func()) (func(), error) {
	return darwinplatform.RegisterGlobalShortcut(s.Modifiers, s.Key, pressed)
}
