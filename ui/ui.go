package ui

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"mud.kristech.io/core/obj/mob/attackable"
)

var lastFocusShift time.Time = time.Now()

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

	ui.Contents.info = newImage(Merchant) // Map & Combat
	ui.Contents.pages = newTextView("Pages | Shop, Inventory, Quests, etc.")

	ui.Contents.location = newTextView(fmt.Sprintf("%s - %s", "go-mud", "Lobby"))
	ui.Contents.tabs = newTextView("Tabs").SetTextAlign(tview.AlignCenter)

	ui.Contents.stats = newTextView("Money, Health, and Stats")
	ui.Contents.chat = newTextView("Chat")

	ui.Contents.quit = tview.NewButton("Q to Quit")
	ui.Contents.input = tview.NewInputField().SetPlaceholder("Press Enter to chat...")

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
func Render(player *attackable.Player, ui *UI, quit *chan bool) {
	ui.SetInputCapture(player, quit)

	ui.App.SetFocus(ui.Contents.input)

	if err := ui.App.EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
