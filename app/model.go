package calc

import (
	"context"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"github.com/gioapp/gel/counter"
	"github.com/gioapp/gel/theme"
	"github.com/w-ingsolutions/c/model"
	"github.com/w-ingsolutions/c/pkg/cache"
	"github.com/w-ingsolutions/c/pkg/translate"
	"github.com/w-ingsolutions/kum/pkg/jdb"
)

type (
	D = layout.Dimensions
	C = layout.Context
)

var (
	//post             = new(model.DuoCMSpost)
	prikazaniElementSumaCena float64
	selected                 int
	projekat                 = &model.WingProjekat{
		//Materijal: make(map[int]model.WingNeophodanMaterijal),
		//Elementi: new(model.WingIzabraniElementi),
	}
	latCyrBool = new(widget.Bool)

	projektantIzbor = new(widget.Enum)
	klijentiIzbor   = new(widget.Enum)
)

type WingCal struct {
	Strana string
	ctx    context.Context
	Jdb    *jdb.JavazacDB
	//LinkoviIzboraVrsteRadova map[int]*widget.Clickable
	EditPolja        *model.EditabilnaPoljaVrsteRadova
	Materijal        map[int]*model.WingMaterijal
	Lica             WingUloge
	Radovi           model.WingVrstaRadova
	Putanja          []string
	IzbornikRadova   map[int]model.ElementMenu
	Transfered       model.WingCalGrupaRadova
	PrikazaniElement *model.WingVrstaRadova
	Suma             *model.WingIzabraniElementi
	Podvrsta         int
	Roditelj         int
	Element          bool
	UI               WingUI
	API              WingAPI
	Jezik            WingJezik
	Kes              cache.Cache
	Podesavanja      *WingPodesavanja
}

type WingUI struct {
	Device      string
	TopSpace    int
	BottomSpace int
	Window      *app.Window
	Tema        *theme.DuoUItheme
	Context     layout.Context
	Ekran       func(gtx layout.Context) layout.Dimensions
	Ops         op.Ops
	Counters    WingCounters
}

type WingAPI struct {
	OK     bool
	Adresa string
}

type WingJezik struct {
	t        translate.Translation
	izabrani string
	dostupni []string
	linkovi  map[string]*widget.Clickable
}

type WingPodesavanja struct {
	Naziv string
	Dir   string
	File  string
	Cyr   bool
}

type WingUloge struct {
	Projektanti []*model.WingPravnoLice
	Investotori []*model.WingPravnoLice
}
type WingCounters struct {
	Kolicina  *counter.DuoUIcounter
	Radovi    *counter.DuoUIcounter
	Materijal *counter.DuoUIcounter
}
