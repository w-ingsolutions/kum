package calc

import (
	"encoding/json"
	"fmt"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/counter"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/c/pkg/latcyr"
	"github.com/w-ingsolutions/c/pkg/lyt"
	"github.com/w-ingsolutions/kum/app/mod"
	"log"
)

func (w *WingCal) PrikazaniElementDugmeDodaj(sumaCena float64) func(gtx C) D {
	return func(gtx C) D {

		//fmt.Println("maxOPED:", gtx.Constraints.Max.X)
		//fmt.Println("minOPED:", gtx.Constraints.Min.X)
		btn := material.Button(w.UI.Tema.T, dodajDugme, latcyr.C("DODAJ", w.Podesavanja.Cyr))
		//btn.FullWidth = true
		//btn.FullHeight = true
		btn.CornerRadius = unit.Dp(0)
		btn.TextSize = unit.Dp(14)
		btn.Inset = layout.Inset{unit.Dp(8), unit.Dp(8), unit.Dp(10), unit.Dp(8)}
		btn.Background = helper.HexARGB("ffb8df42")
		btn.Color = helper.HexARGB(w.UI.Tema.Colors["Dark"])
		var varijacijaRada int

		for dodajDugme.Clicked() {
			//fmt.Println("kolicina.Value", w.UI.Counters.Kolicina.Value)
			if w.UI.Counters.Kolicina.Value > 0 {
				//fmt.Println("w.Suma.Elementiw.Suma.Elementiw.Suma.Elementi", w.Suma.Elementi)

				for _, s := range w.Suma.Elementi {
					if s.Element.ID == w.PrikazaniElement.el.ID {
						varijacijaRada = varijacijaRada + 1
					}
				}
				//fmt.Println("w.PrikazaniElement.IDi", w.PrikazaniElement.ID)

				suma := mod.WingIzabraniElement{
					Sifra:         fmt.Sprint(w.Podvrsta) + "." + fmt.Sprint(w.Roditelj) + "." + fmt.Sprint(w.PrikazaniElement.el.ID) + "." + fmt.Sprint(varijacijaRada+1),
					Kolicina:      w.UI.Counters.Kolicina.Value,
					SumaCena:      sumaCena,
					Element:       w.PrikazaniElement.el,
					DugmeBrisanje: new(widget.Clickable),
				}
				fmt.Println("sumasumasumasumai", suma)

				w.Suma.Elementi = append(w.Suma.Elementi, suma)

				fmt.Println("w.Suma.Elementi", w.Suma.Elementi)

				w.SumaRacunica()
				//w.Strana = "sumaRadovi"
				//fmt.Println("w.Strana", w.Strana)

			}
		}
		return btn.Layout(gtx)
	}
}
func (w *WingCal) SumaElementiPrikaz() {
	//w.Suma.ElementiPrikaz = nil
	//for _, e := range w.Suma.Elementi {
	//	w.Suma.ElementiPrikaz = append(w.Suma.Elementi, e)
	//
	//}
}
func (w *WingCal) PrikazaniElementIzgled() func(gtx C) D {
	return func(gtx C) D {
		neophodanNaslov := material.H6(w.UI.Tema.T, w.text("Neophodan materijal za izvrsenje radova"))
		neophodanNaslov.Color = helper.HexARGB(w.UI.Tema.Colors["Primary"])
		widgets := []layout.Widget{
			material.H5(w.UI.Tema.T, fmt.Sprint(w.Podvrsta)+"."+fmt.Sprint(w.Roditelj.id)+"."+fmt.Sprint(w.PrikazaniElement.el.ID)+" "+w.text(w.PrikazaniElement.el.Struct["Title"].Content.(string))).Layout,
			material.Body1(w.UI.Tema.T, w.text(w.PrikazaniElement.el.Struct["Opis"].Content.(string))).Layout,
			material.Caption(w.UI.Tema.T, w.text(w.PrikazaniElement.el.Struct["Obracun"].Content.(string))).Layout,
			neophodanNaslov.Layout,
			helper.DuoUIline(false, 4, 0, 4, w.UI.Tema.Colors["Secondary"]),
			w.PrikazaniElementStavkeMaterijala(),
			helper.DuoUIline(false, 4, 0, 2, w.UI.Tema.Colors["Primary"]),
			w.RadNeophodanMaterijal(neophodanMaterijalList),
		}
		return elementOpis.Layout(gtx, len(widgets), func(gtx C, i int) D {
			return layout.UniformInset(unit.Dp(8)).Layout(gtx, widgets[i])
		})
	}
}
func (w *WingCal) PrikazaniElementMaterijal() {
	for _, m := range recept(w.PrikazaniElement.el.Struct["NeophodanMaterijal"].Content.(string)) {
		materijal := mod.WingNeophodanMaterijal{
			Id:          m["ID"],
			Koeficijent: float64(m["Koeficijent"]),
			Materijal:   w.Materijal[m["ID"]],
			//materijal.Struct["Materijal"].Content = *w.Materijal[id]
		}
		if materijal.Koeficijent > 0 {
			materijal.Kolicina = materijal.Materijal.Struct["Potrosnja"].Content.(float64) * float64(w.UI.Counters.Kolicina.Value) * materijal.Koeficijent
		}

		materijal.UkupnaCena = w.PrikazaniElement.el.Struct["Cena"].Content.(float64) * float64(materijal.Kolicina)

		materijal.UkupnoPakovanja = int(materijal.Kolicina / float64(materijal.Materijal.Struct["Pakovanje"].Content.(int)))
		//materijal.UkupnoPakovanja = int(materijal.Kolicina / w.PrikazaniElement.el.Struct["Pakovanje"].Content.(float64))
		//fmt.Println("MNMNMN", nm)
		w.PrikazaniElement.mat = append(w.PrikazaniElement.mat, materijal)
	}

}

