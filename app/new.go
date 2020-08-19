package calc

import (
	"context"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioapp/gel/counter"
	"github.com/gioapp/gel/theme"
	"github.com/w-ingsolutions/c/model"
	wapp "github.com/w-ingsolutions/c/pkg/app"
	"github.com/w-ingsolutions/c/pkg/icons"
	"github.com/w-ingsolutions/c/pkg/latcyr"
	"github.com/w-ingsolutions/c/pkg/translate"
	"github.com/w-ingsolutions/kum/db"
	"github.com/w-ingsolutions/kum/pkg/jdb"
	"path/filepath"
)

func NewWingCal() *WingCal {
	w := &WingCal{
		Strana:           "radovi",
		ctx:              context.Background(),
		PrikazaniElement: &model.WingVrstaRadova{},
		Suma: &model.WingIzabraniElementi{
			Elementi:           []*model.WingIzabraniElement{},
			NeophodanMaterijal: make(map[int]model.WingNeophodanMaterijal),
		},
	}
	w.Jdb = jdb.New(w.ctx, "datastore")
	w.Podesavanja = &WingPodesavanja{
		Naziv: "Kalkulator",
		Dir:   wapp.Dir("wing", false),
		Cyr:   false,
	}
	w.Podesavanja.File = filepath.Join(w.Podesavanja.Dir, "conf.json")
	w.Materijal = db.NewMaterijal()
	projektanti, klijenti := db.NewLica()
	w.Lica.Projektanti = projektanti
	w.Lica.Investotori = klijenti
	saStraneMarginom := layout.UniformInset(unit.Dp(0))
	saStraneMarginom.Left = unit.Dp(8)
	saStraneMarginom.Right = unit.Dp(8)
	w.Radovi = model.WingVrstaRadova{
		Id:       0,
		Naziv:    "Radovi",
		Slug:     "radovi",
		Omogucen: false,
		Baza:     false,
		Element:  false,
	}
	w.UI = WingUI{
		Device:      "p",
		TopSpace:    28,
		BottomSpace: 56,
		Tema:        theme.NewDuoUItheme(),
	}
	w.API = WingAPI{
		OK:     true,
		Adresa: "https://wing.marcetin.com/",
	}
	w.Jezik = WingJezik{
		t: translate.Translation{"sr", "sr"},
		//dostupni: []string{"en", "ru", "de", "sl", "gr", "zh", "jp", "it"},
		linkovi: make(map[string]*widget.Clickable),
	}
	w.GenerisanjeDugmicaJezici()
	w.Kes = make(map[string]string)

	w.UI.Window = app.NewWindow(
		app.Size(unit.Dp(999), unit.Dp(1024)),
		app.Title("W-ing "+w.Jezik.t.T(latcyr.C(w.Podesavanja.Naziv, w.Podesavanja.Cyr))),
	)

	counters := WingCounters{
		Kolicina: &counter.DuoUIcounter{
			Value:        1,
			OperateValue: 1,
			From:         1,
			To:           999,
			CounterInput: &widget.Editor{
				Alignment:  text.Middle,
				SingleLine: true,
				Submit:     true,
			},
			PageFunction:    w.PrikazaniElementSumaRacunica(),
			CounterIncrease: new(widget.Clickable),
			CounterDecrease: new(widget.Clickable),
			CounterReset:    new(widget.Clickable),
		},
		Radovi: &counter.DuoUIcounter{
			Value:        100,
			OperateValue: 1,
			From:         1,
			To:           999,
			CounterInput: &widget.Editor{
				Alignment:  text.Middle,
				SingleLine: true,
			},
			PageFunction:    w.ProjekatRacunica(),
			CounterIncrease: new(widget.Clickable),
			CounterDecrease: new(widget.Clickable),
			CounterReset:    new(widget.Clickable),
		},
		Materijal: &counter.DuoUIcounter{
			Value:        100,
			OperateValue: 1,
			From:         1,
			To:           999,
			CounterInput: &widget.Editor{
				Alignment:  text.Middle,
				SingleLine: true,
			},
			PageFunction:    w.ProjekatRacunica(),
			CounterIncrease: new(widget.Clickable),
			CounterDecrease: new(widget.Clickable),
			CounterReset:    new(widget.Clickable),
		},
	}
	w.UI.Counters = counters
	w.UI.Tema.Icons = icons.NewWingUIicons()
	w.Putanja = append(w.Putanja, w.Radovi.Naziv)
	return w
}
