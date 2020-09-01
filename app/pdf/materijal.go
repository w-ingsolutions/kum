package pdf

import (
	"fmt"
)

func (p *WingPrint) materijalSuma() {
	p.pdf.SetFont("Times", "B", 16)
	p.pdf.CellFormat(0, 10, "Lista materijala", "0", 0, "", false, 0, "")
	p.pdf.Ln(20)
	p.pdf.SetFont("Arial", "", 10)
	for _, e := range p.projekat.Suma.NeophodanMaterijal {
		cols := []float64{40, p.d.pagew - p.d.mleft - p.d.mright - 20}
		//rows := [][]string{}
		rows := [][]string{
			[]string{
				fmt.Sprint(e.Id), e.Materijal.Struct["Title"].Content.(string),
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
	}
	p.pdf.SetFont("Times", "B", 16)
	p.pdf.CellFormat(0, 10, "Suma materijal: "+fmt.Sprintf("%.2f", p.projekat.Suma.SumaCenaMaterijal), "0", 0, "", false, 0, "")
}

//func ipList(p *gofpdf.Fpdf, pagew, mleft, mright, marginCell, pageh, mbottom float64, tr func(string) string) {
//	p.pdf.AddPage()
//	ugovor = p.pdf.PageNo()
//
//	w.projektantList(p, pagew, mleft, mright, marginCell, pageh, mbottom, tr)
//	p.pdf.Ln(10)
//
//	p.pdf.SetFont("Arial", "", 8)
//	_, lineHt := p.pdf.GetFontSize()
//	linesA := p.pdf.SplitLines([]byte("Na osnovu člana 128a. Zakona o planiranju i izgradnji objekata (Sl. glasnik Republike Srbije br.72/09, 81/09 – ispravka, 64/10 odluka US, 24/11 i 121/12, 42/13 – odluka US, 50/2013 – odluka US, 98/2013 - odluka US, 132/14 i 145/14, 83/18, 31/19 i 37/19) i odredbi Pravilnika o sadržini, načinu i postupku izrade i način vršenja kontrole tehničke dokumentacije prema klasi i nameni objekta (Sl. glasnik Republike Srbije br.72/2018)"), 200)
//	for _, line := range linesA {
//		p.pdf.CellFormat(190.0, lineHt, string(line), "", 1, "J", false, 0, "")
//	}
//	w.investitorList(p, pagew, mleft, mright, marginCell, pageh, mbottom, tr)
//}
