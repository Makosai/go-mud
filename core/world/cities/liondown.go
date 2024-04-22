package city

import "mud.kristech.io/internals"

// Liondown is a city in the world.
type Liondown struct {
	*City
}

func NewLiondown() *Liondown {
	return &Liondown{
		City: NewCity("Liondown", "A city in the world.", internals.Vector2{X: 0, Y: 0}),
	}
}
