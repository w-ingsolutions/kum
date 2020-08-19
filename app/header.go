package calc

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func header(w *WingCal) func(gtx C) D {
	return func(gtx C) D {
		return container.DuoUIcontainer(w.UI.Tema, 4, w.UI.Tema.Colors["DarkGrayI"]).Layout(gtx, layout.Center, func(gtx C) D {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			return lyt.Format(gtx, "hflexb(middle,r(_),r(_))",
				func(gtx C) D {
					ic := w.UI.Tema.Icons["logo"]
					ic.Color = helper.HexARGB("ffb8df42")
					return ic.Layout(gtx, unit.Dp(32))
				},
				w.headerMenu(),
			)
		})
	}
}
func (w *WingCal) headerMenu() func(gtx C) D {
	r := func(gtx C) D {
		return lyt.Format(gtx, "hflexb(middle,r(_),r(_))",
			w.stranaDugme(radoviDugme, func() {}, w.text("RADOVI"), "radovi"),
			helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["DarkGrayI"]),
			w.stranaDugme(sumaDugme, func() {}, w.text("SUMA"), "sumaRadovi"),
			helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["DarkGrayI"]),
			w.stranaDugme(sumaMaterialDugme, func() {}, w.text("SUMA MATERIJAL"), "sumaMaterijal"),
		)
	}
	if w.UI.Device != "mob" {
		r = func(gtx C) D {
			return lyt.Format(gtx, "hflexb(middle,r(_))",
				w.stranaDugme(radoviDugme, func() {}, w.text("RADOVI"), "radovi"),
			)
		}
	}
	return r
}
