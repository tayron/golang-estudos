package main

import (
	"os"
	"runtime"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type Window struct {
	*widgets.QDialog

	iconGroupBox     *widgets.QGroupBox
	iconLabel        *widgets.QLabel
	iconComboBox     *widgets.QComboBox
	showIconCheckBox *widgets.QCheckBox

	messageGroupBox      *widgets.QGroupBox
	typeLabel            *widgets.QLabel
	durationLabel        *widgets.QLabel
	durationWarningLabel *widgets.QLabel
	titleLabel           *widgets.QLabel
	bodyLabel            *widgets.QLabel
	typeComboBox         *widgets.QComboBox
	durationSpinBox      *widgets.QSpinBox
	titleEdit            *widgets.QLineEdit
	bodyEdit             *widgets.QTextEdit
	showMessageButton    *widgets.QPushButton

	minimizeAction *widgets.QAction
	maximizeAction *widgets.QAction
	restoreAction  *widgets.QAction
	quitAction     *widgets.QAction

	trayIcon     *widgets.QSystemTrayIcon
	trayIconMenu *widgets.QMenu
}

func NewWindow(parent widgets.QWidget_ITF) *Window {
	var window = new(Window)
	window.QDialog = widgets.NewQDialog(parent, 0)

	window.createIconGroupBox()
	window.createMessageGroupBox()

	window.iconLabel.SetMinimumWidth(window.durationLabel.SizeHint().Width())

	window.createActions()
	window.createTrayIcon()

	window.showMessageButton.ConnectClicked(window.showMessage)
	window.showIconCheckBox.ConnectToggled(window.trayIcon.SetVisible)
	window.iconComboBox.ConnectCurrentIndexChanged(window.setIcon)
	window.trayIcon.ConnectMessageClicked(window.messageClicked)
	window.trayIcon.ConnectActivated(window.iconActivated)

	window.ConnectCloseEvent(window.closeEvent)
	// window.ConnectSetVisible(window.setVisible)

	var mainLayout = widgets.NewQVBoxLayout()
	mainLayout.AddWidget(window.iconGroupBox, 0, 0)
	mainLayout.AddWidget(window.messageGroupBox, 0, 0)
	window.SetLayout(mainLayout)

	window.iconComboBox.SetCurrentIndex(1)
	window.trayIcon.Show()

	window.SetWindowTitle("Systray")
	window.Resize2(400, 300)

	return window
}

func (w *Window) createIconGroupBox() {
	w.iconGroupBox = widgets.NewQGroupBox2("Tray Icon", w)

	w.iconLabel = widgets.NewQLabel2("Icon:", w, 0)

	w.iconComboBox = widgets.NewQComboBox(w)
	w.iconComboBox.AddItem2(gui.NewQIcon5(":/images/bad.png"), "Bad", core.NewQVariant())
	w.iconComboBox.AddItem2(gui.NewQIcon5(":/images/heart.png"), "Heart", core.NewQVariant())
	w.iconComboBox.AddItem2(gui.NewQIcon5(":/images/trash.png"), "Trash", core.NewQVariant())

	w.showIconCheckBox = widgets.NewQCheckBox2("Show icon", w)
	w.showIconCheckBox.SetChecked(true)

	var iconLayout = widgets.NewQHBoxLayout()

	iconLayout.AddWidget(w.iconLabel, 0, 0)
	iconLayout.AddWidget(w.iconComboBox, 0, 0)
	iconLayout.AddStretch(0)
	iconLayout.AddWidget(w.showIconCheckBox, 0, 0)
	w.iconGroupBox.SetLayout(iconLayout)
}

func (w *Window) createMessageGroupBox() {
	w.messageGroupBox = widgets.NewQGroupBox2("Balloon Message", w)

	w.typeLabel = widgets.NewQLabel2("Type:", w, 0)

	w.typeComboBox = widgets.NewQComboBox(w)
	w.typeComboBox.AddItem("None", core.NewQVariant1(widgets.QSystemTrayIcon__NoIcon))
	w.typeComboBox.AddItem2(w.Style().StandardIcon(
		widgets.QStyle__SP_MessageBoxInformation, nil, nil),
		"Information",
		core.NewQVariant1(widgets.QSystemTrayIcon__Information))
	w.typeComboBox.AddItem2(w.Style().StandardIcon(
		widgets.QStyle__SP_MessageBoxWarning, nil, nil),
		"Warning",
		core.NewQVariant1(widgets.QSystemTrayIcon__Warning))
	w.typeComboBox.AddItem2(w.Style().StandardIcon(
		widgets.QStyle__SP_MessageBoxCritical, nil, nil),
		"Critical",
		core.NewQVariant1(widgets.QSystemTrayIcon__Critical))
	w.typeComboBox.AddItem2(gui.NewQIcon(),
		"Custom icon",
		core.NewQVariant1(-1))
	w.typeComboBox.SetCurrentIndex(1)

	w.durationLabel = widgets.NewQLabel2("Duration:", w, 0)

	w.durationSpinBox = widgets.NewQSpinBox(w)
	w.durationSpinBox.SetRange(5, 60)
	w.durationSpinBox.SetSuffix(" s")
	w.durationSpinBox.SetValue(15)

	w.durationWarningLabel = widgets.NewQLabel2("(some systems might ignore this hint)", w, 0)
	w.durationWarningLabel.SetIndent(10)

	w.titleLabel = widgets.NewQLabel2("Title:", w, 0)

	w.titleEdit = widgets.NewQLineEdit2("Cannot connect to network", w)

	w.bodyLabel = widgets.NewQLabel2("Body:", w, 0)

	w.bodyEdit = widgets.NewQTextEdit(w)
	w.bodyEdit.SetPlainText("Don't believe me. Honestly, I don't have a " +
		"clue.\nClick this balloon for details.")

	w.showMessageButton = widgets.NewQPushButton2("Show Message", w)
	w.showMessageButton.SetDefault(true)

	var messageLayout = widgets.NewQGridLayout(w)
	messageLayout.AddWidget2(w.typeLabel, 0, 0, 0)
	messageLayout.AddWidget3(w.typeComboBox, 0, 1, 1, 2, 0)
	messageLayout.AddWidget2(w.durationLabel, 1, 0, 0)
	messageLayout.AddWidget2(w.durationSpinBox, 1, 1, 0)
	messageLayout.AddWidget3(w.durationWarningLabel, 1, 2, 1, 3, 0)
	messageLayout.AddWidget2(w.titleLabel, 2, 0, 0)
	messageLayout.AddWidget3(w.titleEdit, 2, 1, 1, 4, 0)
	messageLayout.AddWidget2(w.bodyLabel, 3, 0, 0)
	messageLayout.AddWidget3(w.bodyEdit, 3, 1, 2, 4, 0)
	messageLayout.AddWidget2(w.showMessageButton, 5, 4, 0)
	messageLayout.SetColumnStretch(3, 1)
	messageLayout.SetRowStretch(4, 1)
	w.messageGroupBox.SetLayout(messageLayout)
}

func (w *Window) setVisible(visible bool) {
	w.minimizeAction.SetEnabled(visible)
	w.maximizeAction.SetEnabled(!w.IsMaximized())
	w.restoreAction.SetEnabled(w.IsMaximized() || !visible)
	w.QDialog.SetVisible(visible)
}

func (w *Window) closeEvent(event *gui.QCloseEvent) {
	if runtime.GOOS == "darwin" {
		if !event.Spontaneous() || !w.IsVisible() {
			return
		}
	}

	if w.trayIcon.IsVisible() {
		widgets.QMessageBox_Information(w, "Systray",
			"The program will keep running in the "+
				"system tray. To terminate the program, "+
				"choose <b>Quit</b> in the context menu "+
				"of the system tray entry.",
			widgets.QMessageBox__Ok, widgets.QMessageBox__NoButton)
		w.Hide()
		event.Ignore()
	}
}

func (w *Window) createActions() {
	w.minimizeAction = widgets.NewQAction2("Mi&nimize", w)
	w.minimizeAction.ConnectTriggered(func(bool) { w.Hide() })

	w.maximizeAction = widgets.NewQAction2("Ma&ximize", w)
	w.maximizeAction.ConnectTriggered(func(bool) { w.ShowMaximized() })

	w.restoreAction = widgets.NewQAction2("&Restore", w)
	w.restoreAction.ConnectTriggered(func(bool) { w.ShowNormal() })

	w.quitAction = widgets.NewQAction2("&Quit", w)
	w.quitAction.ConnectTriggered(func(bool) { os.Exit(0) })
}

func (w *Window) createTrayIcon() {
	w.trayIconMenu = widgets.NewQMenu(w)
	w.trayIconMenu.QWidget.AddAction(w.minimizeAction)
	w.trayIconMenu.QWidget.AddAction(w.maximizeAction)
	w.trayIconMenu.QWidget.AddAction(w.restoreAction)
	w.trayIconMenu.AddSeparator()
	w.trayIconMenu.QWidget.AddAction(w.quitAction)

	w.trayIcon = widgets.NewQSystemTrayIcon(w)
	w.trayIcon.SetContextMenu(w.trayIconMenu)
}

func (w *Window) showMessage(bool) {
	w.showIconCheckBox.SetChecked(true)
	var ok bool
	var selectedIcon = w.typeComboBox.ItemData(w.typeComboBox.CurrentIndex(), int(core.Qt__UserRole)).ToInt(&ok)
	var msgIcon = widgets.QSystemTrayIcon__MessageIcon(selectedIcon)

	if selectedIcon == -1 { // custom icon
		var icon = gui.NewQIcon3(w.iconComboBox.ItemIcon(w.iconComboBox.CurrentIndex()))
		w.trayIcon.ShowMessage2(w.titleEdit.Text(), w.bodyEdit.ToPlainText(), icon,
			w.durationSpinBox.Value()*1000)
	} else {
		w.trayIcon.ShowMessage(w.titleEdit.Text(), w.bodyEdit.ToPlainText(), msgIcon,
			w.durationSpinBox.Value()*1000)
	}
}

func (w *Window) setIcon(index int) {
	var icon = w.iconComboBox.ItemIcon(index)
	w.trayIcon.SetIcon(icon)
	w.SetWindowIcon(icon)

	w.trayIcon.SetToolTip(w.iconComboBox.ItemText(index))
}

func (w *Window) messageClicked() {
	widgets.QMessageBox_Information(w, "Systray",
		"Sorry, I already gave what help I could.\n"+
			"Maybe you should try asking a human?",
		widgets.QMessageBox__Ok, widgets.QMessageBox__NoButton)
}

func (w *Window) iconActivated(reason widgets.QSystemTrayIcon__ActivationReason) {
	switch reason {
	case widgets.QSystemTrayIcon__Trigger, widgets.QSystemTrayIcon__DoubleClick:
		println("xxxx")
		w.iconComboBox.SetCurrentIndex((w.iconComboBox.CurrentIndex() + 1) % w.iconComboBox.Count())
	case widgets.QSystemTrayIcon__MiddleClick:
		w.showMessage(true)
	}
}
