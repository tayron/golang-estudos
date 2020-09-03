package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type DigitalClock struct {
	*widgets.QLCDNumber
}

func NewDigitalClock() *DigitalClock {
	var digitalClock = new(DigitalClock)
	digitalClock.SetSegmentStyle(widgets.QLCDNumber__Filled)
	digitalClock.QLCDNumber = widgets.NewQLCDNumber(nil)
	var timer = core.NewQTimer(digitalClock)
	timer.ConnectTimeout(digitalClock.showTime)
	timer.Start(1000)

	digitalClock.showTime()

	digitalClock.SetWindowTitle("Digital Clock")
	digitalClock.Resize(core.NewQSize2(150, 60))

	return digitalClock
}

func (dc *DigitalClock) showTime() {
	var time = core.QTime_CurrentTime()
	var text = time.ToString("hh:mm")
	if time.Second()%2 == 0 {
		var temp = []byte(text)
		temp[2] = ' '
		text = string(temp)
	}
	dc.Display(text)
}
