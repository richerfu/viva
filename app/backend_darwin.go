//go:build !viva_custom_backend

package app

import (
	"github.com/vivaui/viva/internal/platform"
	"github.com/vivaui/viva/internal/platform/darwin"
)

func defaultRunNative(config platform.Config) error { return darwin.Run(config) }
func defaultRunNativeApplication(start platform.ApplicationStart) error {
	return darwin.RunApplication(start)
}
