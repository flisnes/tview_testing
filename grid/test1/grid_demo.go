package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// Create a grid
	grid := tview.NewGrid().
		SetRows(0, 0, 0).
		SetColumns(0, 0, 0).
		SetBorders(true)

	// Create a 3x3 grid of lists
	var lists [3][3]*tview.List
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			list := tview.NewList().
				AddItem("Item 1", "", 0, nil).
				AddItem("Item 2", "", 0, nil).
				AddItem("Item 3", "", 0, nil)
			lists[row][col] = list
			grid.AddItem(list, row, col, 1, 1, 0, 0, false)
		}
	}

	// Variables to keep track of the current row and column
	currentRow, currentCol := 0, 0

	// Set initial focus
	app.SetFocus(lists[currentRow][currentCol])

	// Set input capture for the grid
	grid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			// When Enter is pressed, set focus to the current list
			app.SetFocus(lists[currentRow][currentCol])
		case tcell.KeyUp:
			if currentRow > 0 {
				currentRow--
				app.SetFocus(grid)
			}
		case tcell.KeyDown:
			if currentRow < 2 {
				currentRow++
				app.SetFocus(grid)
			}
		case tcell.KeyLeft:
			if currentCol > 0 {
				currentCol--
				app.SetFocus(grid)
			}
		case tcell.KeyRight:
			if currentCol < 2 {
				currentCol++
				app.SetFocus(grid)
			}
		}
		return event
	})

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
