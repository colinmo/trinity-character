package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func utilityWindow(app *fyne.App) fyne.Window {
	window := (*app).NewWindow("Utility")
	window.SetContent(widget.NewLabel("Hello"))
	return window
}
