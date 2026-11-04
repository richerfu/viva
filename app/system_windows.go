//go:build !viva_custom_backend && windows

package app

import (
	windowsplatform "github.com/vivaui/viva/internal/platform/windows"
	"github.com/vivaui/viva/ui"
)

func defaultNativeClipboardText() (string, error) { return windowsplatform.ClipboardText() }

func defaultNativeHasClipboardImage() bool { return windowsplatform.HasClipboardImage() }

func defaultNativeClipboardImagePNG() ([]byte, error) { return windowsplatform.ClipboardImagePNG() }

func defaultNativeReduceMotion() bool { return windowsplatform.ReduceMotion() }

// nativeTextEngine is the platform's text engine; nil keeps the reference metrics.
func defaultNativeTextEngine() ui.TextEngine { return windowsplatform.TextEngine() }

func defaultNativePreferredLanguages() []string { return windowsplatform.PreferredLanguages() }

func defaultNativeSetClipboardText(value string) error {
	return windowsplatform.SetClipboardText(value)
}

func defaultNativeOnClipboardChange(notify func()) error {
	return windowsplatform.OnClipboardChange(notify)
}
func defaultNativeSetCursor(cursor Cursor) error {
	return windowsplatform.SetCursorShape(uint8(cursor))
}
func defaultNativeOpenFiles(options OpenFileOptions) ([]string, error) {
	filters := make([]windowsplatform.FileFilter, len(options.Filters))
	for index, filter := range options.Filters {
		filters[index] = windowsplatform.FileFilter{Name: filter.Name, Pattern: filter.Pattern}
	}
	paths, canceled, err := windowsplatform.OpenFiles(windowsplatform.OpenFileOptions{
		Title: options.Title, Filters: filters, AllowMultiple: options.AllowMultiple,
	})
	if canceled {
		return nil, ErrCanceled
	}
	return paths, err
}

func defaultNativeOpenDirectory(options OpenDirectoryOptions) (string, error) {
	path, canceled, err := windowsplatform.OpenDirectory(windowsplatform.OpenDirectoryOptions{Title: options.Title})
	if canceled {
		return "", ErrCanceled
	}
	return path, err
}
func defaultNativeSaveFile(options SaveFileOptions) (string, error) {
	filters := make([]windowsplatform.FileFilter, len(options.Filters))
	for index, filter := range options.Filters {
		filters[index] = windowsplatform.FileFilter{Name: filter.Name, Pattern: filter.Pattern}
	}
	path, canceled, err := windowsplatform.SaveFile(windowsplatform.SaveFileOptions{
		Title: options.Title, SuggestedName: options.SuggestedName,
		DefaultExtension: options.DefaultExtension, Filters: filters,
	})
	if canceled {
		return "", ErrCanceled
	}
	return path, err
}
func defaultNativeOpenURL(target string) error { return windowsplatform.OpenURL(target) }

func defaultNativeOpenWith(executable, path string) error {
	return windowsplatform.OpenWith(executable, path)
}
func defaultNativePlaySystemSound(sound SystemSound) error {
	return windowsplatform.PlaySystemSound(uint8(sound))
}
func defaultNativeSendNotification(notification Notification) error {
	return windowsplatform.SendNotification(notification.Title, notification.Body)
}

func defaultNativeStoreSecret(service, account string, secret []byte) error {
	return secretError(windowsplatform.StoreSecret(service, account, secret))
}

func defaultNativeLoadSecret(service, account string) ([]byte, error) {
	secret, err := windowsplatform.LoadSecret(service, account)
	return secret, secretError(err)
}

func defaultNativeDeleteSecret(service, account string) error {
	return secretError(windowsplatform.DeleteSecret(service, account))
}

func defaultNativeEnableCrashDumps(relativeDir string) (string, error) {
	return windowsplatform.EnableCrashDumps(relativeDir)
}

func defaultNativeShowFatalError(title, message string) {
	windowsplatform.ShowFatalError(title, message)
}

func defaultNativeMonospaceFamilies() []string { return windowsplatform.MonospaceFamilies() }
func defaultNativeFontFamilies() []string      { return windowsplatform.FontFamilies() }

func defaultNativeAppearance() Appearance { return appearanceOf(windowsplatform.DarkAppearance()) }

func defaultNativeWatchAppearance() {} // the window procedure hears the settings change

func defaultNativeSetAppearance(appearance Appearance) {
	windowsplatform.SetAppearance(appearance == AppearanceLight, appearance == AppearanceDark)
}

func defaultNativeSetCursorImage(image *ui.CursorImage) error {
	return windowsplatform.SetCursorImage(image)
}
