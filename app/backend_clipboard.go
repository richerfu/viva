package app

import "github.com/vivaui/viva/internal/platform"

// BackendTextClipboardAction and BackendTextClipboardEvent apply the same
// clipboard, secure-input and undo-unit rules as the built-in editors.
func BackendTextClipboardAction(event BackendKeyEvent) byte {
	return platform.TextClipboardAction(event)
}

func BackendTextClipboardEvent(state BackendTextState, secure bool, action byte,
	read func() (string, error), write func(string) error) (BackendTextEvent, bool, error) {
	return platform.TextClipboardEvent(state, secure, action, read, write)
}
