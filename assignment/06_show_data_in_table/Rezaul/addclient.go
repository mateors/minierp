package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showClient(a fyne.App) {

	win := myWindow

	title := canvas.NewText("Please put your client details here", color.White)
	title.TextSize = 20
	title.Alignment = fyne.TextAlignCenter

	nameEntry := widget.NewEntry()
	nameEntry.PlaceHolder = "Enter Client Name"
	mobileEntry := widget.NewEntry()
	mobileEntry.PlaceHolder = "Enter Client Mobile"
	emailEntry := widget.NewEntry()
	emailEntry.PlaceHolder = "Enter Client Email Address"
	addressEntry := widget.NewMultiLineEntry()
	addressEntry.PlaceHolder = "Enter Client Address"

	name1 := widget.NewFormItem("Name", nameEntry)
	mobile1 := widget.NewFormItem("Mobile", mobileEntry)
	email1 := widget.NewFormItem("Email", emailEntry)
	address1 := widget.NewFormItem("Address", addressEntry)
	clientForm := widget.NewForm(name1, mobile1, email1, address1)

	clientForm.SubmitText = "Save"
	clientForm.CancelText = "Cancel"

	clientForm.OnCancel = func() {
		Dashbord(myApp)
	}

	clientForm.OnSubmit = func() {
		name := nameEntry.Text
		mobile := mobileEntry.Text
		email := emailEntry.Text
		address := addressEntry.Text

		id, err := AddClient(name, mobile, email, address)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id)

	}
	win.SetContent(container.NewVBox(title, clientForm))

	win.Show()

}
