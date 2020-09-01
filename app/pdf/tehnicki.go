package pdf

import (
	"fmt"
)

func (p *WingPrint) tehnickiList() {
	p.pdf.AddPage()
	p.pdf.SetFont("Times", "B", 16)
	p.s.tehnicki = p.pdf.PageNo()
	p.pdf.CellFormat(0, 10, "Tehnički list", "0", 0, "", false, 0, "")
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
			[]string{
				"Osobine i namena", e.Materijal.Struct["OsobineNamena"].Content.(string),
			},
			[]string{
				"Nacin rada", e.Materijal.Struct["NacinRada"].Content.(string),
			},

			[]string{
				"Jedinica mere", e.Materijal.Struct["JedinicaPotrosnje"].Content.(string),
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
}
