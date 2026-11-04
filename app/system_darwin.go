//go:build !viva_custom_backend && darwin

package app

import (
	"errors"
	"fmt"

	darwinplatform "github.com/vivaui/viva/internal/platform/darwin"
	"github.com/vivaui/viva/ui"
)

func defaultNativeClipboardText() (string, error) { return darwinplatform.ClipboardText() }

func defaultNativeHasClipboardImage() bool { return darwinplatform.HasClipboardImage() }

func defaultNativeClipboardImagePNG() ([]byte, error) { return darwinplatform.ClipboardImagePNG() }

func defaultNativeReduceMotion() bool { return darwinplatform.ReduceMotion() }

// nativeTextEngine is the platform's text engine; nil keeps the reference metrics.
func defaultNativeTextEngine() ui.TextEngine { return darwinplatform.SharedTextEngine() }

func defaultNativePreferredLanguages() []string { return darwinplatform.PreferredLanguages() }

func defaultNativeOnClipboardChange(notify func()) error {
	return darwinplatform.OnClipboardChange(notify)
}

func defaultNativeSetClipboardText(value string) error {
	return darwinplatform.SetClipboardText(value)
}

func defaultNativeSetCursor(cursor Cursor) error {
	return darwinplatform.SetCursorShape(uint8(cursor))
}

func defaultNativeSetCursorImage(image *ui.CursorImage) error {
	return darwinplatform.SetCursorImage(image)
}

func defaultNativeOpenFiles(options OpenFileOptions) ([]string, error) {
	filters := make([]darwinplatform.FileFilter, len(options.Filters))
	for index, filter := range options.Filters {
		filters[index] = darwinplatform.FileFilter{Name: filter.Name, Pattern: filter.Pattern}
	}
	paths, canceled, err := darwinplatform.OpenFiles(darwinplatform.OpenFileOptions{
		Title: options.Title, Filters: filters, AllowMultiple: options.AllowMultiple,
	})
	if canceled {
		return nil, ErrCanceled
	}
	return paths, err
}

func defaultNativeOpenDirectory(options OpenDirectoryOptions) (string, error) {
	path, canceled, err := darwinplatform.OpenDirectory(darwinplatform.OpenDirectoryOptions{Title: options.Title})
	if canceled {
		return "", ErrCanceled
	}
	return path, err
}

func defaultNativeSaveFile(options SaveFileOptions) (string, error) {
	filters := make([]darwinplatform.FileFilter, len(options.Filters))
	for index, filter := range options.Filters {
		filters[index] = darwinplatform.FileFilter{Name: filter.Name, Pattern: filter.Pattern}
	}
	path, canceled, err := darwinplatform.SaveFile(darwinplatform.SaveFileOptions{
		Title: options.Title, SuggestedName: options.SuggestedName,
		DefaultExtension: options.DefaultExtension, Filters: filters,
	})
	if canceled {
		return "", ErrCanceled
	}
	return path, err
}

func defaultNativeOpenURL(target string) error { return darwinplatform.OpenURL(target) }

func defaultNativeOpenWith(executable, path string) error {
	return darwinplatform.OpenWith(executable, path)
}

func defaultNativePlaySystemSound(sound SystemSound) error {
	return darwinplatform.PlaySystemSound(uint8(sound))
}

func defaultNativeSendNotification(notification Notification) error {
	return darwinNotificationError(darwinplatform.SendNotification(notification.Title, notification.Body))
}

func darwinNotificationError(err error) error {
	if errors.Is(err, darwinplatform.ErrPermissionDenied) {
		return fmt.Errorf("send notification: %w", ErrPermissionDenied)
	}
	return err
}

func defaultNativeStoreSecret(service, account string, secret []byte) error {
	return secretError(darwinplatform.StoreSecret(service, account, secret))
}

func defaultNativeLoadSecret(service, account string) ([]byte, error) {
	secret, err := darwinplatform.LoadSecret(service, account)
	return secret, secretError(err)
}

func defaultNativeDeleteSecret(service, account string) error {
	return secretError(darwinplatform.DeleteSecret(service, account))
}

func defaultNativeEnableCrashDumps(string) (string, error) { return "", ErrUnsupported }

func defaultNativeShowFatalError(title, message string) {
	darwinplatform.ShowFatalError(title, message)
}

func defaultNativeMonospaceFamilies() []string { return darwinplatform.MonospaceFamilies() }
func defaultNativeFontFamilies() []string      { return darwinplatform.FontFamilies() }

func defaultNativeAppearance() Appearance {
	if darwinplatform.DarkAppearance() {
		return AppearanceDark
	}
	return AppearanceLight
}

func defaultNativeWatchAppearance() { darwinplatform.WatchAppearance() }

func defaultNativeSetAppearance(appearance Appearance) {
	darwinplatform.SetAppearance(appearance == AppearanceLight, appearance == AppearanceDark)
}
