package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func ShowDashbord(a fyne.App) {
	// time.Sleep(time.Second * 5)

	win := myWindow
	label := widget.NewLabel("DashBord")
	win.SetContent(label)

	win.Show()

}
