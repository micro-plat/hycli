package data

import (
	"fmt"
	"sort"
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

// fltrColums 过滤用户自定义类型对应的行，自定义行对应的控件按新增模式处理
func fltrColums(tx interface{}, tp string, formName ...string) []*Column {
	t, ok := tx.(*Table)
	if !ok {
		t = tx.(*TTable).Table
	}

	cols := make(Columns, 0, 1)
	tps := getTPS(tp)
	for _, r := range t.Colums {
		if md.HasConstraint(r.RawConsts, tps...) {
			r.ResetCmpnt(tps...) //重置扩展组件
			r.Ext.FormName = types.GetStringByIndex(formName, 0, "form")
			cols = append(cols, r.Index(types.GetStringByIndex(tps, 0), len(cols)))
		}
	}
	sort.Sort(cols)
	return cols
}
func fltrCmpnt(tx interface{}, cmpnt string, tps ...string) []*Column {
	t, ok := tx.(*Table)
	if !ok {
		t = tx.(*TTable).Table
	}
	cols := make(Columns, 0, 1)
	for _, r := range t.Colums {
		if r.allCmpnt.getCmpnt(r.rawRow, tps...).Type == cmpnt {
			r.ResetCmpnt(tps...)
			cols = append(cols, r)
		}
	}
	sort.Sort(cols)
	return cols
}

func resetForm(t *Table) *Table {
	for _, c := range t.Colums {
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
func fltrAssctColums(tx interface{}, colName string) []*Column {
	colums := getColums(tx)
	cols := make(Columns, 0, 1)
	for _, r := range colums {
		if r.Enum.AssctColumn == "" {
			continue
		}
		fmt.Println("r.Enum.AssctColumn", colName)
		if strings.EqualFold(r.Enum.AssctColumn, colName) {
			cols = append(cols, r)
		}
	}
	return cols
}
func getColums(tx interface{}) []*Column {
	if t, ok := tx.(*Table); ok {
		return t.Colums
	}
	if t, ok := tx.(*TTable); ok {
		return t.Colums
	}
	if t, ok := tx.([]*Column); ok {
		return t
	}
	return nil

}
