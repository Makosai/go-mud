package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// UI struct containing the app and all UI elements so they can be accessed from the main thread
type UI struct {
	App      *tview.Application
	Contents *AppContents
}

func NewUI() *UI {
	return &UI{
		App: NewApp(),
	}
}

// Creates a tview application
func NewApp() *tview.Application {
	app := tview.NewApplication().
		SetBeforeDrawFunc(func(screen tcell.Screen) bool {
			screen.Clear()
			return false
		})

	return app
}

// Creates the layout for the UI
func NewLayout(ui *UI) *tview.Grid {
	ui.Contents = new(AppContents)

	// The title uses a global variable
	ui.Contents.info = newTextView("Map & Combat")
	ui.Contents.pages = newTextView("Pages | Shop, Inventory, Quests, etc.")

	ui.Contents.location = newTextView(fmt.Sprintf("%s - %s", "go-mud", "Lobby"))
	ui.Contents.tabs = newTextView("Tabs").SetTextAlign(tview.AlignCenter)

	ui.Contents.stats = newTextView("Money, Health, and Stats")
	ui.Contents.chat = newTextView("Chat")

	ui.Contents.quit = tview.NewButton("Q to Quit")
	ui.Contents.input = newTextView("Press Enter to chat...")

	grid := tview.NewGrid().
		SetRows(0, 1, 0, 1).
		SetColumns(0, 0, 0).
		SetBorders(true)

	// Top
	grid.
		AddItem(ui.Contents.info, 0, 0, 1, 1, 0, 0, false).
		AddItem(ui.Contents.pages, 0, 1, 1, 2, 0, 0, false)

	// Middle
	grid.
		AddItem(ui.Contents.location, 1, 0, 1, 1, 0, 0, false).
		AddItem(ui.Contents.tabs, 1, 1, 1, 2, 0, 0, false)

	// Bottom
	grid.
		AddItem(ui.Contents.stats, 2, 0, 1, 1, 0, 0, false).
		AddItem(ui.Contents.chat, 2, 1, 1, 2, 0, 0, false)

	// Footer
	grid.
		AddItem(ui.Contents.quit, 3, 0, 1, 1, 0, 0, false).
		AddItem(ui.Contents.input, 3, 1, 1, 2, 0, 0, false)

	return grid
}

// Starts the UI
func Render(app *tview.Application, quit *chan bool) {
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

	if err := app.EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
