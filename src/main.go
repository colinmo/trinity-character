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

	funSetup(&mainWindow)
	mainWindow.Show()
	app.Run()
}

func setupMainWindow(app *fyne.App) fyne.Window {
	mainWindow := (*app).NewWindow("Trinity Continuum")
	mainWindow.Resize(fyne.NewSize(800, 1000))
	mainMenu := fyne.NewMainMenu(
		fyne.NewMenu("Trinity"),
		fyne.NewMenu("Characters",
			fyne.NewMenuItem("New", func() {
				pop := newCharacterPrompt(&mainWindow)
				pop.Show()
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
	// (*window).SetContent(MakeEditableCharacterSheet())
	(*window).SetContent(
		container.NewVScroll(
			widget.NewRichTextFromMarkdown(WelcomeText),
		),
	)
}

func newCharacterPrompt(window *fyne.Window) *widget.PopUp {
	var charSelect *widget.PopUp
	charSelect = widget.NewModalPopUp(
		container.NewGridWrap(
			fyne.NewSize(200, 75),
			widget.NewButton("TrinityCore", func() {
				charSelect.Hide()
				(*window).SetContent(MakeCharacterCreationScreen(Trinity))
			}),
			widget.NewButton("Psion", func() {
				charSelect.Hide()
				(*window).SetContent(MakeCharacterCreationScreen(Psion))
			}),
		),
		(*window).Canvas(),
	)
	return charSelect
}

var WelcomeText = `# Welcome

This is the welcome page for the Trinity Character creator. Installed packs:

* Trinity - Talent
* Aeon - Psion, Psiad
* Aberrant - Nova
* Adventure! - Daredevil, Stalwart, Mesmirist
* Aether - Squire, Gog, Magog
`
