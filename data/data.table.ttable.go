package data

import (
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

type Tables []*Table

func NewTables(tbs md.Tables) Tables {
	nt := make([]*Table, 0, 1)
	for _, t := range tbs {
		nt = append(nt, NewTable(t))
	}

	return nt
}
func (name *optrs) GetAssociatedTable(v ...bool) *TTable {
	ut := f_table_find_by_name(name.URL)
	if types.GetBoolByIndex(v, 0, false) {
		ut.UNQ = name.UNQ
	}
	return &TTable{Table: ut, IsTmpl: true}
}
func f_table_find_by_name(uname string) *Table {
	if strings.Contains(uname, "/") {
		uname = strings.Replace(strings.Trim(uname, "/"), "/", "_", -1)
	}
	return Get(uname)
}

type TTable struct {
	*Table
	IsTmpl bool
}
