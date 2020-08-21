package calc

import (
	"encoding/json"
	"fmt"
	"gioui.org/widget"
	"github.com/w-ingsolutions/c/model"
	"github.com/w-ingsolutions/cms/pkg/phi"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (w *WingCal) UcitajRadove(hash, name string) {
	r := w.Jdb.ReadList(hash)
	for _, folder := range r {
		if folder.Name == name {
			w.UcitajPodKategorijuRadova(folder.Cid.String())
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
			//rdv = append(rdv, w.UcitajElement(icon, rad.Cid.String()))
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

func (w *WingCal) UcitajElement(icon *widget.Icon, hash string) (r model.ElementMenu) {
	var item phi.Φ
	w.Jdb.Read(hash, &item)
	fmt.Println("itemTitle,", item.Struct["Title"])
	r = model.ElementMenu{
		Id:        item.ID,
		Title:     item.Struct["Title"].Title,
		Slug:      item.Struct["Slug"].Title,
		Materijal: false,
		Link:      new(widget.Clickable),
		Icon:      icon,
		Hash:      hash,
	}
	return
}

func (w *WingCal) APIpozivElementi(komanda string) {
	radovi := map[int]model.ElementMenu{}
	api, err := w.APIpoziv(w.API.Adresa, komanda)
	if err != nil {
		w.API.OK = false
	} else {
		jsonErr := json.Unmarshal(api, &radovi)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		w.IzbornikRadova = radovi
	}
}

func (w *WingCal) APIpozivElement(komanda string) {
	rad := &model.WingVrstaRadova{}
	api, err := w.APIpoziv(w.API.Adresa, komanda)
	if err != nil {
		w.API.OK = false
	} else {
		jsonErr := json.Unmarshal(api, &rad)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		w.PrikazaniElement = rad
	}
}

func (w *WingCal) APIpoziv(adresa, komanda string) ([]byte, error) {
	var body []byte
	url := adresa + komanda
	fmt.Println("url", url)
	spaceClient := http.Client{
		Timeout: time.Second * 10, // Maximum of 2 secs
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "wing")
	res, err := spaceClient.Do(req)
	if err != nil {
		return nil, err
	} else {
		body, err = ioutil.ReadAll(res.Body)
	}
	if err != nil {
		return nil, err
		//log.Fatal(readErr)
	}
	if body != nil {
		//defer body.Close()
	}
	return body, err
}

func mustIcon(ic *widget.Icon, err error) *widget.Icon {
	if err != nil {
		panic(err)
	}
	return ic
}
