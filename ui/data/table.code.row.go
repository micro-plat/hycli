package data

import (
	"github.com/micro-plat/hycli/md"
)

//CodeRow 查询行
type CodeRow struct {
	*md.Row
	RType         *UIType
	IsQueryField  bool
	IsListField   bool
	IsInsertField bool
	IsUpdateField bool
	IsDeleteField bool
	IsPKField     bool
}

//NewCodeRow 构建查询行数据
func NewCodeRow(t string, r *md.Row) *CodeRow {
	return &CodeRow{
		Row:           r,
		IsQueryField:  md.HasConstraint(r.Constraints, "Q"),
		IsListField:   md.HasConstraint(r.Constraints, "L"),
		IsInsertField: md.HasConstraint(r.Constraints, "C"),
		IsUpdateField: md.HasConstraint(r.Constraints, "U"),
		IsDeleteField: md.HasConstraint(r.Constraints, "D"),
		IsPKField:     md.HasConstraint(r.Constraints, "PK"),
		RType:         NewUIType(t, r),
	}
}
func mergeCodeRow(a []*CodeRow, b []*CodeRow) []*CodeRow {
	rows := make([]*CodeRow, 0, len(a)+len(b))
	cache := map[string]string{}
	for _, r := range a {
		rows = append(rows, r)
		cache[r.Name] = r.Name
	}
	for _, r := range b {
		if _, ok := cache[r.Name]; !ok {
			rows = append(rows, r)
		}
	}
	return rows
}
func fltrRows(t *CodeTable, tp string) []*CodeRow {
	rows := make([]*CodeRow, 0, 1)
	for _, r := range t.Rows {
		if md.HasConstraint(r.Constraints, tp) {
			rows = append(rows, NewCodeRow(tp, r))
		}
	}
	return rows
}
