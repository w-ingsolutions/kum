package calc

import (
	"gioui.org/layout"
	"github.com/w-ingsolutions/c/pkg/lyt"
)

func (w *WingCal) Monitor(izbornikStrana, sumaRadovaStrana, sumaMaterijalStrana func(gtx C) D) func(gtx C) D {
	return func(gtx C) D {
		return layout.Inset{}.Layout(gtx, func(gtx C) D {
			suma := func(gtx C) D {
				return lyt.Format(gtx, "vflexb(middle,f(0.5,_),f(0.5,_))", sumaRadovaStrana, sumaMaterijalStrana)
			}
			if len(w.Suma.Elementi) >= 1 {
				return lyt.Format(gtx, "hflexb(middle,f(0.4,_),f(0.6,_))", izbornikStrana, suma)
			} else {
				return lyt.Format(gtx, "hflexb(middle,f(1,_))", izbornikStrana)
			}
		})
	}
}

//func (w *WingCal) MonitorK(izbornikStrana, sumaRadovaStrana, sumaMaterijalStrana, materijalStrana, projekatStrana func(gtx C) D) func(gtx C) D {
//	return func(gtx C) D {
//		return w.UI.SaMalomMarginom.Layout(gtx, func(gtx C) D {
//			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
//				layout.Flexed(0.3, izbornikStrana),
//				layout.Flexed(0.4, func(gtx C) D {
//					return layout.Flex{
//						Axis: layout.Vertical,
//					}.Layout(gtx, layout.Flexed(0.5, sumaRadovaStrana), layout.Flexed(0.5, sumaMaterijalStrana))
//				}),
//				layout.Flexed(0.4, func(gtx C) D {
//					return layout.Flex{
//						Axis: layout.Vertical,
//					}.Layout(gtx, layout.Flexed(0.5, materijalStrana), layout.Flexed(0.5, projekatStrana))
//				}),
//			)
//		})
//	}
//}
