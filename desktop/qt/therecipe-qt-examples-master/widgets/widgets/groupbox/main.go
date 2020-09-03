package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	var app = widgets.NewQApplication(len(os.Args), os.Args)

	var window = NewWindow(nil)

	window.Show()

	app.Exec()
}
