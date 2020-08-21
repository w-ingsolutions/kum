package mod

import (
	"gioui.org/widget"
	"time"
)

type W struct {
	Id                    string
	Naziv                 string
	DatumNastanka         time.Time
	DatumPoslednjePromene time.Time
	Izmena                *widget.Clickable
	Brisanje              *widget.Clickable
	Data                  interface{}
}

//type Client struct {
//	Socket net.Conn
//	data   chan []byte
//}

type WingCalEcommands map[int]WingCalEcommand

type WingCalEcommand struct {
	Id       string
	Type     string
	Name     string
	Enabled  bool
	CompType string
	SubType  string
	Command  interface{}
	Data     interface{}
}
type EditabilnaPoljaVrsteRadova struct {
	Id        *widget.Editor
	Naziv     *widget.Editor
	Opis      *widget.Editor
	Obracun   *widget.Editor
	Jedinica  *widget.Editor
	Cena      *widget.Editor
	Slug      *widget.Editor
	Omogucen  *widget.Bool
	Materijal map[int]*widget.Editor
}

type ElementMenu struct {
	Id        int
	Title     string
	Slug      string
	Materijal bool
	Link      *widget.Clickable
	Icon      *widget.Icon
	Hash      string
}

type TipSadrzaja struct {
	Naziv        string
	NazivMnozina string
	Slug         string
	SlugMnozina  string
	Struktura    interface{}
	Link         *widget.Clickable
}
