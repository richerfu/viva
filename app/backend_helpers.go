package app

import (
	"github.com/vivaui/viva/gfx"
	"github.com/vivaui/viva/internal/platform"
	"github.com/vivaui/viva/internal/textmetrics"
)

// BackendLabelBaseline and the terminal helpers share the layout conventions
// used by the built-in renderers. They contain no native platform calls.
func BackendLabelBaseline(height int, size float64) int {
	return platform.LabelBaseline(height, size)
}

func BackendTerminalPreeditCellWidth(size float32) float32 {
	return size * textmetrics.MonoAdvance
}

func BackendTerminalPreeditColumn(value string, offset int) int {
	return platform.TerminalPreeditColumn(value, offset)
}

func BackendTerminalPreeditRuns(dst []gfx.TextRun, value string, color gfx.Color) []gfx.TextRun {
	return platform.TerminalPreeditRuns(dst, value, color)
}
