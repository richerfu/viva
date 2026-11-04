//go:build linux && !viva_custom_backend && !vivagtk

package app

import (
	"github.com/vivaui/viva/internal/platform"
	linuxplatform "github.com/vivaui/viva/internal/platform/linux"
)

func defaultRunNative(config platform.Config) error { return linuxplatform.Run(config) }
func defaultRunNativeApplication(start platform.ApplicationStart) error {
	return linuxplatform.RunApplication(start)
}
