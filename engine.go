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
	health := 100

	attack := int16(10)
	strength := int16(10)
	defense := int16(10)

	experience := 0
	mining := int16(1)
	crafting := int16(1)
	woodcutting := int16(1)
	endurance := int16(1)

	return &Client{
		id: 1,
		player: attackable.NewPlayer(
			name,
			health,
			attackable.NewStats(attack, strength, defense),
			attackable.NewPlayerStats(experience, mining, crafting, woodcutting, endurance),
		),
		quit: make(chan bool, 1),
	}
}

func NewClientRemote(id uint8, name string) *Client {
	if id == 0 {
		panic("Client ID cannot initially be 0.")
	}

	if id == 1 {
		panic("Client ID cannot be 1. It is reserved for the local client.")
	}

	health := 100

	attack := int16(10)
	strength := int16(10)
	defense := int16(10)

	experience := 0
	mining := int16(1)
	crafting := int16(1)
	woodcutting := int16(1)
	endurance := int16(1)

	return &Client{
		id: id,
		player: attackable.NewPlayer(
			name,
			health,
			attackable.NewStats(attack, strength, defense),
			attackable.NewPlayerStats(experience, mining, crafting, woodcutting, endurance),
		),
		quit: make(chan bool, 1),
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
