package mod

import (
	"gioui.org/widget"
	"github.com/w-ingsolutions/cms/pkg/phi"
)

type WingVrstaRadova struct {
	Id                 int
	Naziv              string
	Opis               string
	Obracun            string
	Jedinica           string
	Cena               float64
	Slug               string
	Omogucen           bool
	Baza               bool
	Element            bool
	PodvrsteRadova     map[int]WingVrstaRadova
	NeophodanMaterijal map[int]WingNeophodanMaterijal
}

type WingIzabraniElement struct {
	Sifra         string
	Kolicina      int
	SumaCena      float64
	Element       *phi.Φ
	DugmeBrisanje *widget.Clickable
}

type WingSuma struct {
	Id                       string
	SumaCena                 float64
	SumaCenaMaterijal        float64
	Elementi                 []*WingIzabraniElement
	NeophodanMaterijal       map[int]WingNeophodanMaterijal
	NeophodanMaterijalPrikaz map[int]WingNeophodanMaterijal
}

type WingCalGrupaRadova struct {
	Id       string
	Slug     string
	Elementi map[int]WingVrstaRadova
}
