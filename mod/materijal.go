package mod

import "github.com/w-ingsolutions/cms/pkg/phi"

//
//type WingMaterijal struct {
//	Id      int    `json:"id"`
//	Naziv   string `json:"naziv"`
//	Opis    string `json:"opis"`
//	Obracun string `json:"obracun"`
//
//	Proizvodjac       string  `json:"proizvodjac"`
//	OsobineNamena     string  `json:"osobinenamena"`
//	NacinRada         string  `json:"nacinrada"`
//	JedinicaPotrosnje string  `json:"jedinicapotrosnje"`
//	Potrosnja         float64 `json:"potrosnja"`
//	RokUpotrebe       string  `json:"rokupotreba"`
//
//	Jedinica  string  `json:"jedinica"`
//	Pakovanje int     `json:"pakovanje"`
//	Cena      float64 `json:"cena"`
//	Slug      string  `json:"slug"`
//}

type WingNeophodanMaterijal struct {
	Id              int
	Kolicina        float64
	Koeficijent     float64
	UkupnoPakovanja int
	UkupnaCena      float64
	Materijal       phi.Φ
}
