package calc

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/model"
)

func remove(slice []*model.WingIzabraniElement, s int) []*model.WingIzabraniElement {
	return append(slice[:s], slice[s+1:]...)
}

func (w WingCal) brisi(element *model.WingIzabraniElement, i int) func(gtx C) D {
	return func(gtx C) D {
		btn := material.IconButton(w.UI.Tema.T, element.DugmeBrisanje, w.UI.Tema.Icons["Delete"])
		btn.Inset = layout.Inset{unit.Dp(5), unit.Dp(3), unit.Dp(5), unit.Dp(5)}
		btn.Color = helper.HexARGB(w.UI.Tema.Colors["Danger"])
		btn.Size = unit.Dp(16)
		btn.Background = helper.HexARGB(w.UI.Tema.Colors["White"])
		for element.DugmeBrisanje.Clicked() {
			w.Suma.Elementi = remove(w.Suma.Elementi, i)
			w.SumaRacunica()
		}
		return btn.Layout(gtx)
	}
}
