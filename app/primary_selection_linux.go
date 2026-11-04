//go:build linux && !viva_custom_backend

package app

import linuxplatform "github.com/vivaui/viva/internal/platform/linux"

func defaultNativeSetPrimarySelection(text string) error {
	return linuxSystemError(linuxplatform.SetPrimaryText(text))
}

func defaultNativePrimarySelection() (string, error) {
	text, err := linuxplatform.PrimaryText()
	return text, linuxSystemError(err)
}
