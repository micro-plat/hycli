package data

import (
	"sort"

	"github.com/micro-plat/hycli/data/internal/md"
)

type Table struct {
	*md.Table
	Colums         Columns   `json:'cols'`      //所有列
	QColums        Columns   `json:'qcols'`     //查询列 --前后端
	BQColums       Columns   `json:'qcols'`     //查询列--后端
	LColums        Columns   `json:'lcols'`     //列表列
	LEColums       Columns   `json:'lecols'`    //列表扩展列
	LLEColums      Columns   `json:'llecols'`   //列表扩展列
	CColums        Columns   `json:'ccols'`     //创建列
	UColums        Columns   `json:'ucols'`     //更新列
	VColums        Columns   `json:'vcols'`     //预览列
	DColums        Columns   `json:'dcols'`     //删除列
	ExColums       Columns   `json:'dcols'`     //删除列
	PKColums       Columns   `json:'pk_cols'`   //主键列
	EnumColums     Columns   `json:'enum_cols'` //主键列
	PKColum        *Column   //主键列
	ViewOpts       viewOptrs //预览操作
	ListOpts       lstOptrs  //列表操作
	LStatOpts      lstatOptrs
	BatchOpts      batchOptrs
	NeedBatchCheck bool
	Enum           *EnumType
	UNQ            string `json:'unq'`
}

func (t *Table) Sort() {
	sort.Sort(t.Colums)
	sort.Sort(t.QColums)
	sort.Sort(t.BQColums)
	sort.Sort(t.LColums)
	sort.Sort(t.LEColums)
	sort.Sort(t.LLEColums)
	sort.Sort(t.CColums)
	sort.Sort(t.UColums)
	sort.Sort(t.VColums)
	sort.Sort(t.DColums)
}
func NewTable(t *md.Table) *Table {
	if t.Cache != nil {
		return t.Cache.(*Table)
	}
	table := &Table{
		Table:      t,
		Colums:     make(Columns, 0, 1),
		QColums:    make(Columns, 0, 1),
		BQColums:   make(Columns, 0, 1),
		LColums:    make(Columns, 0, 1),
		LEColums:   make(Columns, 0, 1),
		LLEColums:  make(Columns, 0, 1),
		CColums:    make(Columns, 0, 1),
		UColums:    make(Columns, 0, 1),
		VColums:    make(Columns, 0, 1),
		DColums:    make(Columns, 0, 1),
		PKColums:   make(Columns, 0, 1),
		EnumColums: make(Columns, 0, 1),
		PKColum:    &Column{},
		UNQ:        defFids.Next(),
	}
	t.Cache = table
	table.Enum = newEnumType(t.Name.Short, t.Rows)
	for _, row := range t.Rows {
		col := newColum(row)
		table.Colums = append(table.Colums, col)
		if col.Field.IsCreateField {
			table.CColums = append(table.CColums, col.CIndex(len(table.CColums)))
		}
		if col.Field.IsSelectField {
			table.LColums = append(table.LColums, col.LIndex(len(table.LColums)))
			table.LLEColums = append(table.LLEColums, col.LIndex(len(table.LLEColums)))
		}
		if col.Field.IsSelectViewField {
			table.LEColums = append(table.LEColums, col.LeIndex(len(table.LEColums)))
			table.LLEColums = append(table.LLEColums, col.LeIndex(len(table.LLEColums)))
		}
		if col.Field.IsViewField {
			table.VColums = append(table.VColums, col.VIndex(len(table.VColums)))
		}
		if col.Field.IsUpdateField {
			table.UColums = append(table.UColums, col.UIndex(len(table.UColums)))
		}
		if col.Field.IsDeleteField {
			table.DColums = append(table.DColums, col.DIndex(len(table.DColums)))
		}
		if col.Field.IsBackQueryField {
			table.BQColums = append(table.BQColums, col)
		}
		if col.Field.IsFrontQueryField {
			table.QColums = append(table.QColums, col.QIndex(len(table.QColums)))
			table.BQColums = append(table.BQColums, col)
		}
		if col.Field.IsPK {
			table.PKColum = col
			table.PKColums = append(table.PKColums, col)
		}
		if col.Enum.IsEnum {
			table.EnumColums = append(table.EnumColums, col)
		}
		if col.Enum.IsDEColumn {
			table.Enum.DEColums = append(table.Enum.DEColums, col)
		}

	}
	table.ListOpts = createLstOptrs(t.ExtInfo)
	table.ViewOpts = createViewOptrs(t.ExtInfo)
	table.LStatOpts = createLStatOptrs(t.ExtInfo)
	table.BatchOpts = createBatchOptrs(t.ExtInfo)
	table.NeedBatchCheck = table.BatchOpts.NeedCheck(table.Name.Raw)
	if len(table.VColums) > 0 {
		table.ListOpts = append(table.ListOpts, detailOpts)
	}
	if len(table.UColums) > 0 {
		table.ListOpts = append(table.ListOpts, updateOpts)
	}
	if len(table.DColums) > 0 {
		table.ListOpts = append(table.ListOpts, delOpts)
	}
	table.Sort()
	return table
}
