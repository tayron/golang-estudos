package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Box Layout")

	text1 := canvas.NewText("Hello", color.White)
	text2 := canvas.NewText("There", color.White)
	text3 := canvas.NewText("(right)", color.White)
	container := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		text1, text2, layout.NewSpacer(), text3)

	text4 := canvas.NewText("centered", color.White)
	centered := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		layout.NewSpacer(), text4, layout.NewSpacer())
	myWindow.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), container, centered))
	myWindow.ShowAndRun()
}
