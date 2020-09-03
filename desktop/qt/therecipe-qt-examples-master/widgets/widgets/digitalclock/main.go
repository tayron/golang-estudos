package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	var app = widgets.NewQApplication(len(os.Args), os.Args)

	var clock = NewDigitalClock()

	clock.Show()

	app.Exec()
}
