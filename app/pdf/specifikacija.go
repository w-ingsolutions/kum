package pdf

import (
	"fmt"
)

func (p *WingPrint) specifikacijaMaterijalaList() {
	p.pdf.AddPage()
	p.pdf.SetFont("Times", "B", 16)
	p.pdf.CellFormat(0, 10, "Specifikacija materijala", "0", 0, "", false, 0, "")
	p.s.materijal = p.pdf.PageNo()
	p.pdf.Ln(20)

	p.pdf.SetFont("Arial", "", 10)
	for _, e := range p.projekat.Suma.NeophodanMaterijal {
		cols := []float64{40, p.d.pagew - p.d.mleft - p.d.mright - 20}
		//rows := [][]string{}

		rows := [][]string{
			[]string{
				"Šifra", fmt.Sprint(e.Id),
			},
			[]string{
				"Naziv", e.Materijal.Struct["Title"].Content.(string),
			},
			//[]string{
			//	"Osobine i namena", e.Materijal.OsobineNamena,
			//},
			[]string{
				"Jedinica mere", e.Materijal.Struct["JedinicaPotrosnje"].Content.(string),
			},
			[]string{
				"Jedinična cena", e.Materijal.Struct["Cena"].Content.(string),
			},
			[]string{
				"Količina", fmt.Sprint(e.Kolicina),
			},
			[]string{
				"Vrednost materijala", fmt.Sprintf("%.2f", e.UkupnaCena),
			},
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
		p.pdf.Ln(8)
	}

	p.pdf.SetFont("Times", "B", 16)
	p.pdf.CellFormat(0, 10, "Suma materijal: "+fmt.Sprintf("%.2f", p.projekat.Suma.SumaCenaMaterijal), "0", 0, "", false, 0, "")
	p.pdf.Ln(20)

}
