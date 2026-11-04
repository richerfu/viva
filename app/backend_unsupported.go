//go:build !viva_custom_backend && !darwin && !windows && !linux

package app

import "github.com/vivaui/viva/internal/platform"

// Browser painting uses RunCanvas's web path. Other non-desktop targets have
// no built-in native window loop; custom hosts select viva_custom_backend.
func defaultRunNative(platform.Config) error                      { return ErrUnsupported }
func defaultRunNativeApplication(platform.ApplicationStart) error { return ErrUnsupported }
