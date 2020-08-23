package main

import (
	"fmt"
	"gioui.org/app"
	_ "gioui.org/app/permission/storage"
	"gioui.org/io/system"
	"gioui.org/layout"
	"github.com/gioapp/gel/helper"
	"github.com/w-ingsolutions/kum/app"
	"github.com/w-ingsolutions/kum/app/cfg"
	in "github.com/w-ingsolutions/kum/app/cfg/ini"
	"log"
	"os"
)

func main() {

	w := calc.NewWingCal()

	if cfg.Initial {
		fmt.Println("running initial sync")
	}
	in.Init(w.Podesavanja.File)

	w.Ucitaj("QmY7xLQfkxHYLjEEJyT8c9Hbityqb3jho2Wj1qX1kB7PuB")

	//w.GenerisanjeLinkova(w.IzbornikRadova)

	go func() {
		defer os.Exit(0)
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()

}

func loop(w *calc.WingCal) error {
	for {
		select {
		case e := <-w.UI.Window.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				w.UI.Context = layout.NewContext(&w.UI.Ops, e)
				helper.Fill(w.UI.Context, helper.HexARGB(w.UI.Tema.Colors["Light"]))
				if len(w.IzbornikRadova) > 0 {
					w.GlavniEkran(w.UI.Context)
				} else {
					w.GreskaEkran()
				}

				e.Frame(w.UI.Context.Ops)
			}
			w.UI.Window.Invalidate()
		}
	}
}
