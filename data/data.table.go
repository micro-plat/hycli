package data

import (
	"sort"

	"github.com/micro-plat/hycli/data/internal/md"
)

type Table struct {
	*md.Table

	//Columns 所有列
	Columns Columns

	//数据库索引
	DBIdx *DbIdx

	//Optrs 操作栏组件
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

	//是否是临时表
	IsTmpl bool
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
		Table:   t,
		UNQ:     defFids.Next(),
		Columns: columns,
		Enum:    newEnumType(t.Name.Short, columns),
		Optrs:   NewTableOptrs(),
	}
	t.Cache = table
	table.Optrs.ViewExtCmptOpts = &viewExtOptrsMap{}
	table.Optrs.ListOpts = createLstBarOptrs(table, t.ExtInfo)
	table.Optrs.BarOpts, table.NeedBatchCheck = createQBarOptrs(table, t.ExtInfo)
	table.Optrs.ViewOpts, _ = createViewOptrs(table, t.ExtInfo)
	table.Optrs.LStatOpts, table.Optrs.ChartOpts = createLStatChartOptrs(table.Name.Raw, t.ExtInfo)
	table.Optrs.ExtOpts = createExtOptrs(table.Name.Raw, t.ExtInfo)
	table.Optrs.QueryOptrs = createQueryOptrs(table, t.ExtInfo)
	table.DBIdx = NewDbIdx(createUNQIdx(table), createNormalIdx(table))
	table.Conf = newConfig(t.Settings)
	table.Sort()
	return table
}
func (t *Table) LoadExtOptrs() {
	t.Optrs.SetViewExtCmptOpt(getViewExtCmptOptsByTable(t))
	t.Tag = newTag(t)
}
