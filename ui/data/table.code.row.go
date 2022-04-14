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
