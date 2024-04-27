package ui

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"mud.kristech.io/core/obj/mob/attackable"
)

type Element int

const (
	Input Element = iota
)

func (ui *UI) focusShift(element Element) {
	switch element {
	case Input:
		if !ui.Contents.input.HasFocus() {
			lastFocusShift = time.Now()
			ui.App.SetFocus(ui.Contents.input)
		}
	}
}

func (ui *UI) FocusInput() {
	ui.focusShift(Input)
}

// Handles the Input Element Commands
func handleInputCommands(ui *UI, player *attackable.Player) {
	ui.Contents.input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Anything handled here will be executed on the main thread
		switch event.Key() {
		case tcell.KeyEnter:
			// Wait a bit after focusing to prevent sending unintentionally
			if !event.When().After(lastFocusShift.Add(time.Second / 4)) {
				break
			}

			// Send the chat message
			if ui.Contents.input.GetText() != "" {
				ui.Contents.SendChat(player.Name, ui.Contents.input.GetText())
				ui.Contents.input.SetText("")
			}

		// Blur / Unfocus the input field
		case tcell.KeyESC:
			ui.Contents.input.Blur()
		}

		return event
	})
}

// Handles the Root App Commands
func handleAppCommands(ui *UI, quit *chan bool) {
	ui.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Anything handled here will be executed on the main thread
		switch event.Key() {
		// Focus the input field
		case tcell.KeyEnter, tcell.KeyBackspace, tcell.KeyDelete, tcell.KeyDEL:
			ui.FocusInput()

		case tcell.KeyRune:
			switch event.Rune() {
			// Quit the application
			case 'Q', 'q':
				*quit <- true
				ui.App.Stop()
			}
		}

		return event
	})
}

// Sets up the Input Capture for the UI
func (ui *UI) SetInputCapture(player *attackable.Player, quit *chan bool) {
	handleInputCommands(ui, player)
	handleAppCommands(ui, quit)
}
