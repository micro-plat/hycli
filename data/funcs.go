package data

import (
	"sort"
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

var Funcs = map[string]interface{}{
	"flterMainTable":               flterMainTable,
	"fltrNotNullCols":              fltrNotNullRows,
	"getFirstTable":                getFirstTable,
	"getColumn":                    getColumn,
	"IsTmplTb":                     IsTmplTb,
	"fltrBuildTable":               fltrBuildTable,
	"fltrSearchUITable":            fltrSearchUITable,
	"fltrSearchUITableAndResetUNQ": fltrSearchUITableAndResetUNQ,
	"fltrIsMutilValue":             fltrIsMutilValue,
	"fltrNeedBatchCheck":           fltrNeedBatchCheck,
	"fltrSearchTable":              fltrSearchTable,
	"fltrHasConst":                 fltrHasConst,
	"fltrGetConst":                 fltrGetConst,
	"fltrMYSQLType":                fltrMYSQLType,
	"fltrMYSQLDef":                 mySQLDefValue,
	"fltrOptrs":                    fltrOptrs,
	"fltr2CName":                   md.ToCName,
	"fltrOptrsByPUNQ":              fltrOptrsByPUNQ,
	"fltrOptrsByStatic":            fltrOptrsByStatic,
	"fltrOptrsByTag":               fltrOptrsByTag,
	"fltrOptrsByCmd":               fltrOptrsByCmd,
	"fltrOptrsByPosition":          fltrOptrsByPosition,
	"fltrColumns":                  fltrColumns,
	"getFirstColumns":              getFirstColumns,
	"flterJoinColumnNames":         flterJoinColumnNames,
	"mergeOptrs":                   mergeOptrs,
	"getFirstOptr":                 getFirstOptr,
	"fltrAssctColumns":             fltrAssctColumns,
	"fltrTranslate":                fltrTranslate,
	"fltrCmpnt":                    fltrCmpnt,
	"fltrColumnsExcludeExt":        fltrColumnsExcludeExt,
	"fltrOptrPara":                 fltrOptrPara,
	"fltr2Expr":                    fltr2Expr,
	"fltrOptrParaExprs":            fltrOptrParaExprs,
	"fltrStart":                    fltrStart,
	"fltrFindTable":                findTable,
	"fltrTrim":                     fltrTrim,
	"resetForm":                    resetForm,
	"multiply":                     multiply,
	"sjoin":                        sjoin,
	"add":                          fltrAdd,
	"spare":                        spare,
	"bleft":                        bleft,
	"div":                          divide,
	"bright":                       bright,
	"contactTBS":                   contactTables,
	"fltr2Num":                     fltr2Num,
}

func divide(x, y interface{}) int {
	return types.GetInt(x, 0) / types.GetInt(y, 1)
}
func bleft() string {
	return `{-{`
}
func bright() string {
	return `}-}`
}
func fltrAdd(x, y interface{}) int {
	return types.GetInt(x) + types.GetInt(y)
}

func IsTmplTb(o interface{}) bool {
	if v, ok := o.(*TTable); ok {
		return v.IsTmpl
	}
	return false
}

func getFirstTable(ts []*Table) *Table {
	if len(ts) > 0 {
		return ts[0]
	}
	return &Table{}
}
func getColumn(ts []*Column, i int, d *Column) *Column {
	if len(ts) > i {
		return ts[i]
	}
	return d
}
func fltrStart(p string, s string) bool {
	return strings.HasPrefix(p, s)
}
func fltrTrim(p string, s string) string {
	return strings.Trim(p, s)
}
func flterMainTable(tbs []*Table) []*Table {
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
func fltrOptrsByPUNQ(opts []*optrs, punq string, position string) []*optrs {
	lst := make([]*optrs, 0, 1)
	for _, p := range opts {
		if strings.EqualFold(p.ParentUNQ, punq) && strings.EqualFold(p.Position, position) {
			lst = append(lst, p)
		}
	}
	return lst
}
func fltrColumnsExcludeExt(cols []*Column) []*Column {
	vc := make([]*Column, 0, 1)
	for _, v := range cols {
		if !v.Field.IsExtFuncField {
			vc = append(vc, v)
		}
	}
	return vc
}
func fltrHasConst(c *Column, p string) bool {
	return md.HasConstraint(c.Constraints, strings.ToLower(p), strings.ToUpper(p))
}
func fltrGetConst(c *Column, p string, name string, def string) string {
	v, page, _ := md.GetConsByTagIgnorecase(name, c.Constraints...)
	pages := strings.Split(page, "-")
	if len(pages) == 0 || types.StringContains(pages, p) {
		return v
	}
	return def

}
func spare(x int, y int) int {
	return x % y
}

func multiply(v int, b int) int {
	return v * b
}
func sjoin(v ...string) string {
	return strings.Join(v, "")
}

func fltrNotNullRows(rs []*Column) []*Column {
	r := make([]*Column, 0, 1)
	for _, v := range rs {
		if v.Field.Required {
			r = append(r, v)
		}
	}
	return r
}
func fltrOptrs(opts []*optrs, tps string) optrslst {
	nopts := make(optrslst, 0, 1)
	tpn := strings.Split(tps, "-")
	for _, v := range opts {
		for _, tp := range tpn {
			if strings.EqualFold(v.Name, tp) {
				nopts = append(nopts, v)
			}
		}
	}
	sort.Sort(nopts)
	return nopts
}
func mergeOptrs(opts ...[]*optrs) optrslst {
	lst := make(optrslst, 0, 1)
	v := map[string]bool{}
	for _, fopts := range opts {
		for _, opt := range fopts {
			if _, ok := v[opt.UNQ]; !ok {
				lst = append(lst, opt)
				v[opt.UNQ] = true
			}
		}
	}
	sort.Sort(lst)
	return lst
}
func fltrOptrsByTag(opts []*optrs, tag string) optrslst {
	nopts := make(optrslst, 0, 1)
	tags := strings.Split(tag, "-")
	for _, v := range opts {
		for _, tp := range tags {
			if strings.EqualFold(v.Tag, tp) {
				nopts = append(nopts, v)
			}
		}
	}
	sort.Sort(nopts)
	return nopts
}
func fltrOptrsByPosition(opts []*optrs, position string) optrslst {
	nopts := make(optrslst, 0, 1)
	for _, v := range opts {
		if strings.EqualFold(v.Position, position) {
			nopts = append(nopts, v)
		}

	}
	sort.Sort(nopts)
	return nopts
}

func fltrOptrsByCmd(opts []*optrs, cmds string) optrslst {
	nopts := make(optrslst, 0, 1)
	cmd := strings.Split(cmds, "-")
	for _, v := range opts {
		for _, c := range cmd {
			if strings.EqualFold(v.Cmd, c) {
				nopts = append(nopts, v)
			}
		}
	}
	sort.Sort(nopts)
	return nopts
}
func fltrOptrsByStatic(opts *optrs) map[string]string {
	outs := make(map[string]string)
	for k, v := range opts.Params {
		if strings.HasPrefix(k, "@") {
			outs[strings.Trim(k, "@")] = v
		}
	}

	return outs
}
func fltr2Num(t string, def int) int {
	return types.GetInt(t, def)
}
func fltrOptrPara(opt *optrs, name string, def string) string {
	if v, ok := opt.Params[name]; ok {
		return v
	}
	return def
}

type expr struct {
	Name   string
	Field  string
	Symbol string
	Value  string
}

func fltrOptrParaExprs(opt *optrs, name string, def string) []*expr {
	p := fltrOptrPara(opt, name, def)
	return fltr2Expr(p)
}

func fltr2Expr(f string) []*expr {
	prs := md.GetExprs(f)
	lst := make([]*expr, 0, 1)
	for _, pr := range prs {
		lst = append(lst, &expr{
			Name:   types.DecodeString(strings.HasPrefix(pr[0], "@"), true, "", pr[0]),
			Field:  types.DecodeString(strings.HasPrefix(pr[0], "@"), true, strings.Trim(pr[0], "@"), ""),
			Symbol: pr[1],
			Value:  pr[2],
		})
	}
	return lst
}

// fltrColumns 过滤用户自定义类型对应的行，自定义行对应的控件按新增模式处理
func fltrColumns(tx interface{}, tp string, formName ...string) []*Column {
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
func flterJoinColumnNames(tx interface{}, tp string, start string, end ...string) string {
	colums := fltrColumns(tx, tp)
	names := make([]string, 0, 1)
	for _, v := range colums {
		names = append(names, v.Name)
	}
	if len(names) == 0 {
		return ""
	}
	return start + strings.Join(names, types.GetStringByIndex(end, 0)+start) + types.GetStringByIndex(end, 0)
}
func fltrCmpnt(tx interface{}, cmpnt string, tps ...string) []*Column {
	colums := getColumns(tx)
	cols := make(Columns, 0, 1)
	for _, r := range colums {
		if r.allCmpnt.getCmpnt(r.Row, tps...).Type == cmpnt {
			r.ResetCmpnt(tps...)
			cols = append(cols, r)
		}
	}
	sort.Sort(cols)
	return cols
}
func fltrIsMutilValue(tx interface{}, optr *optrs, name string) bool {
	if optr.Params.IsBatchCheck(name) {
		return true
	}
	colums := getColumns(tx)
	for _, r := range colums {
		if strings.HasPrefix(r.Cmpnt.Type, "multi") {
			return true
		}
	}
	return false
}
func fltrNeedBatchCheck(optr *optrs) bool {
	return optr.NeedBatchCheck()
}
func resetForm(t *Table) *Table {
	for _, c := range t.Columns {
		c.Ext.FormName = "form"
	}
	return t
}
func getTPS(tp string) []string {
	is := strings.Split(tp, "-")
	lst := make([]string, 0, 1)
	for _, v := range is {
		lst = append(lst, strings.ToLower(v))
		lst = append(lst, strings.ToUpper(v))
	}
	return lst
}
func fltrAssctColumns(tx interface{}, colName string) []*Column {
	columns := getColumns(tx)
	cols := make(Columns, 0, 1)
	for _, r := range columns {
		if r.Enum.AssctColumn == "" {
			continue
		}
		if strings.EqualFold(r.Enum.AssctColumn, colName) {
			cols = append(cols, r)
		}
	}
	return cols
}
func fltrTranslate(f string, t interface{}) string {
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
func fltrMYSQLType(c *md.Row) string {
	return mySQLType(c.Type.Name, c.Type.Len, c.Type.DLen)
}
func getFirstOptr(opts []*optrs) *optrs {
	if len(opts) > 0 {
		return opts[0]
	}
	return &optrs{}
}
func getFirstColumns(columns []*Column) *Column {
	if len(columns) > 0 {
		return columns[0]
	}
	return &Column{}
}
