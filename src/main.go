package main

import (
	"fmt"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	aboutWindow := setupAboutWindow(&app)
	mainWindow := setupMainWindow(&app, &aboutWindow)

	funSetup(&mainWindow)
	mainWindow.Show()
	app.Run()
}

func setupMainWindow(app *fyne.App, aboutWindow *fyne.Window) fyne.Window {
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
				(*aboutWindow).Show()
			}),
		))
	mainWindow.SetMainMenu(mainMenu)
	mainWindow.SetMaster()

	content := container.NewBorder(nil, nil, nil, nil, widget.NewLabel("Content"))
	mainWindow.SetContent(content)
	return mainWindow
}

func funSetup(window *fyne.Window) {
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
			widget.NewButton("Talent", func() {
				charSelect.Hide()
				(*window).SetContent(MakeCharacterCreationScreen(Talent, window))
			}),
			widget.NewButton("Psion", func() {
				charSelect.Hide()
				(*window).SetContent(MakeCharacterCreationScreen(Psion, window))
			}),
		),
		(*window).Canvas(),
	)
	return charSelect
}

func setupAboutWindow(app *fyne.App) fyne.Window {
	textBox1 := widget.NewRichTextFromMarkdown(AboutText1)
	textBox1.Wrapping = fyne.TextWrapWord
	r, _ := fyne.LoadResourceFromURLString(`https://vonexplaino.com/code/trinity-character-creator/img/OPP-Metal.1.png`)
	image := canvas.NewImageFromResource(r)
	image.FillMode = canvas.ImageFillOriginal
	textBox2 := widget.NewRichTextFromMarkdown(AboutText2)
	textBox2.Wrapping = fyne.TextWrapWord
	aboutWindow := (*app).NewWindow("About")
	aboutWindow.Hide()
	aboutWindow.Resize(fyne.NewSize(460, 550))
	aboutWindow.SetContent(
		container.NewVScroll(
			container.NewVBox(
				textBox1,
				image,
				container.NewCenter(widget.NewRichTextWithText(`2008`)),
				textBox2,
			),
		),
	)
	return aboutWindow
}

var WelcomeText = `# Welcome

This is the welcome page for the Trinity Character creator. Installed packs:

* Trinity - Talent
* Aeon - Psion, Psiad
* Aberrant - Nova
* Adventure! - Daredevil, Stalwart, Mesmirist
* Aether - Squire, Gog, Magog
`

var AboutText1 = `Used under license from Onyx Path Publishing, Inc. No part of this product may be reproduced without the permission of the Licensor.`

var AboutText2 = `Trinity Continuum, Trinity Continuum: Core, and Trinity Continuum: Aeon are trademarks of Onyx Path Publishing. All rights reserved. All characters, names, places and text herein are copyrighted by Onyx Path Publishing.

The mention of or reference to any company or product on this site is not a challenge to the trademark or copyright concerned.

This project and all elements are fiction and intended for entertainment purposes only. Reader discretion is advised.

Check out the Onyx Path at [http://www.theonyxpath.com](http://www.theonyxpath.com).`

type dotSelector struct {
	widget.Select
	Min int
	Max int
	Val int
}

func newDotSelector(min int, max int, val int) *dotSelector {
	selector := &dotSelector{}
	selector.ExtendBaseWidget(selector)
	selector.Min = min
	selector.Max = max
	selector.Val = val
	if val < min || val > max {
		val = min
	}
	optionArray := []string{}
	for x := min; x < max+1; x++ {
		optionArray = append(optionArray, fmt.Sprintf("%d", x))
	}
	selector.Select.Options = optionArray
	selector.Selected = fmt.Sprintf("%d", val)
	selector.OnChanged = func(s string) {
		selector.Val, _ = strconv.Atoi(selector.Selected)
	}
	return selector
}

func (s *dotSelector) Refresh() {
	if s.Val < s.Min || s.Val > s.Max {
		s.Val = s.Min
	}
	optionArray := []string{}
	for x := s.Min; x < s.Max+1; x++ {
		optionArray = append(optionArray, fmt.Sprintf("%d", x))
	}
	s.Select.Options = optionArray
	s.Selected = fmt.Sprintf("%d", s.Val)
}

func (s *dotSelector) SetMin(min int) {
	s.Min = min
	s.Refresh()
}
func (s *dotSelector) SetMax(max int) {
	s.Max = max
	s.Refresh()
}
func (s *dotSelector) SetVal(val int) {
	s.Val = val
	s.Refresh()
}

func (s *dotSelector) MinSize() fyne.Size {
	return fyne.Size{
		Width:  30 * 2,
		Height: 20,
	}
}

func (s *dotSelector) MaxSize() fyne.Size {
	return fyne.Size{
		Width:  30 * 2,
		Height: 20,
	}
}
