// Code generated from the common platform contracts; DO NOT EDIT.
package app

import (
	"github.com/vivaui/viva/internal/input"
	"github.com/vivaui/viva/internal/platform"
	"github.com/vivaui/viva/internal/render"
	"github.com/vivaui/viva/internal/scene"
)

type BackendInputClickCounter = input.ClickCounter
type BackendInputHeldKey = input.HeldKey
type BackendInputKey = input.Key

const BackendInputKeyAltLeft = input.KeyAltLeft
const BackendInputKeyAltRight = input.KeyAltRight
const BackendInputKeyBackspace = input.KeyBackspace
const BackendInputKeyControlLeft = input.KeyControlLeft
const BackendInputKeyControlRight = input.KeyControlRight
const BackendInputKeyDelete = input.KeyDelete
const BackendInputKeyDown = input.KeyDown
const BackendInputKeyEnd = input.KeyEnd
const BackendInputKeyEnter = input.KeyEnter
const BackendInputKeyEscape = input.KeyEscape

type BackendInputKeyEvent = input.KeyEvent

const BackendInputKeyF1 = input.KeyF1
const BackendInputKeyF10 = input.KeyF10
const BackendInputKeyF11 = input.KeyF11
const BackendInputKeyF12 = input.KeyF12
const BackendInputKeyF2 = input.KeyF2
const BackendInputKeyF3 = input.KeyF3
const BackendInputKeyF4 = input.KeyF4
const BackendInputKeyF5 = input.KeyF5
const BackendInputKeyF6 = input.KeyF6
const BackendInputKeyF7 = input.KeyF7
const BackendInputKeyF8 = input.KeyF8
const BackendInputKeyF9 = input.KeyF9
const BackendInputKeyHome = input.KeyHome
const BackendInputKeyInsert = input.KeyInsert
const BackendInputKeyLeft = input.KeyLeft
const BackendInputKeyMetaLeft = input.KeyMetaLeft
const BackendInputKeyMetaRight = input.KeyMetaRight
const BackendInputKeyPageDown = input.KeyPageDown
const BackendInputKeyPageUp = input.KeyPageUp
const BackendInputKeyRight = input.KeyRight
const BackendInputKeyShiftLeft = input.KeyShiftLeft
const BackendInputKeyShiftRight = input.KeyShiftRight
const BackendInputKeySpace = input.KeySpace
const BackendInputKeyTab = input.KeyTab

type BackendInputKeyTracker = input.KeyTracker

const BackendInputKeyUnknown = input.KeyUnknown
const BackendInputKeyUp = input.KeyUp
const BackendInputMaxHeldKeys = input.MaxHeldKeys
const BackendInputModifierAlt = input.ModifierAlt
const BackendInputModifierControl = input.ModifierControl
const BackendInputModifierMeta = input.ModifierMeta
const BackendInputModifierShift = input.ModifierShift

type BackendInputModifiers = input.Modifiers

const BackendInputPressClient = input.PressClient
const BackendInputPressControl = input.PressControl

type BackendInputPressKind = input.PressKind

const BackendInputPressMove = input.PressMove
const BackendInputPressMoveOnDrag = input.PressMoveOnDrag

type BackendInputRect = input.Rect
type BackendInputRegion = input.Region
type BackendInputResult = input.Result
type BackendInputRouter = input.Router
type BackendInputTarget = input.Target

