package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	list := tview.NewList().
		AddItem("Thing 1", "The first thing", 'a', nil).
		AddItem("Thing 2", "The second thing", 'b', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	if err := app.SetRoot(list, true).Run(); err != nil {
		panic(err)
	}
}
