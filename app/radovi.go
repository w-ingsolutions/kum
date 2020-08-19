package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioapp/gel/helper"
	"github.com/gioapp/gel/icontextbtn"
	"github.com/w-ingsolutions/c/model"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func (w *WingCal) IzborVrsteRadova() func(gtx C) D {
	izbornik := w.IzborPodVrsteRadova()
	switch w.UI.Device {
	case "mob":
		izbornik = w.IzborPodVrsteRadova()
	case "tab":
		izbornik = w.IzborVrsteRadovaDveKolone()
	case "lap":
		izbornik = w.IzborVrsteRadovaTriKolone()
	case "des":
		izbornik = w.IzborVrsteRadovaCetiriKolone()
	}
	return izbornik
}
func (w *WingCal) IzborVrsteRadovaDveKolone() func(gtx C) D {
	return w.vBtnLinks(
		w.IzborVrsteRadovaDveKoloneRed(0, 1),
		w.IzborVrsteRadovaDveKoloneRed(2, 3),
		w.IzborVrsteRadovaDveKoloneRed(4, 5),
		w.IzborVrsteRadovaDveKoloneRed(6, 7),
		w.IzborVrsteRadovaDveKoloneRed(8, 9),
		w.IzborVrsteRadovaDveKoloneRed(10, 11),
		w.IzborVrsteRadovaDveKoloneRed(12, 13),
		w.IzborVrsteRadovaDveKoloneRed(14, 15),
		w.IzborVrsteRadovaDveKoloneRed(16, 17),
		w.IzborVrsteRadovaDveKoloneRed(18, 19),
		w.IzborVrsteRadovaDveKoloneRed(20, 21),
		w.IzborVrsteRadovaDveKoloneRed(23, 23),
		w.IzborVrsteRadovaDveKoloneRed(24, 25),
		w.IzborVrsteRadovaDveKoloneRed(26, 27),
		w.IzborVrsteRadovaDveKoloneRed(28, 29),
	)
}
func (w *WingCal) IzborVrsteRadovaDveKoloneRed(a, b int) func(gtx C) D {
	return w.hBtnLinks(w.btnLink(w.IzbornikRadova[a]), w.btnLink(w.IzbornikRadova[b]))

}
func (w *WingCal) hBtnLinks(a, b func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "vflexb(middle,f(0.333,_),f(0.333,_))", a, b)
	}
}

func (w *WingCal) vBtnLinks(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "vflexb(middle,f(0.06,_),f(0.06,_),f(0.06,_),f(0.06,_),f(0.06,_),f(0.06,_),f(0.06,_),f(0.06,_),f(0.06,_),f(0.06,_),f(0.06,_),f(0.06,_),f(0.06,_),f(0.06,_))", a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
	}
}

func (w *WingCal) IzborVrsteRadovaTriKolone() func(gtx C) D {
	return w.tvBtnLinks(
		w.IzborVrsteRadovaTriKoloneRed(0, 1, 2),
		w.IzborVrsteRadovaTriKoloneRed(3, 4, 5),
		w.IzborVrsteRadovaTriKoloneRed(6, 7, 8),
		w.IzborVrsteRadovaTriKoloneRed(9, 10, 11),
		w.IzborVrsteRadovaTriKoloneRed(12, 13, 14),
		w.IzborVrsteRadovaTriKoloneRed(15, 16, 17),
		w.IzborVrsteRadovaTriKoloneRed(18, 19, 20),
		w.IzborVrsteRadovaTriKoloneRed(21, 22, 23),
		w.IzborVrsteRadovaTriKoloneRed(24, 5, 26),
		w.IzborVrsteRadovaTriKoloneRed(27, 28, 29),
	)
}
func (w *WingCal) IzborVrsteRadovaTriKoloneRed(a, b, c int) func(gtx C) D {
	return w.thBtnLinks(w.btnLink(w.IzbornikRadova[a]), w.btnLink(w.IzbornikRadova[b]), w.btnLink(w.IzbornikRadova[c]))

}
func (w *WingCal) thBtnLinks(a, b, c func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "vflexb(middle,f(0.333,_),f(0.333,_),f(0.333,_))", a, b, c)
	}
}

