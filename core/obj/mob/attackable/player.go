package attackable

import "mud.kristech.io/internals"

// A Player is an Attackable that can attack an Enemy and be attacked by an Enemy.
type Player struct {
	Attackable
	PlayerStats *PlayerStats
	Location    *Location
}

func NewPlayer(
	name string,
	health int,
	stats *Stats,
	playerStats *PlayerStats,
) *Player {
	return &Player{
		Attackable: NewAttackable(
			name, "A thing of flesh behind a screen.",
			health,
			stats,
		),
		PlayerStats: playerStats,
		Location: &Location{
			AreaName: "Crescent Fall",
			position: &internals.Vector2{X: 0, Y: 1},
		},
	}
}

type Location struct {
	AreaName string
	position *internals.Vector2
}
