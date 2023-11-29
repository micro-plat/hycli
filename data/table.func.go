package data

import (
	"strings"
)

// GetStaticColumns 获取包含指定组件和包含某个静态字段(从用户信息、指定固定值)的列
//
//	func (tb *Table) GetStaticColumns(tps, prefix string) map[string]*Column {
//		columns := make(map[string]*Column, 1)
//		tplst := strings.Split(tps, "-")
//		for _, c := range tb.Columns {
//			for _, tp := range tplst {
//				f := md.GetFormat(tp, c.Constraints...)
//				if strings.HasPrefix(f, prefix) {
//					columns[f] = c
//				}
//			}
//		}
//		return columns
//	}
//
//	func (tx *Table) GetColumnsByCmpnt(cmpnt string, tps ...string) Columns {
//		return getCmpnt(tx, cmpnt, tps...)
//	}
//
//	func (tx *TTable) GetColumnsByCmpnt(cmpnt string, tps ...string) Columns {
//		return getCmpnt(tx, cmpnt, tps...)
//	}
//
//	func (tx Columns) GetColumnsByCmpnt(cmpnt string, tps ...string) Columns {
//		return getCmpnt(tx, cmpnt, tps...)
//	}
// func (tx *Table) JoinNames(tp string, required bool, start string, end ...string) string {
// 	Columns := tx.Columns.GetAll()
// 	return Columns.JoinNames(tp, required, start, end...)
// }

func getTPS(tp string) []string {
	is := strings.Split(tp, "-")
	lst := make([]string, 0, 1)
	for _, v := range is {
		lst = append(lst, strings.ToLower(v))
		lst = append(lst, strings.ToUpper(v))
	}
	return lst
}
