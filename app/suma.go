package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/model"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

var (
	tabelaSuma = map[int]int{}
)

func (w *WingCal) SumaStranaPrazno() func(gtx C) D {
	return func(gtx C) D {
		return container.DuoUIcontainer(w.UI.Tema, 4, w.UI.Tema.Colors["LightGrayII"]).Layout(gtx, layout.W, func(gtx C) D {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			gtx.Constraints.Min.Y = gtx.Constraints.Max.Y
			prazno := material.H3(w.UI.Tema.T, w.text("Izaberite radove"))
			prazno.Alignment = text.Middle
			return prazno.Layout(gtx)
		})
	}
}

func (w *WingCal) SumaElementi(el []*model.WingIzabraniElement) func(gtx C) D {
	return func(gtx C) D {
		return sumList.Layout(gtx, len(el), func(gtx C, i int) D {
			element := el[i]
			return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
				func(gtx C) D {
					return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx C) D {
						return lyt.Format(gtx, "hflexb(middle,f(0.6,_),r(_),f(0.1,_),r(_),f(0.1,_),r(_),f(0.1,_),r(_),f(0.2,_))",
							w.cell(text.Start, w.text(element.Element.Naziv)),
							helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
							w.cell(text.Middle, fmt.Sprint(element.Element.Cena)),
							helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
							w.cell(text.Middle, fmt.Sprint(element.Kolicina)),
							helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
							w.cell(text.Middle, fmt.Sprintf("%.2f", element.SumaCena)),
							helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
							w.brisi(element, i))
					})
				},
				helper.DuoUIline(false, 0, 0, 1, w.UI.Tema.Colors["Gray"]))
		})
	}
}
