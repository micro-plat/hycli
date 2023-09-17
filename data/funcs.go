package data

import (
	"sort"
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

var Funcs = map[string]interface{}{
	"f_string_contact":   f_string_contact,
	"f_string_join":      f_string_join,
	"f_string_start":     f_string_start,
	"f_string_2cname":    md.ToCName,
	"f_table_cache":      f_table_cache,
	"f_table_first":      f_table_first,
	"f_table_get_ttable": f_table_get_ttable,
	"f_colum_idx":        f_colum_idx,
	"f_table_is_tmp":     f_table_is_tmp,

	"f_mysql_get_type":      f_mysql_get_type,
	"f_mysql_get_def_value": f_mysql_get_def_value,
	"f_colum_first":         f_colum_first,
	// "fltrNotNullCols":    fltrNotNullCols,
	// "fltrSearchUITable": fltrSearchUITable,
	// "fltrSearchUITableAndResetUNQ": fltrSearchUITableAndResetUNQ,
	// "fltrIsMutilValue":        fltrIsMutilValue,
	// "fltrNeedBatchCheck":      fltrNeedBatchCheck,

	// "fltrHasConst":            fltrHasConst,
	// "fltrGetConst":            fltrGetConst,

	// "fltrOptrs":               fltrOptrs,
	// "fltrOptrsByTable":        fltrOptrsByTable,

	// "fltrFindAllExtViewOptrs": fltrFindAllExtViewOptrs,
	// "fltrSelectedOptrs":      fltrSelectedOptrs,

	// "fltrFrontOptrsByStatic": fltrFrontOptrsByStatic,
	// "fltrGetConstrainValues": fltrGetConstrainValues,
	"f_optr_first": f_optr_first,
	// "fltrOptrPara": fltrOptrPara,

	// "fltrColumns":             fltrColumns,

	// "flterJoinColumnNames": flterJoinColumnNames,

	"f_string_translate": f_string_translate,
	// "fltrCmpnt":          fltrCmpnt,
	// "fltrColumnsExcludeExt": fltrColumnsExcludeExt,

	// "fltr2Expr": fltr2Expr,
	// "fltrOptrParaExprs": fltrOptrParaExprs,

	"f_table_find_by_name": f_table_find_by_name,
	"f_string_trim":        f_string_trim,
	// "resetForm":            resetForm,
	"f_num_multiply": f_num_multiply,

	"f_num_add":   f_num_add,
	"f_num_spare": f_num_spare,
	// "bleft":       bleft,
	"f_num_divide": f_num_divide,
	// "bright":      bright,
	// "fltr2Num": fltr2Num,
}

func f_num_divide(x, y interface{}) int {
	return types.GetInt(x, 0) / types.GetInt(y, 1)
}

//	func bleft() string {
//		return `{-{`
//	}
//
//	func bright() string {
//		return `}-}`
//	}
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

//	func fltrNotNullCols(rs []*Column) []*Column {
//		r := make([]*Column, 0, 1)
//		for _, v := range rs {
//			if v.Field.Required {
//				r = append(r, v)
//			}
//		}
//		return r
//	}
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

// func fltrGetConstrainValues(t *Table, tp string) []string {
// 	for _, c := range t.Columns {
// 		if md.HasConstraint(c.Constraints, tp) {
// 			return md.GetConstrainValues(tp, c.Constraints...)
// 		}

// 	}
// 	return nil

// }
// func fltrColumnsExcludeExt(cols []*Column) []*Column {
// 	vc := make([]*Column, 0, 1)
// 	for _, v := range cols {
// 		if !v.Field.IsExtFuncField {
// 			vc = append(vc, v)
// 		}
// 	}
// 	return vc
// }

//	func fltrHasConst(c *Column, p string) bool {
//		return md.HasConstraint(c.Constraints, strings.ToLower(p), strings.ToUpper(p))
//	}
// func fltrGetConst(c *Column, p string, name string, def string) string {
// 	v, page, _ := md.GetConsByTagIgnorecase(name, c.Constraints...)
// 	pages := strings.Split(page, "-")
// 	if len(pages) == 0 || types.StringContains(pages, p) {
// 		return v
// 	}
// 	return def

// }
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

//	func fltrNotNullRows(rs []*Column) []*Column {
//		r := make([]*Column, 0, 1)
//		for _, v := range rs {
//			if v.Field.Required {
//				r = append(r, v)
//			}
//		}
//		return r
//	}
//
//	func fltrOptrs(opts []*optrs, tps string) optrslst {
//		nopts := make(optrslst, 0, 1)
//		tpn := strings.Split(tps, "-")
//		for _, v := range opts {
//			for _, tp := range tpn {
//				if strings.EqualFold(v.Name, tp) {
//					nopts = append(nopts, v)
//				}
//			}
//		}
//		sort.Sort(nopts)
//		return nopts
//	}
// func fltrFindAllExtViewOptrs(optrs viewOptrs) optrslst {
// 	lst := make(optrslst, 0, 1)
// 	for _, op := range optrs {
// 		lst = append(lst, findOptrsByView(op.Table, "lst-lstbar")...)
// 	}
// 	return lst
// }

func findOptrsByView(tbName string, tp string) optrslst {
	tb := f_table_find_by_name(tbName)
	if tb.Name == nil {
		return nil
	}
	tpNames := strings.Split(strings.ToLower(tp), "-")

	nlst := make(optrslst, 0, 1)
	for _, tpName := range tpNames {
		lst := make(optrslst, 0, 1)
		switch tpName {
		case "lst":
			lst = optrslst(tb.ListOpts)
		case "lstbar":
			lst = optrslst(tb.BarOpts)
		}
		for _, v := range lst {
			if v.Params.Get("showView") == "true" {
				nlst = append(nlst, v.Get(tbName))
			}
		}
	}

	return nlst
}

// func fltrOptrsByTag(opts []*optrs, tag string) optrslst {
// 	nopts := make(optrslst, 0, 1)
// 	tags := strings.Split(tag, "-")
// 	for _, v := range opts {
// 		for _, tp := range tags {
// 			if strings.EqualFold(v.Tag, tp) {
// 				nopts = append(nopts, v)
// 			}
// 		}
// 	}
// 	sort.Sort(nopts)
// 	return nopts
// }

// func fltrOptrsByTable(opts []*optrs, tb string) optrslst {
// 	nopts := make(optrslst, 0, 1)
// 	for _, v := range opts {
// 		if strings.EqualFold(v.Table, tb) {
// 			nopts = append(nopts, v)
// 		}
// 	}
// 	sort.Sort(nopts)
// 	return nopts
// }

// func fltrSelectedOptrs(opts []*optrs) *optrs {
// 	for _, optr := range opts {
// 		if v, ok := optr.Params["selected"]; ok && v == "true" {
// 			return optr
// 		}
// 	}
// 	if len(opts) > 0 {
// 		return opts[0]
// 	}
// 	return &optrs{}

// }

//	func fltrOptrsByCmd(opts []*optrs, cmds string) optrslst {
//		nopts := make(optrslst, 0, 1)
//		cmd := strings.Split(cmds, "-")
//		for _, v := range opts {
//			for _, c := range cmd {
//				if strings.EqualFold(v.Cmd, c) {
//					nopts = append(nopts, v)
//				}
//			}
//		}
//		sort.Sort(nopts)
//		return nopts
//	}
// func fltrFrontOptrsByStatic(opts *optrs) map[string]string {
// 	outs := make(map[string]string)
// 	for k, v := range opts.Params {
// 		if strings.HasPrefix(k, "@") {
// 			outs[strings.Trim(k, "@")] = v
// 		}
// 	}

// 	return outs
// }

// func fltrStaticColumn(tb *Table, tps, prefix string) map[string]*Column {
// 	columns := make(map[string]*Column, 1)
// 	tplst := strings.Split(tps, "-")
// 	for _, c := range tb.Columns {
// 		for _, tp := range tplst {
// 			f := md.GetFormat(tp, c.Constraints...)
// 			if strings.HasPrefix(f, prefix) {
// 				columns[f] = c
// 			}
// 		}
// 	}
// 	return columns
// }

// func fltr2Num(t string, def int) int {
// 	return types.GetInt(t, def)
// }

// func fltrOptrPara(opt *optrs, name string, def string) string {
// 	if v, ok := opt.Params[name]; ok {
// 		return v
// 	}
// 	return def
// }

type expr struct {
	Name   string
	Field  string
	Symbol string
	Value  string
}

// func fltrOptrParaExprs(opt *optrs, name string, def string) []*expr {
// 	p := fltrOptrPara(opt, name, def)
// 	return fltr2Expr(p)
// }

// fltrColumns 过滤用户自定义类型对应的行，自定义行对应的控件按新增模式处理
func fltrColumns(tx interface{}, tp string, formName ...string) Columns {
	t, ok := tx.(*Table)
	if !ok {
		t = tx.(*TTable).Table
	}

	cols := make(Columns, 0, 1)
	tps := getTPS(tp)
	for _, r := range t.Columns {
		if md.HasConstraint(r.RawConsts, tps...) {
			r.ResetCmpnt(tps...) //重置扩展组件
			r.Ext.FormName = types.GetStringByIndex(formName, 0, "form")
			cols = append(cols, r.Index(types.GetStringByIndex(tps, 0), len(cols)))
		}
	}
	sort.Sort(cols)
	return cols
}

//	func flterJoinColumnNames(tx interface{}, tp string, start string, end ...string) string {
//		colums := fltrColumns(tx, tp)
//		names := make([]string, 0, 1)
//		for _, v := range colums {
//			names = append(names, v.Name)
//		}
//		if len(names) == 0 {
//			return ""
//		}
//		return start + strings.Join(names, types.GetStringByIndex(end, 0)+start) + types.GetStringByIndex(end, 0)
//	}
// func fltrCmpnt(tx interface{}, cmpnt string, tps ...string) []*Column {
// 	colums := getColumns(tx)
// 	cols := make(Columns, 0, 1)
// 	for _, r := range colums {
// 		if r.allCmpnt.getCmpnt(r.Row, tps...).Type == cmpnt {
// 			r.ResetCmpnt(tps...)
// 			cols = append(cols, r)
// 		}
// 	}
// 	sort.Sort(cols)
// 	return cols
// }

//	func fltrIsMutilValue(tx interface{}, optr *optrs, name string) bool {
//		if optr.Params.IsBatchCheck(name) {
//			return true
//		}
//		colums := getColumns(tx)
//		for _, r := range colums {
//			if strings.HasPrefix(r.Cmpnt.Type, "multi") {
//				return true
//			}
//		}
//		return false
//	}
//
//	func fltrNeedBatchCheck(optr *optrs) bool {
//		return optr.NeedBatchCheck()
//	}
//
//	func resetForm(t *Table) *Table {
//		for _, c := range t.Columns {
//			c.Ext.FormName = "form"
//		}
//		return t
//	}
func getTPS(tp string) []string {
	is := strings.Split(tp, "-")
	lst := make([]string, 0, 1)
	for _, v := range is {
		lst = append(lst, strings.ToLower(v))
		lst = append(lst, strings.ToUpper(v))
	}
	return lst
}

//	func fltrAssctColumns(tx interface{}, colName string) []*Column {
//		columns := getColumns(tx)
//		cols := make(Columns, 0, 1)
//		for _, r := range columns {
//			if r.Enum.AssctColumn == "" {
//				continue
//			}
//			if strings.EqualFold(r.Enum.AssctColumn, colName) {
//				cols = append(cols, r)
//			}
//		}
//		return cols
//	}
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
func fltrDisOptrs(ors []*optrs) []*optrs {
	lst := make([]*optrs, 0, 1)
	vlst := map[string]bool{}
	for _, v := range ors {
		if _, ok := vlst[v.UNQ+v.URL]; !ok {
			lst = append(lst, v)
			vlst[v.UNQ+v.URL] = true
		}
	}
	return lst
}

func getColumns(tx interface{}) []*Column {
	if t, ok := tx.(*Table); ok {
		return t.Columns
	}
	if t, ok := tx.(*TTable); ok {
		return t.Columns
	}
	if t, ok := tx.([]*Column); ok {
		return t
	}
	return nil
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
