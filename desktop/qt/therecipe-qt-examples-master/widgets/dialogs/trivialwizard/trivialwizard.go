package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func createIntroPage() *widgets.QWizardPage {
	var page = widgets.NewQWizardPage(nil)
	page.SetTitle("Introduction")

	var label = widgets.NewQLabel2("This wizard will help you register your copy "+
		"of Super Product Two.", page, 0)
	label.SetWordWrap(true)

	var layout = widgets.NewQVBoxLayout()
	layout.AddWidget(label, 0, 0)
	page.SetLayout(layout)

	return page
}

func createRegistrationPage() *widgets.QWizardPage {
	var page = widgets.NewQWizardPage(nil)
	page.SetTitle("Registration")
	page.SetSubTitle("Please fill both fields.")

	var nameLabel = widgets.NewQLabel2("Name:", page, 0)
	var nameLineEdit = widgets.NewQLineEdit(page)

	var emailLabel = widgets.NewQLabel2("Email address:", page, 0)
	var emailLineEdit = widgets.NewQLineEdit(page)

	var layout = widgets.NewQGridLayout(page)
	layout.AddWidget2(nameLabel, 0, 0, 0)
	layout.AddWidget2(nameLineEdit, 0, 1, 0)
	layout.AddWidget2(emailLabel, 1, 0, 0)
	layout.AddWidget2(emailLineEdit, 1, 1, 0)
	page.SetLayout(layout)

	return page
}

func createConclusionPage() *widgets.QWizardPage {
	var page = widgets.NewQWizardPage(nil)
	page.SetTitle("Conclusion")

	var label = widgets.NewQLabel2("You are now successfully registered. Have a "+
		"nice day!", page, 0)
	label.SetWordWrap(true)

	var layout = widgets.NewQVBoxLayout()
	layout.AddWidget(label, 0, 0)
	page.SetLayout(layout)

	return page
}

func main() {
	var app = widgets.NewQApplication(len(os.Args), os.Args)

	var wizard = widgets.NewQWizard(nil, 0)
	wizard.AddPage(createIntroPage())
	wizard.AddPage(createRegistrationPage())
	wizard.AddPage(createConclusionPage())

	wizard.SetWindowTitle("Trival Winzard")
	wizard.Show()

	app.Exec()
}
