package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
)

func (w *WingCal) Nazad() func(gtx C) D {
	return func(gtx C) D {
		var btn layout.Dimensions
		if len(w.Putanja) > 1 {
			btnNazad := material.Button(w.UI.Tema.T, nazadDugme, w.text("NAZAD"))
			btnNazad.Background = helper.HexARGB(w.UI.Tema.Colors["Secondary"])
			btnNazad.CornerRadius = unit.Dp(0)
			for nazadDugme.Clicked() {
				if w.Element {
					w.Element = false
				} else {
					//komanda := ""
					if len(w.Putanja) == 2 {
						//komanda = "/" + fmt.Sprint(w.Roditelj)
						//podvrstaradova = fmt.Sprint(w.Roditelj)
						fmt.Println("rodddd222222222222221::", w.Roditelj)

					}
					if len(w.Putanja) == 3 {
						//komanda = "/" + fmt.Sprint(w.Roditelj)
						//podvrstaradova = fmt.Sprint(w.Roditelj)
						fmt.Println("roddddditeL1333311::", w.Roditelj)
						w.UcitajRadove(w.Roditelj)
					}
					if len(w.Putanja) == 4 {
						fmt.Println("roddddditeL11444441::", w.Roditelj)

						//komanda = "/" + Podvrstaradova + "/" + fmt.Sprint(w.Roditelj)
					}
					//w.UcitajRadove("QmUn3oue7CxL3ERQH26P8wQMZBmdhxrapHukg2AJwwBGEK", "radovi")
					//w.LinkoviIzboraVrsteRadova = GenerisanjeLinkova(w.IzbornikRadova)
					//w.GenerisanjeLinkova(w.IzbornikRadova)
					w.Putanja = w.Putanja[:len(w.Putanja)-1]
				}
			}
			btn = btnNazad.Layout(gtx)
		}
		return btn
	}
}
