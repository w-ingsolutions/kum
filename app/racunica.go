package calc

import (
	"encoding/json"
	"fmt"
	"github.com/w-ingsolutions/kum/app/mod"
	"log"
)

func (w *WingCal) SumaRacunica() {
	//w.SumaElementiPrikaz()
	s := 0.0
	for _, e := range w.Suma.Elementi {
		s = s + e.SumaCena
	}
	w.Suma.SumaCena = s
	//w.Suma.SumaCenaMaterijal, w.Suma.NeophodanMaterijal = w.NeopodanMaterijal(w.Suma.Elementi)
	w.NeophodanMaterijal()
}

func (w *WingCal) PrikazaniElementSumaRacunica() func() {
	return func() {
		prikazaniElementSumaCena = w.PrikazaniElement.el.Struct["Cena"].Content.(float64) * float64(w.UI.Counters.Kolicina.Value)
	}
}

func (w *WingCal) ProjekatRacunica() func() {
	return func() {
		p := w.Suma
		iz := []*mod.WingIzabraniElement{}
		for _, e := range w.Suma.Elementi {
			ee := e
			ee.SumaCena = e.SumaCena * float64(w.UI.Counters.Radovi.Value) / 100
			iz = append(iz, e)
		}
		p.Elementi = iz
		projekat.Suma = p
		//w.ProjekatSumaRacunica()
	}
}

func (w *WingCal) ProjekatMaterijalSumaRacunica() func() {
	return func() {
		//projekat.Elementi.SumaCena, projekat.Elementi.NeophodanMaterijal = w.NeopodanMaterijal(projekat.Elementi.Elementi)
		//projekat.Elementi.SumaCena = w.Suma.SumaCenaMaterijal + w.Suma.SumaCenaMaterijal*float64(materijal.Value)/100
	}
}

func (w *WingCal) ProjekatSumaRacunica() func() {
	return func() {
		s := 0.0
		for _, e := range projekat.Suma.Elementi {
			s = s + e.SumaCena
		}
		projekat.Suma.SumaCena = s
	}
}

func (w *WingCal) NeophodanMaterijal() {
	ukupanNeophodniMaterijal := make(map[int]mod.WingNeophodanMaterijal)
	unm := make(map[int]mod.WingNeophodanMaterijal)
	sumaCena := 0.0
	for _, e := range w.Suma.Elementi {
		//ukupnaCenaMaterijala := 0.0
		var neophodanMaterijal []map[string]int
		err := json.Unmarshal([]byte(w.PrikazaniElement.el.Struct["NeophodanMaterijal"].Content.(string)), &neophodanMaterijal)
		if err != nil {
			log.Println(err)
		}

		for _, pojedinacniMaterijalSume := range neophodanMaterijal {
			id := pojedinacniMaterijalSume["Id"]
			materijal := mod.WingNeophodanMaterijal{
				Id:        id,
				Materijal: w.Materijal[id-1],
			}
			//k := materijal.Materijal.Struct["Pakovanje"].Content.(float64) * float64(e.Kolicina) * float64(pojedinacniMaterijalSume["Koeficijent"])
			//materijal.Kolicina = ukupanNeophodniMaterijal[id].Kolicina + k
			//ukupnaCena := materijal.Kolicina * materijal.Materijal.Struct["Cena"].Content.(float64)
			//materijal.UkupnaCena = ukupnaCena
			//materijal.UkupnoPakovanja = int(k / materijal.Materijal.Struct["Pakovanje"].Content.(float64))
			ukupanNeophodniMaterijal[id] = materijal
			//ukupnaCenaMaterijala = ukupnaCenaMaterijala + ukupnaCena
		}
		fmt.Println("eeeeeeeeee::", e)

	}
	for _, m := range ukupanNeophodniMaterijal {
		sumaCena = sumaCena + m.UkupnaCena
	}

	fmt.Println("ukupanNeophodniMaterijal::", ukupanNeophodniMaterijal)

	w.Suma.NeophodanMaterijal = ukupanNeophodniMaterijal
	w.Suma.SumaCenaMaterijal = sumaCena
	i := 0
	for _, uuu := range ukupanNeophodniMaterijal {
		unm[i] = uuu
		i++
	}
	w.Suma.NeophodanMaterijalPrikaz = unm
}
