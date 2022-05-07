package data

import "strings"

var Funcs = map[string]interface{}{
	"fltrCodeTable":     fltrCodeTable,     //代码
	"fltrCodeTables":    fltrCodeTables,    //代码
	"fltrSearchUITable": fltrSearchUITable, //全局查找指定表
	"IsTmplTb":          IsTmplTb,
	"flterMainTable":    flterMainTable,
	// "fltrUITable":       fltrUITable,  //代码
	// "fltrUITables":      fltrUITables, //代码
	"fltrUIRows":        fltrUIRows,
	"fltrNotNullRows":   fltrNotNullRows,
	"getFirstCodeTable": getFirstCodeTable,
	"mergeCodeRow":      mergeCodeRow,
	"mergeUIRow":        mergeUIRow,
	"multiply":          multiply,
	"sjoin":             sjoin,
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
