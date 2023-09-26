package data

import (
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

func (cs Columns) GetColumnsBy(name ...string) Columns {
	cols := make(Columns, 0, 1)
	for _, c := range cs {
		if md.HasConstraint(c.RawConsts, name...) {
			cols = append(cols, c)
		}
	}
	return cols
}
func (c *Column) GetOpts(name string) (string, string, string) {
	f, s, t := md.GetConsByTagIgnorecase(name, c.RawConsts...)
	return f, s, t
}
func (c *Column) GetParams(name string) map[string]string {
	s := md.GetParamByTag(name, c.RawConsts...)
	item := strings.Split(s, ";")
	param := make(map[string]string)
	for _, v := range item {
		kv := strings.Split(v, ":")
		if len(kv) == 0 || kv[0] == "" {
			continue
		}
		if len(kv) == 2 {
			param[kv[0]] = kv[1]
			continue
		}
		param[kv[0]] = ""
	}
	return param

}
func (c *Column) GetOpt(name string) string {
	v, _, _ := c.GetOpts(name)
	return v
}
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
func (colums Columns) GetColumnsByColumName(colName string) []*Column {
	return getColumnsByColumName(colums, colName)
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
func (c *Column) HasConst(p string) bool {
	return md.HasConstraint(c.Constraints, strings.ToLower(p), strings.ToUpper(p))
}
func (c *Column) GetConst(p string, name string, def string) string {
	v, page, _ := md.GetConsByTagIgnorecase(name, c.Constraints...)
	pages := strings.Split(page, "-")
	if len(pages) == 0 || types.StringContains(pages, p) {
		return v
	}
	return def

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
