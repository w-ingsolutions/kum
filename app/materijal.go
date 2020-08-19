package calc

import (
	"fmt"
	"gioui.org/text"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func (w *WingCal) MaterijalElementi() func(gtx C) D {
	return func(gtx C) D {
		return materijalList.Layout(gtx, len(w.Materijal), func(gtx C, i int) D {
			m := w.Materijal[i]
			return lyt.Format(gtx, "vflexb(middle,r(inset(4dp,_)),r(_))",
				func(gtx C) D {
					return lyt.Format(gtx, "hflexb(middle,f(0.1,_),r(_),f(0.3,_),r(_),f(0.1,_),r(_),f(0.1,_),r(_),f(0.2,_))",
						w.cell(text.Start, w.text(fmt.Sprint(m.Id))),
						helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
						w.cell(text.Middle, w.text(m.Naziv)),
						helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
						w.cell(text.Middle, w.text(fmt.Sprint(m.Pakovanje))),
						helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
						w.cell(text.Middle, w.text(m.Jedinica)),
						helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
						w.cell(text.End, w.text(fmt.Sprint(m.Cena))),
					)
				},
				helper.DuoUIline(false, 0, 0, 1, w.UI.Tema.Colors["Gray"]))
		})
	}
}

//func (w *WingCal) MaterijalStrana() func(gtx C) D {
//	return func(gtx C) D {
//		return container.DuoUIcontainer(w.UI.Tema, 1, w.UI.Tema.Colors["DarkGrayI"]).Layout(gtx, layout.N, func(gtx C) D {
//			return container.DuoUIcontainer(w.UI.Tema, 0, w.UI.Tema.Colors["White"]).Layout(gtx, layout.N, func(gtx C) D {
//				return materijalList.Layout(gtx, len(w.Materijal), func(gtx C, i int) D {
//					m := w.Materijal[i]
//					return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
//						func(gtx C) D {
//							return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx C) D {
//								return lyt.Format(gtx, "hflexb(middle,f(0.1,_),r(_),f(0.3,_),r(_),f(0.1,_),r(_),f(0.1,_),r(_),f(0.2,_))",
//									w.cell(text.Start, w.text(fmt.Sprint(m.Id))),
//									helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
//									w.cell(text.Middle, w.text(m.Naziv)),
//									helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
//									w.cell(text.Middle, w.text(fmt.Sprint(m.Pakovanje))),
//									helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
//									w.cell(text.Middle, w.text(m.Jedinica)),
//									helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
//									w.cell(text.End, w.text(fmt.Sprint(m.Cena))))
//							})
//						},
//						//layout.Rigid(func(gtx C) D {
//						//	return layout.Flex{
//						//		Axis: layout.Horizontal,
//						//	}.Layout(gtx) //layout.Flexed(0.4, func(gtx C) D {
//						//	return material.Body1(w.UI.Tema.T, m.NacinRada).Layout(gtx)
//						//}),
//						//layout.Flexed(0.2, func(gtx C) D {
//						//	return material.Caption(w.UI.Tema.T, m.OsobineNamena).Layout(gtx)
//						//}),
//						//layout.Flexed(0.1, func(gtx C) D {
//						//	return material.H6(w.UI.Tema.T, fmt.Sprint(m.Potrosnja)).Layout(gtx)
//						//}),
//						//layout.Flexed(0.1, func(gtx C) D {
//						//	return material.H6(w.UI.Tema.T, m.JedinicaPotrosnje).Layout(gtx)
//						//}),
//						//layout.Flexed(0.2, func(gtx C) D {
//						//	return material.H6(w.UI.Tema.T, m.RokUpotrebe).Layout(gtx)
//						//}),
//						//}),
//						helper.DuoUIline(false, 0, 0, 1, w.UI.Tema.Colors["Gray"]))
//				})
//			})
//		})
//	}
//}
