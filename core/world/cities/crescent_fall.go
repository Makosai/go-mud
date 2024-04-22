package city

import "mud.kristech.io/internals"

// Liondown is a city in the world.
type CrescentFall struct {
	*City
}

// NewLiondown creates a new city.
func NewCrescentFall() *CrescentFall {
	return &CrescentFall{
		City: NewCity("Crescent Fall", "A city in the world.", internals.Vector2{X: 0, Y: 1}),
	}
}
