package main

import (
	"net/url"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")

	bugURL, _ := url.Parse("https://github.com/fyne-io/fyne/issues/new")
	myWindow.SetContent(widget.NewHyperlink("Report a bug", bugURL))

	myWindow.ShowAndRun()
}
