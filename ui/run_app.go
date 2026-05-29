package ui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func RunApp() {
	myApp := app.NewWithID("org.Fenrir.app")

	myWindow := myApp.NewWindow("Fenrir xwx<3")

	label := widget.NewLabel("TODO")

	myWindow.SetContent(container.NewVBox(
		label,
	))

	myWindow.ShowAndRun()
}
