package ui

import (
	"bytes"
	_ "embed"
	"image"
	"image/jpeg"

	"github.com/rivo/tview"
)

func newTextView(text string) *tview.TextView {
	return tview.NewTextView().
		SetText(text).
		SetDynamicColors(true)
}

func getPhoto(path Images) image.Image {
	data, _ := images.ReadFile(string(path))
	photo, _ := jpeg.Decode(bytes.NewReader(data))

	return photo
}

func newImage(path Images) *tview.Image {
	image := tview.NewImage()
	image.SetImage(getPhoto(path))

	return image

}
