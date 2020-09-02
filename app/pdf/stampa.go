package pdf

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/w-ingsolutions/c/pkg/pdf"
	"github.com/w-ingsolutions/kum/app/mod"
)

type WingPrint struct {
	naziv    string
	projekat *mod.WingProjekat
	pdf      *gofpdf.Fpdf
	d        dimenzije
	s        sekcije
	tr       func(string) string
}
type dimenzije struct {
	pagew, mleft, mright, marginCell, pageh, mbottom float64
}
type sekcije struct {
	materijal, aktivnosti, tehnicki, ponuda, ugovor, standardi, merenja, uslovi int
}

func KreiranjeNalogaPDF(pr *mod.WingProjekat, nz string) {
	p := WingPrint{
		projekat: pr,
		naziv:    nz,
		pdf:      pdf.P(),
		d:        dimenzije{},
		s:        sekcije{},
	}
	p.tr = p.pdf.UnicodeTranslatorFromDescriptor("")
	p.pdf.SetTopMargin(30)
	p.d.marginCell = 2. // margin of top/bottom of cell
	p.d.pagew, p.d.pageh = p.pdf.GetPageSize()
	p.d.mleft, p.d.mright, _, p.d.mbottom = p.pdf.GetMargins()
	p.pdf.SetHeaderFuncMode(p.pdfHeader(), true)
	p.pdf.SetFooterFunc(p.pdfFooter())
	p.pdf.AliasNbPages("")

	p.ponuda()
	p.ipList()

	p.specifikacijaRadovaList()
	p.specifikacijaMaterijalaList()
	p.tehnickiList()
	//p.novaStrana()
	p.sadrzajList()
	//err := p.pdf.OutputFileAndClose(w.Podesavanja.Dir + "/nalog.pdf")
	err := p.pdf.OutputFileAndClose(p.naziv)
	if err != nil {
	}
}

