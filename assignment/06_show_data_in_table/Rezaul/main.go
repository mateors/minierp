package main

import (
	"database/sql"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

var (
	myApp    fyne.App    = app.New()
	myWindow fyne.Window = myApp.NewWindow("miniERP v0.0.1")
	db       *sql.DB
	err      error
)

func init() {
	dbcon()
}

func main() {
	myWindow.Resize(fyne.NewSize(400, 200))
	myWindow.SetFixedSize(true)

	head := widget.NewLabel("Welcome to miniERP")
	head.Alignment = fyne.TextAlignCenter

	idEntry := widget.NewEntry()
	idEntry.PlaceHolder = "Enter your ID"

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.PlaceHolder = "Enter your password"

	id := widget.NewFormItem("ID", idEntry)
	password := widget.NewFormItem("Password", passwordEntry)

	loginForm := widget.NewForm(id, password)

	loginForm.SubmitText = "Login"
	loginForm.CancelText = "Quit"

	loginForm.OnCancel = func() {
		myWindow.Close()
	}

	loginForm.OnSubmit = func() {

		appEmail := ""
		appPass := ""
		if idEntry.Text == appEmail && passwordEntry.Text == appPass {
			Dashbord(myApp)
		} else {
			dialog.NewInformation("Warning !", "ID or Password invalid..", myWindow).Show()
		}

	}

	myWindow.SetContent(

		container.NewVBox(head, loginForm),
	)
	myWindow.ShowAndRun()
}
