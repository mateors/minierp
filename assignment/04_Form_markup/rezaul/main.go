package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {

	App := app.New()
	Window := App.NewWindow("My Form")

	Window.Resize(fyne.NewSize(600, 400))

	title := canvas.NewText("Welcome to miniERP", color.White)
	title.TextSize = 20
	title.Alignment = fyne.TextAlignCenter

	name := widget.NewEntry()
	message := widget.NewMultiLineEntry()

	row1 := widget.NewFormItem("Name", name)
	row2 := widget.NewFormItem("Message", message)

	wform := widget.NewForm(row1, row2)
	wform.SubmitText = "Save"
	wform.OnSubmit = func() {
		name := name.Text
		message := message.Text
		myData := fmt.Sprintf("Welcome %v \n %v", name, message)
		dialog.NewInformation("Confirmation", myData, Window).Show()

		// Welcome "name"
		// message
	}
	wform.OnCancel = func() {}

	Window.SetContent(container.NewVBox(
		title,
		wform,
	))
	Window.ShowAndRun()

}
