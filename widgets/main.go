package widgets

import (
	"github.com/jroimartin/gocui"
	"github.com/OdaDaisuke/go-ftp-client/store"
	"fmt"
	"strings"
	"github.com/OdaDaisuke/go-ftp-client/lib"
)

func RenderMainMenu(g *gocui.Gui) error {
	fc := lib.NewFileClient()
	store.FtpConnectionList = fc.ReadAll()
	initY := store.InitY

	for i, conn := range store.FtpConnectionList {
		y0, y1 := ( 3 * i ) + initY, ( i * 3 + 2 ) + initY
		if i == 0 {
			y0 = initY
		}

		v, err := g.SetView(conn.Name, 0, y0, 20, y1)
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, conn.Name)
	}

	if _, err := g.SetCurrentView(store.FtpConnectionList[0].Name); err != nil {
		return err
	}
	return nil
}

func RenderConnDetail(g *gocui.Gui) error {
	maxX, _ := g.Size()
	v, err := g.SetView(store.DetailViewName, 21, store.InitY, maxX, 18)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		curView := store.FtpConnectionList[store.CurView]
		fmt.Fprintln(v, "[Detail]")
		fmt.Fprintf(v, "Host:     %s\n", curView.Name)
		fmt.Fprintf(v, "User:     %s\n", curView.User)
		fmt.Fprintf(v, "Port:     %d\n", curView.Port)
		fmt.Fprintf(v, "Password: %s\n\n", strings.Repeat("*", len(curView.Password)))
		fmt.Fprintln(v, "[Action]")
		fmt.Fprintln(v, "> 1. Connect")
		fmt.Fprintln(v, "> 2. Update")
		fmt.Fprintln(v, "> 3. Delete")
	}
	return nil
}