func (w *WingCal) PrikazaniElementSuma() func(gtx C) D {
	return func(gtx C) D {
		return container.DuoUIcontainer(w.UI.Tema, 0, w.UI.Tema.Colors["Gray"]).Layout(gtx, layout.NW, func(gtx C) D {
			gtx.Constraints.Min.X = gtx.Constraints.Max.X
			sumaCena := float64(w.UI.Counters.Kolicina.Value) * w.PrikazaniElement.el.Struct["Cena"].Content.(float64)
			return lyt.Format(gtx, "hflexb(middle,f(1,_),r(_))",
				func(gtx C) D {
					return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
						func(gtx C) D {
							return container.DuoUIcontainer(w.UI.Tema, 10, w.UI.Tema.Colors["LightGrayII"]).Layout(gtx, layout.NW, func(gtx C) D {
								gtx.Constraints.Min.X = gtx.Constraints.Max.X
								return material.Body2(w.UI.Tema.T, w.text("Cena:")+fmt.Sprint(w.PrikazaniElement.el.Struct["Cena"].Content.(float64))).Layout(gtx)
							})
						},
						helper.DuoUIline(false, 0, 0, 1, w.UI.Tema.Colors["Dark"]),
						func(gtx C) D {
							return container.DuoUIcontainer(w.UI.Tema, 10, w.UI.Tema.Colors["LightGrayII"]).Layout(gtx, layout.NW, func(gtx C) D {
								gtx.Constraints.Min.X = gtx.Constraints.Max.X
								return material.Body2(w.UI.Tema.T, latcyr.C("Suma:", w.Podesavanja.Cyr)+fmt.Sprintf("%.2f", prikazaniElementSumaCena)).Layout(gtx)

							})
						})
				},
				func(gtx C) D {
					return layout.Inset{}.Layout(gtx, func(gtx C) D {
						return lyt.Format(gtx, "vflexb(middle,r(_),r(_))",
							counter.DuoUIcounterSt(w.UI.Tema, w.UI.Counters.Kolicina).Layout(gtx, w.UI.Tema.T, latcyr.C("KOLIČINA", w.Podesavanja.Cyr), fmt.Sprint(w.UI.Counters.Kolicina.Value)),
							w.PrikazaniElementDugmeDodaj(sumaCena),
						)
					})
				})
		})
	}
}

func recept(s string) (m []map[string]int) {
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		log.Println(err)
	}
	return
}
