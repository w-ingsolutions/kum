package calc

import (
	"fmt"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/skratchdot/open-golang/open"
	"github.com/w-ingsolutions/kum/app/pdf"
)

var (
	materijal, aktivnosti, tehnicki, ponuda, ugovor, standardi, merenja, uslovi int
)

func (w *WingCal) Stampa() func(gtx C) D {
	return func(gtx C) D {
		btn := material.Button(w.UI.Tema.T, stampajDugme, w.text("Štampaj"))
		btn.CornerRadius = unit.Dp(0)
		if len(w.Suma.Elementi) != 0 {
			for stampajDugme.Clicked() {
				fmt.Println("proj::", projekat.Investitor)
				pdf.KreiranjeNalogaPDF(projekat, w.Suma, "nalog.pdf")

				open.Run("nalog.pdf")

			}
		}
		return btn.Layout(gtx)
	}
}
