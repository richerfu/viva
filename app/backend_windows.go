//go:build !viva_custom_backend

package app

import (
	"github.com/vivaui/viva/internal/platform"
	"github.com/vivaui/viva/internal/platform/windows"
)

func defaultRunNative(config platform.Config) error { return windows.Run(config) }
func defaultRunNativeApplication(start platform.ApplicationStart) error {
	return windows.RunApplication(start)
}
