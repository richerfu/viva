//go:build !viva_custom_backend && !windows && !darwin && !linux && !(js && wasm)

package app

import (
	"os"

	"github.com/vivaui/viva/ui"
)

func defaultNativeReduceMotion() bool                                 { return false }
func defaultNativeTextEngine() ui.TextEngine                          { return nil }
func defaultNativePreferredLanguages() []string                       { return languagesFromEnvironment(os.Getenv) }
func defaultNativeClipboardText() (string, error)                     { return "", ErrUnsupported }
func defaultNativeSetClipboardText(string) error                      { return ErrUnsupported }
func defaultNativeHasClipboardImage() bool                            { return false }
func defaultNativeClipboardImagePNG() ([]byte, error)                 { return nil, ErrUnsupported }
func defaultNativeOnClipboardChange(func()) error                     { return ErrUnsupported }
func defaultNativeSetCursor(Cursor) error                             { return ErrUnsupported }
func defaultNativeSetCursorImage(*ui.CursorImage) error               { return ErrUnsupported }
func defaultNativeOpenFiles(OpenFileOptions) ([]string, error)        { return nil, ErrUnsupported }
func defaultNativeOpenDirectory(OpenDirectoryOptions) (string, error) { return "", ErrUnsupported }
func defaultNativeSaveFile(SaveFileOptions) (string, error)           { return "", ErrUnsupported }
func defaultNativeOpenURL(string) error                               { return ErrUnsupported }
func defaultNativePlaySystemSound(SystemSound) error                  { return ErrUnsupported }
func defaultNativeSendNotification(Notification) error                { return ErrUnsupported }
func defaultNativeStoreSecret(string, string, []byte) error           { return ErrSecretStoreUnavailable }
func defaultNativeLoadSecret(string, string) ([]byte, error)          { return nil, ErrSecretStoreUnavailable }
func defaultNativeDeleteSecret(string, string) error                  { return ErrSecretStoreUnavailable }

func defaultNativeEnableCrashDumps(string) (string, error) { return "", ErrUnsupported }

func defaultNativeShowFatalError(string, string) {}

func defaultNativeMonospaceFamilies() []string { return nil }
func defaultNativeFontFamilies() []string      { return nil }

func defaultNativeAppearance() Appearance { return AppearanceUnknown }

func defaultNativeWatchAppearance() {}

func defaultNativeSetAppearance(Appearance) {}
