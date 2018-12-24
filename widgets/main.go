package widgets

import (
	"github.com/jroimartin/gocui"
	"github.com/OdaDaisuke/go-ftp-client/store"
	"fmt"
	"strings"
	"github.com/OdaDaisuke/go-ftp-client/lib"
)

type Menu struct {
}

func NewMenu() *Menu { return &Menu{} }

func (m *Menu) Layout(g *gocui.Gui) error {
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

type ConnDetail struct {
	CurView lib.ConnectionJsonModel
	FtpStatus string
}

func NewConnDetail() *ConnDetail { return &ConnDetail{FtpStatus: ""} }

func (d *ConnDetail) Layout(g *gocui.Gui) error {
	maxX, _ := g.Size()
	v, err := g.SetView(store.DetailViewName, 21, store.InitY, maxX, 20)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		d.CurView = store.FtpConnectionList[store.CurView]
		fmt.Fprintln(v, "[Detail]")
		fmt.Fprintf(v, "Name:     %s\n", d.CurView.Name)
		fmt.Fprintf(v, "Host:     %s\n", d.CurView.Host)
		fmt.Fprintf(v, "User:     %s\n", d.CurView.User)
		fmt.Fprintf(v, "Port:     %d\n", d.CurView.Port)
		fmt.Fprintf(v, "Password: %s\n\n", strings.Repeat("*", len(d.CurView.Password)))
		if d.FtpStatus == "" {
			fmt.Fprintln(v, "[Action]")
			fmt.Fprintln(v, "> 1. Connect")
			fmt.Fprintln(v, "> 2. Update")
			fmt.Fprintln(v, "> 3. Delete")
		}
	}
	return nil
}