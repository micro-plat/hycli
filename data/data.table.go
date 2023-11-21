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

	//Optrs 所有操作栏
	Optrs *tableOptrs

	//NeedBatchCheck 是否需要批量选择
	NeedBatchCheck bool

	//Conf 配置
	Conf *config

	//Enum 枚举类型
	Enum *EnumType

	//UNQ 唯一标识
	UNQ string

	//Tag 系统功能标签
	Tag *Tag
}

func (t *Table) Sort() {
	sort.Sort(t.Columns)
	t.Optrs.Sort()
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
		Optrs:       NewTableOptrs(),
	}
	t.Cache = table
	table.Optrs.ViewExtCmptOpts = &viewExtOptrsMap{}
	table.Optrs.ListOpts = createLstBarOptrs(table, t.ExtInfo)
	table.Optrs.BarOpts, table.NeedBatchCheck = createQBarOptrs(table, t.ExtInfo)
	table.Optrs.ViewOpts, _ = createViewOptrs(table, t.ExtInfo)
	table.Optrs.LStatOpts, table.Optrs.ChartOpts = createLStatChartOptrs(table.Name.Raw, t.ExtInfo)
	table.Optrs.ExtOpts = createExtOptrs(table.Name.Raw, t.ExtInfo)
	table.Optrs.QueryOptrs = createQueryOptrs(table, t.ExtInfo)
	table.NormalIdx = createNormalIdx(table)
	table.UNQIndex = createUNQIdx(table)
	table.Conf = newConfig(t.Settings)
	table.Sort()
	return table
}
func (t *Table) LoadExtOptrs() {
	t.Optrs.SetViewExtCmptOpt(getViewExtCmptOptsByTable(t))
	t.Tag = newTag(t)
}
