package data

import (
	"sort"
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

func (colums Columns) JoinNames(tp string, required bool, start string, end ...string) string {
	names := make([]string, 0, 1)
	for _, v := range colums {
		if required && v.Field.Required || !required {
			names = append(names, v.Name)
		}
	}
	if len(names) == 0 {
		return ""
	}
	return start + strings.Join(names, types.GetStringByIndex(end, 0)+start) + types.GetStringByIndex(end, 0)
}

func (colums Columns) GetValidColumns() Columns {
	vc := make([]*Column, 0, 1)
	for _, v := range colums {
		if !v.Field.IsExtFuncField {
			vc = append(vc, v)
		}
	}
	return vc
}
func (rs Columns) GetRequiedColums() []*Column {
	r := make([]*Column, 0, 1)
	for _, v := range rs {
		if v.Field.Required {
			r = append(r, v)
		}
	}
	return r
}

func (cs Columns) GetColumnsBy(name ...string) Columns {
	cols := make(Columns, 0, 1)
	for _, c := range cs {
		if md.HasConstraint(c.RawConsts, name...) {
			cols = append(cols, c)
		}
	}
	return cols
}
func (cs Columns) GetColumnsByTriggerChangeEvent(colName string) Columns {
	cls := make(Columns, 0, 1)
	for _, c := range cs {
		if c.NeedTriggerChangeEvent(colName) {
			cls = append(cls, c)
		}
	}
	return cls
}

func (c Columns) GetPKColumns() Columns {
	cols := make(Columns, 0, 1)
	for _, v := range c {
		if v.Field.IsPK {
			cols = append(cols, v)
		}
	}
	return cols
}
func (c Columns) GetEnumColumns() Columns {
	cols := make(Columns, 0, 1)
	for _, v := range c {
		if v.Enum.IsEnum {
			cols = append(cols, v)
		}
	}
	return cols
}
func (c Columns) GetEnumDelColumns() Columns {
	cols := make(Columns, 0, 1)
	for _, v := range c {
		if v.Enum.IsDEColumn {
			cols = append(cols, v)
		}
	}
	return cols
}
func (c Columns) GetQueryColums() Columns {
	cols := make(Columns, 0, 1)
	tps := []string{"q"}
	for _, r := range c {
		if md.HasConstraint(r.RawConsts, tps...) {
			r.ResetCmpnt(tps...) //重置扩展组件
			cols = append(cols, r.Index(types.GetStringByIndex(tps, 0), len(cols)))
		}
	}
	sort.Sort(cols)
	return cols
}
