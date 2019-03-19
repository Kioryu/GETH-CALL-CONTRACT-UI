package GUI

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"log"
)

type Window struct {
	main       *ui.Window
	title      string
	width      int
	height     int
	hasMenubar bool
}

func NewWindow() Window {
	return Window{
		main:       nil,
		title:      "GETH-CALL-CONTRACT-UI",
		width:      600,
		height:     600,
		hasMenubar: false,
	}
}

func Main(gui func()) {
	ui.Main(gui)
}

func Gui() {
	win := NewWindow()

	win.main = ui.NewWindow(win.title, win.width, win.height, win.hasMenubar)

	win.main.OnClosing(func(window *ui.Window) bool {
		ui.Quit()
		log.Println("OnClosing")
		return true
	})

	ui.OnShouldQuit(func() bool {
		win.main.Destroy()
		log.Println("OnShouldQuit")
		return true
	})

	win.main.Show()
}
