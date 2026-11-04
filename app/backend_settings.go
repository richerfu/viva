package app

import (
	"github.com/vivaui/viva/internal/platform"
	"github.com/vivaui/viva/ui"
)

type BackendSetting = platform.DesktopSetting

const (
	BackendSettingFonts  = platform.SettingFonts
	BackendSettingMotion = platform.SettingMotion
)

// BackendAppearanceChanged reports an observed system change to the normal
// application watchers. Backends call it on their serial UI executor.
func BackendAppearanceChanged() { platform.NotifyAppearanceChanged() }

func watchDesktopSettings(c *ui.Canvas, w *platform.Waker) func() {
	if b := selectedBackend(); b != nil {
		if watch := b.Services().WatchSettings; watch != nil {
			stop := watch(func(changed BackendSetting) {
				if changed&BackendSettingMotion != 0 {
					c.SetReduceMotion(nativeReduceMotion())
				}
				if changed&BackendSettingFonts != 0 {
					c.SetTextEngine(nativeTextEngine())
				}
				if w != nil {
					w.Wake()
				}
			})
			if stop != nil {
				return stop
			}
		}
		return func() {}
	}
	return defaultWatchDesktopSettings(c, w)
}
