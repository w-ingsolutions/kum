package pdf

func (p *WingPrint) ponuda() {
	p.pdf.AddPage()
	p.pdf.SetFont("Times", "B", 18)
	p.s.ponuda = p.pdf.PageNo()
	p.pdf.CellFormat(0, 10, "Ponuda", "0", 0, "", false, 0, "")
	p.pdf.Ln(10)
	p.investitorList()
	p.pdf.Ln(20)
	p.aktivnostiSuma()
	p.pdf.Ln(20)
	//p.materijalSuma()
}
func (p *WingPrint) ipList() {
	p.pdf.AddPage()
	p.s.ugovor = p.pdf.PageNo()

	p.projektantList()
	p.pdf.Ln(10)

	p.pdf.SetFont("Arial", "", 8)
	_, lineHt := p.pdf.GetFontSize()
	linesA := p.pdf.SplitLines([]byte("Na osnovu člana 128a. Zakona o planiranju i izgradnji objekata (Sl. glasnik Republike Srbije br.72/09, 81/09 – ispravka, 64/10 odluka US, 24/11 i 121/12, 42/13 – odluka US, 50/2013 – odluka US, 98/2013 - odluka US, 132/14 i 145/14, 83/18, 31/19 i 37/19) i odredbi Pravilnika o sadržini, načinu i postupku izrade i način vršenja kontrole tehničke dokumentacije prema klasi i nameni objekta (Sl. glasnik Republike Srbije br.72/2018)"), 200)
	for _, line := range linesA {
		p.pdf.CellFormat(190.0, lineHt, string(line), "", 1, "J", false, 0, "")
	}
	p.investitorList()
}
