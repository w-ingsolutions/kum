package db

import (
	"github.com/w-ingsolutions/c/model"
)

func NewLica() (p []*model.WingPravnoLice, k []*model.WingPravnoLice) {
	for _, l := range lica() {
		if l.Projektant {
			p = append(p, l)
		} else {
			k = append(k, l)
		}
	}
	return
}
func lica() map[int]*model.WingPravnoLice {
	return map[int]*model.WingPravnoLice{
		0: &model.WingPravnoLice{
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
			Racuni: []model.WingBankaRacun{
				model.WingBankaRacun{
					Banka: "Erste Bank a.d. Novi Sad",
					Racun: "340-0000013011331-95",
				},
				model.WingBankaRacun{
					Banka: "Erste Bank a.d. Novi Sad",
					Racun: "340-0000013002280-88",
				},
				model.WingBankaRacun{
					Banka: "Erste Bank a.d. Novi Sad",
					Racun: "340-0000011005651-31",
				},
				model.WingBankaRacun{
					Banka: "Erste Bank a.d. Novi Sad",
					Racun: "340-0000010008608-68",
				},
			},
			Email:        "adresa@mail.com",
			BrojTelefona: "063/0000000",
		},
		1: &model.WingPravnoLice{
			Id:             0,
			DugiNaziv:      "PREDUZEĆE ZA PROIZVODNJU, PROMET I USLUGE ENERGOTEHNA DOO, NOVI SAD",
			Naziv:          "ENERGOTEHNA DOO NOVI SAD",
			Adresa:         "KIJEVSKA 24",
			Grad:           "21000 Novi Sad",
			MB:             "20113278",
			PIB:            "104202125",
			DatumOsnivanja: "28.11.2005.",
			Delatnost:      "Izgradnja cevovoda",
			Racuni: []model.WingBankaRacun{
				model.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0002024302784-84",
				},
				model.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0000000017631-17",
				},
				model.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0000000005685-62",
				},
				model.WingBankaRacun{
					Banka: "Halkbank a.d. Beograd",
					Racun: "155-1000000031626-42",
				},
				model.WingBankaRacun{
					Banka: "Halkbank a.d. Beograd",
					Racun: "155-0000000021603-94",
				},
			},
			Email:        "adresa@mail.com",
			BrojTelefona: "063/0000000",
		},
		2: &model.WingPravnoLice{
			Id:             0,
			Naziv:          "DAVID NICHOLAS VENNIK PR ZORA MREŽA",
			DugiNaziv:      "DAVID NICHOLAS VENNIK PR RAČUNARSKO PROGRAMIRANJE ZORA MREŽA NOVI SAD",
			MB:             "ccccccccccccc",
			PIB:            "ccccccccccccc",
			Adresa:         "PAJE MARKOVIĆA ADAMOVA 11 24",
			Grad:           "21000 Novi Sad",
			DatumOsnivanja: "15.5.2017.",
			Delatnost:      "Računarsko programiranje",
			Racuni: []model.WingBankaRacun{
				model.WingBankaRacun{
					Banka: "Erste Bank a.d. Novi Sad",
					Racun: "340-0000013019687-53",
				},
				model.WingBankaRacun{
					Banka: "Erste Bank a.d. Novi Sad",
					Racun: "340-0000011419079-86",
				},
				model.WingBankaRacun{
					Banka: "Erste Bank a.d. Novi Sad",
					Racun: "340-0000010025142-33",
				},
			},
			Email:        "adresa@mail.com",
			BrojTelefona: "063/0000000",
		},
		3: &model.WingPravnoLice{
			Id:             0,
			Naziv:          "PROMONT-PRODANOVIĆ",
			DugiNaziv:      "SZR PROMONT-PRODANOVIĆ PRODANOVIĆ RADIVOJ PREDUZETNIK NOVI SAD",
			MB:             "54516711",
			PIB:            "101648529",
			Adresa:         "Partizanskih Baza 3",
			Grad:           "21000 Novi Sad",
			DatumOsnivanja: "25.5.1999.",
			Delatnost:      "Proizvodnja ostalih mašina i aparata opšte namene",
			Racuni: []model.WingBankaRacun{
				model.WingBankaRacun{
					Banka: "NLB banka a.d. Beograd",
					Racun: "310-0000000222227-19",
				},
				model.WingBankaRacun{
					Banka: "NLB banka a.d. Beograd",
					Racun: "310-0000000007977-47",
				},
			},
			Email:        "adresa@mail.com",
			BrojTelefona: "063/0000000",
		},
		4: &model.WingPravnoLice{
			Id:             0,
			Naziv:          "CESLA IB INVEST DOO NOVI SAD",
			DugiNaziv:      "CESLA IB INVEST DOO ZA TRGOVINU, GRAĐEVINARSTVO I USLUGE, NOVI SAD",
			MB:             "8801533",
			PIB:            "103146555",
			Adresa:         "RADNIČKA 28",
			Grad:           "21000 Novi Sad",
			DatumOsnivanja: "11.11.2003.",
			Delatnost:      "Ostali nepomenuti specifični građevinski radovi",
			Racuni: []model.WingBankaRacun{
				model.WingBankaRacun{
					Banka: "ProCredit Bank a.d. Beograd",
					Racun: "220-2058000000030-91",
				},
				model.WingBankaRacun{
					Banka: "ProCredit Bank a.d. Beograd",
					Racun: "220-0000000147006-36",
				},
				model.WingBankaRacun{
					Banka: "ProCredit Bank a.d. Beograd",
					Racun: "220-0000000012130-77",
				},
				model.WingBankaRacun{
					Banka: "Banca Intesa a.d. Beograd",
					Racun: "160-0053600000825-49",
				},
				model.WingBankaRacun{
					Banka: "Banca Intesa a.d. Beograd",
					Racun: "160-0000000109180-08",
				},
			},
			Email:        "adresa@mail.com",
			BrojTelefona: "063/0000000",
		},
		5: &model.WingPravnoLice{
			Id:             0,
			Naziv:          "TEL-ING DOO NOVI SAD",
			DugiNaziv:      "TEL-ING DOO ZA TRGOVINU I USLUGE NOVI SAD",
			MB:             "20351306",
			PIB:            "105277334",
			Adresa:         "KISAČKA 64 A",
			Grad:           "21000 Novi Sad",
			DatumOsnivanja: "9.11.2007.",
			Delatnost:      "Postavljanje električnih instalacija",
			Racuni: []model.WingBankaRacun{
				model.WingBankaRacun{
					Banka: "Unicredit Bank Srbija a.d. Beograd",
					Racun: "170-0030018578000-27",
				},

				model.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0007010458727-31",
				},
				model.WingBankaRacun{
					Banka: "Banca Intesa a.d. Beograd",
					Racun: "60-6000000224495-35",
				},
				model.WingBankaRacun{
					Banka: "Banca Intesa a.d. Beograd",
					Racun: "160-0050100237978-54",
				},
				model.WingBankaRacun{
					Banka: "Banca Intesa a.d. Beograd",
					Racun: "160-0000000331525-42",
				},
			},
			Email:        "adresa@mail.com",
			BrojTelefona: "063/0000000",
		},
		6: &model.WingPravnoLice{
			Id:             0,
			Naziv:          "DON DON DOO BEOGRAD",
			DugiNaziv:      "PRIVREDNO DRUŠTVO ZA PROIZVODNJU HLEBA I PECIVA DON DON DOO BEOGRAD",
			MB:             "20383399",
			PIB:            "105425574",
			Adresa:         "BULEVAR ZORANA ĐINĐIĆA 144B",
			Grad:           "11070 NOVI BEOGRAD",
			DatumOsnivanja: "13.2.2008.",
			Delatnost:      "Proizvodnja hleba, svežeg peciva i kolača",
			Racuni: []model.WingBankaRacun{
				model.WingBankaRacun{
					Banka: "Societe Generale banka Srbija a.d. Beograd",
					Racun: "275-0010221286932-18",
				},
				model.WingBankaRacun{
					Banka: "Societe Generale banka Srbija a.d. Beograd",
					Racun: "275-0010221286916-66",
				},
				model.WingBankaRacun{
					Banka: "Unicredit Bank Srbija a.d. Beograd",
					Racun: "170-0030013281320-12",
				},
				model.WingBankaRacun{
					Banka: "Unicredit Bank Srbija a.d. Beograd",
					Racun: "170-0030013281000-02",
				},
				model.WingBankaRacun{
					Banka: "Agroindustrijska komercijalna banka - AIK banka a.d. Niš",
					Racun: "105-0510120009916-12",
				},
			},
			Email:        "adresa@mail.com",
			BrojTelefona: "063/0000000",
		},
		7: &model.WingPravnoLice{
			Id:             0,
			Naziv:          "DOO ŠUŠA VETERNIK",
			DugiNaziv:      "DOO ŠUŠA ZA GRAĐEVINARSTVO, TRGOVINU I USLUGE VETERNIK",
			MB:             "8825106",
			PIB:            "103567021",
			Adresa:         "DODOLSKA 4",
			Grad:           "21203 Veternik",
			DatumOsnivanja: "ccccccccccccc",
			Delatnost:      "ccccccccccccc",
			Racuni: []model.WingBankaRacun{
				model.WingBankaRacun{
					Banka: "Ministarstvo finansija - Uprava za trezor",
					Racun: "840-0000023864763-84",
				},
				model.WingBankaRacun{
					Banka: "Komercijalna banka a.d. Beograd",
					Racun: "205-0070100423764-67",
				},
				model.WingBankaRacun{
					Banka: "Komercijalna banka a.d. Beograd",
					Racun: "205-0000000182955-49",
				},
				model.WingBankaRacun{
					Banka: "Banca Intesa a.d. Beograd",
					Racun: "160-6000000223443-87",
				},
				model.WingBankaRacun{
					Banka: "Banca Intesa a.d. Beograd",
					Racun: "160-0050170161774-90",
				},
			},
			Email:        "adresa@mail.com",
			BrojTelefona: "063/0000000",
		},
		8: &model.WingPravnoLice{
			Id:             0,
			Naziv:          "URBANS INVEST DOO NOVI SAD",
			DugiNaziv:      "URBANS INVEST DOO ZA GRAĐEVINARSTVO, UNUTRAŠNJU I SPOLJNU TRGOVINU NOVI SAD",
			MB:             "20422998",
			PIB:            "105625475",
			Adresa:         "Železnička 39/10",
			Grad:           "21000 Novi Sad",
			DatumOsnivanja: "20.5.2008.",
			Delatnost:      "Izgradnja stambenih i nestambenih zgrada",
			Racuni: []model.WingBankaRacun{
				model.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0007011600535-65",
				},
				model.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0007011534435-97",
				},
				model.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0007011534408-81",
				},
				model.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0007011534378-74",
				},
				model.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0002024305088-59",
				},
			},
			Email:        "adresa@mail.com",
			BrojTelefona: "063/0000000",
		},
		9: &model.WingPravnoLice{
			Id:             0,
			Naziv:          "DREAM HOME REAL ESTATE DOO NOVI SAD",
			DugiNaziv:      "DREAM HOME REAL ESTATE DOO NOVI SAD",
			MB:             "21278815",
			PIB:            "109979047",
			Adresa:         "TONE HADŽIĆA 12",
			Grad:           "21000 Novi Sad",
			DatumOsnivanja: "4.4.2017.",
			Delatnost:      "Izgradnja stambenih i nestambenih zgrada",
			Racuni: []model.WingBankaRacun{
				model.WingBankaRacun{
					Banka: "Raiffeisen banka a.d. Beograd",
					Racun: "265-1000000186611-05",
				},
				model.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0007009210509-97",
				},
				model.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0007009210479-90",
				},
				model.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0007009210444-98",
				},
				model.WingBankaRacun{
					Banka: "Addiko Bank a.d. Beograd",
					Racun: "165-0007009210417-82",
				},
			},
			Email:        "adresa@mail.com",
			BrojTelefona: "063/0000000",
		},
	}
}
