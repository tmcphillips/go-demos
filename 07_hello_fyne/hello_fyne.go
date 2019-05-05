package main

import "fyne.io/fyne/app"

func main() {

	app := app.New()

	win := app.NewWindow("Hello World")

	win.ShowAndRun()
}
