package calc

import (
	"fmt"
	"gioui.org/widget"
	"github.com/w-ingsolutions/cms/pkg/phi"
	"github.com/w-ingsolutions/kum/app/mod"
)

func (w *WingCal) Ucitaj(hash string) {
	r := w.Jdb.ReadList(hash)
	for _, folder := range r {

		switch folder.Name {
		case "radovi":
			go w.UcitajRadove(folder.Cid.String())
		case "materijali":
			go w.UcitajMaterijal(folder.Cid.String())
		}
	}
}

func (w *WingCal) UcitajRadove(hash string) {
	var rdv []mod.ElementMenu
	radovi := make(map[int]mod.ElementMenu)
	radoviDb := w.Jdb.ReadList(hash)
	for _, rad := range radoviDb {
		rdv = append(rdv, w.UcitajVrsteRadova(rad.Cid.String()))
	}
	for _, r := range rdv {
		if r.Id > 0 {
			radovi[r.Id-1] = r
		}
	}
	w.IzbornikRadova = radovi
	//
	//fmt.Println("IzbornikRadova--------><", w.IzbornikRadova)
	//fmt.Println("--------------------------------------------------------------------------------------------------------------")
	//fmt.Println("radovi--------><", radovi)
	return
}

func (w *WingCal) UcitajElemente(icon *widget.Icon, hash string) {
	var rdv []mod.ElementMenu
	radovi := make(map[int]mod.ElementMenu)
	radoviDb := w.Jdb.ReadList(hash)

	for _, rad := range radoviDb {
		if rad.Name != "φ" {
			//fff := w.UcitajElementMenu(icon, rad.Cid.String())
			//fmt.Println("fffffff--------><", fff)
			//fmt.Println("--------><")
			rdv = append(rdv, w.UcitajElementMenu(icon, rad.Cid.String()))
		}
	}
	for _, r := range rdv {
		if r.Id > 0 {
			radovi[r.Id-1] = r
		}
	}
	w.IzbornikRadova = radovi
	return
}

func (w *WingCal) UcitajVrsteRadova(hash string) (r mod.ElementMenu) {
	vrstaRadovaDb := w.Jdb.ReadList(hash)
	for _, vrstaRadova := range vrstaRadovaDb {
		if vrstaRadova.Name == "φ" {
			var item phi.C
			w.Jdb.Read(vrstaRadova.Cid.String(), &item)
			//fmt.Println("itemTitle,", item.Title)
			r = mod.ElementMenu{
				Id:        item.ID,
				Title:     item.Title,
				Slug:      item.Slug,
				Materijal: false,
				Link:      new(widget.Clickable),
				Icon:      mustIcon(widget.NewIcon(item.Icon)),
				Hash:      hash,
			}
		}
	}
	return
}

func (w *WingCal) UcitajElementMenu(icon *widget.Icon, hash string) (r mod.ElementMenu) {
	var item phi.Φ
	fmt.Println("hash--------><", hash)
	w.Jdb.Read(hash, &item)
	fmt.Println("itemTitle,", item.Struct["Title"])
	r = mod.ElementMenu{
		Id:        item.ID,
		Title:     (item.Struct["Title"].Content).(string),
		Slug:      (item.Struct["Slug"].Content).(string),
		Materijal: false,
		Link:      new(widget.Clickable),
		Icon:      icon,
		Hash:      hash,
	}
	return
}

func (w *WingCal) UcitajElement(icon *widget.Icon, hash string) {
	var rad *phi.Φ
	w.Jdb.Read(hash, &rad)
	w.PrikazaniElement.el = rad
	fmt.Println("radr33333333333333adrad", rad)
	return
}

func (w *WingCal) UcitajMaterijal(hash string) {
	//var mat []phi.Φ
	materijali := make(map[int]phi.Φ)
	materijaliDb := w.Jdb.ReadList(hash)

	for _, materijal := range materijaliDb {
		if materijal.Name != "φ" {
			var m phi.Φ
			w.Jdb.Read(materijal.Cid.String(), &m)
			materijali[m.ID] = m

			//fmt.Println("00000000000000")
			//fmt.Println("00000000000000")
			//fmt.Println("mmmmmmmmmmID", m.ID)
			//fmt.Println("mmmmmmmmmmISSSSS", m.Struct)
			//fmt.Println("00000000000000")
			//fmt.Println("00000000000000")

		}
	}
	w.Materijal = materijali
	return
}

func mustIcon(ic *widget.Icon, err error) *widget.Icon {
	if err != nil {
		panic(err)
	}
	return ic
}
