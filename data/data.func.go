package data

import (
	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

// fltrColums 过滤用户自定义类型对应的行，自定义行对应的控件按新增模式处理
func fltrColums(t *Table, tp string, formName ...string) []*Column {
	cols := make([]*Column, 0, 1)
	for _, r := range t.Colums {
		if md.HasConstraint(r.RawConsts, tp) {
			r.Ext.FormName = types.GetStringByIndex(formName, 0)
			cols = append(cols, r)
		}
	}
	return cols
}
