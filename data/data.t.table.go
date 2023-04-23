package data

import (
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
)

type Tables []*Table

func NewTables(tbs md.Tables) Tables {
	nt := make([]*Table, 0, 1)
	for _, t := range tbs {
		nt = append(nt, NewTable(t))
	}
	return nt
}
func fltrSearchUITable(t *Table, name *optrs) *TTable {
	return fltrSearchTable(t, name.URL)
}
func fltrBuildTable(t interface{}) *TTable {
	if v, ok := t.(*TTable); ok {
		return v
	}
	return &TTable{Table: t.(*Table), Current: t.(*Table), IsTmpl: true}
}
func fltrSearchUITableAndResetUNQ(tb *Table, name *optrs) *TTable {
	t := fltrSearchTable(tb, name.URL)
	t.UNQ = name.UNQ
	return t
}
func findTable(uname string) *Table {
	if strings.Contains(uname, "/") {
		uname = strings.Replace(strings.Trim(uname, "/"), "/", "_", -1)
	}
	return Get(uname)
}

func fltrSearchTable(t *Table, uname string) *TTable {
	ut := findTable(uname)
	return &TTable{Table: ut, Current: t, IsTmpl: true}
}

type TTable struct {
	*Table
	Current *Table
	IsTmpl  bool
}
