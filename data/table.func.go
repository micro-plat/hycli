package data

import (
	"sort"
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

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
func (tb *Table) GetStaticColumn(tps, prefix string) map[string]*Column {
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
func (tx *Table) JoinNames(tp string, start string, end ...string) string {
	colums := tx.GetColumnsByName(tp)
	return colums.JoinNames(tp, start, end...)
}
func (tx *TTable) JoinNames(tp string, start string, end ...string) string {
	colums := tx.GetColumnsByName(tp)
	return colums.JoinNames(tp, start, end...)
}

func getCmpnt(tx interface{}, cmpnt string, tps ...string) []*Column {
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

func (tx *Table) GetColumnsByName(tp string, formName ...string) Columns {
	return getColumnsByName(tx, tp, false, false, formName...)
}
func (tx *TTable) GetColumnsByName(tp string, formName ...string) Columns {
	return getColumnsByName(tx, tp, false, false, formName...)
}
func (tx *Table) GetRequiredColumnsByName(tp string, formName ...string) Columns {
	return getColumnsByName(tx, tp, true, false, formName...)
}
func (tx *TTable) GetRequiredColumnsByName(tp string, formName ...string) Columns {
	return getColumnsByName(tx, tp, true, false, formName...)
}
func (tx *Table) GetValidColumnsByName(tp string, formName ...string) Columns {
	return getColumnsByName(tx, tp, false, true, formName...)
}

func getColumnsByName(tx interface{}, tp string, required bool, excludedField bool, formName ...string) Columns {
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

func (tx *Table) GetColumnsByColumName(colName string) []*Column {
	return getColumnsByColumName(tx.Columns, colName)
}
func (tx *TTable) GetColumnsByColumName(colName string) []*Column {
	return getColumnsByColumName(tx.Columns, colName)
}
func getColumnsByColumName(columns Columns, colName string) []*Column {
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
