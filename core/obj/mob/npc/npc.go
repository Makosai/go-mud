package npc

import "mud.kristech.io/core/obj/mob"

// Npc is a non-player character.
type Npc struct {
	mob.Mob
}

func NewNpc(name, description string) Npc {
	return Npc{
		Mob: mob.NewMob(name, description),
	}
}
