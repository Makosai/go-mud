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
	ui.Contents.title = newTextView(fmt.Sprintf("%s - %s", "go-mud", "Lobby"))
	ui.Contents.titleCombo = newTextView(fmt.Sprintf("[SWAP] %s - %s", "go-mud", "Lobby"))
	ui.Contents.tabs = newTextView("Tabs").SetTextAlign(tview.AlignCenter)
	ui.Contents.progress = newTextView("Money")

	ui.Contents.chatBox = newTextView("Chat")
	ui.Contents.menu = newTextView("Inventory & Commands")
	ui.Contents.graphics = newImage("images/merchant.jpg")
	ui.Contents.comboBox = newTextView("ComboBox | Chat or Graphics + Inventory & Commands")
	ui.Contents.info = newTextView("Stats & Shopping Info")

	ui.Contents.chatInput = newTextView("Press Enter to chat...")

	ui.Contents.quitBtn = tview.NewButton("Q to Quit")

	grid := tview.NewGrid().
		SetRows(1, 0, 0, 0, 1).
		SetColumns(40, 0, 30).
		SetBorders(true)

	// Small Screen Support
	grid.
		AddItem(ui.Contents.titleCombo, 0, 0, 1, 1, 0, 0, false).
		AddItem(ui.Contents.comboBox, 1, 0, 3, 2, 0, 0, false)

	// Topbar
	grid.
		AddItem(ui.Contents.title, 0, 0, 1, 1, 20, 100, false).
		AddItem(ui.Contents.tabs, 0, 1, 1, 1, 0, 0, false).
		AddItem(ui.Contents.progress, 0, 2, 1, 1, 0, 0, false)

	// Center
	grid.
		AddItem(ui.Contents.chatBox, 1, 1, 3, 1, 20, 100, false).
		AddItem(ui.Contents.graphics, 1, 0, 1, 1, 20, 100, false).
		AddItem(ui.Contents.menu, 2, 0, 2, 1, 20, 100, false).
		AddItem(ui.Contents.info, 1, 2, 3, 1, 0, 0, false)

	// Bottombar
	grid.AddItem(ui.Contents.chatInput, 4, 0, 1, 3, 0, 0, false)

	grid.AddItem(ui.Contents.quitBtn, 4, 2, 1, 1, 0, 0, false)

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
