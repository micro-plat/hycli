package data

import (
	"sort"
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

func (Columns Columns) JoinNames(tp string, required bool, start string, end ...string) string {
	names := make([]string, 0, 1)
	tpColumn := Columns.GetColumns(tp)
	for _, v := range tpColumn {
		field := v
		alias := v.GetOtherCmpntValue("alias")
		name := v.Field.Name
		if alias != "" {
			field = Columns.GetColumnByFieldName(alias)
		}
		if !required || (field != nil && required && field.Field.Required) {
			names = append(names, name)
		}
		continue

	}
	if len(names) == 0 {
		return ""
	}
	return start + strings.Join(names, types.GetStringByIndex(end, 0)+start) + types.GetStringByIndex(end, 0)
}
func (Columns Columns) GetColumnByFieldName(name string) *Column {
	for _, c := range Columns {
		if strings.EqualFold(c.Field.Name, name) {
			return c
		}
	}
	return nil
}
func (Columns Columns) First() *Column {
	if len(Columns) > 0 {
		return Columns[0]
	}
	return &Column{}
}
func (Columns Columns) GetValidColumns() Columns {
	vc := make([]*Column, 0, 1)
	for _, v := range Columns {
		if !v.Field.IsExtFuncField {
			vc = append(vc, v)
		}
	}
	return vc
}
func (rs Columns) GetRequiedColumns() Columns {
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
func (c Columns) GetAll() Columns {
	return c.GetColumns("q-bq-c-bc-fc-l-le-u-bu-d-v")
}
func (c Columns) GetFontQueryColumns() Columns {
	return c.GetColumns("q")
}
func (c Columns) GetAllQueryColumns() Columns {
	return c.GetColumns("q-bq")
}
func (c Columns) GetFontCreateColumns() Columns {
	return c.GetColumns("c")
}
func (c Columns) GetFontBackCreateColumns() Columns {
	return c.GetColumns("fc")
}
func (c Columns) GetAllCreateColumns() Columns {
	return c.GetColumns("c-bc-fc")
}
func (c Columns) GetOnlyListColumns() Columns {
	return c.GetColumns("l")
}
func (c Columns) GetOnlyListExtColumns() Columns {
	return c.GetColumns("le")
}
func (c Columns) GetFontListColumns() Columns {
	return c.GetColumns("l-le")
}
func (c Columns) GetAllListColumns() Columns {
	return c.GetColumns("q-l-le-bl")
}
func (c Columns) GetDelColumns() Columns {
	return c.GetColumns("D")
}
func (c Columns) GetFontUpdateColumns() Columns {
	return c.GetColumns("U")
}
func (c Columns) GetAllUpdateColumns() Columns {
	return c.GetColumns("U-BU")
}
func (c Columns) GetViewColumns() Columns {
	return c.GetColumns("v")
}

func (c Columns) GetColumns(tp string, formName ...string) Columns {
	cols := make(Columns, 0, 1)
	tps := getTPS(tp)
	for _, r := range c {
		if md.HasConstraint(r.RawConsts, tps...) {
			r.ResetCmpnt(tps...) //重置扩展组件
			r.Ext.FormName = types.GetStringByIndex(formName, 0, "form")
			r.Ext.TP = tps
			cols = append(cols, r.Index(types.GetStringByIndex(tps, 0), len(cols)))
		}
	}
	sort.Sort(cols)
	return cols
}
func (c Columns) GetByCmpnt(cmpnt string) Columns {
	cols := make(Columns, 0, 1)
	cmpnts := strings.Split(cmpnt, "-")
	for _, r := range c {
		cmp := r.allCmpnt.getCmpnt(r.Row, r.Ext.TP...)
		if types.StringContains(cmpnts, cmp.Type) {
			r.ResetCmpnt(r.Ext.TP...)
			cols = append(cols, r)
		}
	}
	sort.Sort(cols)
	return cols
}
func (cs Columns) GetStaticColumns() Columns {
	prefix := "#"
	return cs.GetStartColumns(prefix)
}
func (cs Columns) GetCmpntValues(tp string) []string {
	for _, c := range cs {
		if md.HasConstraint(c.Constraints, tp) {
			return md.GetConstrainValues(tp, c.Constraints...)
		}
	}
	return nil
}
func (cs Columns) HasCmpnt(tp string) bool {
	for _, c := range cs {
		if md.HasConstraint(c.Constraints, tp) {
			return true
		}
	}
	return false
}
func (cs Columns) GetStartColumns(prefix string) Columns {
	cls := make(Columns, 0, 1)
	for _, c := range cs {
		for _, tp := range c.Ext.TP {
			f := md.GetFormat(tp, c.Constraints...)
			if strings.HasPrefix(f, prefix) {
				c.Ext.Name = strings.Trim(f, prefix)
				cls = append(cls, c)
			}
		}
	}
	return cls
}
