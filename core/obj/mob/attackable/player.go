package attackable

import (
	"mud.kristech.io/core/obj"
	"mud.kristech.io/core/obj/mob"
)

type Player struct {
	Attackable
}

func NewPlayer(name string) *Player {
	return &Player{
		Attackable: Attackable{
			Mob: mob.Mob{
				Obj: obj.Obj{
					Name: name,
				},
			},
		},
	}
}
