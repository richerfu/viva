//go:build linux && !viva_custom_backend

package app

import (
	"errors"
	"fmt"
	"os"

	"github.com/vivaui/viva/gfx"
	linuxplatform "github.com/vivaui/viva/internal/platform/linux"
)

// nativeReduceMotion is the desktop's preference for less animation: the settings portal,
// the XSETTINGS manager, or KDE's animation speed, whichever the desktop has.
func defaultNativeReduceMotion() bool { return linuxplatform.ReduceMotion() }

// nativeTextEngine is the platform's text engine; nil keeps the reference metrics.
func defaultNativeTextEngine() gfx.TextEngine { return linuxplatform.TextEngine() }

func defaultNativePreferredLanguages() []string { return languagesFromEnvironment(os.Getenv) }

func defaultNativeClipboardText() (string, error) {
	value, err := linuxplatform.ClipboardText()
	return value, linuxSystemError(err)
}

func defaultNativeOnClipboardChange(notify func()) error {
	return linuxSystemError(linuxplatform.OnClipboardChange(notify))
}

func defaultNativeSetClipboardText(value string) error {
	return linuxSystemError(linuxplatform.SetClipboardText(value))
}

func defaultNativeHasClipboardImage() bool { return linuxplatform.HasClipboardImage() }

func defaultNativeClipboardFiles() ([]string, error) {
	paths, err := linuxplatform.ClipboardFiles()
	return paths, linuxSystemError(err)
}

func defaultNativeSetClipboardFiles(paths []string) error {
	return linuxSystemError(linuxplatform.SetClipboardFiles(paths))
}

func defaultNativeClipboardImagePNG() ([]byte, error) {
	png, err := linuxplatform.ClipboardImagePNG()
	return png, linuxSystemError(err)
}

func defaultNativeSetCursor(cursor Cursor) error {
	return linuxSystemError(linuxplatform.SetCursorShape(uint8(cursor)))
}

func defaultNativeSetCursorImage(image *gfx.CursorImage) error {
	return linuxSystemError(linuxplatform.SetCursorImage(image))
}

func defaultNativeOpenFiles(options OpenFileOptions) ([]string, error) {
	filters := make([]linuxplatform.FileFilter, len(options.Filters))
	for index, filter := range options.Filters {
		filters[index] = linuxplatform.FileFilter{Name: filter.Name, Pattern: filter.Pattern}
	}
	paths, canceled, err := linuxplatform.OpenFiles(linuxplatform.OpenFileOptions{
		Title: options.Title, Filters: filters, AllowMultiple: options.AllowMultiple,
	})
	if canceled {
		return nil, ErrCanceled
	}
	return paths, linuxSystemError(err)
}

func defaultNativeOpenDirectory(options OpenDirectoryOptions) (string, error) {
	path, canceled, err := linuxplatform.OpenDirectory(linuxplatform.OpenDirectoryOptions{Title: options.Title})
	if canceled {
		return "", ErrCanceled
	}
	return path, linuxSystemError(err)
}

func defaultNativeSaveFile(options SaveFileOptions) (string, error) {
	filters := make([]linuxplatform.FileFilter, len(options.Filters))
	for index, filter := range options.Filters {
		filters[index] = linuxplatform.FileFilter{Name: filter.Name, Pattern: filter.Pattern}
	}
	path, canceled, err := linuxplatform.SaveFile(linuxplatform.SaveFileOptions{
		Title: options.Title, SuggestedName: options.SuggestedName,
		DefaultExtension: options.DefaultExtension, Filters: filters,
	})
	if canceled {
		return "", ErrCanceled
	}
	return path, linuxSystemError(err)
}

func defaultNativeSaveFileWithCheckbox(options SaveFileOptions, checkbox SaveFileCheckbox) (SaveFileSelection, error) {
	filters := make([]linuxplatform.FileFilter, len(options.Filters))
	for index, filter := range options.Filters {
		filters[index] = linuxplatform.FileFilter{Name: filter.Name, Pattern: filter.Pattern}
	}
	path, checked, canceled, err := linuxplatform.SaveFileWithCheckbox(linuxplatform.SaveFileOptions{
		Title: options.Title, SuggestedName: options.SuggestedName,
		DefaultExtension: options.DefaultExtension, Filters: filters,
	}, checkbox.Label, checkbox.Checked)
	if canceled {
		return SaveFileSelection{}, ErrCanceled
	}
	if err != nil {
		return SaveFileSelection{}, linuxSystemError(err)
	}
	return SaveFileSelection{Path: path, Checked: checked}, nil
}

func defaultNativeOpenURL(target string) error {
	return linuxSystemError(linuxplatform.OpenURL(target))
}

func defaultNativeRevealPath(path string) error {
	return linuxSystemError(linuxplatform.RevealPath(path))
}

func defaultNativeOpenWith(executable, path string) error {
	return linuxSystemError(linuxplatform.OpenWith(executable, path))
}

func defaultNativePlaySystemSound(sound SystemSound) error {
	return linuxSystemError(linuxplatform.PlaySystemSound(uint8(sound)))
}

func defaultNativeSendNotification(notification Notification) error {
	return linuxSystemError(linuxplatform.SendNotification(notification.Title, notification.Body))
}

func linuxSystemError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, linuxplatform.ErrUnsupported) {
		return fmt.Errorf("%w: %s", ErrUnsupported, err)
	}
	return err
}

func defaultNativeStoreSecret(service, account string, secret []byte) error {
	return secretError(linuxplatform.StoreSecret(service, account, secret))
}

func defaultNativeLoadSecret(service, account string) ([]byte, error) {
	secret, err := linuxplatform.LoadSecret(service, account)
	return secret, secretError(err)
}

func defaultNativeDeleteSecret(service, account string) error {
	return secretError(linuxplatform.DeleteSecret(service, account))
}

func defaultNativeEnableCrashDumps(string) (string, error) { return "", ErrUnsupported }

func defaultNativeShowFatalError(title, message string) { linuxplatform.ShowFatalError(title, message) }

func defaultNativeMonospaceFamilies() []string { return linuxplatform.MonospaceFamilies() }
func defaultNativeFontFamilies() []string      { return linuxplatform.FontFamilies() }

func defaultNativeAppearance() Appearance { return appearanceOf(linuxplatform.DarkAppearance()) }

func defaultNativeWatchAppearance() {} // the event loop listens while an observer is set (platform.DesktopWatched)

func defaultNativeSetAppearance(Appearance) {} // Viva draws the window buttons
