package mwd

import (
	"log"

	"github.com/lxn/win"

	"github.com/lxn/walk"
	// . "github.com/lxn/walk/declarative"
)

type MyWindow struct {
	*walk.MainWindow
	HWnd win.HWND
	Ni   *walk.NotifyIcon
}

func (mw *MyWindow) SetMinimizeBox(flg bool) {
	if flg {
		mw.addStyle(win.WS_MINIMIZEBOX)
		return
	}
	mw.removeStyle(^win.WS_MINIMIZEBOX)
}

func (mw *MyWindow) SetMaximizeBox(flg bool) {
	if flg {
		mw.addStyle(win.WS_MAXIMIZEBOX)
		return
	}
	mw.removeStyle(^win.WS_MAXIMIZEBOX)
}

func (mw *MyWindow) SetSizePersistent(flg bool) {
	if flg {
		mw.addStyle(win.WS_SIZEBOX)
		return
	}
	mw.removeStyle(^win.WS_SIZEBOX)
}

func (mw *MyWindow) addStyle(style int32) {
	currStyle := win.GetWindowLong(mw.HWnd, win.GWL_STYLE)
	win.SetWindowLong(mw.HWnd, win.GWL_STYLE, currStyle|style)
}

func (mw *MyWindow) removeStyle(style int32) {
	currStyle := win.GetWindowLong(mw.HWnd, win.GWL_STYLE)
	win.SetWindowLong(mw.HWnd, win.GWL_STYLE, currStyle&style)
}

func (mw *MyWindow) SetCloseBox(flg bool) {
	if flg {
		win.GetSystemMenu(mw.HWnd, true)
		return
	}
	hMenu := win.GetSystemMenu(mw.HWnd, false)
	win.RemoveMenu(hMenu, win.SC_CLOSE, win.MF_BYCOMMAND)
}

func (mw *MyWindow) AddNotifyIcon() {
	var err error
	mw.Ni, err = walk.NewNotifyIcon(mw)
	if err != nil {
		log.Fatal(err)
	}
	icon, err := walk.Resources.Image("./img/favicon.ico")
	if err != nil {
		log.Fatal(err)
	}
	mw.SetIcon(icon)
	mw.Ni.SetIcon(icon)
	mw.Ni.SetVisible(true)

	mw.Ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button == walk.LeftButton {
			mw.Show()
			win.ShowWindow(mw.Handle(), win.SW_RESTORE)
		}
	})

}
