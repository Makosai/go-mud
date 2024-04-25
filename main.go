package main

import (
	"mud.kristech.io/ui"
)

func main() {
	client := NewClient("Mako")

	go ui.Render(client.player.Location)

	client.Run()
}
