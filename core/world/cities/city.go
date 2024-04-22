package city

import (
	"mud.kristech.io/core/world"
	"mud.kristech.io/internals"
)

type City struct {
	*world.Area
}

// NewCity creates a new city.
func NewCity(name, description string, pos internals.Vector2) *City {
	return &City{
		Area: world.NewArea(name, description, pos),
	}
}
