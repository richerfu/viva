package app

import (
	"github.com/vivaui/viva/internal/platform"
	"github.com/vivaui/viva/ui"
)

// BackendLabelBaseline and the terminal helpers share the layout conventions
// used by the built-in renderers. They contain no native platform calls.
func BackendLabelBaseline(height int, size float64) int {
	return platform.LabelBaseline(height, size)
}

func BackendTerminalPreeditCellWidth(size float32) float32 {
	return platform.TerminalPreeditCellWidth(size)
}

func BackendTerminalPreeditColumn(value string, offset int) int {
	return platform.TerminalPreeditColumn(value, offset)
}

func BackendTerminalPreeditRuns(dst []ui.TextRun, value string, color ui.Color) []ui.TextRun {
	return platform.TerminalPreeditRuns(dst, value, color)
}
