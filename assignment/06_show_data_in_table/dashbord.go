package main

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Dashbord(a fyne.App) {
	w := myWindow
	myWindow.Resize(fyne.NewSize(900, 400))

	headtext := canvas.NewText("Welcome to miniERP", color.White)
	headtext.TextSize = 20
	headtext.Alignment = fyne.TextAlignCenter

	btn1 := widget.NewButton("Create New Client", func() {
		showClient(myApp)
	})
	btn2 := widget.NewButton("All Client", func() {
		ShowData(myApp)
	})

	totalClient := processAllClientData()
	totalClientCount := strconv.Itoa((len(totalClient) - 1))

	card1 := widget.NewCard(
		"Total Client",
		totalClientCount,
		canvas.NewRectangle(color.Black),
	)

	w.SetContent(
		container.NewVBox(
			container.NewGridWithColumns(1,
				container.NewGridWithColumns(1,
					container.NewCenter(headtext),
				),
				container.NewGridWithColumns(3,
					container.NewVBox(btn1, btn2),
					container.NewVBox(card1),
				), container.NewGridWithColumns(3,
					container.NewVBox(),
				),
			),
		),
	)

	w.Show()
}
