package data

type Scene struct {
	Columns Columns
	UNQ     string
}

func (t *Table) NewScene(c Columns) *Scene {
	return &Scene{
		Columns: c,
		UNQ:     t.UNQ,
	}
}
func (c Columns) NewScene(unq string) *Scene {
	return &Scene{
		Columns: c,
		UNQ:     unq,
	}
}
func (t *optrs) NewScene(c Columns) *Scene {
	return &Scene{
		Columns: c,
		UNQ:     t.UNQ,
	}
}
