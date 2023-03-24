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
	// uname := name.URL
	// if strings.Contains(name.URL, "/") {
	// 	uname = strings.Replace(strings.Trim(name.URL, "/"), "/", "_", -1)
	// }
	// ut := Get(uname)
	// return &TTable{Table: ut, IsTmpl: true}
	return fltrSearchTable(name.URL)
}
func fltrSearchTable(uname string) *TTable {
	if strings.Contains(uname, "/") {
		uname = strings.Replace(strings.Trim(uname, "/"), "/", "_", -1)
	}
	ut := Get(uname)
	return &TTable{Table: ut, IsTmpl: true}
}

type TTable struct {
	*Table
	IsTmpl bool
}
