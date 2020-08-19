package calc

import (
	"github.com/jung-kurt/gofpdf"
)

func (w *WingCal) novaStrana(p *gofpdf.Fpdf, pagew, mleft, mright, marginCell, pageh, mbottom float64, tr func(string) string) {
	p.AddPage()

	p.SetFont("Arial", "", 12)
	standardi = p.PageNo()

	_, lineHt := p.GetFontSize()

	p.SetFont("Times", "B", 16)
	p.CellFormat(0, 10, w.text("MOLERSKO-FARBARSKI RADOVI"), "0", 0, "", false, 0, "")
	p.Ln(10)
	p.SetFont("Arial", "", 10)
	p.CellFormat(0, 10, w.text("O p š t i  o p i s"), "0", 0, "", false, 0, "")

	p.Ln(10)

	p.SetFont("Times", "B", 16)
	p.CellFormat(0, 10, w.text("standardi"), "0", 0, "", false, 0, "")
	p.Ln(10)
	p.SetFont("Arial", "", 12)
	p.CellFormat(0, 10, w.text("1.Relevantni standardi"), "0", 0, "", false, 0, "")

	p.Ln(30)

	standardi := [][][]byte{
		p.SplitLines([]byte("SRPS U.F2.013 (1978) Završni radovi u građevinarstvu - Tehnički uslovi za izvođenje molerskih radova"), 200),
		p.SplitLines([]byte("SRPS EN 13300:2009 Boje I lakovi - Vodorastvorni materijali za prevlačenje i sistemi prevlaka za unutrašnje zidove i plafone - Klasifikacija"), 200),
		p.SplitLines([]byte("SRPS EN ISO 11998:2010 Boje i lakovi - Određivanje otpornosti prema 'vlažnom ribanju' i čišćenju prevlaka"), 200),
		p.SplitLines([]byte("SRPS EN ISO 2814:2010 Boje i lakovi - Upoređivanje odnosa kontrasta (pokrivne moći) boja istog tipa i boje"), 200),
		p.SplitLines([]byte("SRPS EN ISO 3668:2006 Boje i lakovi - Vizuelno poređenje boje boja"), 200),
		p.SplitLines([]byte("SRPS H.C8.054:1975 Boje i lakovi - Određivanje pokrivne moći (metoda šahovskog polja)"), 200),
	}
	for _, standard := range standardi {
		for _, s := range standard {
			p.CellFormat(190.0, lineHt, string(s), "", 1, "J", false, 0, "")
		}
		p.Ln(10)
	}

	p.AddPage()

	merenja = p.PageNo()
	p.SetFont("Times", "B", 16)
	p.CellFormat(0, 10, w.text("MERENJA"), "0", 0, "", false, 0, "")
	p.Ln(10)
	p.SetFont("Arial", "", 10)
	p.CellFormat(0, 10, w.text("Način vršenja kontrole i merenja"), "0", 0, "", false, 0, "")
	p.Ln(30)
	p.SetFont("Arial", "", 10)
	merenja := [][][]byte{
		p.SplitLines([]byte("Izvršeni radovi se obračunavaju po 1m2 površine ili po komadu, mere se uzimaju na licu mesta."), 200),
		p.SplitLines([]byte("Za radove koji se obračunavaju po komadu sa naznačenim dimenzijama, odstupanje do +/-5cm od jedne mere navedene u opisu ne uzima se u obzir. Za veća odstupanja (makar i jedne dimenzije) tolerancija se ne uzima u obzir i cena komada se menja procentualno u odnosu na promenu obrađene površine."), 200),
		p.SplitLines([]byte("Radovi u prostorijama sa stepenastim ili kosim podom, sa neravnim kosim plafonom ili svodom (prostori stepeništa, dvorane i slično) i radovi u prostorijama višim od 4m obračunavaju se posebno."), 200),
		p.SplitLines([]byte("Bočne površine podvlaka, greda i površine zidnih pojasa i ispada dodaju se kvadraturi plafona ako su obrađene u istoj tehnici. Kosi plafoni, podgledi stepenišnih krakova i slično - po m2 prema stvarnoj površini."), 200),
		p.SplitLines([]byte("Visina zidova se meri od poda ili gornje ivice podnožja do gornje granice zida. Ako je podnožje izrađeno od drugog materijala ili izvedeno u drugoj tehnici bojenja, visini zidne površine iznad podnožja dodaje se 20% visine podnožja. Ukoliko je visina podnožja koje je izvedeno od drugog materijala manja od 25cm, onda se visina zida meri od poda do gornje granice zida."), 200),
		p.SplitLines([]byte("Svodovi se obračunavaju po m2 i to se za visine temena svoda jednake 1/10 raspona svoda meri svetla vodoravna površina između zidova a za visine temena svoda iznad 1/10 raspona, meri se kao prostorija sa ravnim plafonom, s tim što se za obračun zidova uzima visina temena svoda."), 200),
		p.SplitLines([]byte("Ispadi i udubljenja (špalete, niše i slično) obračunavaju se posebno, iako su obrađeni u istoj tehnici kao plafon i zidovi."), 200),
		p.SplitLines([]byte("Zidovi stepenišnih i sličnih prostorija mere se u celoj visini od najnižeg nivoa poda do najviše granice plafona (zida) ako podnožje zida nije više od 25cm. Ukoliko se u takvom prostoru na zidovima nalazi podnožje obrađeno od drugog materijala, visine veće od 25cm, od ukupne visine zida odbija se zbir visina podnožja, umanjen za 20%."), 200),
		p.SplitLines([]byte("Dekorativna obrada ili ukrasno slikanje plafona, svodova i zidova prostorija obračunava se paušalno, po komadu ili po jednom poslu."), 200),
		p.SplitLines([]byte("Otvori veličine do 3m2 ne odbijaju se od izmerene površine. Za veće otvore odbija se razlika veća od 3m2. Kao otvori smatraju se prozori, vrata, ugrađeni ormani, pregrade i sl"), 200),
	}
	for _, merenje := range merenja {
		for _, m := range merenje {
			p.CellFormat(190.0, lineHt, string(m), "", 1, "J", false, 0, "")
		}
		p.Ln(5)
	}

	p.AddPage()
	uslovi = p.PageNo()
	p.SetFont("Times", "B", 16)
	p.CellFormat(0, 10, w.text("04-00   MOLERSKO-FARBARSKI RADOVI (GN 531)"), "0", 0, "", false, 0, "")
	p.Ln(30)
	p.SetFont("Times", "", 10)

	radovi := [][][]byte{
		p.SplitLines([]byte("Molersko-farbarski radovi se moraju izvesti stručno i kvalitetno a u svemu prema tehničkim uslovima za izvođenje molerskih radova (SRPS U.F2.013) i tehničkim uslovima za izvođenje farbarskih radova (SRPS U.F2.012)."), 200),
		p.SplitLines([]byte("Molersko-farbareske i tapetarske radove može da obavlja samo specijalizovano preduzeće ili pogoni, prema tehničkim uslovima u skladu sa SRPS U.F2.013, SRPS U.F2.014. Sav upotrebljeni materijal mora da odgovara zahtevima SRPS-a, a za materijale koji  nisu obuhvaćeni potrebno je pribaviti ateste. "), 200),
		p.SplitLines([]byte("Izvođač je dužan da primeni materijal koji odgovara mestu i uslovima ugradnje, boje i pigmenti moraju da budu otporni na svetlost. Svi spoljni premazi otporni na atmosferilije. Ukoliko izvođač upotrebi materijal za koji se atestom pokaže da nije kvalitetan, dužanje da ukloni loše izveden rad i o svom trošku izvede radove odgovarajućim, kvalitetnim materijalom.  "), 200),
		p.SplitLines([]byte("Sve pozicije  molersko-farbarskih  radova  moraju  biti izvedene  stručno  i kvalitetno,  sa materijalima  koji u svemu odgovaraju tehničkim propisima, normativima i važećim standardima, i to u onim prostorijama gde je to predviđeno izvođačkim projektom. Materijali  se mogu  ugrađivati  i primenjivati  samo  na onim površinama  za koje su prema  svojim fizičko- hemijskim  i mehaničkim  osobinama  i namenjeni.  "), 200),
		p.SplitLines([]byte("Materijali  koji nisu obuhvaćeni  standardima  moraju biti najboljeg kvaliteta i za ove materijale izvođač je dužan da dostavi ateste o izvršenom ispitivanju. Izvođač je obavezan da pre početka radova dostavi naručiocu ateste za sve materijale koje ugrađuje. Atesti moraju biti izdati od organizacija ovlašćenih za ovu vrstu poslova i ne smeju biti stariji od jedne godine, računajući  od  dana  izdavanja  atesta  do  dana  početka  izvođenja  radova  na  objektu.  "), 200),
		p.SplitLines([]byte("Gotovi,  fabrički proizvedeni materijali moraju se upotrebiti u svemu prema uputstvu proizvođača. Obojene površine moraju biti čiste, bez tragova četki i valjaka. Boja i ton moraju biti potpuno ujednačenog intenziteta, bez mrlja. Boja mora da prekrije podlogu u potpunosti, svi završeci obojenih površina moraju biti ravni i pravilni, kao i sastavi sa vratima, prozorima i sl. Izvođač je obavezan da pre početka radova dobro očisti podlogu od mehaničkih nečistoća, prašine i masnoće."), 200),
		p.SplitLines([]byte("Posne  i emulzivne,  odnosno  fasadne,  poludisperzivne,  kao  i lakovi,  boje  i zaštita  drveta,  ne smeju  se  ljuštiti  i  moraju  biti  otporne  na  otiranje  ukoliko  prema  uputstvu  proizvođača  posle  roka  za vezivanje mogu da se brišu lakim trljanjem krpom. "), 200),
		p.SplitLines([]byte("Disperzivne boje, uljni i bezuljni lakovi, uljane boje i mat uljane boje moraju biti postojani na pranje ukoliko prema uputstvu proizvođača posle roka za vezivanje mogu da se peru mekim sunđerom i vodom, sa malim dodatkom (oko 1%) neutralnog sredstva za pranje, a da se voda pritom ne oboji."), 200),
		p.SplitLines([]byte("Obojene površine moraju biti otporne na svetlost, uticaj temperature, razne hemijske i mehaničke uticaje, kao i na atmosferilije.  "), 200),
		p.SplitLines([]byte("Uljane boje ne smeju da se mreškaju  i da pucaju.  Za sve vrste premaza  upotrebiti  boje sa pigmentima otpornim na svetlost. Izbor boja vrši projektant, naručilac radova, ili odgovorni predstavnik naručioca, po dogovoru. "), 200),
		p.SplitLines([]byte("Izvođač je obavezan da podnese ton karte za odgovarajuće  materijale. Izvođač je obavezan da uradi probne uzorke veličine  1,0  m2  za  svaku  vrstu  bojenja  i mođe  da  pristupi  finalnom  bojenju  tek  po  dobijanju  pismene saglasnosti lica određenog da izvrši izbor boja. "), 200),
		p.SplitLines([]byte("Za vreme izvođenja radova izvođač ne sme da nepažnjom svojih radnika uprlja već izvedene druge vrste radova drugih izvođača. U protivnom, izvođač je dužan da prizna naručiocu vrednost izvršenih popravki na tim radovima. Obračun  izvedenih  radova  izvršiće  se  u  skladu  sa  normama  za  izvođenje  završnih  radova  u građevinarstvu. "), 200),
		p.SplitLines([]byte("Malterisane površine i elemente od gipsa ostaviti da se „suše“ najmanje 30 dana pre početka izvođenja molerskih radova. Maksimalna dopuštena vlažnost maltera zavisi od vrste materijala koji se primenjuje. Neophodno je da nadzor odobri početak izvođenja molerskih radova. "), 200),
		p.SplitLines([]byte("Površine koje se obrađuju treba da su bez metalnih delova, zavrtnjeva, ankera i sl. ili da su oni identifikovani pre obrade. Pre obrade površine očistiti od prašine i drugih prljavština, kao što su smola, ulje, mast, rđa, cementni malter i sl. i reparirati sva minorna oštećenja, pukotine, spojeve, rupe. "), 200),
		p.SplitLines([]byte("Stare premaze koji nisu čvrsti i podesni kao podloga treba skinuti odgovarajućim postupkom, kao što je struganje, pranje, brušenje i sl."), 200),
		p.SplitLines([]byte("Dozvoljeno je obrađivati samo suvu i pripremljenu podlogu, bez nedostataka kao što su: -malter koji sadrži razne aktivne soli, neugašene čestice kreča (kokice), čestice uglja i druge organske materije koje su topive u vodi i ulju;  -mekani i slabi malteri koji se drobe (lome) ili udubljuju na pritisak prsta; -trošan, smrznut, pregašen, ispucao ili vlažan malter;  -beton ili cementni malter koji nije dovoljno suv i očišćen od ulja i masti. "), 200),
		p.SplitLines([]byte("Svi elektro fitinzi treba da se pre početka izvođenja radova demontiraju i posle završnog bojenja ponovo vrate."), 200),
		p.SplitLines([]byte("Zaštititi od boje sve površine koje ne podležu bojenju, podove, stolariju, sanitarnu opremu i sl. papirom, folijama i/ili  krep trakom. Molerski radovi treba da se izvode na temperaturi većoj od 10˚C i manjoj od 35˚C ukoliko nije drugačije određeno katalogom proizvođača ili dopušteno od strane nadzora. Period sušenja materijala između dve „ruke“ ili period sušenja nanetog materijala do nanošenja drugog materijala treba da je u skladu sa preporukama proizvođača materijala. "), 200),
		p.SplitLines([]byte("Armiranje spojeva različitih materijala ili spojeva tabli gips kartona vršiti trakama za armiranje sa staklenim vlaknima (fiberglas) min. širine 5cm, koje se postavljaju u debljinu sloja gleta 1mm. Dozvoljena su dva načina montaže, ili se trake utapaju u sveži sloj gleta, ili se lepe na izgletovanu površinu. U oba slučaja nanosi se još jedan sloj gleta tako da se obrazuju slojevi glet-traka-glet. Trake treba da dobro provode vlagu. "), 200),
		p.SplitLines([]byte("Za armiranje spojeva tabli gipskartona koriste se samolepljive trake. Spoljašnje uglove zaštititi aluminijumskim ugaonim profilima. Glet masu nanositi ručno ili mašinski. Voditi računa da svuda masa bude jednake debljine (oko 2-3 mm) i da nema neravnina i linija od krajeva gleterice. Kada su se zidovi osušili, sitnom šmirglom (ručno ili mašinski) lagano prešmirglati ogletovane površine."), 200),
		p.SplitLines([]byte("Ako su i dalje ostale neke neravnine poravnati ih šmirglom. Zatim naneti drugi sloj koji je nešto tanji nego prethodni (oko 2mm). Kada se osuši drugi sloj ponoviti postupak šmirglanja. Ukoliko nadzor zahteva naneti i treći sloj koji je nešto tanji nego prethodni (oko 1mm). Pošto se osušio treći sloj veoma sitnom šmirglom preći lagano preko zida. Nijansa boje, za svaki od zidova ili plafona, biće opredeljena u glavnom projektu. "), 200),
		p.SplitLines([]byte("Za prostorije za koje se ne radi projekat eneterijera, boju određuje projektant uz obaveznu saglasnost naručioca. Probni premazi se moraju po želji naručioca izvesti za sve premaze, različite po tonu i načinu izvođenja. Boju nanositi u najmanje dva sloja, ručno (mikrofiber valjkom, sa što manjom upotrebom četke) ili mašinski (pištoljem). "), 200),
		p.SplitLines([]byte("Premazi boje moraju da odaju ujednačenu površinu, bez tragova četke ili valjka. Boja mora biti ujednačenog intenziteta (bez mrlja) i da potpuno pokriva podlogu. Svi završeci obojenih površina moraju biti ravni i pravilni. "), 200),
		p.SplitLines([]byte("Obaveze izvođača molerskih radova, koje se ne obračunavaju i ne naplaćuju posebno su: -Dovoz materijala, čak i u slučaju kada ga daje naručilac, od skladišta na gradilištu do mesta ugradnje i njegovo eventualno vraćanje; -Nabavka, montaža, korišćenje i demontaža i odvoz skele sa radnim postoljem do visine 2,0m; 	-Popravljanje manjih neravnina podloge i kpljenje rupica od eksera i vijaka; -Izrada probnih uzoraka."), 200),
	}
	for _, rad := range radovi {
		for _, r := range rad {
			p.CellFormat(190.0, lineHt, string(r), "", 1, "J", false, 0, "")
		}
		p.Ln(5)
	}
}
