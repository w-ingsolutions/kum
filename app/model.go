package calc

import (
	"context"
	"github.com/w-ingsolutions/cms/pkg/phi"
	"github.com/w-ingsolutions/kum/app/mod"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"github.com/gioapp/gel/counter"
	"github.com/gioapp/gel/theme"
	"github.com/w-ingsolutions/c/pkg/cache"
	"github.com/w-ingsolutions/c/pkg/translate"
	"github.com/w-ingsolutions/kum/pkg/jdb"
)

type (
	D = layout.Dimensions
	C = layout.Context
)

var (
	//post             = new(mod.DuoCMSpost)
	prikazaniElementSumaCena float64
	selected                 int
	projekat                 = &mod.WingProjekat{
		//Materijal: make(map[int]mod.WingNeophodanMaterijal),
		//Elementi: new(mod.WingIzabraniElementi),
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
	EditPolja *mod.EditabilnaPoljaVrsteRadova
	Materijal map[int]phi.Φ
	Lica      WingUloge
	//Radovi         mod.WingVrstaRadova
	Putanja        []string
	IzbornikRadova map[int]mod.ElementMenu

	Transfered       mod.WingCalGrupaRadova
	PrikazaniElement *WingPrikazaniElement
	Suma             *mod.WingSuma
	Podvrsta         int
	Roditelj         WingRoditelj
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

type WingRoditelj struct {
	id   int
	hash string
}
type WingPrikazaniElement struct {
	el  *phi.Φ
	mat []mod.WingNeophodanMaterijal
}

//
//type WingRadoviIzbornik struct {
//	VrstaRadova   int
//	PodKategorija int
//	Element       int
//}
//
//func (r *WingRadoviIzbornik)RadoviIzbornik(){
//
//}

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
	Projektanti []*mod.WingPravnoLice
	Investotori []*mod.WingPravnoLice
}
type WingCounters struct {
	Kolicina  *counter.DuoUIcounter
	Radovi    *counter.DuoUIcounter
	Materijal *counter.DuoUIcounter
}
