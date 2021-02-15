package theme

import "fyne.io/fyne/v2"

type (
	InterfaceLibTheme interface {
		SetDarkTheme(a fyne.App)
		SetLightTheme(a fyne.App)
	}
	ObjLibTheme struct {
		DarkTheme
		LightTheme
	}
)

func (p ObjLibTheme) SetDarkTheme(a fyne.App) {
	a.Settings().SetTheme(p.DarkTheme)
}
func (p ObjLibTheme) SetLightTheme(a fyne.App) {
	a.Settings().SetTheme(p.LightTheme)
}
