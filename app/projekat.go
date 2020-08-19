package calc

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
	"github.com/gioapp/gel/counter"
	"github.com/w-ingsolutions/c/pkg/latcyr"
	"strconv"
)

func (w *WingCal) ProjekatStrana() func(gtx C) D {
	return func(gtx C) D {
		return container.DuoUIcontainer(w.UI.Tema, 1, w.UI.Tema.Colors["DarkGrayI"]).Layout(gtx, layout.N, func(gtx C) D {
			return container.DuoUIcontainer(w.UI.Tema, 0, w.UI.Tema.Colors["White"]).Layout(gtx, layout.N, func(gtx C) D {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Flexed(1, func(gtx C) D {
						return projekatList.Layout(gtx, len(w.projekat()), func(gtx C, i int) D {
							return layout.Flex{}.Layout(gtx,
								layout.Flexed(0.5, w.projekat()[i]))
						})
					}),
					layout.Rigid(w.Stampa()))
			})
		})
	}
}

func (w *WingCal) projekat() []func(gtx C) D {
	return []func(gtx C) D{
		func(gtx C) D { return material.H6(w.UI.Tema.T, w.text("Projektant")).Layout(gtx) },
		w.projektant(),
		func(gtx C) D { return material.H6(w.UI.Tema.T, w.text("Investitor")).Layout(gtx) },
		w.investitori(),
		func(gtx C) D { return material.H6(w.UI.Tema.T, w.text("Radovi")).Layout(gtx) },
		w.radovi(),
		func(gtx C) D { return material.H6(w.UI.Tema.T, w.text("Radovi")).Layout(gtx) },
		w.SumaElementi(projekat.Elementi.Elementi),
		func(gtx C) D { return material.H6(w.UI.Tema.T, w.text("Materijal")).Layout(gtx) },
		w.materijal(),
		func(gtx C) D { return material.H6(w.UI.Tema.T, w.text("Materijal")).Layout(gtx) },
		w.UkupanNeophodanMaterijal(projekat.Elementi.NeophodanMaterijal),
	}
}

func (w *WingCal) projektant() func(gtx C) D {
	return func(gtx C) D {
		return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["Gray"]).Layout(gtx, layout.Center, func(gtx C) D {
			id, _ := strconv.Atoi(projektantIzbor.Value)
			projekat.Projektant = w.Lica.Projektanti[id]

			return layout.Flex{}.Layout(gtx,
				layout.Flexed(0.4, func(gtx C) D {
					return projektantiList.Layout(gtx, len(w.Lica.Projektanti), func(gtx C, i int) D {
						p := w.Lica.Projektanti[i]
						return container.DuoUIcontainer(w.UI.Tema, 0, w.UI.Tema.Colors["White"]).Layout(gtx, layout.Center, func(gtx C) D {
							return material.RadioButton(w.UI.Tema.T, projektantIzbor, strconv.Itoa(i), p.Naziv).Layout(gtx)
						})
					})
				}),
				layout.Flexed(0.6, func(gtx C) D {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx C) D {
							return material.Caption(w.UI.Tema.T, w.text("Naziv: ")+projekat.Projektant.Naziv).Layout(gtx)
						}),
						layout.Rigid(func(gtx C) D {
							return material.Caption(w.UI.Tema.T, w.text("Adresa: ")+":"+projekat.Projektant.Adresa+" "+projekat.Projektant.Grad).Layout(gtx)
							//}),
							//layout.Rigid(func(gtx C) D {
							//return material.Caption(w.UI.Tema.T, w.text("Matičmi broj: ")+":"+w.text(projektant.JMBG)).Layout(gtx)
							//}),
							//layout.Rigid(func(gtx C) D {
							//	return material.Caption(w.UI.Tema.T, w.text("Email: ")+":"+w.text(projektant.Email)).Layout(gtx)
						}))
				}))
		})
	}
}

func (w *WingCal) investitori() func(gtx C) D {
	return func(gtx C) D {
		return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["Gray"]).Layout(gtx, layout.Center, func(gtx C) D {
			id, _ := strconv.Atoi(klijentiIzbor.Value)
			projekat.Investitor = w.Lica.Investotori[id]
			return layout.Flex{}.Layout(gtx,
				layout.Flexed(0.4, func(gtx C) D {
					return klijentiList.Layout(gtx, len(w.Lica.Investotori), func(gtx C, i int) D {
						inv := w.Lica.Investotori[i]
						return container.DuoUIcontainer(w.UI.Tema, 0, w.UI.Tema.Colors["White"]).Layout(gtx, layout.Center, func(gtx C) D {
							return material.RadioButton(w.UI.Tema.T, klijentiIzbor, strconv.Itoa(i), inv.Naziv).Layout(gtx)
						})
					})
				}),
				layout.Flexed(0.6, w.lice()))
		})
	}
}

