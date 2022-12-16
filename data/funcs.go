package data

import "strings"

var Funcs = map[string]interface{}{
	"fltrCodeTable":     fltrCodeTable,     //代码
	"fltrCodeTables":    fltrCodeTables,    //代码
	"fltrSearchUITable": fltrSearchUITable, //全局查找指定表
	"IsTmplTb":          IsTmplTb,
	"flterMainTable":    flterMainTable,
	"fltrForginKey":     fltrForginKey,
	"fltrUIRows":        fltrUIRows,
	"fltrNotNullRows":   fltrNotNullRows,
	"getFirstCodeTable": getFirstCodeTable,
	"mergeCodeRow":      mergeCodeRow,
	"mergeUIRow":        mergeUIRow,
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
	if v, ok := o.(*TUITable); ok {
		return v.IsTmpl
	}
	return false
}

func getFirstCodeTable(ts []*CodeTable) *CodeTable {
	if len(ts) > 0 {
		return ts[0]
	}
	return &CodeTable{}
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
func fltrNotNullRows(rs []*CodeRow) []*CodeRow {
	r := make([]*CodeRow, 0, 1)
	for _, v := range rs {
		if !v.AllowNull {
			r = append(r, v)
		}
	}
	return r
}
func fltrForginKey(t *TUITable, n *UITable, fiedName string) *UIRow {
	//根据sl配置匹配
	for _, r := range t.QRows {
		if r.RType.IsSelect && strings.HasSuffix(n.Name.Raw, r.RType.SelectName) {
			return r
		}
	}

	//根据字段名称匹配
	for _, r := range t.QRows {
		if strings.EqualFold(r.Name, fiedName) {
			return r
		}
	}
	return &UIRow{}
}
