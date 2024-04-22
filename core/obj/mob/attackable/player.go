package attackable

// A Player is an Attackable that can attack an Enemy and be attacked by an Enemy.
type Player struct {
	Attackable
	PlayerStats *PlayerStats
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
	}
}
