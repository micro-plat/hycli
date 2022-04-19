package data

var Funcs = map[string]interface{}{
	"fltrCodeTable":  fltrCodeTable,  //代码
	"fltrCodeTables": fltrCodeTables, //代码
	"fltrUITable":    fltrUITable,    //代码
	"fltrUITables":   fltrUITables,   //代码
	"fltrRows":       fltrRows,
	"mergeCodeRow":   mergeCodeRow,
}

// func fltrIsInput(t string) bool {
// 	return strings.EqualFold(t, "input") || strings.EqualFold(t, "link") ||
// 		strings.EqualFold(t, "progress") || strings.EqualFold(t, "tag")
// }
