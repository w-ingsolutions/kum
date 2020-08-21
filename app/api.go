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
			w.UcitajRadoveZ(folder.Cid.String())
		}
		//fmt.Println("folder>>>>", folder.Cid.String())
	}
	//fmt.Println("radoviradoviradovi", radovi)

}

func (w *WingCal) UcitajRadoveZ(hash string) {
	radovi := map[int]model.ElementMenu{}
	radoviDb := w.Jdb.ReadList(hash)
	for _, rad := range radoviDb {
		vrstaRadovaDb := w.Jdb.ReadList(rad.Cid.String())
		for _, vrstaRadova := range vrstaRadovaDb {
			fmt.Println("vrstaRadsssova", vrstaRadova.Name)
			if vrstaRadova.Name == "φ" {
				var item phi.C
				w.Jdb.Read(vrstaRadova.Cid.String(), &item)
				radovi[item.ID-1] = model.ElementMenu{
					Id:        item.ID,
					Title:     item.Title,
					Slug:      item.Slug,
					Materijal: false,
					Link:      new(widget.Clickable),
					Icon:      mustIcon(widget.NewIcon(item.Icon)),
					Hash:      rad.Cid.String(),
				}
			}
		}
	}
	w.IzbornikRadova = radovi
	return
}

func (w *WingCal) UcitajVrsteRadova(hash string) map[int]model.ElementMenu {
	radovi := map[int]model.ElementMenu{}
	vrstaRadovaDb := w.Jdb.ReadList(hash)
	for _, vrstaRadova := range vrstaRadovaDb {
		fmt.Println("vrstaRadova", vrstaRadova.Name)
		if vrstaRadova.Name == "φ" {
			var item phi.C
			w.Jdb.Read(vrstaRadova.Cid.String(), &item)
			radovi[item.ID-1] = model.ElementMenu{
				Id:        item.ID,
				Title:     item.Title,
				Slug:      item.Slug,
				Materijal: false,
				Link:      new(widget.Clickable),
				Icon:      mustIcon(widget.NewIcon(item.Icon)),
				//Hash:      rad.Cid.String(),
			}
		}
	}
	return radovi
}

func (w *WingCal) UcitajRadovePodKategorija(hash string) {
	//radovi := map[int]model.ElementMenu{}
	//vrstaRadovaDb := w.Jdb.ReadList(hash)

	//for _, vrstaRadova := range vrstaRadovaDb {
	//	//	//fmt.Println("vrstaRadova", vrstaRadova.Name)
	//	//	if vrstaRadova.Name == "φ" {
	//	//		var item phi.C
	//	//		w.Jdb.Read(vrstaRadova.Cid.String(), &item)
	//	//		radovi[item.ID-1] = model.ElementMenu{
	//	//			Id:        item.ID,
	//	//			Title:     item.Title,
	//	//			Slug:      item.Slug,
	//	//			Materijal: false,
	//	//			Link:      new(widget.Clickable),
	//	//			Icon:      mustIcon(widget.NewIcon(item.Icon)),
	//	//			Hash:      vrstaRadova.Cid.String(),
	//	//		}
	//	fmt.Println("rad>>>>", vrstaRadova.Name)
	fmt.Println("radCCCC>>>>", hash)
	//	//	}
	//}
	w.UcitajRadoveZ(hash)

	//fmt.Println("radoviradoviradovi", vrstaRadovaDb)
	//w.IzbornikRadova = radovi
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
