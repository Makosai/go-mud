package main

import (
	"time"

	"mud.kristech.io/core/obj/mob/attackable"
)

type Client struct {
	id     uint8
	player *attackable.Player
	quit   chan bool
}

func NewClient(name string) *Client {
	return &Client{
		id:     1,
		player: attackable.NewPlayer(name),
		quit:   make(chan bool, 1),
	}
}

func NewClientRemote(id uint8, name string) *Client {
	if id == 0 {
		panic("Client ID cannot initially be 0.")
	}

	if id == 1 {
		panic("Client ID cannot be 1. It is reserved for the local client.")
	}

	return &Client{
		id:     id,
		player: attackable.NewPlayer(name),
		quit:   make(chan bool, 1),
	}
}

func (c *Client) Run() {
	tick := time.Tick(16 * time.Millisecond)

Engine:
	for range tick {
		println("boom")
		select {
		case <-c.quit:
			println("pow!")
			break Engine

		case <-tick:
			println("tick!")

		default:
			println("bang!")
		}
	}

	if c.id == 0 {
		println("Goodbye, my c.quit lover!")
	} else {
		println("Client", c.id, "quit.")
	}
}

func (c *Client) Call(opt string) {
	switch opt {
	case "quit":
		c.quit <- true
	}
}
