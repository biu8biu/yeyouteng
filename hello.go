package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Everyone!")
	w.SetContent(widget.NewVBox(
		hello,
		widget.NewButton("Hey!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
