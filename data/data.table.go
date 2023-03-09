package data

import (
	"github.com/micro-plat/hycli/data/internal/md"
)

type Table struct {
	*md.Table
	Colums     []*Column `json:'cols'`      //所有列
	QColums    []*Column `json:'qcols'`     //查询列 --前后端
	BQColums   []*Column `json:'qcols'`     //查询列--后端
	LColums    []*Column `json:'lcols'`     //列表列
	LEColums   []*Column `json:'lecols'`    //列表扩展列
	LLEColums  []*Column `json:'llecols'`   //列表扩展列
	CColums    []*Column `json:'ccols'`     //创建列
	UColums    []*Column `json:'ucols'`     //更新列
	VColums    []*Column `json:lum'vcols'`  //预览列
	DColums    []*Column `json:'dcols'`     //删除列
	ExColums   []*Column `json:'dcols'`     //删除列
	PKColums   []*Column `json:'pk_cols'`   //主键列
	EnumColums []*Column `json:'enum_cols'` //主键列
	PKColum    *Column   //主键列
	ViewOpts   viewOptrs //预览操作
	ListOpts   lstOptrs  //列表操作
	Enum       *EnumType
	UNQ        string `json:'unq'`
}

func NewTable(t *md.Table) *Table {
	if t.Cache != nil {
		return t.Cache.(*Table)
	}
	table := &Table{
		Table:      t,
		Colums:     make([]*Column, 0, 1),
		QColums:    make([]*Column, 0, 1),
		BQColums:   make([]*Column, 0, 1),
		LColums:    make([]*Column, 0, 1),
		LEColums:   make([]*Column, 0, 1),
		LLEColums:  make([]*Column, 0, 1),
		CColums:    make([]*Column, 0, 1),
		UColums:    make([]*Column, 0, 1),
		VColums:    make([]*Column, 0, 1),
		DColums:    make([]*Column, 0, 1),
		PKColums:   make([]*Column, 0, 1),
		EnumColums: make([]*Column, 0, 1),
		PKColum:    &Column{},
		UNQ:        defFids.Next(),
	}
	t.Cache = table
	table.Enum = newEnumType(t.Name.Short, t.Rows)
	for _, row := range t.Rows {
		col := newColum(row)
		table.Colums = append(table.Colums, col)
		if col.Field.IsCreateField {
			table.CColums = append(table.CColums, col)
		}
		if col.Field.IsSelectField {
			table.LColums = append(table.LColums, col)
		}
		if col.Field.IsSelectViewField {
			table.LEColums = append(table.LEColums, col)
		}
		if col.Field.IsSelectField || col.Field.IsSelectViewField {
			table.LLEColums = append(table.LLEColums, col)
		}
		if col.Field.IsViewField {
			table.VColums = append(table.VColums, col)
		}
		if col.Field.IsUpdateField {
			table.UColums = append(table.UColums, col)
		}
		if col.Field.IsDeleteField {
			table.DColums = append(table.DColums, col)
		}
		if col.Field.IsBackQueryField {
			table.BQColums = append(table.BQColums, col)
		}
		if col.Field.IsFrontQueryField {
			table.QColums = append(table.QColums, col)
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
	if len(table.VColums) > 0 {
		table.ListOpts = append(table.ListOpts, detailOpts)
	}
	if len(table.UColums) > 0 {
		table.ListOpts = append(table.ListOpts, updateOpts)
	}
	if len(table.DColums) > 0 {
		table.ListOpts = append(table.ListOpts, delOpts)
	}

	return table
}
