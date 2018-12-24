package main

import (
	"fmt"
	"log"
	"github.com/jroimartin/gocui"
	"github.com/OdaDaisuke/go-ftp-client/store"
	"github.com/OdaDaisuke/go-ftp-client/keybinds"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorRed

	g.SetManagerFunc(renderLayout)

	if err := keybinds.InitKeybindings(g); err != nil {
		log.Panicln(err)
	}
	if err := renderMainMenu(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func renderLayout(g *gocui.Gui) error {
	maxX, _ := g.Size()
	v, err := g.SetView("help", 0, 0, maxX-1, store.InitY - 1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Choose your ftp connections.\n")
		fmt.Fprintln(v, "Backspace: Delete Connection")
		fmt.Fprintln(v, "n: Create new connection")
		fmt.Fprintln(v, "^C: Exit")
	}
	return nil
}

func renderMainMenu(g *gocui.Gui) error {
	initY := store.InitY
	// TODO: fileClientから読み込んだものを描画
	for i, name := range store.FtpConnectionList {
		y0, y1 := ( 3 * i ) + initY, ( i * 3 + 2 ) + initY
		if i == 0 {
			y0 = initY
		}

		v, err := g.SetView(name, 0, y0, 30, y1)
		fmt.Fprintln(v, name)
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	if _, err := g.SetCurrentView(store.FtpConnectionList[0]); err != nil {
		return err
	}
	return nil
}
