package gui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func Run() {
	a := app.New()
	w := a.NewWindow("Sambal V0.1")

	w.SetContent(widget.NewLabel("Sambal GUI comming soon..."))
	w.ShowAndRun()
}
