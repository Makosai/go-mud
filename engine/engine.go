package engine

import "time"

type Client struct {
	id   int
	quit chan bool
}

func NewClient(id int) *Client {
	return &Client{
		id:   id,
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
