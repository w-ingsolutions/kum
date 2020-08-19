package calc

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
)

func (w *WingCal) PodesavanjaStrana() func(gtx C) D {
	return func(gtx C) D {
		return container.DuoUIcontainer(w.UI.Tema, 1, w.UI.Tema.Colors["DarkGrayI"]).Layout(gtx, layout.N, func(gtx C) D {
			return container.DuoUIcontainer(w.UI.Tema, 0, w.UI.Tema.Colors["White"]).Layout(gtx, layout.N, func(gtx C) D {

				return podesavanjaList.Layout(gtx, len(w.podesavanja()), func(gtx C, i int) D {
					return layout.Flex{}.Layout(gtx,
						layout.Flexed(0.5, w.podesavanja()[i]))
				})

			})
		})
	}
}

func (w *WingCal) podesavanja() []func(gtx C) D {
	return []func(gtx C) D{
		//func(gtx C) D {
		//	return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["White"]).Layout(gtx, layout.Center, func(gtx C) D {
		//		return jezicilList.Layout(gtx, len(w.Jezik.dostupni), func(gtx C, i int) D {
		//			return w.jezikDugme(gtx, w.Jezik.linkovi[w.Jezik.dostupni[i]], w.Jezik.dostupni[i])
		//		})
		//	})
		//},
		func(gtx C) D {
			return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["White"]).Layout(gtx, layout.Center, func(gtx C) D {
				latCyr := material.CheckBox(w.UI.Tema.T, latCyrBool, "Ћирилица/Latinica")
				if latCyrBool.Changed() {
					if latCyrBool.Value {
						w.Podesavanja.Cyr = true
					} else {
						w.Podesavanja.Cyr = false
					}
				}
				return latCyr.Layout(gtx)
			})
		},
	}
}
