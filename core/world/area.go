package world

import "mud.kristech.io/internals"

type Area struct {
	name        string
	description string
	pos         internals.Vector2
}

// NewArea
func NewArea(name, description string, pos internals.Vector2) *Area {
	return &Area{
		name:        name,
		description: description,
		pos:         pos,
	}
}
