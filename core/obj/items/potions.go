package items

import (
	"strconv"

	"mud.kristech.io/core/obj"
)

// A potion that can be consumed by a player.
type HealthPotion struct {
	Consumable
	Health int
}

type SmallHealthPotion struct {
	HealthPotion
}

func NewSmallHealthPotion() *SmallHealthPotion {
	name := "Small Health Potion"
	description := "A potion that restores a small amount of health."
	health := 10

	return &SmallHealthPotion{
		HealthPotion: HealthPotion{
			Consumable: Consumable{
				Item: Item{
					Obj: obj.Obj{
						Name:        name,
						Description: description,
					},
				},
			},
			Health: health,
		},
	}
}

// Use consumes the potion.
func (potion *HealthPotion) Use() {
	println("Drank " + potion.Name + ".")
	println("Restored " + strconv.Itoa(potion.Health) + " health.")
}
