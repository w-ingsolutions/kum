package calc

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func footer(w *WingCal) func(gtx C) D {
	return func(gtx C) D {
		return container.DuoUIcontainer(w.UI.Tema, 4, w.UI.Tema.Colors["DarkGrayI"]).Layout(gtx, layout.Center, func(gtx C) D {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			return lyt.Format(gtx, "hflexb(middle,r(_),r(_),r(_))",
				w.iconLink(podesavanjeDugme, "podesavanja", "settingsIcon"),
				w.iconLink(pomocDugme, "pomoc", "settingsIcon"),
				w.footerMenu())
		})
	}
}

func (w *WingCal) footerMenu() func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "hflexb(middle,r(_),r(_),r(_))",
			w.stranaDugme(projekatDugme, w.ProjekatRacunica(), w.text("PROJEKAT"), "projekat"),
			helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["DarkGrayI"]),
			w.stranaDugme(materijalDugme, func() {}, w.text("MATERIJAL"), "materijal"),
		)
	}
}

func (w *WingCal) iconLink(b *widget.Clickable, s, i string) func(gtx C) D {
	return func(gtx C) D {
		return layout.Center.Layout(gtx, func(gtx C) D {
			btn := material.IconButton(w.UI.Tema.T, b, w.UI.Tema.Icons[i])
			btn.Color = helper.HexARGB(w.UI.Tema.Colors["Danger"])
			btn.Size = unit.Dp(16)
			btn.Inset = layout.Inset{
				Top:    unit.Dp(4),
				Right:  unit.Dp(4),
				Bottom: unit.Dp(4),
				Left:   unit.Dp(4),
			}
			btn.Background = helper.HexARGB(w.UI.Tema.Colors["White"])
			for b.Clicked() {
				w.Strana = s
			}
			return btn.Layout(gtx)
		})
	}
}
