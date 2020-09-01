package pdf

import (
	"time"
)

func (p *WingPrint) pdfHeader() func() {
	return func() {
		currentTime := time.Now()
		//pdf.Image("/usr/home/marcetin/Public/wingcal/NOVOGUI/pdfheader.png", 5, 5, 200, 25, false, "", 0, "")
		//pdf.SetDrawColor(200,200,200)
		p.pdf.SetFillColor(200, 200, 200)
		p.pdf.Rect(5, 5, 200, 20, "F")
		p.pdf.SetY(5)
		p.pdf.SetFont("Arial", "", 8)
		p.pdf.CellFormat(47, 6, "MB:20701005", "0", 0, "L", false, 0, "")
		p.pdf.SetFont("Arial", "B", 10)
		p.pdf.CellFormat(47, 10, p.projekat.Projektant.Naziv, "0", 0, "R", false, 0, "")
		p.pdf.SetFont("Arial", "", 8)
		p.pdf.CellFormat(47, 8, "     SIFRA PROJEKTA", "0", 0, "L", false, 0, "")
		p.pdf.SetFont("Arial", "", 8)
		p.pdf.CellFormat(47, 8, "fhe38833", "0", 0, "R", false, 0, "")
		p.pdf.Ln(5)
		p.pdf.SetFont("Arial", "", 8)
		p.pdf.CellFormat(47, 6, "PIB:106892584", "0", 0, "L", false, 0, "")
		p.pdf.CellFormat(47, 8, "Bulevar Oslobodjenja 30A", "0", 0, "R", false, 0, "")
		p.pdf.SetFont("Arial", "", 8)
		p.pdf.CellFormat(47, 8, "     NAZIV PROJEKTA", "0", 0, "L", false, 0, "")
		p.pdf.SetFont("Arial", "", 8)
		p.pdf.CellFormat(47, 8, "projekat za evidentiranje", "0", 0, "R", false, 0, "")
		p.pdf.Ln(5)
		p.pdf.SetFont("Arial", "", 8)
		p.pdf.CellFormat(47, 6, "tel:069/222-44-33", "0", 0, "L", false, 0, "")
		p.pdf.CellFormat(47, 8, "21000 Novi Sad", "0", 0, "R", false, 0, "")
		p.pdf.SetFont("Arial", "", 8)
		p.pdf.CellFormat(47, 8, "     DATUM PROJEKTA", "0", 0, "L", false, 0, "")
		p.pdf.SetFont("Arial", "", 8)
		p.pdf.CellFormat(47, 10, currentTime.Format("06-Jan-02"), "0", 0, "R", false, 0, "")
		p.pdf.Ln(5)
		p.pdf.SetFont("Arial", "", 8)
		p.pdf.CellFormat(47, 6, "email:vukobrat.cedomir@gmail.com", "0", 0, "L", false, 0, "")
	}
}
