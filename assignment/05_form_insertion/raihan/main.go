package main

import (
	"database/sql"
	"fmt"
	"image/color"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var err error

func init() {

	// Connect to database
	db, err = sql.Open("sqlite3", "./minierp.db")
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(1)
	log.Println("db connection successful")

}

func main() {
	defer db.Close()
	App := app.New()
	Window := App.NewWindow("My Form")

	Window.Resize(fyne.NewSize(600, 400))

	title := canvas.NewText("Welcome to miniERP", color.White)
	title.TextSize = 20
	title.Alignment = fyne.TextAlignCenter

	name := widget.NewEntry()
	name.PlaceHolder = "Enter Client name"
	mobile := widget.NewEntry()
	mobile.PlaceHolder = "Enter Client Mobile Number"
	email := widget.NewEntry()
	email.PlaceHolder = "Enter Client Email Address"
	address := widget.NewMultiLineEntry()
	address.PlaceHolder = "Type your Client address"

	row1 := widget.NewFormItem("Name", name)
	row2 := widget.NewFormItem("mobile", mobile)
	row3 := widget.NewFormItem("email", email)
	row4 := widget.NewFormItem("address", address)

	// save item //
	wform := widget.NewForm(row1, row2, row3, row4)
	wform.SubmitText = "Save"

	wform.OnSubmit = func() {
		name := strings.ToUpper(name.Text)
		phone := mobile.Text
		emailID := strings.ToLower(email.Text)
		address := address.Text

		id, err := AddClient(name, phone, emailID, address)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(id)

	}

	wform.OnCancel = func() {}

	Window.SetContent(container.NewVBox(
		title,
		wform,
	))
	Window.ShowAndRun()

}
