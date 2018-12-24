package main

import (
	"fmt"
	"log"
	"github.com/OdaDaisuke/go-ftp-client/widgets"
	"github.com/jroimartin/gocui"
	"github.com/OdaDaisuke/go-ftp-client/store"
	"github.com/OdaDaisuke/go-ftp-client/keybinds"
	"github.com/OdaDaisuke/go-ftp-client/lib"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorRed

	mainLayout := NewMainLayout()
	mainMenu := widgets.NewMenu()
	connDetail := widgets.NewConnDetail()

	ftpClient := lib.NewFTPClient("", "", "", 0)
	keybindClient := keybinds.NewKeyBinds(ftpClient, connDetail)

	g.SetManager(mainLayout, mainMenu, connDetail)

	if err := keybindClient.InitKeybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

type MainLayout struct {
}

func NewMainLayout() *MainLayout {
	return &MainLayout{}
}

func (l *MainLayout) Layout(g *gocui.Gui) error {
	maxX, _ := g.Size()
	v, err := g.SetView("help", 0, 0, maxX-1, store.InitY - 1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Choose your ftp connections.\n")
		fmt.Fprintln(v, "n: Create new connection")
		fmt.Fprintln(v, "^C: Exit")
	}
	return nil
}