//go:build viva_custom_backend

package app

import (
	"github.com/vivaui/viva/internal/platform"
	"github.com/vivaui/viva/ui"
)

func defaultNativeReduceMotion() bool                                  { return false }
func defaultNativeTextEngine() ui.TextEngine                           { return nil }
func defaultNativePreferredLanguages() []string                        { return nil }
func defaultNativeClipboardText() (string, error)                      { return "", ErrUnsupported }
func defaultNativeSetClipboardText(value string) error                 { return ErrUnsupported }
func defaultNativeHasClipboardImage() bool                             { return false }
func defaultNativeClipboardImagePNG() ([]byte, error)                  { return nil, ErrUnsupported }
func defaultNativeOnClipboardChange(notify func()) error               { return ErrUnsupported }
func defaultNativeSetCursor(cursor Cursor) error                       { return ErrUnsupported }
func defaultNativeSetCursorImage(image *ui.CursorImage) error          { return ErrUnsupported }
func defaultNativeOpenFiles(options OpenFileOptions) ([]string, error) { return nil, ErrUnsupported }
func defaultNativeOpenDirectory(options OpenDirectoryOptions) (string, error) {
	return "", ErrUnsupported
}
func defaultNativeSaveFile(options SaveFileOptions) (string, error) { return "", ErrUnsupported }
func defaultNativeOpenURL(target string) error                      { return ErrUnsupported }
func defaultNativeOpenWith(executable, path string) error           { return ErrUnsupported }
func defaultNativePlaySystemSound(sound SystemSound) error          { return ErrUnsupported }
func defaultNativeSendNotification(notification Notification) error { return ErrUnsupported }
func defaultNativeStoreSecret(service, account string, secret []byte) error {
	return ErrSecretStoreUnavailable
}
func defaultNativeLoadSecret(service, account string) ([]byte, error) {
	return nil, ErrSecretStoreUnavailable
}
func defaultNativeDeleteSecret(service, account string) error         { return ErrSecretStoreUnavailable }
func defaultNativeEnableCrashDumps(directory string) (string, error)  { return "", ErrUnsupported }
func defaultNativeShowFatalError(title, message string)               {}
func defaultNativeMonospaceFamilies() []string                        { return nil }
func defaultNativeFontFamilies() []string                             { return nil }
func defaultNativeAppearance() Appearance                             { return AppearanceUnknown }
func defaultNativeWatchAppearance()                                   {}
func defaultNativeSetAppearance(appearance Appearance)                {}
func defaultNativeSetPrimarySelection(value string) error             { return ErrUnsupported }
func defaultNativePrimarySelection() (string, error)                  { return "", ErrUnsupported }
func defaultNativeSchemeHandlerAvailable(scheme string) (bool, error) { return false, ErrUnsupported }
func defaultNativeFileIcon(path string, pixels int) ([]byte, error)   { return nil, ErrUnsupported }
func defaultNativeHighContrast() bool                                 { return false }

func defaultRunNative(platform.Config) error                      { return ErrUnsupported }
func defaultRunNativeApplication(platform.ApplicationStart) error { return ErrUnsupported }
