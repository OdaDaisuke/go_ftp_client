package keybinds

import (
	"github.com/jroimartin/gocui"
	"github.com/OdaDaisuke/go-ftp-client/store"
	"github.com/OdaDaisuke/go-ftp-client/widgets"
	"github.com/OdaDaisuke/go-ftp-client/lib"
)

type AppKeybind struct {
	ftpClient *lib.FTPClient
	connDetail *widgets.ConnDetail
}

func NewKeyBinds(ftpClient *lib.FTPClient, connDetail *widgets.ConnDetail) *AppKeybind {
	return &AppKeybind{ftpClient, connDetail}
}

func (k *AppKeybind) InitKeybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return gocui.ErrQuit
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyBackspace, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return k.DelView(g)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return k.NextView(g)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return k.NextView(g)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return k.PrevView(g)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return k.FocusConnList(g)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return k.FocusConnDetail(g)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return k.FocusConnDetail(g)
		}); err != nil {
			return err
	}
	if err := g.SetKeybinding("", '1', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			curConn := store.FtpConnectionList[store.CurView]
			k.ftpClient.SetConf(curConn.Host, curConn.User, curConn.Password, curConn.Port)
			err := k.ftpClient.Connect()
			if err != nil {
				return err
			}
			k.reloadConnDetail(g)
			return nil
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", 'e', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			err := k.ToggleEditView(v)
			return err
		}); err != nil {
		return err
	}
	return nil
}

func (k *AppKeybind) NextView(g *gocui.Gui) error {
	next := store.CurView + 1
	if next > len(store.FtpConnectionList)-1 {
		next = 0
	}

	if _, err := g.SetCurrentView(store.FtpConnectionList[next].Name); err != nil {
		return err
	}

	store.CurView = next
	k.reloadConnDetail(g)
	return nil
}

func (k *AppKeybind) PrevView(g *gocui.Gui) error {
	next := store.CurView - 1
	if next < 0 {
		next = len(store.FtpConnectionList) - 1
	}

	if _, err := g.SetCurrentView(store.FtpConnectionList[next].Name); err != nil {
		return err
	}

	store.CurView = next
	k.reloadConnDetail(g)
	return nil
}

func (k *AppKeybind) FocusConnList(g *gocui.Gui) error {
	if _, err := g.SetCurrentView(store.FtpConnectionList[store.CurView].Name); err != nil {
		return err
	}
	return nil
}

func (k *AppKeybind) FocusConnDetail(g *gocui.Gui) error {
	if _, err := g.SetCurrentView(store.DetailViewName); err != nil {
		return err
	}
	return nil
}

func (k *AppKeybind) ToggleEditView(v *gocui.View) error {
	v.Editable = !v.Editable
	return nil
}

func (k *AppKeybind) DelView(g *gocui.Gui) error {
	if len(store.FtpConnectionList) <= 1 {
		return nil
	}

	if err := g.DeleteView(store.FtpConnectionList[store.CurView].Name); err != nil {
		return err
	}
	curView := store.CurView
	store.FtpConnectionList = append(store.FtpConnectionList[:curView], store.FtpConnectionList[curView+1:]...)

	return k.NextView(g)
}

func (k *AppKeybind) reloadConnDetail(g *gocui.Gui) {
	g.DeleteView(store.DetailViewName)
	curView := store.FtpConnectionList[store.CurView]
	k.connDetail.CurView.Name = curView.Name
	k.connDetail.CurView.Host = curView.Host
	k.connDetail.CurView.User = curView.User
	k.connDetail.CurView.Port = curView.Port
	k.connDetail.CurView.Password = curView.Password
}