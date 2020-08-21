package calc

import (
	"fmt"
	"gioui.org/widget"
	"github.com/w-ingsolutions/c/model"
	"github.com/w-ingsolutions/cms/pkg/phi"
)

func (w *WingCal) UcitajRadove(hash, name string) {
	fmt.Println("hashhash--------><", hash)

	r := w.Jdb.ReadList(hash)
	fmt.Println("rrrrrrrrrrrrrrr--------><", r)

	for _, folder := range r {
		if folder.Name == name {
			w.UcitajPodKategorijuRadova(folder.Cid.String())

			fmt.Println("ffffffffff--------><", folder.Cid.String())
			fmt.Println("--------------------------------------------------------------------------------------------------------------")
		}
	}
}

func (w *WingCal) UcitajPodKategorijuRadova(hash string) {
	var rdv []model.ElementMenu
	radovi := make(map[int]model.ElementMenu)
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
	var rdv []model.ElementMenu
	radovi := make(map[int]model.ElementMenu)
	radoviDb := w.Jdb.ReadList(hash)
	for _, rad := range radoviDb {
		if rad.Name != "φ" {
			fmt.Println("Izbffffffffff--------><", rad.Name)
			fmt.Println("IzbffCidCidCidCidfff--------><", rad.Cid.String())
			rdv = append(rdv, w.UcitajElementMenu(icon, rad.Cid.String()))
		}
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

func (w *WingCal) UcitajVrsteRadova(hash string) (r model.ElementMenu) {
	vrstaRadovaDb := w.Jdb.ReadList(hash)
	for _, vrstaRadova := range vrstaRadovaDb {
		if vrstaRadova.Name == "φ" {
			var item phi.C
			w.Jdb.Read(vrstaRadova.Cid.String(), &item)
			fmt.Println("itemTitle,", item.Title)
			r = model.ElementMenu{
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

func (w *WingCal) UcitajElementMenu(icon *widget.Icon, hash string) (r model.ElementMenu) {
	var item phi.Φ
	w.Jdb.Read(hash, &item)
	fmt.Println("itemTitle,", item.Struct["Title"])
	r = model.ElementMenu{
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
	var rad phi.Φ
	w.Jdb.Read(hash, &rad)
	w.PrikazaniElement = rad
	return
}

func mustIcon(ic *widget.Icon, err error) *widget.Icon {
	if err != nil {
		panic(err)
	}
	return ic
}
