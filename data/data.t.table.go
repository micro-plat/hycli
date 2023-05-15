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
func fltrSearchUITable(name *optrs) *TTable {
	return fltrSearchTable(name.URL)
}
func fltrSearchUITableAndResetUNQ(name *optrs) *TTable {
	t := fltrSearchTable(name.URL)
	t.UNQ = name.UNQ
	return t
}
func findTable(uname string) *Table {
	if strings.Contains(uname, "/") {
		uname = strings.Replace(strings.Trim(uname, "/"), "/", "_", -1)
	}
	return Get(uname)
}

func fltrSearchTable(uname string) *TTable {
	ut := findTable(uname)
	return &TTable{Table: ut, IsTmpl: true}
}

type TTable struct {
	*Table
	IsTmpl bool
}
