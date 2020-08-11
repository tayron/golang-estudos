package main

import (
	"os"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func buttonAction(labelButton *widget.Label) {
	labelButton.SetText("Welcome :)")
}

func setDarkTheme() {
	os.Setenv("FYNE_THEME", "dark")
}

func setLightTheme() {
	os.Setenv("FYNE_THEME", "light")
}

func init() {
	setLightTheme()
}

func main() {
	application := app.New()
	window := application.NewWindow("Hello")
	labelButton := widget.NewLabel("Hello Fyne!")
	buttonCliqueHere := widget.NewButton("Click here!!", func() {
		buttonAction(labelButton)
	})

	box := widget.NewVBox(labelButton, buttonCliqueHere)
	window.SetContent(box)
	window.ShowAndRun()
}
