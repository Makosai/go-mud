package attackable

type Stats struct {
	Attack   int16
	Strength int16
	Defense  int16
}

func NewStats(attack, strength, defense int16) *Stats {
	return &Stats{
		Attack:   attack,
		Strength: strength,
		Defense:  defense,
	}
}

type PlayerStats struct {
	Experience  int
	Mining      int16
	Crafting    int16
	Woodcutting int16
	Endurance   int16
}

func NewPlayerStats(
	experience int,
	mining,
	crafting,
	woodcutting,
	endurance int16,
) *PlayerStats {
	return &PlayerStats{
		Experience:  experience,
		Mining:      mining,
		Crafting:    crafting,
		Woodcutting: woodcutting,
		Endurance:   endurance,
	}
}
