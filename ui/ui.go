package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"mud.kristech.io/core/obj/mob/attackable"
)

func NewApp() *tview.Application {
	app := tview.NewApplication()

	return app
}

func newPrimitive(text string) tview.Primitive {
	return tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(text)
}

func Render(app *tview.Application, location *attackable.Location, quit *chan bool) {
	// The title uses a global variable
	title := newPrimitive(fmt.Sprintf("go-mud - %s", location.AreaName))
	titleCombo := newPrimitive(fmt.Sprintf("[SWAP COMBO BOX] go-mud - %s", location.AreaName))
	tabs := newPrimitive("Tabs")
	progress := newPrimitive("Money")

	chatBox := newPrimitive("Chat")
	menu := newPrimitive("Inventory & Commands")
	graphics := newPrimitive("Graphics")
	comboBox := newPrimitive("ComboBox | Chat or Graphics + Inventory & Commands")
	info := newPrimitive("Stats & Shopping Info")

	chatInput := newPrimitive("Press Enter to chat...")

	grid := tview.NewGrid().
		SetRows(1, 0, 0, 0, 1).
		SetColumns(40, 0, 30).
		SetBorders(true)

	// Small Screen Support
	grid.
		AddItem(titleCombo, 0, 0, 1, 1, 0, 0, false).
		AddItem(comboBox, 1, 0, 3, 2, 0, 0, false)

	// Topbar
	grid.
		AddItem(title, 0, 0, 1, 1, 20, 100, false).
		AddItem(tabs, 0, 1, 1, 1, 0, 0, false).
		AddItem(progress, 0, 2, 1, 1, 0, 0, false)

	// Center
	grid.
		AddItem(chatBox, 1, 1, 3, 1, 20, 100, false).
		AddItem(graphics, 1, 0, 1, 1, 20, 100, false).
		AddItem(menu, 2, 0, 2, 1, 20, 100, false).
		AddItem(info, 1, 2, 3, 1, 0, 0, false)

	// Bottombar
	grid.AddItem(chatInput, 4, 0, 1, 3, 0, 0, false)

	grid.AddItem(tview.NewButton("Q to Quit").SetSelectedFunc(func() { app.Stop() }), 4, 2, 1, 1, 0, 0, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Anything handled here will be executed on the main thread
		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case 'Q', 'q':
				*quit <- true
				app.Stop()
			}
		}

		return event
	})

	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
