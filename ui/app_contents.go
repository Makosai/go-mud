package ui

import (
	"github.com/rivo/tview"
)

type AppContents struct {
	info     *tview.TextView
	pages    *tview.TextView
	location *tview.TextView

	tabs  *tview.TextView
	stats *tview.TextView
	chat  *tview.TextView

	quit  *tview.Button
	input *tview.TextView
}

func (contents *AppContents) UpdateTitle(title string) {
	contents.location.SetText(title)
}
