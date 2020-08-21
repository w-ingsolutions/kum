package db

import (
	"github.com/w-ingsolutions/kum/mod"
)

func NewLica() (p []*mod.WingPravnoLice, k []*mod.WingPravnoLice) {
	for _, l := range lica() {
		if l.Projektant {
			p = append(p, l)
		} else {
			k = append(k, l)
		}
	}
	return
}
func lica() map[int]*mod.WingPravnoLice {
	return map[int]*mod.WingPravnoLice{
		0: &mod.WingPravnoLice{
			Id:             0,
			Projektant:     true,
			Naziv:          "W-ING SOLUTIONS DOO",
			DugiNaziv:      "W-ING SOLUTIONS DOO NOVI SAD",
			MB:             "20701005",
			PIB:            "106892584",
			Adresa:         "BULEVAR OSLOBOĐENJA 30 A",
			Grad:           "21000 Novi Sad",
			DatumOsnivanja: "12/29/2010",
			Delatnost:      "Engineering activities and related technical consultancy",
			Racuni: []mod.WingBankaRacun{
				mod.WingBankaRacun{
					Banka: "Erste Bank a.d. Novi Sad",
					Racun: "340-0000013011331-95",
				},
				mod.WingBankaRacun{
					Banka: "Erste Bank a.d. Novi Sad",
					Racun: "340-0000013002280-88",
				},
				mod.WingBankaRacun{
					Banka: "Erste Bank a.d. Novi Sad",
					Racun: "340-0000011005651-31",
				},
				mod.WingBankaRacun{
					Banka: "Erste Bank a.d. Novi Sad",
					Racun: "340-0000010008608-68",
				},
			},
			Email:        "adresa@mail.com",
			BrojTelefona: "063/0000000",
		},
		1: &mod.WingPravnoLice{
			Id:             0,
			DugiNaziv:      "PREDUZEĆE ZA PROIZVODNJU, PROMET I USLUGE ENERGOTEHNA DOO, NOVI SAD",
			Naziv:          "ENERGOTEHNA DOO NOVI SAD",
			Adresa:         "KIJEVSKA 24",
			Grad:           "21000 Novi Sad",
			MB:             "20113278",
			PIB:            "104202125",
			DatumOsnivanja: "28.11.2005.",
			Delatnost:      "Izgradnja cevovoda",
			Racuni: []mod.WingBankaRacun{
				mod.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0002024302784-84",
				},
				mod.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0000000017631-17",
				},
				mod.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0000000005685-62",
				},
				mod.WingBankaRacun{
					Banka: "Halkbank a.d. Beograd",
					Racun: "155-1000000031626-42",
				},
				mod.WingBankaRacun{
					Banka: "Halkbank a.d. Beograd",
					Racun: "155-0000000021603-94",
				},
			},
			Email:        "adresa@mail.com",
			BrojTelefona: "063/0000000",
		},
		2: &mod.WingPravnoLice{
			Id:             0,
			Naziv:          "DAVID NICHOLAS VENNIK PR ZORA MREŽA",
			DugiNaziv:      "DAVID NICHOLAS VENNIK PR RAČUNARSKO PROGRAMIRANJE ZORA MREŽA NOVI SAD",
			MB:             "ccccccccccccc",
			PIB:            "ccccccccccccc",
			Adresa:         "PAJE MARKOVIĆA ADAMOVA 11 24",
			Grad:           "21000 Novi Sad",
			DatumOsnivanja: "15.5.2017.",
			Delatnost:      "Računarsko programiranje",
			Racuni: []mod.WingBankaRacun{
				mod.WingBankaRacun{
					Banka: "Erste Bank a.d. Novi Sad",
					Racun: "340-0000013019687-53",
				},
				mod.WingBankaRacun{
					Banka: "Erste Bank a.d. Novi Sad",
					Racun: "340-0000011419079-86",
				},
				mod.WingBankaRacun{
					Banka: "Erste Bank a.d. Novi Sad",
					Racun: "340-0000010025142-33",
				},
			},
			Email:        "adresa@mail.com",
			BrojTelefona: "063/0000000",
		},
		3: &mod.WingPravnoLice{
			Id:             0,
			Naziv:          "PROMONT-PRODANOVIĆ",
			DugiNaziv:      "SZR PROMONT-PRODANOVIĆ PRODANOVIĆ RADIVOJ PREDUZETNIK NOVI SAD",
			MB:             "54516711",
			PIB:            "101648529",
			Adresa:         "Partizanskih Baza 3",
			Grad:           "21000 Novi Sad",
			DatumOsnivanja: "25.5.1999.",
			Delatnost:      "Proizvodnja ostalih mašina i aparata opšte namene",
			Racuni: []mod.WingBankaRacun{
				mod.WingBankaRacun{
					Banka: "NLB banka a.d. Beograd",
					Racun: "310-0000000222227-19",
				},
				mod.WingBankaRacun{
					Banka: "NLB banka a.d. Beograd",
					Racun: "310-0000000007977-47",
				},
			},
			Email:        "adresa@mail.com",
			BrojTelefona: "063/0000000",
		},
		4: &mod.WingPravnoLice{
			Id:             0,
			Naziv:          "CESLA IB INVEST DOO NOVI SAD",
			DugiNaziv:      "CESLA IB INVEST DOO ZA TRGOVINU, GRAĐEVINARSTVO I USLUGE, NOVI SAD",
			MB:             "8801533",
			PIB:            "103146555",
			Adresa:         "RADNIČKA 28",
			Grad:           "21000 Novi Sad",
			DatumOsnivanja: "11.11.2003.",
			Delatnost:      "Ostali nepomenuti specifični građevinski radovi",
			Racuni: []mod.WingBankaRacun{
				mod.WingBankaRacun{
					Banka: "ProCredit Bank a.d. Beograd",
					Racun: "220-2058000000030-91",
				},
				mod.WingBankaRacun{
					Banka: "ProCredit Bank a.d. Beograd",
					Racun: "220-0000000147006-36",
				},
				mod.WingBankaRacun{
					Banka: "ProCredit Bank a.d. Beograd",
					Racun: "220-0000000012130-77",
				},
				mod.WingBankaRacun{
					Banka: "Banca Intesa a.d. Beograd",
					Racun: "160-0053600000825-49",
				},
				mod.WingBankaRacun{
					Banka: "Banca Intesa a.d. Beograd",
					Racun: "160-0000000109180-08",
				},
			},
			Email:        "adresa@mail.com",
			BrojTelefona: "063/0000000",
		},
		5: &mod.WingPravnoLice{
			Id:             0,
			Naziv:          "TEL-ING DOO NOVI SAD",
			DugiNaziv:      "TEL-ING DOO ZA TRGOVINU I USLUGE NOVI SAD",
			MB:             "20351306",
			PIB:            "105277334",
			Adresa:         "KISAČKA 64 A",
			Grad:           "21000 Novi Sad",
			DatumOsnivanja: "9.11.2007.",
			Delatnost:      "Postavljanje električnih instalacija",
			Racuni: []mod.WingBankaRacun{
				mod.WingBankaRacun{
					Banka: "Unicredit Bank Srbija a.d. Beograd",
					Racun: "170-0030018578000-27",
				},

				mod.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0007010458727-31",
				},
				mod.WingBankaRacun{
					Banka: "Banca Intesa a.d. Beograd",
					Racun: "60-6000000224495-35",
				},
				mod.WingBankaRacun{
					Banka: "Banca Intesa a.d. Beograd",
					Racun: "160-0050100237978-54",
				},
				mod.WingBankaRacun{
					Banka: "Banca Intesa a.d. Beograd",
					Racun: "160-0000000331525-42",
				},
			},
			Email:        "adresa@mail.com",
			BrojTelefona: "063/0000000",
		},
		6: &mod.WingPravnoLice{
			Id:             0,
			Naziv:          "DON DON DOO BEOGRAD",
			DugiNaziv:      "PRIVREDNO DRUŠTVO ZA PROIZVODNJU HLEBA I PECIVA DON DON DOO BEOGRAD",
			MB:             "20383399",
			PIB:            "105425574",
			Adresa:         "BULEVAR ZORANA ĐINĐIĆA 144B",
			Grad:           "11070 NOVI BEOGRAD",
			DatumOsnivanja: "13.2.2008.",
			Delatnost:      "Proizvodnja hleba, svežeg peciva i kolača",
			Racuni: []mod.WingBankaRacun{
				mod.WingBankaRacun{
					Banka: "Societe Generale banka Srbija a.d. Beograd",
					Racun: "275-0010221286932-18",
				},
				mod.WingBankaRacun{
					Banka: "Societe Generale banka Srbija a.d. Beograd",
					Racun: "275-0010221286916-66",
				},
				mod.WingBankaRacun{
					Banka: "Unicredit Bank Srbija a.d. Beograd",
					Racun: "170-0030013281320-12",
				},
				mod.WingBankaRacun{
					Banka: "Unicredit Bank Srbija a.d. Beograd",
					Racun: "170-0030013281000-02",
				},
				mod.WingBankaRacun{
					Banka: "Agroindustrijska komercijalna banka - AIK banka a.d. Niš",
					Racun: "105-0510120009916-12",
				},
			},
			Email:        "adresa@mail.com",
			BrojTelefona: "063/0000000",
		},
		7: &mod.WingPravnoLice{
			Id:             0,
			Naziv:          "DOO ŠUŠA VETERNIK",
			DugiNaziv:      "DOO ŠUŠA ZA GRAĐEVINARSTVO, TRGOVINU I USLUGE VETERNIK",
			MB:             "8825106",
			PIB:            "103567021",
			Adresa:         "DODOLSKA 4",
			Grad:           "21203 Veternik",
			DatumOsnivanja: "ccccccccccccc",
			Delatnost:      "ccccccccccccc",
			Racuni: []mod.WingBankaRacun{
				mod.WingBankaRacun{
					Banka: "Ministarstvo finansija - Uprava za trezor",
					Racun: "840-0000023864763-84",
				},
				mod.WingBankaRacun{
					Banka: "Komercijalna banka a.d. Beograd",
					Racun: "205-0070100423764-67",
				},
				mod.WingBankaRacun{
					Banka: "Komercijalna banka a.d. Beograd",
					Racun: "205-0000000182955-49",
				},
				mod.WingBankaRacun{
					Banka: "Banca Intesa a.d. Beograd",
					Racun: "160-6000000223443-87",
				},
				mod.WingBankaRacun{
					Banka: "Banca Intesa a.d. Beograd",
					Racun: "160-0050170161774-90",
				},
			},
			Email:        "adresa@mail.com",
			BrojTelefona: "063/0000000",
		},
		8: &mod.WingPravnoLice{
			Id:             0,
			Naziv:          "URBANS INVEST DOO NOVI SAD",
			DugiNaziv:      "URBANS INVEST DOO ZA GRAĐEVINARSTVO, UNUTRAŠNJU I SPOLJNU TRGOVINU NOVI SAD",
			MB:             "20422998",
			PIB:            "105625475",
			Adresa:         "Železnička 39/10",
			Grad:           "21000 Novi Sad",
			DatumOsnivanja: "20.5.2008.",
			Delatnost:      "Izgradnja stambenih i nestambenih zgrada",
			Racuni: []mod.WingBankaRacun{
				mod.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0007011600535-65",
				},
				mod.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0007011534435-97",
				},
				mod.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0007011534408-81",
				},
				mod.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0007011534378-74",
				},
				mod.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0002024305088-59",
				},
			},
			Email:        "adresa@mail.com",
			BrojTelefona: "063/0000000",
		},
		9: &mod.WingPravnoLice{
			Id:             0,
			Naziv:          "DREAM HOME REAL ESTATE DOO NOVI SAD",
			DugiNaziv:      "DREAM HOME REAL ESTATE DOO NOVI SAD",
			MB:             "21278815",
			PIB:            "109979047",
			Adresa:         "TONE HADŽIĆA 12",
			Grad:           "21000 Novi Sad",
			DatumOsnivanja: "4.4.2017.",
			Delatnost:      "Izgradnja stambenih i nestambenih zgrada",
			Racuni: []mod.WingBankaRacun{
				mod.WingBankaRacun{
					Banka: "Raiffeisen banka a.d. Beograd",
					Racun: "265-1000000186611-05",
				},
				mod.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0007009210509-97",
				},
				mod.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0007009210479-90",
				},
				mod.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0007009210444-98",
				},
				mod.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0007009210417-82",
				},
			},
			Email:        "adresa@mail.com",
			BrojTelefona: "063/0000000",
		},
	}
}
