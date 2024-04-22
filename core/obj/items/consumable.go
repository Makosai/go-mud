package items

// A consumable is an object that can be consumed by a player.
type Consumable struct {
	Item
}

// Interface to use a consumable.
type Useable interface {
	Use()
}

// Use consumes the consumable.
func (c *Consumable) Use() {
	println("Consumed " + c.Name + ".")
}
