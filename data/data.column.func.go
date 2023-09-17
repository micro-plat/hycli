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
func (c *Column) GetOpt(name string) string {
	v, _, _ := c.GetOpts(name)
	return v
}
func (colums Columns) JoinNames(tp string, start string, end ...string) string {
	names := make([]string, 0, 1)
	for _, v := range colums {
		names = append(names, v.Name)
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
