package data

import (
	"sort"

	"github.com/micro-plat/hycli/data/internal/md"
)

type Table struct {
	*md.Table

	//Columns 所有列
	Columns Columns

	//PKColumns 主键列
	PKColumns Columns

	//UNQIndex 唯一索引
	UNQIndex dbIdxs

	//NormalIdx 普通索引
	NormalIdx dbIdxs

	//EnumColumns 枚举列
	EnumColumns Columns

	//ViewOpts 预览操作
	ViewOpts viewOptrs

	//ViewExtCmptOpts 预览页扩展组件
	ViewExtCmptOpts viewExtCmptOpts

	//ListOpts 列表操作
	ListOpts lstOptrs

	//LStatOpts 统计操作
	LStatOpts lstatOptrs

	//ChartOpts 图表组件
	ChartOpts chartOptrs

	ExtOpts extOptrs

	//BarOpts 工具栏操作
	BarOpts barOptrs

	QueryOptrs queryOptrs

	//NeedBatchCheck 是否需要批量选择
	NeedBatchCheck bool

	//Enum 枚举类型
	Enum *EnumType

	//UNQ 唯一标识
	UNQ string

	//Tag 系统功能标签
	Tag *Tag
}

func (t *Table) Sort() {
	sort.Sort(t.Columns)
	sort.Sort(optrslst(t.BarOpts))
	sort.Sort(optrslst(t.ViewExtCmptOpts))
	sort.Sort(optrslst(t.ListOpts))
	sort.Sort(optrslst(t.LStatOpts))
	sort.Sort(optrslst(t.ChartOpts))
	sort.Sort(optrslst(t.ExtOpts))
	sort.Sort(optrslst(t.BarOpts))
}
func NewTable(t *md.Table) *Table {
	if t.Cache != nil {
		return t.Cache.(*Table)
	}

	//构建列
	columns := newColumns(t.Rows)
	table := &Table{
		Table:       t,
		UNQ:         defFids.Next(),
		Columns:     columns,
		PKColumns:   columns.GetPKColumns(),
		EnumColumns: columns.GetEnumColumns(),
		Enum:        newEnumType(t.Name.Short, t.Rows, columns.GetEnumDelColumns()),
	}
	t.Cache = table
	table.ListOpts = createLstOptrs(table, t.ExtInfo)
	table.BarOpts, table.NeedBatchCheck = createBarOptrs(table, t.ExtInfo)
	table.ViewOpts, table.ViewExtCmptOpts = createViewOptrs(table.Name.Raw, t.ExtInfo)
	table.LStatOpts, table.ChartOpts = createLStatChartOptrs(table.Name.Raw, t.ExtInfo)
	table.ExtOpts = createExtOptrs(table.Name.Raw, t.ExtInfo)
	table.QueryOptrs = createQueryOptrs(table, t.ExtInfo)
	table.NormalIdx = createNormalIdx(table)
	table.UNQIndex = createUNQIdx(table)
	table.Tag = newTag(table)
	table.Sort()
	return table
}