//func aktivnostiSuma(p *gofpdf.Fpdf, pagew, mleft, mright, marginCell, pageh, mbottom float64, tr func(string) string) {
//	p.pdf.SetFont("Times", "B", 16)
//	p.pdf.CellFormat(0, 10, tr(w.text("Lista aktivnosti")), "0", 0, "", false, 0, "")
//	p.pdf.Ln(20)
//	p.pdf.SetFont("Arial", "", 10)
//	for _, e := range w.Suma.Elementi {
//		cols := []float64{40, pagew - mleft - mright - 20}
//		rows := [][]string{
//			[]string{
//				e.Sifra, e.Element.Struct["Naziv"].Content.(string),
//			},
//		}
//		for _, row := range rows {
//			curx, y := p.pdf.GetXY()
//			x := curx
//			height := 0.
//			_, lineHt := p.pdf.GetFontSize()
//			for i, txt := range row {
//				lines := p.pdf.SplitLines([]byte(txt), cols[i])
//				h := float64(len(lines))*lineHt + marginCell*float64(len(lines))
//				if h > height {
//					height = h
//				}
//			}
//			// add a new page if the height of the row doesn't fit on the page
//			if p.pdf.GetY()+height > pageh-mbottom {
//				p.pdf.AddPage()
//				y = p.pdf.GetY()
//			}
//			for i, txt := range row {
//				width := cols[i]
//				//pdf.Rect(x, y, width, height, "")
//				if i < 1 {
//					p.pdf.SetFont("Arial", "B", 10)
//				} else {
//					p.pdf.SetFont("Arial", "", 10)
//				}
//				//fmt.Println("Col::", i)
//
//				p.pdf.MultiCell(width, lineHt+marginCell, tr(txt), "", "", false)
//				x += width
//				p.pdf.SetXY(x, y)
//			}
//			p.pdf.SetXY(curx, y+height)
//		}
//	}
//
//	p.pdf.SetFont("Times", "B", 16)
//	p.pdf.CellFormat(0, 10, w.text("Suma: ")+fmt.Sprintf("%.2f", w.Suma.SumaCena), "0", 0, "", false, 0, "")
//}
//
//func specifikacijaMaterijalaList(p *gofpdf.Fpdf, pagew, mleft, mright, marginCell, pageh, mbottom float64, tr func(string) string) {
//	p.pdf.AddPage()
//	p.pdf.SetFont("Times", "B", 16)
//	p.pdf.CellFormat(0, 10, w.text("Specifikacija materijala"), "0", 0, "", false, 0, "")
//	materijal = p.pdf.PageNo()
//	p.pdf.Ln(20)
//
//	p.pdf.SetFont("Arial", "", 10)
//	for _, e := range w.Suma.NeophodanMaterijal {
//		cols := []float64{40, pagew - mleft - mright - 20}
//		//rows := [][]string{}
//
//		rows := [][]string{
//			[]string{
//				"Šifra", fmt.Sprint(e.Id),
//			},
//			[]string{
//				"Naziv", e.Materijal.Struct["Naziv"].Content.(string),
//			},
//			//[]string{
//			//	"Osobine i namena", e.Materijal.OsobineNamena,
//			//},
//			[]string{
//				"Jedinica mere", e.Materijal.Struct["JedinicaPotrosnje"].Content.(string),
//			},
//			[]string{
//				"Jedinična cena", e.Materijal.Struct["Cena"].Content.(string),
//			},
//			[]string{
//				"Količina", fmt.Sprint(e.Kolicina),
//			},
//			[]string{
//				"Vrednost materijala", fmt.Sprintf("%.2f", e.UkupnaCena),
//			},
//		}
//		for _, row := range rows {
//			curx, y := p.pdf.GetXY()
//			x := curx
//			height := 0.
//			_, lineHt := p.pdf.GetFontSize()
//			for i, txt := range row {
//				lines := p.pdf.SplitLines([]byte(txt), cols[i])
//				h := float64(len(lines))*lineHt + marginCell*float64(len(lines))
//				if h > height {
//					height = h
//				}
//			}
//			// add a new page if the height of the row doesn't fit on the page
//			if p.pdf.GetY()+height > pageh-mbottom {
//				p.pdf.AddPage()
//				y = p.pdf.GetY()
//			}
//			for i, txt := range row {
//				width := cols[i]
//				//pdf.Rect(x, y, width, height, "")
//				p.pdf.MultiCell(width, lineHt+marginCell, tr(txt), "", "", false)
//				x += width
//				p.pdf.SetXY(x, y)
//			}
//			p.pdf.SetXY(curx, y+height)
//		}
//		p.pdf.Ln(8)
//	}
//
//	p.pdf.SetFont("Times", "B", 16)
//	p.pdf.CellFormat(0, 10, w.text("Suma materijal: ")+fmt.Sprintf("%.2f", p.projekat.Elementi.SumaCenaMaterijal), "0", 0, "", false, 0, "")
//	p.pdf.Ln(20)
//
//}
//
//func materijalSuma(p *gofpdf.Fpdf, pagew, mleft, mright, marginCell, pageh, mbottom float64, tr func(string) string) {
//	p.pdf.SetFont("Times", "B", 16)
//	p.pdf.CellFormat(0, 10, w.text("Lista materijala"), "0", 0, "", false, 0, "")
//	p.pdf.Ln(20)
//	p.pdf.SetFont("Arial", "", 10)
//	for _, e := range w.Suma.NeophodanMaterijal {
//		cols := []float64{40, pagew - mleft - mright - 20}
//		//rows := [][]string{}
//		rows := [][]string{
//			[]string{
//				fmt.Sprint(e.Id), e.Materijal.Struct["Naziv"].Content.(string),
//			},
//		}
//		for _, row := range rows {
//			curx, y := p.pdf.GetXY()
//			x := curx
//			height := 0.
//			_, lineHt := p.pdf.GetFontSize()
//			for i, txt := range row {
//				lines := p.pdf.SplitLines([]byte(txt), cols[i])
//				h := float64(len(lines))*lineHt + marginCell*float64(len(lines))
//				if h > height {
//					height = h
//				}
//			}
//			// add a new page if the height of the row doesn't fit on the page
//			if p.pdf.GetY()+height > pageh-mbottom {
//				p.pdf.AddPage()
//				y = p.pdf.GetY()
//			}
//			for i, txt := range row {
//				width := cols[i]
//				//pdf.Rect(x, y, width, height, "")
//				p.pdf.MultiCell(width, lineHt+marginCell, tr(txt), "", "", false)
//				x += width
//				p.pdf.SetXY(x, y)
//			}
//			p.pdf.SetXY(curx, y+height)
//		}
//	}
//	p.pdf.SetFont("Times", "B", 16)
//	p.pdf.CellFormat(0, 10, w.text("Suma materijal: ")+fmt.Sprintf("%.2f", p.projekat.Elementi.SumaCenaMaterijal), "0", 0, "", false, 0, "")
//}
//
//func sadrzajList(p *gofpdf.Fpdf, pagew, mleft, mright, marginCell, pageh, mbottom float64) {
//	p.pdf.AddPage()
//	p.pdf.SetFont("Times", "B", 16)
//	p.pdf.CellFormat(0, 10, w.text("Sadržaj"), "0", 0, "", false, 0, "")
//	p.pdf.Ln(30)
//	p.pdf.SetFont("Arial", "", 12)
//	p.pdf.CellFormat(0, 10, fmt.Sprintf("Ponuda %d", ponuda), "0", 0, "", false, 0, "")
//	p.pdf.Ln(10)
//	p.pdf.CellFormat(0, 10, fmt.Sprintf("Ugovor %d", ugovor), "0", 0, "", false, 0, "")
//	p.pdf.Ln(10)
//	p.pdf.CellFormat(0, 10, fmt.Sprintf("Specifikacija radova %d", aktivnosti), "0", 0, "", false, 0, "")
//	p.pdf.Ln(10)
//	p.pdf.CellFormat(0, 10, fmt.Sprintf("Specifikacija materijala %d", materijal), "0", 0, "", false, 0, "")
//	p.pdf.Ln(10)
//	p.pdf.CellFormat(0, 10, fmt.Sprintf("Tehnički list %d", tehnicki), "0", 0, "", false, 0, "")
//	p.pdf.Ln(10)
//	p.pdf.CellFormat(0, 10, fmt.Sprintf("Standardi %d", standardi), "0", 0, "", false, 0, "")
//	p.pdf.Ln(10)
//	p.pdf.CellFormat(0, 10, fmt.Sprintf("Merenja %d", merenja), "0", 0, "", false, 0, "")
//	p.pdf.Ln(10)
//	p.pdf.CellFormat(0, 10, fmt.Sprintf("Uslovi %d", uslovi), "0", 0, "", false, 0, "")
//	p.pdf.Ln(10)
//}
//

