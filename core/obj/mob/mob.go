package mob

import "mud.kristech.io/core/obj"

// A movable object in the game.
type Mob struct {
	obj.Obj
}

func NewMob(name, description string) Mob {
	return Mob{
		Obj: obj.NewObj(name, description),
	}
}
