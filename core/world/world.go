package world

type World struct {
	// x and y coordinates of Cities, Forests, and Dungeons
	Area map[uint8]map[uint8]Area
}
