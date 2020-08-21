package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/text"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"github.com/w-ingsolutions/cms/pkg/phi"
	"github.com/w-ingsolutions/kum/mod"
)

func (w *WingCal) RadNeophodanMaterijal(l *layout.List) func(gtx C) D {
	return func(gtx C) D {
		//var materijal mod.WingNeophodanMaterijal
		nm := w.PrikazaniElement.Struct["NeophodanMaterijal"].Content.([]phi.Φ)
		//width := gtx.Constraints.Max.X
		return l.Layout(gtx, len(nm), func(gtx C, i int) D {
			//materijal := nm[i]
			//id := materijal.ID - 1
			//materijal.Struct["Koeficijent"].Content = materijal.Struct["Koeficijent"].Content.(string)
			//materijal.Struct["Materijal"].Content = *w.Materijal[id]
			//if materijal.Struct["Koeficijent"].Content.(int) > 0 {
			//	materijal.Struct["Kolicina"].Content.(int) = materijal.Materijal.Potrosnja * float64(w.UI.Counters.Kolicina.Value) * materijal.Struct["Koeficijent"].Content.(int)
			//}
			//materijal.UkupnaCena = materijal.Materijal.Cena * float64(materijal.Kolicina)
			//materijal.UkupnoPakovanja = int(materijal.Kolicina / float64(materijal.Materijal.Pakovanje))

			//gtx.Constraints.Min.X = width
			return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
				//return lyt.Format(gtx, "vflexb(middle,r(_))",
				func(gtx C) D {
					//return lyt.Format(gtx, "hflexb(middle,f(0.4,_),r(_),f(0.15,_),r(_),f(0.15,_),r(_),f(0.15,_),r(_),f(0.15,_))",
					//	w.cell(text.Start, w.text(materijal.Materijal.Naziv)),
					//	helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"]),
					//	w.cell(text.Middle, fmt.Sprintf("%.2f", materijal.Materijal.Potrosnja)),
					//	helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"]),
					//	w.cell(text.Middle, fmt.Sprint(materijal.Koeficijent)),
					//	helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"]),
					//	w.cell(text.Middle, fmt.Sprintf("%.2f", materijal.Kolicina)),
					//	helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"]),
					//	w.cell(text.End, fmt.Sprintf("%.2f", materijal.UkupnaCena)),
					//)
					return D{}
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
						w.cell(text.Start, w.text(materijal.Materijal.Struct["Naziv"].Content.(string))),
						helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"]),
						w.cell(text.Middle, w.text(materijal.Materijal.Struct["Cena"].Content.(string))),
						helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"]),
						w.cell(text.Middle, fmt.Sprintf("%.2f", materijal.Kolicina)),
						helper.DuoUIline(true, 0, 2, 2, w.UI.Tema.Colors["Gray"]),
						w.cell(text.End, fmt.Sprintf("%.2f", materijal.UkupnaCena)))
				},
				helper.DuoUIline(false, 0, 0, 1, w.UI.Tema.Colors["Gray"]))
		})
	}
}
