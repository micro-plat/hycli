package data

import (
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

var Funcs = map[string]interface{}{

	"flterMainTable":    flterMainTable,
	"fltrNotNullCols":   fltrNotNullRows,
	"getFirstTable":     getFirstTable,
	"fltrUICols":        fltrUICols,
	"IsTmplTb":          IsTmplTb,
	"fltrSearchUITable": fltrSearchUITable, //全局查找指定表??
	"fltrColums":        fltrColums,
	"resetForm":         resetForm,
	"multiply":          multiply,
	"sjoin":             sjoin,
	"add":               fltrAdd,
	"spare":             spare,
	"bleft":             bleft,
	"bright":            bright,
	"contactTBS":        contactTables,
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

func spare(x int, y int) int {
	return x % y
}

func multiply(v int, b int) int {
	return v * b
}
func sjoin(v ...string) string {
	return strings.Join(v, "")
}

// fltrUICols 过滤用户自定义类型对应的行，自定义行对应的控件按新增模式处理
func fltrUICols(t *Table, tp string, formName ...string) []*Column {
	rows := make([]*Column, 0, 1)
	for _, r := range t.Colums {
		if md.HasConstraint(r.RawConsts, tp) {
			r.Ext.FormName = types.GetStringByIndex(formName, 0)
			rows = append(rows, r)
		}
	}
	return rows
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
