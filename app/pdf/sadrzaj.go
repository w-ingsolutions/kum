package pdf

import (
	"fmt"
)

func (p *WingPrint) sadrzajList() {
	p.pdf.AddPage()
	p.pdf.SetFont("Times", "B", 16)
	p.pdf.CellFormat(0, 10, "Sadržaj", "0", 0, "", false, 0, "")
	p.pdf.Ln(30)
	p.pdf.SetFont("Arial", "", 12)
	p.pdf.CellFormat(0, 10, fmt.Sprintf("Ponuda %d", p.s.ponuda), "0", 0, "", false, 0, "")
	p.pdf.Ln(10)
	p.pdf.CellFormat(0, 10, fmt.Sprintf("Ugovor %d", p.s.ugovor), "0", 0, "", false, 0, "")
	p.pdf.Ln(10)
	p.pdf.CellFormat(0, 10, fmt.Sprintf("Specifikacija radova %d", p.s.aktivnosti), "0", 0, "", false, 0, "")
	p.pdf.Ln(10)
	p.pdf.CellFormat(0, 10, fmt.Sprintf("Specifikacija materijala %d", p.s.materijal), "0", 0, "", false, 0, "")
	p.pdf.Ln(10)
	p.pdf.CellFormat(0, 10, fmt.Sprintf("Tehnički list %d", p.s.tehnicki), "0", 0, "", false, 0, "")
	p.pdf.Ln(10)
	p.pdf.CellFormat(0, 10, fmt.Sprintf("Standardi %d", p.s.standardi), "0", 0, "", false, 0, "")
	p.pdf.Ln(10)
	p.pdf.CellFormat(0, 10, fmt.Sprintf("Merenja %d", p.s.merenja), "0", 0, "", false, 0, "")
	p.pdf.Ln(10)
	p.pdf.CellFormat(0, 10, fmt.Sprintf("Uslovi %d", p.s.uslovi), "0", 0, "", false, 0, "")
	p.pdf.Ln(10)
}
