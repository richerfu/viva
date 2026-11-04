//go:build !viva_custom_backend && darwin

package app

import darwinplatform "github.com/vivaui/viva/internal/platform/darwin"

func defaultNativeSchemeHandlerAvailable(scheme string) (bool, error) {
	return darwinplatform.SchemeHandlerAvailable(scheme)
}