func (w *WingCal) lice() func(gtx C) D {
	return func(gtx C) D {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				return material.Caption(w.UI.Tema.T, w.text("Naziv: ")+w.text(projekat.Investitor.Naziv)).Layout(gtx)
			}),
			layout.Rigid(func(gtx C) D {
				return material.Caption(w.UI.Tema.T, w.text("Adresa: ")+":"+w.text(projekat.Investitor.Adresa)+" "+w.text(projekat.Investitor.Grad)).Layout(gtx)
			}),
			layout.Rigid(func(gtx C) D {
				return material.Caption(w.UI.Tema.T, w.text("PIB: ")+":"+w.text(projekat.Investitor.PIB)).Layout(gtx)
			}),
			layout.Rigid(func(gtx C) D {
				return material.Caption(w.UI.Tema.T, w.text("Matičmi broj: ")+":"+w.text(projekat.Investitor.MB)).Layout(gtx)
			}),
			layout.Rigid(func(gtx C) D {
				return material.Caption(w.UI.Tema.T, w.text("Datum osnivanja: ")+":"+w.text(projekat.Investitor.DatumOsnivanja)).Layout(gtx)
			}),
			layout.Rigid(func(gtx C) D {
				return material.Caption(w.UI.Tema.T, w.text("Delatnost: ")+":"+w.text(projekat.Investitor.Delatnost)).Layout(gtx)
			}),
			layout.Rigid(func(gtx C) D {
				return material.Caption(w.UI.Tema.T, w.text("Računi: ")).Layout(gtx)
			}),
			layout.Rigid(func(gtx C) D {
				return racuniList.Layout(gtx, len(projekat.Investitor.Racuni), func(gtx C, i int) D {
					racuni := projekat.Investitor.Racuni[i]
					return container.DuoUIcontainer(w.UI.Tema, 0, w.UI.Tema.Colors["White"]).Layout(gtx, layout.Center, func(gtx C) D {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text(racuni.Banka)).Layout(gtx)
							}),
							layout.Rigid(func(gtx C) D {
								return material.Caption(w.UI.Tema.T, w.text(racuni.Racun)).Layout(gtx)
							}))
					})
				})
			}),
			layout.Rigid(func(gtx C) D {
				return material.Caption(w.UI.Tema.T, w.text("Email: ")+":"+w.text(projekat.Investitor.Email)).Layout(gtx)
			}))
	}
}

func (w *WingCal) radovi() func(gtx C) D {
	return func(gtx C) D {
		return container.DuoUIcontainer(w.UI.Tema, 0, w.UI.Tema.Colors["Gray"]).Layout(gtx, layout.Center, func(gtx C) D {
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				layout.Flexed(0.5, func(gtx C) D {
					return material.H6(w.UI.Tema.T, w.text("Ukupna cena radova")).Layout(gtx)
				}),
				layout.Flexed(0.5, func(gtx C) D {
					return material.H6(w.UI.Tema.T, fmt.Sprintf("%.2f", w.Suma.SumaCena)).Layout(gtx)
				}),
				layout.Rigid(counter.DuoUIcounterSt(w.UI.Tema, w.UI.Counters.Radovi).Layout(gtx, w.UI.Tema.T, "%", fmt.Sprint(w.UI.Counters.Radovi.Value))))
		})
	}
}

func (w *WingCal) materijal() func(gtx C) D {
	return func(gtx C) D {
		return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["Gray"]).Layout(gtx, layout.Center, func(gtx C) D {
			return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
				layout.Flexed(0.5, func(gtx C) D {
					return material.H6(w.UI.Tema.T, w.text("Ukupna cena materijala")).Layout(gtx)
				}),
				layout.Flexed(0.5, func(gtx C) D {
					return material.H6(w.UI.Tema.T, fmt.Sprintf("%.2f", w.Suma.SumaCenaMaterijal)).Layout(gtx)
				}),
				layout.Rigid(counter.DuoUIcounterSt(w.UI.Tema, w.UI.Counters.Materijal).Layout(gtx, w.UI.Tema.T, latcyr.C("%", w.Podesavanja.Cyr), fmt.Sprint(w.UI.Counters.Materijal.Value))))
		})
	}
}
