package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type Window struct {
	*widgets.QWidget
}

func NewWindow(parent widgets.QWidget_ITF) *Window {
	var window = new(Window)
	window.QWidget = widgets.NewQWidget(parent, 0)

	var grid = widgets.NewQGridLayout(window)
	grid.AddWidget2(window.createFirstExclusiveGroup(), 0, 0, core.Qt__AlignLeft)
	grid.AddWidget2(window.createSecondExclusiveGroup(), 1, 0, core.Qt__AlignLeft)
	grid.AddWidget2(window.createNonExclusiveGroup(), 0, 1, core.Qt__AlignLeft)
	grid.AddWidget2(window.createPushButtonGroup(), 1, 1, core.Qt__AlignLeft)
	window.SetLayout(grid)

	window.SetWindowTitle("Group Boxes")
	window.Resize2(480, 320)

	return window
}

func (w *Window) createFirstExclusiveGroup() *widgets.QGroupBox {
	var groupBox = widgets.NewQGroupBox2("Exclusive Radio Buttons", w)

	var radio1 = widgets.NewQRadioButton2("&Radio button 1", groupBox)
	var radio2 = widgets.NewQRadioButton2("R&adio button 2", groupBox)
	var radio3 = widgets.NewQRadioButton2("Ra&dio button 3", groupBox)

	radio1.SetChecked(true)

	var vbox = widgets.NewQVBoxLayout()
	vbox.AddWidget(radio1, 0, core.Qt__AlignLeft)
	vbox.AddWidget(radio2, 0, core.Qt__AlignLeft)
	vbox.AddWidget(radio3, 0, core.Qt__AlignLeft)
	vbox.AddStretch(1)
	groupBox.SetLayout(vbox)

	return groupBox
}

func (w *Window) createSecondExclusiveGroup() *widgets.QGroupBox {
	var groupBox = widgets.NewQGroupBox2("E&xclusive Radio Buttons", w)
	groupBox.SetCheckable(true)
	groupBox.SetChecked(false)

	var radio1 = widgets.NewQRadioButton2("Rad&io button 1", groupBox)
	var radio2 = widgets.NewQRadioButton2("Radi&o button 2", groupBox)
	var radio3 = widgets.NewQRadioButton2("Radio &button 3", groupBox)
	radio1.SetChecked(true)
	var checkBox = widgets.NewQCheckBox2("Ind&ependent checkbox", groupBox)
	checkBox.SetChecked(true)

	var vbox = widgets.NewQVBoxLayout()
	vbox.AddWidget(radio1, 0, core.Qt__AlignLeft)
	vbox.AddWidget(radio2, 0, core.Qt__AlignLeft)
	vbox.AddWidget(radio3, 0, core.Qt__AlignLeft)
	vbox.AddWidget(checkBox, 0, core.Qt__AlignLeft)
	vbox.AddStretch(1)
	groupBox.SetLayout(vbox)

	return groupBox
}

func (w *Window) createNonExclusiveGroup() *widgets.QGroupBox {
	var groupBox = widgets.NewQGroupBox2("Non-Exclusive Checkboxes", w)
	groupBox.SetFlat(true)

	var checkBox1 = widgets.NewQCheckBox2("&Checkbox 1", groupBox)
	var checkBox2 = widgets.NewQCheckBox2("C&heckbox 2", groupBox)
	checkBox2.SetChecked(true)
	var tristateBox = widgets.NewQCheckBox2("Tri-&state button", groupBox)
	tristateBox.SetTristate(true)

	tristateBox.SetCheckState(core.Qt__PartiallyChecked)

	var vbox = widgets.NewQVBoxLayout()
	vbox.AddWidget(checkBox1, 0, core.Qt__AlignLeft)
	vbox.AddWidget(checkBox2, 0, core.Qt__AlignLeft)
	vbox.AddWidget(tristateBox, 0, core.Qt__AlignLeft)
	vbox.AddStretch(1)
	groupBox.SetLayout(vbox)

	return groupBox
}

func (w *Window) createPushButtonGroup() *widgets.QGroupBox {
	var groupBox = widgets.NewQGroupBox2("&Push Buttons", w)
	groupBox.SetCheckable(true)
	groupBox.SetChecked(true)

	var pushButton = widgets.NewQPushButton2("&Normal Button", groupBox)
	var toggleButton = widgets.NewQPushButton2("&Toggle Button", groupBox)
	toggleButton.SetCheckable(true)
	toggleButton.SetChecked(true)
	var flatButton = widgets.NewQPushButton2("&Flat Button", groupBox)
	flatButton.SetFlat(true)

	var popupButton = widgets.NewQPushButton2("Pop&up Button", groupBox)
	var menu = widgets.NewQMenu(w)
	menu.AddAction("&First Item")
	menu.AddAction("&Second Item")
	menu.AddAction("&Third Item")
	menu.AddAction("F&ourth Item")
	popupButton.SetMenu(menu)

	var newAction = menu.AddAction("Submenu")
	var subMenu = widgets.NewQMenu2("Popup Submenu", w)
	subMenu.AddAction("Item 1")
	subMenu.AddAction("Item 2")
	subMenu.AddAction("Item 3")
	newAction.SetMenu(subMenu)

	var vbox = widgets.NewQVBoxLayout()
	vbox.AddWidget(pushButton, 0, core.Qt__AlignLeft)
	vbox.AddWidget(toggleButton, 0, core.Qt__AlignLeft)
	vbox.AddWidget(flatButton, 0, core.Qt__AlignLeft)
	vbox.AddWidget(popupButton, 0, core.Qt__AlignLeft)
	vbox.AddStretch(1)
	groupBox.SetLayout(vbox)

	return groupBox
}
