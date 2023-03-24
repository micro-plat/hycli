package data

import (
	"fmt"
	"sort"

	"github.com/micro-plat/hycli/data/internal/md"
)

type Table struct {
	*md.Table

	//Colums 所有列
	Colums Columns

	//PKColums 主键列
	PKColums Columns

	//EnumColums 枚举列
	EnumColums Columns

	//ViewOpts 预览操作
	ViewOpts viewOptrs

	//ListOpts 列表操作
	ListOpts lstOptrs

	//LStatOpts 统计操作
	LStatOpts lstatOptrs

	ChartOpts chartOptrs

	//BarOpts 工具栏操作
	BarOpts barOptrs

	ExtCmptOpts extCmptOpts

	//NeedBatchCheck 是否需要批量选择
	NeedBatchCheck bool

	//Enum 枚举类型
	Enum *EnumType

	//UNQ 唯一标识
	UNQ string
}

func (t *Table) Sort() {
	sort.Sort(t.Colums)
}
func NewTable(t *md.Table) *Table {
	if t.Cache != nil {
		return t.Cache.(*Table)
	}

	//构建列
	colums := newColums(t.Rows)
	enum := newEnumType(t.Name.Short, t.Rows)
	enum.DEColums = colums.GetEnumDelColumns()
	barOpts := createBarOptrs(t.ExtInfo)
	table := &Table{
		Table:          t,
		UNQ:            defFids.Next(),
		Enum:           enum,
		ListOpts:       createLstOptrs(t.ExtInfo),
		ViewOpts:       createViewOptrs(t.ExtInfo),
		LStatOpts:      createLStatOptrs(t.ExtInfo),
		ChartOpts:      createChartOptrs(t.ExtInfo),
		ExtCmptOpts:    make(extCmptOpts, 0, 1),
		BarOpts:        barOpts,
		NeedBatchCheck: barOpts.NeedCheck(t.Name.Raw),
		Colums:         colums,
		PKColums:       colums.GetPKColumns(),
		EnumColums:     colums.GetEnumColumns(),
	}
	t.Cache = table

	//构建操作
	if len(fltrColums(table, VIEW_COLUMN)) > 0 {
		table.ListOpts = append(table.ListOpts, detailOpts)
	}
	if len(fltrColums(table, UPDATE_COLUMN)) > 0 {
		table.ListOpts = append(table.ListOpts, updateOpts)
	}
	if len(fltrColums(table, DELETE_COLUMN)) > 0 {
		table.ListOpts = append(table.ListOpts, delOpts)
	}

	table.ExtCmptOpts = creatExtCmptOpts(table.ViewOpts)
	if len(table.ExtCmptOpts) > 0 {
		fmt.Println("ext.cmpt.opts:", table.Name, table.ExtCmptOpts[0].URL)
	}

	table.Sort()
	return table
}
