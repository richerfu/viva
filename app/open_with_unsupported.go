//go:build !viva_custom_backend && !windows && !darwin && !linux

package app

import "fmt"

func defaultNativeOpenWith(_, _ string) error {
	return fmt.Errorf("%w: OpenWith", ErrUnsupported)
}
