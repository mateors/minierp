package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("My Form")

	myWindow.Resize(fyne.NewSize(600, 400))

	name := widget.NewEntry()
	message := widget.NewMultiLineEntry()

	row1 := widget.NewFormItem("Name", name)
	row2 := widget.NewFormItem("Message", message)

	wform := widget.NewForm(row1, row2)
	wform.SubmitText = "Save"
	wform.OnSubmit = func() {
		name := name.Text
		message := message.Text
		myData := fmt.Sprintf(`%v = %v`, name, message)
		dialog.NewInformation("Confirmation", myData, myWindow).Show()

	}
	wform.OnCancel = func() {}

	//labelName := widget.NewLabel("Name")
	//inputName := widget.NewEntry()
	//fContainer := container.NewBorder(widget.NewLabel("My Form"), widget.NewLabel("Footer"), labelName, nil, inputName)

	myWindow.SetContent(wform)
	myWindow.ShowAndRun()

}
