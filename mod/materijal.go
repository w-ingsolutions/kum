package mod

type WingMaterijal struct {
	Id      int    `json:"id"`
	Naziv   string `json:"naziv"`
	Opis    string `json:"opis"`
	Obracun string `json:"obracun"`

	Proizvodjac       string  `json:"proizvodjac"`
	OsobineNamena     string  `json:"osobinenamena"`
	NacinRada         string  `json:"nacinrada"`
	JedinicaPotrosnje string  `json:"jedinicapotrosnje"`
	Potrosnja         float64 `json:"potrosnja"`
	RokUpotrebe       string  `json:"rokupotreba"`

	Jedinica  string  `json:"jedinica"`
	Pakovanje int     `json:"pakovanje"`
	Cena      float64 `json:"cena"`
	Slug      string  `json:"slug"`
}

type WingNeophodanMaterijal struct {
	Id              int           `json:"id"`
	Kolicina        float64       `json:"kolicina"`
	Koeficijent     float64       `json:"koeficijent"`
	UkupnoPakovanja int           `json:"ukupnopakovanja"`
	UkupnaCena      float64       `json:"ukupnacena"`
	Materijal       WingMaterijal `json:"materijal"`
}
