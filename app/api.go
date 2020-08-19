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
	radovi := map[int]model.ElementMenu{}
	r := w.Jdb.ReadList(hash)
	for _, folder := range r {
		if folder.Name == name {
			radoviDb := w.Jdb.ReadList(folder.Cid.String())
			for _, rad := range radoviDb {
				vrstaRadovaDb := w.Jdb.ReadList(rad.Cid.String())
				for _, vrstaRadova := range vrstaRadovaDb {
					//fmt.Println("vrstaRadova", vrstaRadova.Name)
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
						//fmt.Println("rad>>>>", rad.Name)
					}
					//fmt.Println("vrstaRadova>>>>", vrstaRadova.Cid.String())
				}
				fmt.Println("rad>>>>", rad.Cid.String())
				fmt.Println("radNNNNNNN>>>", rad.Name)
			}
		}
		//fmt.Println("folder>>>>", folder.Cid.String())
	}
	//fmt.Println("radoviradoviradovi", radovi)
	w.IzbornikRadova = radovi
}

func (w *WingCal) UcitajRadovePodKategorija(hash, name string) {
	//radovi := map[int]model.ElementMenu{}
	vrstaRadovaDb := w.Jdb.ReadList(hash)

	for _, vrstaRadova := range vrstaRadovaDb {
		//	//fmt.Println("vrstaRadova", vrstaRadova.Name)
		//	if vrstaRadova.Name == "φ" {
		//		var item phi.C
		//		w.Jdb.Read(vrstaRadova.Cid.String(), &item)
		//		radovi[item.ID-1] = model.ElementMenu{
		//			Id:        item.ID,
		//			Title:     item.Title,
		//			Slug:      item.Slug,
		//			Materijal: false,
		//			Link:      new(widget.Clickable),
		//			Icon:      mustIcon(widget.NewIcon(item.Icon)),
		//			Hash:      vrstaRadova.Cid.String(),
		//		}
		fmt.Println("rad>>>>", vrstaRadova.Name)
		fmt.Println("radCCCC>>>>", vrstaRadova.Cid)
		//	}
	}
	fmt.Println("radoviradoviradovi", vrstaRadovaDb)
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
