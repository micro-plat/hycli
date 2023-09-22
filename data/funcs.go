package data

import (
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

var Funcs = map[string]interface{}{
	"f_string_contact":   f_string_contact,
	"f_string_join":      f_string_join,
	"f_string_contains":  f_string_contains,
	"f_string_start":     f_string_start,
	"f_string_2cname":    md.ToCName,
	"f_string_translate": f_string_translate,
	"f_string_trim":      f_string_trim,

	"f_mysql_get_type":      f_mysql_get_type,
	"f_mysql_get_def_value": f_mysql_get_def_value,

	"f_num_add":      f_num_add,
	"f_num_spare":    f_num_spare,
	"f_num_divide":   f_num_divide,
	"f_num_multiply": f_num_multiply,

	"f_table_cache":        f_table_cache,
	"f_table_first":        f_table_first,
	"f_table_get_ttable":   f_table_get_ttable,
	"f_table_find_by_name": f_table_find_by_name,
	"f_table_is_tmp":       f_table_is_tmp,

	"f_colum_idx":   f_colum_idx,
	"f_colum_first": f_colum_first,
	"f_optr_first":  f_optr_first,
}

func f_num_divide(x, y interface{}) int {
	return types.GetInt(x, 0) / types.GetInt(y, 1)
}

func f_num_add(x, y interface{}) int {
	return types.GetInt(x) + types.GetInt(y)
}

func f_table_is_tmp(o interface{}) bool {
	if v, ok := o.(*TTable); ok {
		return v.IsTmpl
	}
	if v, ok := o.(*BUITable); ok {
		return v.IsTmpl
	}
	return false
}

func f_table_first(ts []*Table) *Table {
	if len(ts) > 0 {
		return ts[0]
	}
	return &Table{}
}
func f_colum_idx(ts []*Column, i int, d *Column) *Column {
	if len(ts) > i {
		return ts[i]
	}
	return d
}
func f_string_start(p string, s string) bool {
	return strings.HasPrefix(p, s)
}
func f_string_contains(p string, s string) bool {
	return strings.Contains(p, s)
}
func f_string_trim(p string, s string) string {
	return strings.Trim(p, s)
}
func f_table_cache(tbs []*Table) []*Table {
	cache := map[string]bool{}
	ntbs := make([]*Table, 0, len(tbs))
	for _, t := range tbs {
		if _, ok := cache[t.Name.Main]; !ok {
			cache[t.Name.Main] = true
			ntbs = append(ntbs, t)
		}
	}
	return ntbs
}

func f_num_spare(x int, y int) int {
	return x % y
}

func f_num_multiply(v int, b int) int {
	return v * b
}
func f_string_contact(v ...string) string {
	return strings.Join(v, "")
}
func f_string_join(v []string, p string) string {
	return strings.Join(v, p)
}

type expr struct {
	Name   string
	Field  string
	Symbol string
	Value  string
}

func f_string_translate(f string, t interface{}) string {
	tb := getTable(t)
	return types.Translate(f,
		"name", tb.Name,
		"prefix", tb.Name.Prefix,
		"main", tb.Name.Main,
		"mainPath", strings.ToLower(tb.Name.MainPath))
}
func getTable(tx interface{}) *Table {
	if t, ok := tx.(*Table); ok {
		return t
	}
	if t, ok := tx.(*TTable); ok {
		return t.Table
	}
	return &Table{}
}

func f_mysql_get_type(c *md.Row) string {
	return mySQLType(c.Type.Name, c.Type.Len, c.Type.DLen)
}
func f_optr_first(opts []*optrs) *optrs {
	if len(opts) > 0 {
		return opts[0]
	}
	return &optrs{}
}
func f_colum_first(columns []*Column) *Column {
	if len(columns) > 0 {
		return columns[0]
	}
	return &Column{}
}
func f_table_get_ttable(uname string) *TTable {
	ut := f_table_find_by_name(uname)
	return &TTable{Table: ut, IsTmpl: true}
}
