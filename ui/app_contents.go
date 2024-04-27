package ui

import (
	"fmt"

	"github.com/rivo/tview"
)

type AppContents struct {
	title      *tview.TextView
	titleCombo *tview.TextView
	tabs       *tview.TextView
	progress   *tview.TextView

	chatBox   *tview.TextView
	menu      *tview.TextView
	graphics  *tview.Image
	container *tview.TextView
	info      *tview.TextView

	chatInput *tview.TextView
	quitBtn   *tview.Button
}

func (contents *AppContents) UpdateTitle(title string) {
	contents.title.SetText(title)
	contents.titleCombo.SetText(fmt.Sprintf("[SWAP] %s", title))
}
