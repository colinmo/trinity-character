package main

import (
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	mainWindow := setupMainWindow(&app)

	funSetup(&mainWindow)
	mainWindow.Show()
	app.Run()
}

func setupMainWindow(app *fyne.App) fyne.Window {
	mainWindow := (*app).NewWindow("Trinity Continuum")
	mainWindow.Resize(fyne.NewSize(640, 960))
	mainMenu := fyne.NewMainMenu(
		fyne.NewMenu("Trinity"),
		fyne.NewMenu("Characters",
			fyne.NewMenuItem("New", func() {
				log.Println("New character")
			}),
			fyne.NewMenuItem("Load", func() {
				log.Println("Load character")
			})),
		fyne.NewMenu("Help",
			fyne.NewMenuItem("About", func() {
				log.Println("Display help")
			}),
		))
	mainWindow.SetMainMenu(mainMenu)
	mainWindow.SetMaster()

	content := container.NewBorder(nil, nil, nil, nil, widget.NewLabel("Content"))
	mainWindow.SetContent(content)
	return mainWindow
}

func funSetup(window *fyne.Window) {
	mep0 := MakeDotButtons(5, 1)
	mep1 := MakeDotButtons(5, 1)
	mep2 := MakeDotButtons(5, 1)
	mep3 := widget.NewSelect(
		[]string{
			"●oooo",
			"●●ooo",
			"●●●oo",
			"●●●●o",
			"●●●●●",
		},
		func(s string) { fmt.Printf("Changed\n") },
	)
	statLine := container.New(
		layout.NewFormLayout(),
		canvas.NewText("Might", color.Black),
		mep0.ToCanvas(),
		canvas.NewText("Dexterity", color.Black),
		mep1.ToCanvas(),
		canvas.NewText("Resilience", color.Black),
		mep2.ToCanvas(),
		canvas.NewText("Intelligence", color.Black),
		mep3,
	)

	(*window).SetContent(statLine)
}
