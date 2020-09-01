package pdf

func (p *WingPrint) investitorList() {
	p.pdf.SetFont("Times", "B", 16)
	p.pdf.CellFormat(0, 10, "Investitor", "0", 0, "", false, 0, "")
	p.pdf.Ln(10)
	p.pdf.SetFont("Arial", "", 10)
	cols := []float64{40, p.d.pagew - p.d.mleft - p.d.mright - 20}
	rows := [][]string{
		[]string{
			"MB", p.projekat.Investitor.MB,
		},
		[]string{
			"PIB", p.projekat.Investitor.PIB,
		},
		[]string{
			"Kratak Naziv", p.projekat.Investitor.Naziv,
		},
		[]string{
			"Dugi Naziv", p.projekat.Investitor.DugiNaziv,
		},
		[]string{
			"Delatnost", p.projekat.Investitor.Delatnost,
		},
		[]string{
			"Adresa", p.projekat.Investitor.Adresa,
		},
		[]string{
			"Grad", p.projekat.Investitor.Grad,
		},
		[]string{
			"Email", p.projekat.Investitor.Email,
		},
		[]string{
			"Broj telefona", p.projekat.Investitor.BrojTelefona,
		},
		[]string{
			"Datum Osnivanja", p.projekat.Investitor.DatumOsnivanja,
		},
		//[]string{
		//	"Racuni", p.projekat.Investitor.Racuni,
		//},
	}
	for _, row := range rows {
		curx, y := p.pdf.GetXY()
		x := curx
		height := 0.
		_, lineHt := p.pdf.GetFontSize()
		for i, txt := range row {
			lines := p.pdf.SplitLines([]byte(txt), cols[i])
			h := float64(len(lines))*lineHt + p.d.marginCell*float64(len(lines))
			if h > height {
				height = h
			}
		}
		// add a new page if the height of the row doesn't fit on the page
		if p.pdf.GetY()+height > p.d.pageh-p.d.mbottom {
			p.pdf.AddPage()
			y = p.pdf.GetY()
		}
		for i, txt := range row {
			width := cols[i]
			//pdf.Rect(x, y, width, height, "")
			p.pdf.MultiCell(width, lineHt+p.d.marginCell, p.tr(txt), "", "", false)
			x += width
			p.pdf.SetXY(x, y)
		}
		p.pdf.SetXY(curx, y+height)
	}
}