//func ipList(p *gofpdf.Fpdf, pagew, mleft, mright, marginCell, pageh, mbottom float64, tr func(string) string) {
//	p.pdf.AddPage()
//	ugovor = p.pdf.PageNo()
//
//	w.projektantList()
//	p.pdf.Ln(10)
//
//	p.pdf.SetFont("Arial", "", 8)
//	_, lineHt := p.pdf.GetFontSize()
//	linesA := p.pdf.SplitLines([]byte("Na osnovu člana 128a. Zakona o planiranju i izgradnji objekata (Sl. glasnik Republike Srbije br.72/09, 81/09 – ispravka, 64/10 odluka US, 24/11 i 121/12, 42/13 – odluka US, 50/2013 – odluka US, 98/2013 - odluka US, 132/14 i 145/14, 83/18, 31/19 i 37/19) i odredbi Pravilnika o sadržini, načinu i postupku izrade i način vršenja kontrole tehničke dokumentacije prema klasi i nameni objekta (Sl. glasnik Republike Srbije br.72/2018)"), 200)
//	for _, line := range linesA {
//		p.pdf.CellFormat(190.0, lineHt, string(line), "", 1, "J", false, 0, "")
//	}
//	w.investitorList()
//}
//
//func investitorList(p *gofpdf.Fpdf, pagew, mleft, mright, marginCell, pageh, mbottom float64, tr func(string) string) {
//	p.pdf.SetFont("Times", "B", 16)
//	p.pdf.CellFormat(0, 10, w.text("Investitor"), "0", 0, "", false, 0, "")
//	p.pdf.Ln(10)
//	p.pdf.SetFont("Arial", "", 10)
//	cols := []float64{40, pagew - mleft - mright - 20}
//	rows := [][]string{
//		[]string{
//			"MB", p.projekat.Investitor.MB,
//		},
//		[]string{
//			"PIB", p.projekat.Investitor.PIB,
//		},
//		[]string{
//			"Kratak Naziv", p.projekat.Investitor.Naziv,
//		},
//		[]string{
//			"Dugi Naziv", p.projekat.Investitor.DugiNaziv,
//		},
//		[]string{
//			"Delatnost", p.projekat.Investitor.Delatnost,
//		},
//		[]string{
//			"Adresa", p.projekat.Investitor.Adresa,
//		},
//		[]string{
//			"Grad", p.projekat.Investitor.Grad,
//		},
//		[]string{
//			"Email", p.projekat.Investitor.Email,
//		},
//		[]string{
//			"Broj telefona", p.projekat.Investitor.BrojTelefona,
//		},
//		[]string{
//			"Datum Osnivanja", p.projekat.Investitor.DatumOsnivanja,
//		},
//		//[]string{
//		//	"Racuni", p.projekat.Investitor.Racuni,
//		//},
//	}
//	for _, row := range rows {
//		curx, y := p.pdf.GetXY()
//		x := curx
//		height := 0.
//		_, lineHt := p.pdf.GetFontSize()
//		for i, txt := range row {
//			lines := p.pdf.SplitLines([]byte(txt), cols[i])
//			h := float64(len(lines))*lineHt + marginCell*float64(len(lines))
//			if h > height {
//				height = h
//			}
//		}
//		// add a new page if the height of the row doesn't fit on the page
//		if p.pdf.GetY()+height > pageh-mbottom {
//			p.pdf.AddPage()
//			y = p.pdf.GetY()
//		}
//		for i, txt := range row {
//			width := cols[i]
//			//pdf.Rect(x, y, width, height, "")
//			p.pdf.MultiCell(width, lineHt+marginCell, tr(txt), "", "", false)
//			x += width
//			p.pdf.SetXY(x, y)
//		}
//		p.pdf.SetXY(curx, y+height)
//	}
//}
//
//func projektantList(p *gofpdf.Fpdf, pagew, mleft, mright, marginCell, pageh, mbottom float64, tr func(string) string) {
//	p.pdf.SetFont("Times", "B", 16)
//	p.pdf.CellFormat(0, 10, w.text("Nadzor"), "0", 0, "", false, 0, "")
//	p.pdf.Ln(10)
//	p.pdf.SetFont("Arial", "", 10)
//	cols := []float64{40, pagew - mleft - mright - 20}
//	rows := [][]string{
//		[]string{
//			"MB", p.projekat.Projektant.MB,
//		},
//		[]string{
//			"PIB", p.projekat.Projektant.PIB,
//		},
//		[]string{
//			"Kratak Naziv", p.projekat.Projektant.Naziv,
//		},
//		[]string{
//			"DugiNaziv", p.projekat.Projektant.DugiNaziv,
//		},
//		[]string{
//			"Delatnost", p.projekat.Projektant.Delatnost,
//		},
//		[]string{
//			"Adresa", p.projekat.Projektant.Adresa,
//		},
//		[]string{
//			"Grad", p.projekat.Projektant.Grad,
//		},
//		[]string{
//			"Email", p.projekat.Projektant.Email,
//		},
//		[]string{
//			"Broj telefona", p.projekat.Projektant.BrojTelefona,
//		},
//		[]string{
//			"Datum Osnivanja", p.projekat.Projektant.DatumOsnivanja,
//		},
//	}
//	for _, row := range rows {
//		curx, y := p.pdf.GetXY()
//		x := curx
//		height := 0.
//		_, lineHt := p.pdf.GetFontSize()
//		for i, txt := range row {
//			lines := p.pdf.SplitLines([]byte(txt), cols[i])
//			h := float64(len(lines))*lineHt + marginCell*float64(len(lines))
//			if h > height {
//				height = h
//			}
//		}
//		// add a new page if the height of the row doesn't fit on the page
//		if p.pdf.GetY()+height > pageh-mbottom {
//			p.pdf.AddPage()
//			y = p.pdf.GetY()
//		}
//		for i, txt := range row {
//			width := cols[i]
//			//pdf.Rect(x, y, width, height, "")
//			p.pdf.MultiCell(width, lineHt+marginCell, tr(txt), "", "", false)
//			x += width
//			p.pdf.SetXY(x, y)
//		}
//		p.pdf.SetXY(curx, y+height)
//	}
//}
