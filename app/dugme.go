package calc

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/helper"
)

var (
	latcyrDugme       = new(widget.Clickable)
	radoviDugme       = new(widget.Clickable)
	materijalDugme    = new(widget.Clickable)
	sumaDugme         = new(widget.Clickable)
	sumaMaterialDugme = new(widget.Clickable)

	podesavanjeDugme    = new(widget.Clickable)
	pomocDugme          = new(widget.Clickable)
	projekatDugme       = new(widget.Clickable)
	dodajDugme          = new(widget.Clickable)
	nazadDugme          = new(widget.Clickable)
	zatvoriElementDugme = new(widget.Clickable)
	stampajDugme        = new(widget.Clickable)
)

func (w *WingCal) stranaDugme(b *widget.Clickable, f func(), t, p string) func(gtx C) D {
	return func(gtx C) D {
		btn := material.Button(w.UI.Tema.T, b, t)
		btn.CornerRadius = unit.Dp(0)
		btn.TextSize = unit.Dp(12)
		btn.Inset = layout.Inset{unit.Dp(8), unit.Dp(8), unit.Dp(10), unit.Dp(8)}
		btn.Background = helper.HexARGB(w.UI.Tema.Colors["Secondary"])
		for b.Clicked() {
			f()
			w.Strana = p
		}
		return btn.Layout(gtx)
	}
}

func (w *WingCal) latDugme() func(gtx C) D {
	return func(gtx C) D {
		latcyr := "Ћирилица"
		if w.Podesavanja.Cyr {
			latcyr = "Latinica"
		}
		btn := material.Button(w.UI.Tema.T, latcyrDugme, latcyr)
		btn.CornerRadius = unit.Dp(0)
		btn.TextSize = unit.Dp(10)
		btn.Background = helper.HexARGB(w.UI.Tema.Colors["Warning"])
		btn.Color = helper.HexARGB(w.UI.Tema.Colors["Dark"])
		for latcyrDugme.Clicked() {
			if w.Podesavanja.Cyr {
				w.Podesavanja.Cyr = false
			} else {
				w.Podesavanja.Cyr = true
			}
		}
		return btn.Layout(gtx)
	}
}

//
//func (w *WingCal) jezikDugme(gtx C, d *widget.Clickable, jezik string) D {
//	btn := material.Button(w.UI.Tema.T, d, jezik)
//	btn.CornerRadius = unit.Dp(0)
//	btn.TextSize = unit.Dp(10)
//	btn.Background = helper.HexARGB(w.UI.Tema.Colors["Warning"])
//	btn.Color = helper.HexARGB(w.UI.Tema.Colors["Dark"])
//	for d.Clicked() {
//		w.Jezik.t = translate.Translation{"sr", jezik}
//	}
//	return btn.Layout(gtx)
//}

//func (w *WingCal) GenerisanjeLinkova(radovi map[int]mod.ElementMenu) {
//	for i, rad := range radovi {
//		rad.Link = new(widget.Clickable)
//		w.IzbornikRadova[i] = rad
//	}
//	return
//}

func (w *WingCal) GenerisanjeDugmicaJezici() {
	for _, jezik := range w.Jezik.dostupni {
		w.Jezik.linkovi[jezik] = new(widget.Clickable)
	}
	return
}
