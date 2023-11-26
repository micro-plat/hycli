package data

type SceneTable struct {
	Table   *Table
	Columns Columns
	FormUNQ string
}

func (t *Table) NewScene(c Columns) *SceneTable {
	return &SceneTable{
		Table:   t,
		Columns: c,
	}
}
func (t *Table) NewSceneByList(c Columns) *SceneTable {
	return &SceneTable{
		Table:   t,
		Columns: c,
		FormUNQ: t.UNQ,
	}
}
