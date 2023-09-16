package data

type BUITable struct {
	Current interface{}
	Main    interface{}
	IsTmpl  bool
}

func (t *Table) Contact(x interface{}) *BUITable {
	return &BUITable{
		Current: t,
		Main:    x,
		IsTmpl:  t != x,
	}
}
