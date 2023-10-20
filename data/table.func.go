package data

import (
	"sort"
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

// HasStaticColumn 是否包含静态字段(从用户信息、指定固定值)取值
func (tb *Table) HasStaticColumn(tps, prefix string) bool {
	tplst := strings.Split(tps, "-")
	for _, c := range tb.Columns {
		for _, tp := range tplst {
			f := md.GetFormat(tp, c.Constraints...)
			if strings.HasPrefix(f, prefix) {
				return true
			}
		}
	}
	return false
}

// GetStaticColumns 获取包含指定组件和包含某个静态字段(从用户信息、指定固定值)的列
func (tb *Table) GetStaticColumns(tps, prefix string) map[string]*Column {
	columns := make(map[string]*Column, 1)
	tplst := strings.Split(tps, "-")
	for _, c := range tb.Columns {
		for _, tp := range tplst {
			f := md.GetFormat(tp, c.Constraints...)
			if strings.HasPrefix(f, prefix) {
				columns[f] = c
			}
		}
	}
	return columns
}
func (tx *Table) GetColumsByCmpnt(cmpnt string, tps ...string) []*Column {
	return getCmpnt(tx, cmpnt, tps...)
}
func (tx *TTable) GetColumsByCmpnt(cmpnt string, tps ...string) []*Column {
	return getCmpnt(tx, cmpnt, tps...)
}
func (tx Columns) GetColumsByCmpnt(cmpnt string, tps ...string) []*Column {
	return getCmpnt(tx, cmpnt, tps...)
}
func (tx *Table) JoinNames(tp string, required bool, start string, end ...string) string {
	colums := tx.GetColumnsByTPName(tp)
	return colums.JoinNames(tp, required, start, end...)
}
func (tx *TTable) JoinNames(tp string, required bool, start string, end ...string) string {
	colums := tx.GetColumnsByTPName(tp)
	return colums.JoinNames(tp, required, start, end...)
}

func getCmpnt(tx interface{}, cmpnt string, tps ...string) []*Column {
	colums := getColumns(tx)
	cols := make(Columns, 0, 1)
	cmpnts := strings.Split(cmpnt, "-")
	for _, r := range colums {
		if !md.HasConstraint(r.Constraints, tps...) {
			continue
		}
		cmp := r.allCmpnt.getCmpnt(r.Row, tps...)
		if types.StringContains(cmpnts, cmp.Type) {
			r.ResetCmpnt(tps...)
			cols = append(cols, r)
		}

	}
	sort.Sort(cols)
	return cols
}

func (tx *Table) GetColumnsByTPName(tp string, formName ...string) Columns {
	return getColumnsByTPName(tx, tp, false, false, formName...)
}
func (tx *Table) GetColumnsByName(colName string) *Column {
	for _, c := range tx.Columns {
		if strings.EqualFold(c.Name, colName) {
			return c
		}
	}
	return &Column{}
}
func (tx *TTable) GetColumnsByTPName(tp string, formName ...string) Columns {
	return getColumnsByTPName(tx, tp, false, false, formName...)
}
func (tx *Table) GetRequiredColumnsByName(tp string, formName ...string) Columns {
	return getColumnsByTPName(tx, tp, true, false, formName...)
}
func (tx *TTable) GetRequiredColumnsByName(tp string, formName ...string) Columns {
	return getColumnsByTPName(tx, tp, true, false, formName...)
}
func (tx *Table) GetValidColumnsByName(tp string, formName ...string) Columns {
	return getColumnsByTPName(tx, tp, false, true, formName...)
}

func getColumnsByTPName(tx interface{}, tp string, required bool, excludedField bool, formName ...string) Columns {
	colums := getColumns(tx)

	cols := make(Columns, 0, 1)
	tps := getTPS(tp)
	for _, r := range colums {
		if excludedField && r.Field.IsExtFuncField {
			continue
		}
		if required && !r.Field.Required {
			continue
		}
		if md.HasConstraint(r.RawConsts, tps...) {
			r.ResetCmpnt(tps...) //重置扩展组件
			r.Ext.FormName = types.GetStringByIndex(formName, 0, "form")
			cols = append(cols, r.Index(types.GetStringByIndex(tps, 0), len(cols)))
		}
	}
	sort.Sort(cols)
	return cols
}

func (t *Table) IsMutilValue(optr *optrs, name string) bool {
	if optr.Params.IsBatchCheck(name) {
		return true
	}
	for _, r := range t.Columns {
		if strings.HasPrefix(r.Cmpnt.Type, "multi") {
			return true
		}
	}
	return false
}
func (t *Table) GetKeyParams(tp string) []string {
	for _, c := range t.Columns {
		if md.HasConstraint(c.Constraints, tp) {
			return md.GetConstrainValues(tp, c.Constraints...)
		}

	}
	return nil

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
func getTPS(tp string) []string {
	is := strings.Split(tp, "-")
	lst := make([]string, 0, 1)
	for _, v := range is {
		lst = append(lst, strings.ToLower(v))
		lst = append(lst, strings.ToUpper(v))
	}
	return lst
}
