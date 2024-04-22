package inventory

import (
	"mud.kristech.io/core/obj/items"
)

// Inventory is a collection of items.
type Inventory struct {
	ID    int
	Items [16]items.Item
}
