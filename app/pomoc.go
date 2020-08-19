package calc

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/gioapp/gel/container"
)

func (w *WingCal) PomocStrana() func(gtx C) D {
	return func(gtx C) D {
		return container.DuoUIcontainer(w.UI.Tema, 1, w.UI.Tema.Colors["DarkGrayI"]).Layout(gtx, layout.N, func(gtx C) D {
			return container.DuoUIcontainer(w.UI.Tema, 0, w.UI.Tema.Colors["White"]).Layout(gtx, layout.N, func(gtx C) D {

				return pomocList.Layout(gtx, len(w.pomoc()), func(gtx C, i int) D {
					return layout.Flex{}.Layout(gtx,
						layout.Flexed(0.5, w.pomoc()[i]))
				})

			})
		})
	}
}

func (w *WingCal) pomoc() []func(gtx C) D {
	return []func(gtx C) D{
		func(gtx C) D {
			return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["White"]).Layout(gtx, layout.Center, func(gtx C) D {
				t := material.H6(w.UI.Tema.T, "Nesto objasni")
				t.Alignment = text.Start
				return t.Layout(gtx)
			})
		},
		func(gtx C) D {
			return container.DuoUIcontainer(w.UI.Tema, 8, w.UI.Tema.Colors["White"]).Layout(gtx, layout.Center, func(gtx C) D {
				t := material.Body1(w.UI.Tema.T, "Nesto objasni Nesto objasni opisi duze i vise dalje Nesto objasni Nesto objasni opisi duze i vise dalje Nesto objasni Nesto objasni opisi duze i vise dalje Nesto objasni Nesto objasni opisi duze i vise dalje Nesto objasni Nesto objasni opisi duze i vise dalje Nesto objasni Nesto objasni opisi duze i vise dalje Nesto objasni Nesto objasni opisi duze i vise dalje Nesto objasni Nesto objasni opisi duze i vise dalje ")
				t.Alignment = text.Start
				return t.Layout(gtx)
			})
		},
	}
}
