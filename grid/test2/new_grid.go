package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// Create the grid
	grid := tview.NewGrid().
		SetRows(0, 0, 0).
		SetColumns(0, 0, 0).
		SetBorders(true)

	// Create lists for each grid cell
	lists := make([]*tview.List, 9)
	for i := range lists {
		lists[i] = tview.NewList().
			AddItem("Item 1", "Description 1", 'a', nil).
			AddItem("Item 2", "Description 2", 'b', nil).
			AddItem("Item 3", "Description 3", 'c', nil)
	}

	// Add lists to the grid
	grid.
		AddItem(lists[0], 0, 0, 1, 1, 0, 0, false).
		AddItem(lists[1], 0, 1, 1, 1, 0, 0, false).
		AddItem(lists[2], 0, 2, 1, 1, 0, 0, false).
		AddItem(lists[3], 1, 0, 1, 1, 0, 0, false).
		AddItem(lists[4], 1, 1, 1, 1, 0, 0, false).
		AddItem(lists[5], 1, 2, 1, 1, 0, 0, false).
		AddItem(lists[6], 2, 0, 1, 1, 0, 0, false).
		AddItem(lists[7], 2, 1, 1, 1, 0, 0, false).
		AddItem(lists[8], 2, 2, 1, 1, 0, 0, false)

	// Track the currently focused cell
	currentRow, currentCol := 0, 0

	// Function to update the grid's focus
	updateFocus := func() {
		for row := 0; row < 3; row++ {
			for col := 0; col < 3; col++ {
				cell := lists[row*3+col]
				if row == currentRow && col == currentCol {
					cell.SetBorder(true).SetTitle("Focused")
				} else {
					cell.SetBorder(false).SetTitle("")
				}
			}
		}
		app.SetFocus(grid)
	}

	// Initial focus on the grid
	updateFocus()

	// Keybindings to navigate the grid
	grid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:
			if currentRow > 0 {
				currentRow--
			}
		case tcell.KeyDown:
			if currentRow < 2 {
				currentRow++
			}
		case tcell.KeyLeft:
			if currentCol > 0 {
				currentCol--
			}
		case tcell.KeyRight:
			if currentCol < 2 {
				currentCol++
			}
		case tcell.KeyEnter:
			app.SetFocus(lists[currentRow*3+currentCol])
			return nil
		}
		updateFocus()
		return event
	})

	// Keybinding to move focus back to the grid from the list
	for _, list := range lists {
		list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Rune() == 'q' {
				updateFocus()
				return nil
			}
			return event
		})
	}

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
