package mod

type WingProjekat struct {
	Id             int
	Naziv          string
	Opis           string
	IdProjekta     string
	BrojDokumenta  string
	Sveska         string
	VrstaDokumenta string
	DatumDokumenta string
	Objekti        []*WingObjekat
	Investitor     *WingPravnoLice
	Projektant     *WingPravnoLice
	Elementi       *WingIzabraniElementi
	RabatRadovi    int
	RabatMaterijal int

	//OdgovorniProjektant *WingOdgovorniProjektant
	//Dokumntacija        WingDokumentacija
}

type WingObjekat struct {
	BrojObjekta          string
	KategorijaObjekta    string
	KlasifikacijaObjekta string
	Funkcija             string
	Gradjenje            string
	Spratnost            string
	Lokacija             string
	Ulica                string
	Broj                 string
	KP                   string
	KO                   string
}

type WingPravnoLice struct {
	Id             int
	Naziv          string
	Projektant     bool
	DugiNaziv      string
	BrojLicence    string
	Adresa         string
	Grad           string
	Region         string
	PIB            string
	MB             string
	DatumOsnivanja string
	Delatnost      string
	Racuni         []WingBankaRacun
	Email          string
	BrojTelefona   string
}

type WingFizickoLice struct {
	Id           int
	Naziv        string
	Projektant   bool
	Admin        bool
	Ime          string
	Prezime      string
	JMBG         string
	BrojLicence  string
	Adresa       string
	Grad         string
	Region       string
	Delatnost    string
	Racuni       []WingBankaRacun
	Email        string
	BrojTelefona string
}

type WingBankaRacun struct {
	Banka string
	Racun string
}

type WingDokumentacija struct {
	Tekstualna string
	Numericka  string
	Graficka   string
}
