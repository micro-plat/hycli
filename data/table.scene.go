package data

type Scene struct {
	Columns Columns
	FormUNQ string
}

func (t *Table) NewScene(c Columns) *Scene {
	return &Scene{
		Columns: c,
		FormUNQ: t.UNQ,
	}
}
func (t *Table) NewSceneByList(c Columns) *Scene {
	return &Scene{
		Columns: c,
		FormUNQ: t.UNQ,
	}
}
