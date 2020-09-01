package pdf

import (
	"fmt"
)

func (p *WingPrint) pdfFooter() func() {
	return func() {
		p.pdf.SetY(-30)
		p.pdf.SetFont("Arial", "I", 8)
		p.pdf.CellFormat(0, 10, fmt.Sprintf("Strana %d/{nb}", p.pdf.PageNo()), "", 0, "C", false, 0, "")
		p.pdf.Ln(10)
		p.pdf.SetFillColor(200, 200, 200)
		p.pdf.Rect(5, 275, 200, 20, "F")
		p.pdf.SetY(-22)
		p.pdf.SetFont("Arial", "", 8)
		p.pdf.CellFormat(47, 6, "MB:"+p.projekat.Investitor.MB, "0", 0, "L", false, 0, "")
		p.pdf.SetFont("Arial", "B", 10)
		p.pdf.CellFormat(47, 10, p.projekat.Investitor.Naziv, "0", 0, "R", false, 0, "")
		p.pdf.SetFont("Arial", "", 8)
		p.pdf.CellFormat(47, 8, "     SIFRA DOKUMENTA", "0", 0, "L", false, 0, "")
		p.pdf.SetFont("Arial", "", 8)
		p.pdf.CellFormat(47, 8, "fhe38833", "0", 0, "R", false, 0, "")
		p.pdf.Ln(5)
		p.pdf.SetFont("Arial", "", 8)
		p.pdf.CellFormat(47, 6, "PIB:"+p.projekat.Investitor.PIB, "0", 0, "L", false, 0, "")
		p.pdf.CellFormat(47, 8, p.projekat.Investitor.Adresa, "0", 0, "R", false, 0, "")
		p.pdf.SetFont("Arial", "", 8)
		p.pdf.CellFormat(47, 8, "     NAZIV DOKUMENTA", "0", 0, "L", false, 0, "")
		p.pdf.SetFont("Arial", "", 8)
		p.pdf.CellFormat(47, 8, "dokument za evidentiranje", "0", 0, "R", false, 0, "")
		p.pdf.Ln(5)
		p.pdf.SetFont("Arial", "", 8)
		p.pdf.CellFormat(47, 6, "tel:069/222-44-33", "0", 0, "L", false, 0, "")
		p.pdf.CellFormat(47, 8, "21000 Novi Sad", "0", 0, "R", false, 0, "")
		p.pdf.SetFont("Arial", "", 8)
		p.pdf.CellFormat(47, 8, "     DATUM DOKUMENTA", "0", 0, "L", false, 0, "")
		p.pdf.SetFont("Arial", "", 8)
		p.pdf.CellFormat(47, 10, "mart 2020", "0", 0, "R", false, 0, "")
		p.pdf.Ln(5)
		p.pdf.SetFont("Arial", "", 8)
		p.pdf.CellFormat(47, 6, "email:vukobrat.cedomir@gmail.com", "0", 0, "L", false, 0, "")

	}
}
