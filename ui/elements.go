package ui

import (
	"bytes"
	_ "embed"
	"image"
	"image/jpeg"

	"github.com/rivo/tview"
)

func newTextView(text string) *tview.TextView {
	title := tview.NewTextView().
		SetText(text).
		SetDynamicColors(true)

	return title
}

func getPhoto(path string) image.Image {
	data, _ := images.ReadFile(path)
	photo, _ := jpeg.Decode(bytes.NewReader(data))

	return photo
}

func newImage(path string) *tview.Image {
	image := tview.NewImage()
	image.SetImage(getPhoto(path))

	return image

}
