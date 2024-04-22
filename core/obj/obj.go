package obj

// An Obj that contains an ID, a name, and a description.
type Obj struct {
	Name        string
	Description string
}

func NewObj(name, description string) Obj {
	return Obj{
		Name:        name,
		Description: description,
	}
}
