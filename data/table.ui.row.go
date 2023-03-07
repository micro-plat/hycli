package data

import (
	"github.com/micro-plat/hycli/md"
	"github.com/micro-plat/lib4go/types"
)

// UIRow 查询行
type UIRow struct {
	*md.Row
	RType       *UIType
	RefedRows   []*UIRow
	Placeholder string
	Model       string
	Label       string
	Required    bool
	Ext         *Ext
	UNQ         string
}

// NewUIRow 构建查询行数据
func NewUIRow(t string, r *md.Row, formName ...string) *UIRow {
	row := &UIRow{
		Row:         r,
		Placeholder: r.Desc.Name,
		Label:       r.Desc.Name,
		RefedRows:   make([]*UIRow, 0),
		Model:       r.Name,
		Required:    !r.AllowNull,
		RType:       NewUIType(t, r),
		Ext:         newExt(t, r, types.GetStringByIndex(formName, 0)),
		UNQ:         defFids.Next(),
	}

	return row
}
