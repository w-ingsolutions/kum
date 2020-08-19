package calc

import (
	"gioui.org/text"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func (w *WingCal) SumaStavkeMaterijala() func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return lyt.Format(gtx, "hflexb(middle,f(0.5,_),r(_),f(0.15,_),r(_),f(0.15,_),r(_),f(0.2,_))",
			w.cell(text.Start, w.text("Naziv")),
			helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
			w.cell(text.Middle, w.text("Pojedinačna cena")),
			helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
			w.cell(text.Middle, w.text("Količina")),
			helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
			w.cell(text.End, w.text("Ukupna cena")))
	}
}

func (w *WingCal) PrikazaniElementStavkeMaterijala() func(gtx C) D {
	return func(gtx C) D {
		//gtx.Constraints.Min.X = width
		return lyt.Format(gtx, "hflexb(middle,f(0.4,_),r(_),f(0.15,_),r(_),f(0.15,_),r(_),f(0.15,_),r(_),f(0.15,_))",
			w.cell(text.Start, w.text("Naziv")),
			helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
			w.cell(text.Middle, w.text("Potrosnja")),
			helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
			w.cell(text.Middle, w.text("Koeficijent")),
			helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
			w.cell(text.Middle, w.text("Merodavna potrosnja")),
			helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
			w.cell(text.End, w.text("Cena materijala")))
	}
}

func (w *WingCal) SumaRadoviStavke() func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return lyt.Format(gtx, "hflexb(middle,f(0.6,_),r(_),f(0.15,_),r(_),f(0.1,_),r(_),f(0.1,_),r(_),f(0.05,_))",
			w.cell(text.Start, w.text("Naziv")),
			helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
			w.cell(text.Middle, w.text("Pojedinačna cena")),
			helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
			w.cell(text.Middle, w.text("Količina")),
			helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
			w.cell(text.Middle, w.text("Cena")),
			helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
			w.cell(text.Middle, w.text("briši")))
	}
}
func (w *WingCal) MaterijalStavke() func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints.Min.X = gtx.Constraints.Max.X
		return lyt.Format(gtx, "hflexb(middle,f(0.1,_),r(_),f(0.3,_),r(_),f(0.1,_),r(_),f(0.1,_),r(_),f(0.2,_))",
			w.cell(text.Start, w.text("Id")),
			helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
			w.cell(text.Middle, w.text("Naziv")),
			helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
			w.cell(text.Middle, w.text("Pakovanje")),
			helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
			w.cell(text.Middle, w.text("Jedinica mere")),
			helper.DuoUIline(true, 0, 8, 2, w.UI.Tema.Colors["Gray"]),
			w.cell(text.End, w.text("Cena")))
	}
}