func (w *WingCal) tvBtnLinks(a, b, c, d, e, f, g, h, i, j func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "vflexb(middle,f(0.1,_),f(0.1,_),f(0.1,_),f(0.1,_),f(0.1,_),f(0.1,_),f(0.1,_),f(0.1,_),f(0.1,_),f(0.1,_))", a, b, c, d, e, f, g, h, i, j)
	}
}

func (w *WingCal) IzborVrsteRadovaCetiriKolone() func(gtx C) D {
	return w.fvBtnLinks(
		w.IzborVrsteRadovaCetiriKoloneRed(0, 1, 2, 3),
		w.IzborVrsteRadovaCetiriKoloneRed(4, 5, 6, 7),
		w.IzborVrsteRadovaCetiriKoloneRed(8, 9, 10, 11),
		w.IzborVrsteRadovaCetiriKoloneRed(12, 13, 14, 15),
		w.IzborVrsteRadovaCetiriKoloneRed(16, 17, 18, 19),
		w.IzborVrsteRadovaCetiriKoloneRed(20, 21, 22, 23),
		w.IzborVrsteRadovaCetiriKoloneRed(24, 25, 26, 27),
		w.IzborVrsteRadovaCetiriKoloneRed(28, 29, 30, 29),
	)
}

func (w *WingCal) IzborVrsteRadovaCetiriKoloneRed(a, b, c, d int) func(gtx C) D {
	return w.fhBtnLinks(w.btnLink(w.IzbornikRadova[a]), w.btnLink(w.IzbornikRadova[b]), w.btnLink(w.IzbornikRadova[c]), w.btnLink(w.IzbornikRadova[d]))
}

func (w *WingCal) fhBtnLinks(a, b, c, d func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "hflexb(f(0.25,_),f(0.25,_),f(0.25,_),f(0.25,_))", a, b, c, d)
	}
}

func (w *WingCal) fvBtnLinks(a, b, c, d, e, f, g, h func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return lyt.Format(gtx, "vflexb(f(0.125,_),f(0.125,_),f(0.125,_),f(0.125,_),f(0.125,_),f(0.125,_),f(0.125,_),f(0.125,_))", a, b, c, d, e, f, g, h)
	}
}

func (w *WingCal) btnLink(r model.ElementMenu) func(gtx C) D {
	return func(gtx C) D {
		btn := icontextbtn.IconTextBtn(w.UI.Tema.T, r.Link, w.UI.Tema.Icons[r.Slug], unit.Dp(48), w.UI.Tema.Colors["Light"], fmt.Sprint(r.Id)+". "+w.text(r.Title))
		switch w.UI.Device {
		case "mob":
			btn.Axis = layout.Horizontal
		case "tab":
			btn.Axis = layout.Horizontal
		case "lap":
			btn.IconSize = unit.Dp(64)
			btn.TextSize = unit.Dp(12)
			btn.Axis = layout.Vertical

		case "des":
			btn.IconSize = unit.Dp(72)
			btn.TextSize = unit.Dp(12)
			btn.Axis = layout.Vertical
		}
		btn.CornerRadius = unit.Dp(0)
		btn.Background = helper.HexARGB(w.UI.Tema.Colors["Gray"])
		//btn.Inset = layout.Inset{
		//	Top:    unit.Dp(2),
		//	Right:  unit.Dp(2),
		//	Bottom: unit.Dp(2),
		//	Left:   unit.Dp(2),
		//}
		if r.Materijal {
			//btn.Background = helper.HexARGB(w.UI.Tema.Colors["DarkGray"])
		}
		w.LinkoviIzboraVrsteRadovaKlik(r)
		return btn.Layout(gtx)
	}
}
