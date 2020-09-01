package pdf

import (
	"fmt"
)

func (p *WingPrint) aktivnostiSuma() {
	p.pdf.SetFont("Times", "B", 16)
	p.pdf.CellFormat(0, 10, p.tr("Lista aktivnosti"), "0", 0, "", false, 0, "")
	p.pdf.Ln(20)
	p.pdf.SetFont("Arial", "", 10)
	for _, e := range p.projekat.Suma.Elementi {
		cols := []float64{40, p.d.pagew - p.d.mleft - p.d.mright - 20}
		rows := [][]string{
			[]string{
				e.Sifra, e.Element.Struct["Title"].Content.(string),
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
				if i < 1 {
					p.pdf.SetFont("Arial", "B", 10)
				} else {
					p.pdf.SetFont("Arial", "", 10)
				}
				//fmt.Println("Col::", i)

				p.pdf.MultiCell(width, lineHt+p.d.marginCell, p.tr(txt), "", "", false)
				x += width
				p.pdf.SetXY(x, y)
			}
			p.pdf.SetXY(curx, y+height)
		}
	}

	p.pdf.SetFont("Times", "B", 16)
	p.pdf.CellFormat(0, 10, "Suma: "+fmt.Sprintf("%.2f", p.projekat.Suma.SumaCena), "0", 0, "", false, 0, "")
}
