package app

import "image"

// BackendOffscreenConfig describes an independent drawing target. An offscreen
// backend owns its device and synchronization; it need not create a window.
type BackendOffscreenConfig struct {
	Width, Height int
	Scale         float64
	ClearColor    BackendColor
}

type BackendOffscreen interface {
	Set(BackendFrame) error
	Image() (*image.RGBA, error)
	Close()
}

// OffscreenBackend optionally implements headless snapshots with the same
// frames and text input representation used by native windows.
type OffscreenBackend interface {
	NewOffscreen(BackendOffscreenConfig) (BackendOffscreen, error)
}

func NewBackendOffscreen(config BackendOffscreenConfig) (BackendOffscreen, error) {
	if b, ok := selectedBackend().(OffscreenBackend); ok {
		return b.NewOffscreen(config)
	}
	return nil, ErrUnsupported
}
