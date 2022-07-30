package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	mainWindow := setupMainWindow(&app)

	mainWindow.Show()
	app.Run()
}

func setupMainWindow(app *fyne.App) fyne.Window {
	mainWindow := (*app).NewWindow("Trinity Continuum")
	mainWindow.Resize(fyne.NewSize(640, 480))
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

	content := container.NewBorder(nil, nil, nil, nil, widget.NewLabel("Content"))
	mainWindow.SetContent(content)
	return mainWindow
}
