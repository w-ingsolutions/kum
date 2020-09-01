package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"github.com/w-ingsolutions/kum/app/mod"
)

func (w *WingCal) RadNeophodanMaterijal(l *layout.List) func(gtx C) D {
	return func(gtx C) D {
		//var materijal mod.WingNeophodanMaterijal

		//width := gtx.Constraints.Max.X

		return l.Layout(gtx, len(w.PrikazaniElement.mat), func(gtx C, i int) D {
			m := w.PrikazaniElement.mat[i]

			//gtx.Constraints.Min.X = width
			return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
				//return lyt.Format(gtx, "vflexb(middle,r(_))",
				func(gtx C) D {
					return lyt.Format(gtx, "hflexb(middle,f(0.4,_),r(_),f(0.15,_),r(_),f(0.15,_),r(_),f(0.15,_),r(_),f(0.15,_))",
						w.cell(text.Start, w.text(m.Materijal.Struct["Title"].Content.(string))),
						helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"]),
						w.cell(text.Middle, fmt.Sprintf("%.2f", m.Materijal.Struct["Potrosnja"].Content)),
						helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"]),
						w.cell(text.Middle, fmt.Sprint(m.Koeficijent)),
						helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"]),
						w.cell(text.Middle, fmt.Sprintf("%.2f", m.Kolicina)),
						helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"]),
						w.cell(text.End, fmt.Sprintf("%.2f", m.UkupnaCena)))
					//return D{}
				},
				helper.DuoUIline(false, 1, 1, 1, w.UI.Tema.Colors["Gray"]))
		})
	}
}
func (w *WingCal) UkupanNeophodanMaterijal(unm map[int]mod.WingNeophodanMaterijal) func(gtx C) D {
	return func(gtx C) D {
		width := gtx.Constraints.Max.X
		return ukupanNeophodanMaterijalList.Layout(gtx, len(unm), func(gtx C, i int) D {
			//materijal := unm[i]
			materijal := w.Suma.NeophodanMaterijalPrikaz[i]
			gtx.Constraints.Min.X = width
			return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
				func(gtx C) D {
					return lyt.Format(gtx, "hflexb(middle,f(0.5,_),r(_),f(0.15,_),r(_),f(0.15,_),r(_),f(0.2,_))",
						w.cell(text.Start, w.text(materijal.Materijal.Struct["Title"].Content.(string))),
						helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"]),
						w.cell(text.End, fmt.Sprintf("%.2f", materijal.Materijal.Struct["Cena"].Content.(float64))),
						helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"]),
						w.cell(text.Middle, fmt.Sprintf("%.2f", materijal.Kolicina)),
						helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"]),
						w.cell(text.End, fmt.Sprintf("%.2f", materijal.UkupnaCena)))
				},
				helper.DuoUIline(false, 0, 0, 1, w.UI.Tema.Colors["Gray"]))
		})
	}
}
