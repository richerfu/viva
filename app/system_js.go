//go:build !viva_custom_backend && js && wasm

package app

import (
	"errors"
	"fmt"

	webplatform "github.com/vivaui/viva/internal/platform/web"
	"github.com/vivaui/viva/ui"
)

func defaultNativeReduceMotion() bool { return false }

// nativeTextEngine is the platform's text engine; nil keeps the reference metrics.
func defaultNativeTextEngine() ui.TextEngine { return nil }

func defaultNativePreferredLanguages() []string { return webplatform.PreferredLanguages() }

func defaultNativeClipboardText() (string, error) {
	value, err := webplatform.ClipboardText()
	return value, webSystemError("read clipboard", err)
}

func defaultNativeOnClipboardChange(notify func()) error {
	if notify == nil {
		return nil
	}
	return webSystemError("watch clipboard", webplatform.ErrUnavailable)
}

func defaultNativeSetClipboardText(value string) error {
	return webSystemError("set clipboard", webplatform.SetClipboardText(value))
}

func defaultNativeHasClipboardImage() bool { return false }

func defaultNativeClipboardImagePNG() ([]byte, error) {
	return nil, webSystemError("read clipboard image", webplatform.ClipboardImagePNG())
}

func defaultNativeSetCursor(cursor Cursor) error {
	return webSystemError("set cursor", webplatform.SetCursorShape(uint8(cursor)))
}

func defaultNativeSetCursorImage(image *ui.CursorImage) error {
	return webSystemError("set cursor image", webplatform.SetCursorImage(image))
}

func defaultNativeOpenFiles(OpenFileOptions) ([]string, error) {
	paths, err := webplatform.OpenFiles()
	return paths, webSystemError("open files", err)
}

func defaultNativeOpenDirectory(OpenDirectoryOptions) (string, error) {
	return "", ErrUnsupported
}

func defaultNativeSaveFile(SaveFileOptions) (string, error) {
	path, err := webplatform.SaveFile()
	return path, webSystemError("save file", err)
}

func defaultNativeOpenURL(target string) error {
	return webSystemError("open external URL", webplatform.OpenURL(target))
}

func defaultNativePlaySystemSound(sound SystemSound) error {
	return webSystemError("play system sound", webplatform.PlaySystemSound(uint8(sound)))
}

func defaultNativeSendNotification(notification Notification) error {
	return webSystemError("send notification", webplatform.SendNotification(notification.Title, notification.Body))
}

func webSystemError(action string, err error) error {
	if err == nil {
		return nil
	}
	switch {
	case errors.Is(err, webplatform.ErrAsyncRequired):
		return fmt.Errorf("%s: %w: %v", action, ErrAsyncRequired, err)
	case errors.Is(err, webplatform.ErrUserGestureRequired):
		return fmt.Errorf("%s: %w: %v", action, ErrUserGestureRequired, err)
	case errors.Is(err, webplatform.ErrPermissionDenied):
		return fmt.Errorf("%s: %w: %v", action, ErrPermissionDenied, err)
	case errors.Is(err, webplatform.ErrUnsupported), errors.Is(err, webplatform.ErrUnavailable):
		return fmt.Errorf("%s: %w: %v", action, ErrUnsupported, err)
	default:
		return fmt.Errorf("app: %s: %w", action, err)
	}
}

// A browser has no store for credentials that is not the page's own storage.
func defaultNativeStoreSecret(string, string, []byte) error  { return ErrSecretStoreUnavailable }
func defaultNativeLoadSecret(string, string) ([]byte, error) { return nil, ErrSecretStoreUnavailable }
func defaultNativeDeleteSecret(string, string) error         { return ErrSecretStoreUnavailable }

// A page cannot list the fonts of the machine it runs on.
func defaultNativeEnableCrashDumps(string) (string, error) { return "", ErrUnsupported }

func defaultNativeShowFatalError(string, string) {}

func defaultNativeMonospaceFamilies() []string { return nil }
func defaultNativeFontFamilies() []string      { return nil }

func defaultNativeAppearance() Appearance { return appearanceOf(webplatform.DarkAppearance()) }

func defaultNativeWatchAppearance() { webplatform.WatchAppearance() }

func defaultNativeSetAppearance(Appearance) {}
