// Package autostart starts the application when the user logs in. It writes
// what each system reads at login (a launch agent, an autostart entry, a Run
// registry value) and runs nothing itself.
package autostart

import (
	"errors"
	"fmt"
	"strings"
)

// ErrUnsupported means this system has no login-start mechanism this package
// drives.
var ErrUnsupported = errors.New("autostart: not supported on this system")

// Entry is what starts at login.
type Entry struct {
	// ID names the entry; it is also the file or value name, so it is restricted
	// to letters, digits, dot, dash and underscore.
	ID string
	// Name is what the system's startup list shows, where it shows one (the
	// Linux autostart entry's Name). Empty means the ID.
	Name string
	// Executable is the absolute path of the program.
	Executable string
	Args       []string
}

// ValidateID checks the identifier before dispatching to a login-start service.
func ValidateID(id string) error {
	if id == "" || strings.Trim(id, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.-_") != "" {
		return fmt.Errorf("autostart: invalid id %q", id)
	}
	return nil
}

// ValidateArgs checks command arguments before dispatching to a login-start service.
func ValidateArgs(args []string) error {
	for _, arg := range args {
		if strings.ContainsAny(arg, "\x00\r\n") {
			return errors.New("autostart: the command contains a control character")
		}
	}
	return nil
}

func (entry Entry) validate() error {
	if err := ValidateID(entry.ID); err != nil {
		return err
	}
	if entry.Executable == "" {
		return errors.New("autostart: an executable is required")
	}
	return ValidateArgs(append([]string{entry.Executable}, entry.Args...))
}

// Set turns starting at login on or off for the entry.
func Set(entry Entry, enabled bool) error {
	if err := entry.validate(); err != nil {
		return err
	}
	if enabled {
		return enable(entry)
	}
	return disable(entry.ID)
}

// Enabled reports whether an entry with the id is installed.
func Enabled(id string) (bool, error) {
	if err := ValidateID(id); err != nil {
		return false, err
	}
	return enabled(id)
}
