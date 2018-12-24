package keybinds

import (
	"github.com/jroimartin/gocui"
	"github.com/OdaDaisuke/go-ftp-client/store"
	"github.com/OdaDaisuke/go-ftp-client/widgets"
)

func InitKeybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return gocui.ErrQuit
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyBackspace, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return DelView(g)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return NextView(g)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return NextView(g)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return PrevView(g)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return FocusConnList(g)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return FocusConnDetail(g)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return FocusConnDetail(g)
		}); err != nil {
			return err
	}
	if err := g.SetKeybinding("", 'e', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			err := ToggleEditView(v)
			return err
		}); err != nil {
		return err
	}
	return nil
}

func NextView(g *gocui.Gui) error {
	next := store.CurView + 1
	if next > len(store.FtpConnectionList)-1 {
		next = 0
	}

	if _, err := g.SetCurrentView(store.FtpConnectionList[next].Name); err != nil {
		return err
	}

	store.CurView = next
	reloadConnDetail(g)
	return nil
}

func PrevView(g *gocui.Gui) error {
	next := store.CurView - 1
	if next < 0 {
		next = len(store.FtpConnectionList) - 1
	}

	if _, err := g.SetCurrentView(store.FtpConnectionList[next].Name); err != nil {
		return err
	}

	store.CurView = next
	reloadConnDetail(g)
	return nil
}

func FocusConnList(g *gocui.Gui) error {
	if _, err := g.SetCurrentView(store.FtpConnectionList[store.CurView].Name); err != nil {
		return err
	}
	return nil
}

func FocusConnDetail(g *gocui.Gui) error {
	if _, err := g.SetCurrentView(store.DetailViewName); err != nil {
		return err
	}
	return nil
}

func ToggleEditView(v *gocui.View) error {
	v.Editable = !v.Editable
	return nil
}

func DelView(g *gocui.Gui) error {
	if len(store.FtpConnectionList) <= 1 {
		return nil
	}

	if err := g.DeleteView(store.FtpConnectionList[store.CurView].Name); err != nil {
		return err
	}
	curView := store.CurView
	store.FtpConnectionList = append(store.FtpConnectionList[:curView], store.FtpConnectionList[curView+1:]...)

	return NextView(g)
}

func reloadConnDetail(g *gocui.Gui) {
	g.DeleteView(store.DetailViewName)
	widgets.RenderConnDetail(g)
}