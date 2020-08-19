package calc

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func (w *WingCal) Panel(naslov string, stavke, sadrzaj, footer func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return container.DuoUIcontainer(w.UI.Tema, 1, w.UI.Tema.Colors["LightGrayIII"]).Layout(gtx, layout.Center, func(gtx C) D {
			return container.DuoUIcontainer(w.UI.Tema, 0, w.UI.Tema.Colors["White"]).Layout(gtx, layout.Center, func(gtx C) D {
				return lyt.Format(gtx, "vflexb(middle,r(_),r(_),f(1,_),r(_))",
					func(gtx C) D {
						return container.DuoUIcontainer(w.UI.Tema, 4, w.UI.Tema.Colors["Primary"]).Layout(gtx, layout.W, func(gtx C) D {
							gtx.Constraints.Min.X = gtx.Constraints.Max.X
							suma := material.H6(w.UI.Tema.T, naslov)
							suma.Alignment = text.Start
							return suma.Layout(gtx)
						})
					},
					func(gtx C) D {
						return lyt.Format(gtx, "vflexb(middle,r(_))",
							func(gtx C) D {
								return container.DuoUIcontainer(w.UI.Tema, 4, w.UI.Tema.Colors["Gray"]).Layout(gtx, layout.W, stavke)
							})
					},
					sadrzaj,
					func(gtx C) D {
						gtx.Constraints.Min.X = gtx.Constraints.Max.X
						return container.DuoUIcontainer(w.UI.Tema, 0, w.UI.Tema.Colors["LightGrayII"]).Layout(gtx, layout.SE, footer)
					})
			})
		})
	}
}
