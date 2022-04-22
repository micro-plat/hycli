package data

import "strings"

var Funcs = map[string]interface{}{
	"fltrCodeTable":  fltrCodeTable,  //代码
	"fltrCodeTables": fltrCodeTables, //代码
	"fltrUITable":    fltrUITable,    //代码
	"fltrUITables":   fltrUITables,   //代码
	"fltrUIRows":     fltrUIRows,
	"mergeCodeRow":   mergeCodeRow,
	"multiply":       multiply,
	"sjoin":          sjoin,
}

func multiply(v int, b int) int {
	return v * b
}
func sjoin(v ...string) string {
	return strings.Join(v, "")
}
