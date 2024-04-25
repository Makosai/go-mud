package ui

import (
	"fmt"

	"github.com/rivo/tview"
	"mud.kristech.io/core/obj/mob/attackable"
)

func newPrimitive(text string) tview.Primitive {
	return tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(text)
}

func Render(location *attackable.Location) {
	// The title uses a global variable
	title := newPrimitive(fmt.Sprintf("go-mud - %s", location.AreaName))
	tabs := newPrimitive("Tabs")
	progress := newPrimitive("Money")

	chatBox := newPrimitive("Chat")
	menu := newPrimitive("Inventory & Commands")
	graphics := newPrimitive("Graphics")
	info := newPrimitive("Stats & Shopping Info")

	chatInput := newPrimitive("Press Enter to chat...")

	app := tview.NewApplication()

	grid := tview.NewGrid().
		SetRows(1, 0, 0, 0, 1).
		SetColumns(40, 0, 30).
		SetBorders(true)

	// Topbar
	grid.
		AddItem(title, 0, 0, 1, 1, 0, 0, false).
		AddItem(tabs, 0, 1, 1, 1, 0, 0, false).
		AddItem(progress, 0, 2, 1, 1, 0, 0, false)

	// Center
	grid.
		AddItem(chatBox, 1, 1, 3, 1, 0, 0, false).
		AddItem(graphics, 1, 0, 1, 1, 0, 0, false).
		AddItem(menu, 2, 0, 2, 1, 0, 0, false).
		AddItem(info, 1, 2, 3, 1, 0, 0, false)

	// Bottombar
	grid.AddItem(chatInput, 4, 0, 1, 3, 0, 0, false)

	grid.AddItem(tview.NewButton("Quit").SetSelectedFunc(func() { app.Stop() }), 4, 2, 1, 1, 0, 0, false)

	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