const BackendA11yRoleApplication = platform.A11yRoleApplication
const BackendA11yRoleButton = platform.A11yRoleButton
const BackendA11yRoleCheckbox = platform.A11yRoleCheckbox
const BackendA11yRoleCombobox = platform.A11yRoleCombobox
const BackendA11yRoleGrid = platform.A11yRoleGrid
const BackendA11yRoleGridCell = platform.A11yRoleGridCell
const BackendA11yRoleGroup = platform.A11yRoleGroup
const BackendA11yRoleHeading = platform.A11yRoleHeading
const BackendA11yRoleImage = platform.A11yRoleImage
const BackendA11yRoleLink = platform.A11yRoleLink
const BackendA11yRoleListBox = platform.A11yRoleListBox
const BackendA11yRoleMenu = platform.A11yRoleMenu
const BackendA11yRoleMenuItem = platform.A11yRoleMenuItem
const BackendA11yRoleMenubar = platform.A11yRoleMenubar
const BackendA11yRoleOption = platform.A11yRoleOption
const BackendA11yRoleRadio = platform.A11yRoleRadio
const BackendA11yRoleSeparator = platform.A11yRoleSeparator
const BackendA11yRoleSlider = platform.A11yRoleSlider
const BackendA11yRoleStatus = platform.A11yRoleStatus
const BackendA11yRoleSwitch = platform.A11yRoleSwitch
const BackendA11yRoleTab = platform.A11yRoleTab
const BackendA11yRoleText = platform.A11yRoleText
const BackendA11yRoleTextField = platform.A11yRoleTextField

type BackendAccessibilityAction = platform.AccessibilityAction
type BackendAccessibilityChange = platform.AccessibilityChange

const BackendAccessibilityDecrement = platform.AccessibilityDecrement
const BackendAccessibilityFocus = platform.AccessibilityFocus

type BackendAccessibilityHandler = platform.AccessibilityHandler

const BackendAccessibilityIncrement = platform.AccessibilityIncrement

type BackendAccessibilityNode = platform.AccessibilityNode
type BackendAccessibilityPatterns = platform.AccessibilityPatterns

const BackendAccessibilityPress = platform.AccessibilityPress

type BackendAccessibilityRequest = platform.AccessibilityRequest
type BackendAccessibilitySelection = platform.AccessibilitySelection

const BackendAccessibilitySetNumber = platform.AccessibilitySetNumber
const BackendAccessibilitySetSelection = platform.AccessibilitySetSelection
const BackendAccessibilitySetText = platform.AccessibilitySetText

type BackendAccessibilityTextUnit = platform.AccessibilityTextUnit

const BackendBackdropAcrylic = platform.BackdropAcrylic

type BackendBackdropKind = platform.BackdropKind

const BackendBackdropMica = platform.BackdropMica
const BackendBackdropMicaAlt = platform.BackdropMicaAlt
const BackendBackdropNone = platform.BackdropNone

type BackendBackdropStyle = platform.BackdropStyle

const BackendBaselineBelowCenter = platform.BaselineBelowCenter
const BackendChangedBounds = platform.ChangedBounds
const BackendChangedDescription = platform.ChangedDescription
const BackendChangedEnabled = platform.ChangedEnabled
const BackendChangedExpand = platform.ChangedExpand
const BackendChangedFocus = platform.ChangedFocus
const BackendChangedHint = platform.ChangedHint
const BackendChangedLevel = platform.ChangedLevel
const BackendChangedLive = platform.ChangedLive
const BackendChangedName = platform.ChangedName
const BackendChangedPassword = platform.ChangedPassword
const BackendChangedPosition = platform.ChangedPosition
const BackendChangedRange = platform.ChangedRange
const BackendChangedReadOnly = platform.ChangedReadOnly
const BackendChangedRole = platform.ChangedRole
const BackendChangedSelected = platform.ChangedSelected
const BackendChangedSelection = platform.ChangedSelection
const BackendChangedToggle = platform.ChangedToggle
const BackendChangedValue = platform.ChangedValue

type BackendCloseRequestHandler = platform.CloseRequestHandler
type BackendColor = platform.Color

const BackendDefaultTitlebarHeight = platform.DefaultTitlebarHeight

type BackendDesktopSetting = platform.DesktopSetting
type BackendDeviceRecovery = platform.DeviceRecovery
type BackendDeviceRecoveryPlan = platform.DeviceRecoveryPlan

