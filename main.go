package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var db *sql.DB
var err error

func init() {

	// Connect to database
	db, err = sql.Open("sqlite3", "/home/mostain/go/src/master_academy/minierp.db")
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(1)
	log.Println("db connection successful")

}

func main() {

	// addClient("RAIHAN", "01672710028", "raihan@gmail.com", "House# 01, Road# 02")
	// os.Exit(1)

	defer db.Close()
	myApp := app.New()
	myWindow := myApp.NewWindow("My Form")

	myWindow.Resize(fyne.NewSize(600, 400))

	name := widget.NewEntry()
	name.PlaceHolder = "Enter customer name"

	mobile := widget.NewEntry()
	mobile.PlaceHolder = "Enter customer number"

	email := widget.NewEntry()
	email.PlaceHolder = "Enter customer email"

	address := widget.NewMultiLineEntry()
	address.PlaceHolder = "Type customer address"

	row1 := widget.NewFormItem("Name", name)
	row2 := widget.NewFormItem("Mobile", mobile)
	row3 := widget.NewFormItem("Email", email)
	row4 := widget.NewFormItem("Address", address)

	wform := widget.NewForm(row1, row2, row3, row4)
	wform.SubmitText = "Save"

	wform.OnSubmit = func() {

		name := strings.ToUpper(name.Text)
		phone := mobile.Text
		emailID := strings.ToLower(email.Text)
		address := address.Text

		//myData := fmt.Sprintf(`%s %v %s %s`, name, phone, emailID, address)
		id, err := addClient(name, phone, emailID, address)
		if err != nil {
			log.Println(err)
			return
		}
		myData := fmt.Sprintf(`Client added New client ID # %d`, id)
		dialog.NewInformation("Confirmation", myData, myWindow).Show()

	}
	wform.OnCancel = func() {}

	//labelName := widget.NewLabel("Name")
	//inputName := widget.NewEntry()
	//fContainer := container.NewBorder(widget.NewLabel("My Form"), widget.NewLabel("Footer"), labelName, nil, inputName)

	myWindow.SetContent(wform)
	myWindow.ShowAndRun()

}
