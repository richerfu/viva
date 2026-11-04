//go:build linux && !viva_custom_backend

package app

import linuxplatform "github.com/vivaui/viva/internal/platform/linux"

func defaultNativeSchemeHandlerAvailable(scheme string) (bool, error) {
	available, err := linuxplatform.SchemeHandlerAvailable(scheme)
	return available, linuxSystemError(err)
}