const BackendDeviceRecoveryWindow = platform.DeviceRecoveryWindow
const BackendEdgeBorder = platform.EdgeBorder
const BackendEdgeBottom = platform.EdgeBottom
const BackendEdgeBottomLeft = platform.EdgeBottomLeft
const BackendEdgeBottomRight = platform.EdgeBottomRight
const BackendEdgeCorner = platform.EdgeCorner
const BackendEdgeIndexBottom = platform.EdgeIndexBottom
const BackendEdgeIndexLeft = platform.EdgeIndexLeft
const BackendEdgeIndexRight = platform.EdgeIndexRight
const BackendEdgeIndexTop = platform.EdgeIndexTop
const BackendEdgeLeft = platform.EdgeLeft
const BackendEdgeNone = platform.EdgeNone
const BackendEdgeRight = platform.EdgeRight
const BackendEdgeTop = platform.EdgeTop
const BackendEdgeTopLeft = platform.EdgeTopLeft
const BackendEdgeTopRight = platform.EdgeTopRight

type BackendFault = platform.Fault

const BackendFileDropDrop = platform.FileDropDrop

type BackendFileDropEvent = platform.FileDropEvent
type BackendFileDropHandler = platform.FileDropHandler

const BackendFileDropLeave = platform.FileDropLeave
const BackendFileDropOver = platform.FileDropOver

type BackendFileDropPhase = platform.FileDropPhase
type BackendFocusLostHandler = platform.FocusLostHandler
type BackendFrameChangeHandler = platform.FrameChangeHandler
type BackendKeyHandler = platform.KeyHandler

const BackendMacInset = platform.MacInset
const BackendMaxClipboardImageBytes = platform.MaxClipboardImageBytes
const BackendMaxClipboardRawBytes = platform.MaxClipboardRawBytes

type BackendMaximizeHandler = platform.MaximizeHandler
type BackendMenuEntry = platform.MenuEntry
type BackendMenuItem = platform.MenuItem
type BackendMinimizeHandler = platform.MinimizeHandler
type BackendNativeView = platform.NativeView

const BackendPanelGap = platform.PanelGap
const BackendPanelMargin = platform.PanelMargin
const BackendPatternExpandCollapse = platform.PatternExpandCollapse
const BackendPatternInvoke = platform.PatternInvoke
const BackendPatternRangeValue = platform.PatternRangeValue
const BackendPatternSelection = platform.PatternSelection
const BackendPatternSelectionItem = platform.PatternSelectionItem
const BackendPatternText = platform.PatternText
const BackendPatternToggle = platform.PatternToggle
const BackendPatternValue = platform.PatternValue

type BackendPointerEventKind = platform.PointerEventKind
type BackendPointerHandler = platform.PointerHandler

const BackendPressClient = platform.PressClient
const BackendPressControl = platform.PressControl

type BackendPressKind = platform.PressKind

const BackendPressMove = platform.PressMove
const BackendPressMoveOnDrag = platform.PressMoveOnDrag
const BackendRecoverGiveUp = platform.RecoverGiveUp
const BackendRecoverHardware = platform.RecoverHardware
const BackendRecoverSoftware = platform.RecoverSoftware

type BackendRedrawHandler = platform.RedrawHandler
type BackendResizeHandler = platform.ResizeHandler
type BackendScreenRect = platform.ScreenRect
type BackendScrollHandler = platform.ScrollHandler
type BackendShortcutHandler = platform.ShortcutHandler
type BackendStatusClick = platform.StatusClick
type BackendStatusIcon = platform.StatusIcon

const BackendTextDeleteToLineEnd = platform.TextDeleteToLineEnd
const BackendTextDeleteToLineStart = platform.TextDeleteToLineStart
const BackendTextDeleteWordBackward = platform.TextDeleteWordBackward
const BackendTextDeleteWordForward = platform.TextDeleteWordForward

type BackendTextHandler = platform.TextHandler
type BackendTextInputUpdate = platform.TextInputUpdate

