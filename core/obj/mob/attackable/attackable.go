package attackable

import "mud.kristech.io/core/obj/mob"

// An attackable mob in the game that has Health and Stats.
type Attackable struct {
	mob.Mob
	Health int
	Stats  *Stats
}

func NewAttackable(
	name, description string,
	health int,
	stats *Stats,
) Attackable {
	return Attackable{
		Mob:    mob.NewMob(name, description),
		Health: health,
		Stats:  stats,
	}
}

func (a *Attackable) TakeDamage(damage int) {
	a.Health -= damage

	if a.Health < 0 {
		a.Health = 0
	}
}

func (a *Attackable) IsDead() bool {
	return a.Health <= 0
}

func (a *Attackable) Attack(target *Attackable) {
	target.TakeDamage(int(a.Stats.Attack - target.Stats.Defense))
}

func (a *Attackable) Heal(amount int) {
	a.Health += amount
}
