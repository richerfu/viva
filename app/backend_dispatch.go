package app

import (
	"github.com/vivaui/viva/internal/platform"
	"github.com/vivaui/viva/ui"
)

func backendSetAutostart(entry Autostart, enabled bool) (bool, error) {
	if b := selectedBackend(); b != nil {
		if f := b.Services().SetAutostart; f != nil {
			return true, f(entry, enabled)
		}
		return true, ErrUnsupported
	}
	return false, nil
}

func backendAutostartEnabled(id string) (bool, bool, error) {
	if b := selectedBackend(); b != nil {
		if f := b.Services().AutostartEnabled; f != nil {
			enabled, err := f(id)
			return enabled, true, err
		}
		return false, true, ErrUnsupported
	}
	return false, false, nil
}

func nativeReduceMotion() bool {
	if b := selectedBackend(); b != nil {
		if f := b.Services().ReduceMotion; f != nil {
			return f()
		}
		return false
	}
	return defaultNativeReduceMotion()
}
func nativeTextEngine() ui.TextEngine {
	if b := selectedBackend(); b != nil {
		return b.TextEngine()
	}
	return defaultNativeTextEngine()
}
func nativePreferredLanguages() []string {
	if b := selectedBackend(); b != nil {
		if f := b.Services().PreferredLanguages; f != nil {
			return f()
		}
		return nil
	}
	return defaultNativePreferredLanguages()
}
func nativeClipboardText() (string, error) {
	if b := selectedBackend(); b != nil {
		if f := b.Services().ClipboardText; f != nil {
			return f()
		}
		return "", ErrUnsupported
	}
	return defaultNativeClipboardText()
}
func nativeSetClipboardText(value string) error {
	if b := selectedBackend(); b != nil {
		if f := b.Services().SetClipboardText; f != nil {
			return f(value)
		}
		return ErrUnsupported
	}
	return defaultNativeSetClipboardText(value)
}
func nativeHasClipboardImage() bool {
	if b := selectedBackend(); b != nil {
		if f := b.Services().HasClipboardImage; f != nil {
			return f()
		}
		return false
	}
	return defaultNativeHasClipboardImage()
}
func nativeClipboardImagePNG() ([]byte, error) {
	if b := selectedBackend(); b != nil {
		if f := b.Services().ClipboardImagePNG; f != nil {
			return f()
		}
		return nil, ErrUnsupported
	}
	return defaultNativeClipboardImagePNG()
}
func nativeOnClipboardChange(notify func()) error {
	if b := selectedBackend(); b != nil {
		if f := b.Services().OnClipboardChange; f != nil {
			return f(notify)
		}
		return ErrUnsupported
	}
	return defaultNativeOnClipboardChange(notify)
}
func nativeSetCursor(cursor Cursor) error {
	if b := selectedBackend(); b != nil {
		if f := b.Services().SetCursor; f != nil {
			return f(cursor)
		}
		return ErrUnsupported
	}
	return defaultNativeSetCursor(cursor)
}
func nativeSetCursorImage(image *ui.CursorImage) error {
	if b := selectedBackend(); b != nil {
		if f := b.Services().SetCursorImage; f != nil {
			return f(image)
		}
		return ErrUnsupported
	}
	return defaultNativeSetCursorImage(image)
}
func nativeOpenFiles(options OpenFileOptions) ([]string, error) {
	if b := selectedBackend(); b != nil {
		if f := b.Services().OpenFiles; f != nil {
			return f(options)
		}
		return nil, ErrUnsupported
	}
	return defaultNativeOpenFiles(options)
}
func nativeOpenDirectory(options OpenDirectoryOptions) (string, error) {
	if b := selectedBackend(); b != nil {
		if f := b.Services().OpenDirectory; f != nil {
			return f(options)
		}
		return "", ErrUnsupported
	}
	return defaultNativeOpenDirectory(options)
}
func nativeSaveFile(options SaveFileOptions) (string, error) {
	if b := selectedBackend(); b != nil {
		if f := b.Services().SaveFile; f != nil {
			return f(options)
		}
		return "", ErrUnsupported
	}
	return defaultNativeSaveFile(options)
}
func nativeOpenURL(target string) error {
	if b := selectedBackend(); b != nil {
		if f := b.Services().OpenURL; f != nil {
			return f(target)
		}
		return ErrUnsupported
	}
	return defaultNativeOpenURL(target)
}
func nativeOpenWith(executable, path string) error {
	if b := selectedBackend(); b != nil {
		if f := b.Services().OpenWith; f != nil {
			return f(executable, path)
		}
		return ErrUnsupported
	}
	return defaultNativeOpenWith(executable, path)
}
func nativePlaySystemSound(sound SystemSound) error {
	if b := selectedBackend(); b != nil {
		if f := b.Services().PlaySystemSound; f != nil {
			return f(sound)
		}
		return ErrUnsupported
	}
	return defaultNativePlaySystemSound(sound)
}
func nativeSendNotification(notification Notification) error {
	if b := selectedBackend(); b != nil {
		if f := b.Services().SendNotification; f != nil {
			return f(notification)
		}
		return ErrUnsupported
	}
	return defaultNativeSendNotification(notification)
}
func nativeStoreSecret(service, account string, secret []byte) error {
	if b := selectedBackend(); b != nil {
		if f := b.Services().StoreSecret; f != nil {
			return f(service, account, secret)
		}
		return ErrSecretStoreUnavailable
	}
	return defaultNativeStoreSecret(service, account, secret)
}
func nativeLoadSecret(service, account string) ([]byte, error) {
	if b := selectedBackend(); b != nil {
		if f := b.Services().LoadSecret; f != nil {
			return f(service, account)
		}
		return nil, ErrSecretStoreUnavailable
	}
	return defaultNativeLoadSecret(service, account)
}
func nativeDeleteSecret(service, account string) error {
	if b := selectedBackend(); b != nil {
		if f := b.Services().DeleteSecret; f != nil {
			return f(service, account)
		}
		return ErrSecretStoreUnavailable
	}
	return defaultNativeDeleteSecret(service, account)
}
func nativeEnableCrashDumps(directory string) (string, error) {
	if b := selectedBackend(); b != nil {
		if f := b.Services().EnableCrashDumps; f != nil {
			return f(directory)
		}
		return "", ErrUnsupported
	}
	return defaultNativeEnableCrashDumps(directory)
}
func nativeShowFatalError(title, message string) {
	if b := selectedBackend(); b != nil {
		if f := b.Services().ShowFatalError; f != nil {
			f(title, message)
			return
		}
		return
	}
	defaultNativeShowFatalError(title, message)
}
func nativeMonospaceFamilies() []string {
	if b := selectedBackend(); b != nil {
		if f := b.Services().MonospaceFamilies; f != nil {
			return f()
		}
		return nil
	}
	return defaultNativeMonospaceFamilies()
}
func nativeFontFamilies() []string {
	if b := selectedBackend(); b != nil {
		if f := b.Services().FontFamilies; f != nil {
			return f()
		}
		return nil
	}
	return defaultNativeFontFamilies()
}
func nativeAppearance() Appearance {
	if b := selectedBackend(); b != nil {
		if f := b.Services().Appearance; f != nil {
			return f()
		}
		return AppearanceUnknown
	}
	return defaultNativeAppearance()
}
func nativeWatchAppearance() {
	if b := selectedBackend(); b != nil {
		if f := b.Services().WatchAppearance; f != nil {
			f()
			return
		}
		return
	}
	defaultNativeWatchAppearance()
}
func nativeSetAppearance(appearance Appearance) {
	if b := selectedBackend(); b != nil {
		if f := b.Services().SetAppearance; f != nil {
			f(appearance)
			return
		}
		return
	}
	defaultNativeSetAppearance(appearance)
}
func nativeSetPrimarySelection(value string) error {
	if b := selectedBackend(); b != nil {
		if f := b.Services().SetPrimarySelection; f != nil {
			return f(value)
		}
		return ErrUnsupported
	}
	return defaultNativeSetPrimarySelection(value)
}
func nativePrimarySelection() (string, error) {
	if b := selectedBackend(); b != nil {
		if f := b.Services().PrimarySelection; f != nil {
			return f()
		}
		return "", ErrUnsupported
	}
	return defaultNativePrimarySelection()
}
func nativeSchemeHandlerAvailable(scheme string) (bool, error) {
	if b := selectedBackend(); b != nil {
		if f := b.Services().SchemeHandlerAvailable; f != nil {
			return f(scheme)
		}
		return false, ErrUnsupported
	}
	return defaultNativeSchemeHandlerAvailable(scheme)
}
func nativeFileIcon(path string, pixels int) ([]byte, error) {
	if b := selectedBackend(); b != nil {
		if f := b.Services().FileIcon; f != nil {
			return f(path, pixels)
		}
		return nil, ErrUnsupported
	}
	return defaultNativeFileIcon(path, pixels)
}
func nativeHighContrast() bool {
	if b := selectedBackend(); b != nil {
		if f := b.Services().HighContrast; f != nil {
			return f()
		}
		return false
	}
	return defaultNativeHighContrast()
}

func runNative(c platform.Config) error {
	if b := selectedBackend(); b != nil {
		return b.Run(c)
	}
	return defaultRunNative(c)
}
func runNativeApplication(start platform.ApplicationStart) error {
	if b := selectedBackend(); b != nil {
		return b.RunApplication(start)
	}
	return defaultRunNativeApplication(start)
}

func nativeRequestClipboardImage(done func([]byte, error)) {
	if b := selectedBackend(); b != nil {
		if f := b.Services().RequestClipboardImage; f != nil {
			f(done)
			return
		}
		png, err := nativeClipboardImagePNG()
		done(png, err)
		return
	}
	defaultNativeRequestClipboardImage(done)
}
