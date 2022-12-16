package data

import (
	"github.com/micro-plat/hycli/md"
	"github.com/micro-plat/lib4go/types"
)

// CodeRow 查询行
type CodeRow struct {
	*md.Row
	RType         *UIType
	IsQueryField  bool
	IsListField   bool
	IsInsertField bool
	IsUpdateField bool
	IsDeleteField bool
	IsPKField     bool
	BQDefValue    string
}

// NewCodeRow 构建查询行数据
func NewCodeRow(t string, r *md.Row) *CodeRow {
	return &CodeRow{
		Row: r,
		IsQueryField: md.HasConstraint(r.Constraints, "Q") ||
			md.HasConstraint(r.Constraints, "BQ"),
		BQDefValue:    md.GetValue("BQ", r.Constraints...),
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
func mergeUIRow(a []*UIRow, b []*UIRow) []*UIRow {
	rows := make([]*UIRow, 0, len(a)+len(b))
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

// fltrUIRows 过滤用户自定义类型对应的行，自定义行对应的控件按新增模式处理
func fltrUIRows(t *UITable, tp string, formName ...string) []*UIRow {
	rows := make([]*UIRow, 0, 1)
	for _, r := range t.Rows {
		if md.HasConstraint(r.Constraints, tp) {
			rows = append(rows, NewUIRow("C", r, types.GetStringByIndex(formName, 0)))
		}
	}
	return rows
}
