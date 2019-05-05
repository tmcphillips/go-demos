package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {

	app := app.New()

	win := app.NewWindow("Hello World")

	win.SetContent(widget.NewVBox(
		widget.NewLabel("Hello World!"),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	win.ShowAndRun()
}
