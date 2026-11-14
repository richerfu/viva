//go:build viva_custom_backend || (!darwin && !windows && !linux && !(js && wasm))

package app

func nativePreventSleep(string) (func(), error) { return nil, ErrSleepPreventionUnavailable }
