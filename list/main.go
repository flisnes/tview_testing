package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	list := tview.NewList()

	list.
		AddItem("Thing 1", "The first thing", 'a', func() {
			showTextWindow(app, list, "This text will be displayed in the new window.")
		}).
		AddItem("Thing 2", "The second thing", 'b', func() {
			showInputWindow(app, list, 1)
		}).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	if err := app.SetRoot(list, true).Run(); err != nil {
		panic(err)
	}
}

func showTextWindow(app *tview.Application, list *tview.List, text string) {
	modal := tview.NewModal().
		SetText(text).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				app.SetRoot(list, true) // Return to the list
			}
		})

	app.SetRoot(modal, true)
}

func showInputWindow(app *tview.Application, list *tview.List, index int) {
	inputField := tview.NewInputField().
		SetLabel("Enter new description: ").
		SetFieldWidth(30)

	form := tview.NewForm().
		AddFormItem(inputField).
		AddButton("Submit", func() {
			newText := inputField.GetText()
			list.SetItemText(index, "Thing 2", newText)
			app.SetRoot(list, true)
		}).
		AddButton("Cancel", func() {
			app.SetRoot(list, true)
		})

	form.SetBorder(true).SetTitle("Input Form").SetTitleAlign(tview.AlignCenter)

	app.SetRoot(form, true)
}
