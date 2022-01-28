package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func processAllClientData() [][]string {
	tableData := [][]string{
		{"Id", "Name", "Mobile", "Email", "Address"},
	}
	rows := GetClient()
	for i := 0; i < len(rows); i++ {
		var tempRow []string
		tempRow = append(tempRow, fmt.Sprintf("%v", rows[i]["id"]))
		tempRow = append(tempRow, fmt.Sprintf("%v", rows[i]["name"]))
		tempRow = append(tempRow, fmt.Sprintf("%v", rows[i]["mobile"]))
		tempRow = append(tempRow, fmt.Sprintf("%v", rows[i]["email"]))
		tempRow = append(tempRow, fmt.Sprintf("%v", rows[i]["address"]))

		tableData = append(tableData, tempRow)
	}

	return tableData
}

func ShowData(a fyne.App) {
	win := myWindow
	btnHead := widget.NewButton("< Back to Dashbord", func() {
		Dashbord(myApp)
	})

	btnHead.Resize(fyne.NewSize(200, 40))
	btnHead.Move(fyne.NewPos(10, 0))

	data := processAllClientData()
	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	list.Resize(fyne.NewSize(1500, 800))
	list.Move(fyne.NewPos(10, 50))

	list.SetColumnWidth(0, 60.0)
	list.SetColumnWidth(1, 150.0)
	list.SetColumnWidth(3, 210.0)

	myWindow.SetContent(
		container.NewWithoutLayout(btnHead, list),
	)
	win.Show()
}
