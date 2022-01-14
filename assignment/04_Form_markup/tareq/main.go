package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	myApp := app.New()

	myWindow := myApp.NewWindow("Add new client")

	myWindow.Resize(fyne.NewSize(400, 400))

	title := canvas.NewText("Add new client", color.White)

	title.Resize(fyne.NewSize(300, 35)) // 300 is width & 35 is height
	title.Move(fyne.NewPos(50, 10))

	//client name
	entry_name := widget.NewEntry()
	entry_name.SetPlaceHolder("Client's Name")
	entry_name.Resize(fyne.NewSize(300, 35))
	entry_name.Move(fyne.NewPos(50, 50))

	// mobile
	entry_mobile := widget.NewEntry()
	entry_mobile.SetPlaceHolder("Mobile Number")
	entry_mobile.Resize(fyne.NewSize(300, 35))
	entry_mobile.Move(fyne.NewPos(50, 100))

	// Email address
	entry_email := widget.NewEntry()
	entry_email.SetPlaceHolder("Email Address")
	entry_email.Resize(fyne.NewSize(300, 35))
	entry_email.Move(fyne.NewPos(50, 150))

	// Address
	entry_addr := widget.NewEntry()
	entry_addr.SetPlaceHolder("Address")
	entry_addr.Resize(fyne.NewSize(300, 35))
	entry_addr.Move(fyne.NewPos(50, 200))

	// Submit button
	submit_btn := widget.NewButton("SUBMIT", nil)
	submit_btn.Resize(fyne.NewSize(80, 30))

	submit_btn.Move(fyne.NewPos(50, 355))
	// setup content
	// we are going to use container without layout// my favorite
	myWindow.SetContent(
		container.NewWithoutLayout(
			title,
			entry_name,
			entry_mobile,
			entry_email,
			entry_addr,
			submit_btn,
		),
	)
	// Show and run
	myWindow.ShowAndRun()
}
