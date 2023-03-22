package data

import (
	"strings"

	"github.com/micro-plat/lib4go/types"
)

var Funcs = map[string]interface{}{

	"flterMainTable":       flterMainTable,
	"fltrNotNullCols":      fltrNotNullRows,
	"getFirstTable":        getFirstTable,
	"IsTmplTb":             IsTmplTb,
	"fltrSearchUITable":    fltrSearchUITable,
	"fltrOptrs":            fltrOptrs,
	"fltrColums":           fltrColums,
	"fltrColumsExcludeExt": fltrColumsExcludeExt,
	"fltrOptrPara":         fltrOptrPara,
	"resetForm":            resetForm,
	"multiply":             multiply,
	"sjoin":                sjoin,
	"add":                  fltrAdd,
	"spare":                spare,
	"bleft":                bleft,
	"div":                  divide,
	"bright":               bright,
	"contactTBS":           contactTables,
	"fltr2Num":             fltr2Num,
}

func divide(x, y int) int {
	return x / y
}
func bleft() string {
	return `{{`
}
func bright() string {
	return `}}`
}
func fltrAdd(x int, y int) int {
	return x + y
}

func IsTmplTb(o interface{}) bool {
	if v, ok := o.(*TTable); ok {
		return v.IsTmpl
	}
	return false
}

func getFirstTable(ts []*Table) *Table {
	if len(ts) > 0 {
		return ts[0]
	}
	return &Table{}
}
func flterMainTable(tbs []*Table) []*Table {
	cache := map[string]bool{}
	ntbs := make([]*Table, 0, len(tbs))
	for _, t := range tbs {
		if _, ok := cache[t.Name.Main]; !ok {
			cache[t.Name.Main] = true
			ntbs = append(ntbs, t)
		}
	}
	return ntbs
}
func fltrColumsExcludeExt(cols []*Column) []*Column {
	vc := make([]*Column, 0, 1)
	for _, v := range cols {
		if !v.Field.IsExtFuncField {
			vc = append(vc, v)
		}
	}
	return vc
}
func spare(x int, y int) int {
	return x % y
}

func multiply(v int, b int) int {
	return v * b
}
func sjoin(v ...string) string {
	return strings.Join(v, "")
}

func fltrNotNullRows(rs []*Column) []*Column {
	r := make([]*Column, 0, 1)
	for _, v := range rs {
		if v.Field.Required {
			r = append(r, v)
		}
	}
	return r
}
func fltrOptrs(opts []*optrs, tps string) []*optrs {
	nopts := make([]*optrs, 0, 1)
	tpn := strings.Split(tps, "-")
	for _, v := range opts {
		for _, tp := range tpn {
			if strings.EqualFold(v.Name, tp) {
				nopts = append(nopts, v)
			}
		}

	}
	return nopts
}
func fltr2Num(t string, def int) int {
	return types.GetInt(t, def)
}
func fltrOptrPara(opt *optrs, name string, def string) string {
	if v, ok := opt.Params[name]; ok {
		return v
	}
	return def
}
