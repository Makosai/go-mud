package main

import (
	"mud.kristech.io/engine"
)

func main() {
	client := engine.NewClient(0)
	client.Run()
}
