//go:build (!linux && !darwin && !windows) || (darwin && !cgo)

package filewatch

func openBackend(string, *Watch) (backend, error) { return nil, ErrUnsupported }