const BackendTextInsertLineBreak = platform.TextInsertLineBreak
const BackendTextMoveDocumentEnd = platform.TextMoveDocumentEnd
const BackendTextMoveDocumentEndAndModifySelection = platform.TextMoveDocumentEndAndModifySelection
const BackendTextMoveDocumentStart = platform.TextMoveDocumentStart
const BackendTextMoveDocumentStartAndModifySelection = platform.TextMoveDocumentStartAndModifySelection
const BackendTextMoveLeftAndModifySelection = platform.TextMoveLeftAndModifySelection
const BackendTextMoveLineEnd = platform.TextMoveLineEnd
const BackendTextMoveLineEndAndModifySelection = platform.TextMoveLineEndAndModifySelection
const BackendTextMoveLineStart = platform.TextMoveLineStart
const BackendTextMoveLineStartAndModifySelection = platform.TextMoveLineStartAndModifySelection
const BackendTextMoveRightAndModifySelection = platform.TextMoveRightAndModifySelection
const BackendTextMoveWordLeft = platform.TextMoveWordLeft
const BackendTextMoveWordLeftAndModifySelection = platform.TextMoveWordLeftAndModifySelection
const BackendTextMoveWordRight = platform.TextMoveWordRight
const BackendTextMoveWordRightAndModifySelection = platform.TextMoveWordRightAndModifySelection
const BackendTextPaste = platform.TextPaste
const BackendTextPasteNoText = platform.TextPasteNoText
const BackendTextRedo = platform.TextRedo
const BackendTextUndo = platform.TextUndo
const BackendTextUnitCharacter = platform.TextUnitCharacter
const BackendTextUnitDocument = platform.TextUnitDocument
const BackendTextUnitLine = platform.TextUnitLine
const BackendTextUnitParagraph = platform.TextUnitParagraph
const BackendTextUnitWord = platform.TextUnitWord

type BackendWaker = platform.Waker

const BackendWebBridgeScript = platform.WebBridgeScript

type BackendWebHost = platform.WebHost
type BackendWebLoad = platform.WebLoad

const BackendWebLoadCommitted = platform.WebLoadCommitted
const BackendWebLoadFinished = platform.WebLoadFinished

type BackendWebLoadPhase = platform.WebLoadPhase

const BackendWebLoadStarted = platform.WebLoadStarted

type BackendWebScheme = platform.WebScheme
type BackendWebView = platform.WebView
type BackendWebViewConfig = platform.WebViewConfig

const BackendWheelNotch = platform.WheelNotch
const BackendWheelPageScroll = platform.WheelPageScroll

type BackendWindowActions = platform.WindowActions

const BackendWindowControlWidth = platform.WindowControlWidth
const BackendWindowCustomTitlebar = platform.WindowCustomTitlebar

type BackendWindowEdge = platform.WindowEdge

const BackendWindowFrameless = platform.WindowFrameless

type BackendWindowIcon = platform.WindowIcon
type BackendWindowMaterials = platform.WindowMaterials
type BackendWindowMode = platform.WindowMode

const BackendWindowPanel = platform.WindowPanel
const BackendWindowTitled = platform.WindowTitled

type BackendRenderBatch = render.Batch
type BackendRenderRectangle = render.Rectangle
type BackendRenderScissor = render.Scissor
type BackendRenderStore = render.Store
type BackendRenderUpdate = render.Update
type BackendSceneChange = scene.Change
type BackendSceneColor = scene.Color
type BackendSceneID = scene.ID
type BackendSceneKind = scene.Kind

const BackendSceneKindGroup = scene.KindGroup
const BackendSceneKindRectangle = scene.KindRectangle
const BackendSceneKindShadow = scene.KindShadow
const BackendSceneKindTriangle = scene.KindTriangle

type BackendSceneOperation = scene.Operation

const BackendSceneOperationDelete = scene.OperationDelete
const BackendSceneOperationUpsert = scene.OperationUpsert

type BackendScenePoint = scene.Point
type BackendSceneRect = scene.Rect
type BackendSceneScene = scene.Scene
type BackendSceneShade = scene.Shade
type BackendSceneShadow = scene.Shadow
