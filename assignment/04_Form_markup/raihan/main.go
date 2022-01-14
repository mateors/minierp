package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()
var myWindow fyne.Window = myApp.NewWindow("Hello")

func main() {

	myWindow.Resize(fyne.NewSize(500, 500))
	emailEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()

	email := widget.NewFormItem("Name", emailEntry)
	password := widget.NewFormItem("Password", passwordEntry)

	clientForm := widget.NewForm(email, password)

	clientForm.SubmitText = "Login"
	messageLabel := widget.NewLabel("")

	clientForm.OnSubmit = func() {
		appEmail := "admin@abc.com"
		appPass := "12345678"
		if emailEntry.Text == appEmail && passwordEntry.Text == appPass {
			ShowDashbord(myApp)
		} else {
			dialog.NewInformation("Invalid", "Email or Passwor invalid", myWindow).Show()
		}

	}

	myWindow.SetContent(
		container.NewVBox(clientForm, messageLabel),
	)
	myWindow.ShowAndRun()
}
