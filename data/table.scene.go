package data

type SceneTable struct {
	Table   *Table
	Cols    Columns
	FormUNQ string
}

func (t *Table) NewScene(c Columns) *SceneTable {
	return &SceneTable{
		Table: t,
		Cols:  c,
	}
}
func (t *Table) NewSceneByList(c Columns) *SceneTable {
	return &SceneTable{
		Table:   t,
		Cols:    c,
		FormUNQ: t.UNQ,
	}
}
