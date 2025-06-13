package ui

import (
	"github.com/rivo/tview"
)

type AppContents struct {
	info  *tview.Image
	pages *tview.TextView

	location *tview.TextView
	tabs     *tview.TextView

	stats *tview.TextView
	chat  *tview.TextView

	quit  *tview.Button
	input *tview.InputField

	grid *tview.Grid
}

func (contents *AppContents) SetInfo(path Images) {
	contents.info.SetImage(getPhoto(path))
}
func (contents *AppContents) SetPages(pages string) {
	contents.pages.SetText(pages)
}

func (contents *AppContents) SetLocation(location string) {
	contents.location.SetText(location)
}
func (contents *AppContents) SetTabs(tabs string) {
	contents.tabs.SetText(tabs)
}

func (contents *AppContents) SetStats(stats string) {
	contents.stats.SetText(stats)
}
func (contents *AppContents) SetChat(chat string) {
	contents.chat.SetText(chat)
}
func (contents *AppContents) AppendChat(chat string) {
	contents.chat.SetText(contents.chat.GetText(false) + "\n" + chat)
}
func (contents *AppContents) SendChat(sender, chat string) {
	contents.AppendChat(sender + ": " + chat)
}

func (contents *AppContents) SetQuit(quit string) {
	contents.quit.SetLabel(quit)
}
func (contents *AppContents) SetInput(input string) {
	contents.input.SetText(input)
}
