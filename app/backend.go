package app

import (
	"errors"
	"sync"

	"github.com/vivaui/viva/internal/platform"
	"github.com/vivaui/viva/internal/render"
	"github.com/vivaui/viva/internal/scene"
	"github.com/vivaui/viva/ui"
)

// Backend is a native platform implementation. A host installs it before
// running the application. The regular painter, controls and input router
// produce Config and Frame; a backend owns native windows and their renderer.
// It must consume borrowed frames before invoking another framework callback.
type Backend interface {
	Run(BackendConfig) error
	RunApplication(BackendApplicationStart) error
	TextEngine() ui.TextEngine
	Services() BackendServices
}

// HostedBackend is a backend whose native host owns the event loop. Run and
// RunApplication register windows and return; each window must invoke OnClosed
// exactly once when its native session ends. Desktop backends block normally.
type HostedBackend interface {
	Backend
	ExternallyDriven() bool
}

// BackendServices are optional native services. Missing operations report
// ErrUnsupported rather than using a simulated clipboard or desktop backend.
// Functions run on the backend's serial event-loop thread unless documented
// otherwise by the corresponding public app API.
type BackendServices struct {
	RequestClipboardImage  func(func([]byte, error))
	ReduceMotion           func() bool
	PreferredLanguages     func() []string
	ClipboardText          func() (string, error)
	SetClipboardText       func(value string) error
	HasClipboardImage      func() bool
	ClipboardImagePNG      func() ([]byte, error)
	OnClipboardChange      func(notify func()) error
	SetCursor              func(cursor Cursor) error
	SetCursorImage         func(image *ui.CursorImage) error
	OpenFiles              func(options OpenFileOptions) ([]string, error)
	OpenDirectory          func(options OpenDirectoryOptions) (string, error)
	SaveFile               func(options SaveFileOptions) (string, error)
	OpenURL                func(target string) error
	OpenWith               func(executable, path string) error
	PlaySystemSound        func(sound SystemSound) error
	SendNotification       func(notification Notification) error
	StoreSecret            func(service, account string, secret []byte) error
	LoadSecret             func(service, account string) ([]byte, error)
	DeleteSecret           func(service, account string) error
	EnableCrashDumps       func(directory string) (string, error)
	ShowFatalError         func(title, message string)
	MonospaceFamilies      func() []string
	FontFamilies           func() []string
	Appearance             func() Appearance
	WatchAppearance        func()
	SetAppearance          func(appearance Appearance)
	SetPrimarySelection    func(value string) error
	PrimarySelection       func() (string, error)
	SchemeHandlerAvailable func(scheme string) (bool, error)
	FileIcon               func(path string, pixels int) ([]byte, error)
	HighContrast           func() bool
	WatchSettings          func(func(BackendSetting)) func()
}

var installedBackend struct {
	sync.RWMutex
	value Backend
}

// UseBackend selects a native backend before RunCanvas or RunApplication.
// Passing nil restores the default platform backend. Installations must not
// overlap a live native application; they are process-level startup wiring.
func UseBackend(backend Backend) error {
	if backend != nil && backend.TextEngine() == nil {
		return errors.New("app: a native backend must provide its text engine")
	}
	installedBackend.Lock()
	installedBackend.value = backend
	installedBackend.Unlock()
	return nil
}

func selectedBackend() Backend {
	installedBackend.RLock()
	b := installedBackend.value
	installedBackend.RUnlock()
	return b
}

func externallyDriven() bool {
	b, ok := selectedBackend().(HostedBackend)
	return ok && b.ExternallyDriven()
}

// A host without a cursor service still renders and routes pointer events.
// The public SetCursor API keeps reporting the missing capability to callers.
func optionalNativeCursor(cursor Cursor) error {
	err := nativeSetCursor(cursor)
	if errors.Is(err, ErrUnsupported) {
		return nil
	}
	return err
}

// The backend contracts are exported aliases of the contracts used by Viva's
// built-in platforms, so an extension receives the same production frames and
// callbacks without a second scene representation or an internal import.
type BackendConfig = platform.Config
type BackendFrame = platform.Frame
type BackendRectangle = render.Rectangle
type BackendTextLabel = platform.TextLabel
type BackendTextInput = platform.TextInput
type BackendTextState = platform.TextState
type BackendTextEvent = platform.TextEvent
type BackendTextCommand = platform.TextCommand
type BackendTextEventKind = platform.TextEventKind
type BackendPointerEvent = platform.PointerEvent
type BackendKeyEvent = platform.KeyEvent
type BackendScrollEvent = platform.ScrollEvent
type BackendSize = platform.Size
type BackendApplication = platform.Application
type BackendApplicationStart = platform.ApplicationStart
type BackendWindow = platform.Window
type BackendStatusItem = platform.StatusItem
type BackendStatusItemConfig = platform.StatusItemConfig

const (
	BackendPointerDown        = platform.PointerDown
	BackendPointerMove        = platform.PointerMove
	BackendPointerUp          = platform.PointerUp
	BackendPointerExit        = platform.PointerExit
	BackendTextSetSelection   = platform.TextSetSelection
	BackendTextSelectAll      = platform.TextSelectAll
	BackendTextInsert         = platform.TextInsert
	BackendTextSetMarked      = platform.TextSetMarked
	BackendTextUnmark         = platform.TextUnmark
	BackendTextCommandEvent   = platform.TextCommandEvent
	BackendTextDeleteBackward = platform.TextDeleteBackward
	BackendTextDeleteForward  = platform.TextDeleteForward
	BackendTextMoveLeft       = platform.TextMoveLeft
	BackendTextMoveRight      = platform.TextMoveRight
	BackendTextInsertNewline  = platform.TextInsertNewline
	BackendTextInsertTab      = platform.TextInsertTab
	BackendRectangleKind      = scene.KindRectangle
	BackendTriangleKind       = scene.KindTriangle
	BackendShadowKind         = scene.KindShadow
)

// BackendTextEditingCommand and BackendTextHistoryCommand share Viva's native
// editor shortcut rules with external platform backends.
func BackendTextEditingCommand(e BackendKeyEvent) (BackendTextCommand, bool) {
	return platform.TextEditingCommand(e)
}
func BackendTextHistoryCommand(e BackendKeyEvent) (BackendTextCommand, bool) {
	return platform.TextHistoryCommand(e)
}
