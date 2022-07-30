package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	mainWindow := app.NewWindow("Trinity Continuum")
	setupMainWindow(mainWindow)

	mainWindow.Show()
	app.Run()
	tidyUp()
}

func setupMainWindow(mainWindow *app.window) {

}

func tidyUp(
	fmt.Println("Completed")
)