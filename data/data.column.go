package data

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

const (
	VIEW_COLUMN   = "V"
	VIEW_TAG      = "view" //11/26
	ADD_COLUMN    = "C"
	ADD_TAG       = "ADD"
	UPDATE_COLUMN = "U"
	UPDATE_TAG    = "UPDATE"
	DELETE_COLUMN = "D"
	DELETE_TAG    = "DEL"
	TAB_TAG       = "TAB"
	QUERY_COLUMN  = "Q"
	QUERY_TAG     = "QBAR"
	CNFRM         = "CNFRM"
)

type BaseColumn struct {
	*md.Row
	Name      string //字段英文名
	Label     string //字段中文名
	UNQ       string //字段唯一标
	RawConsts []string
	index     int //顺序编号
}

type Column struct { //字段基础息
	BaseColumn
	Field    fieldType     //字段类型
	Enum     enumType      //枚举类型
	Cmpnt    displayCmpnt  //扩展参数组件
	Style    displayStyle  //显示样式
	allCmpnt displayCmpnts //所有组件
	Ext      ext
}

type Columns []*Column

func (s Columns) Len() int           { return len(s) }
func (s Columns) Less(i, j int) bool { return s[i].index < s[j].index }
func (s Columns) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Columns) SLen() int {
	return len(s) - 1
}
func newColumns(r []*md.Row) Columns {
	cs := make([]*Column, 0, 1)
	for _, v := range r {
		cs = append(cs, newColum(v))
	}
	return cs
}
func (col *Column) NeedTriggerChangeEvent(currentColName string) bool {
	return col.Enum.AssctColumn == currentColName ||
		md.HasConstraint(col.Constraints, "@change") && col.Name == currentColName
}
func (col *Column) IsAssctColumn(assctName string) bool {
	return col.Enum.AssctColumn != col.Name && strings.EqualFold(col.Enum.AssctColumn, assctName)
}

func newColum(r *md.Row) *Column {
	cmpnts := newLstCmpnt(r)
	return &Column{
		BaseColumn: BaseColumn{
			Name:      r.Name,
			Label:     r.Desc.Name,
			UNQ:       defFids.Next(),
			RawConsts: r.Constraints,
			Row:       r,
			index:     99,
		},
		allCmpnt: cmpnts,
		Enum:     createEnumType(r),
		Field:    createFieldType(r),
		Style:    createStyle(r),
		Ext:      ext{FormName: "form"},
	}
}

func (c *Column) ResetCmpnt(t ...string) {
	c.Cmpnt = c.allCmpnt.getCmpnt(c.Row, t...)
}
func (c *Column) Index(t string, i int) *Column {
	ncol := *c
	fc, _, _ := md.GetConsByTagIgnorecase(fmt.Sprintf("%sidx", t), c.RawConsts...)
	ncol.index = types.GetInt(fc, i+100+2)
	return &ncol
}
func (c *Column) NewIndex(i int) *Column {
	ncol := *c
	ncol.index = i
	return &ncol
}
