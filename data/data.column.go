package data

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/jsons"
	"github.com/micro-plat/lib4go/types"
)

const (
	VIEW_COLUMN   = "V"
	VIEW_TAG      = "VIEW"
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
type fieldType struct {
	Name           string //字段原名称
	Type           string //string,um,decimal,date
	Desc           string
	Required       bool
	Len            int    //字段原长度
	DLen           int    //字段第二长度g
	DefaultValue   string //默认值
	IsDate         bool
	IsNumber       bool
	IsPK           bool   //是否是主键
	IsSEQ          bool   //是否自增字段
	IsExtFuncField bool   //扩展功能字段
	IsLongText     bool   //是否是长字符
	TMYSQL         string //mysql 数据库类型
	VMYSQL         string //mysql 默认值
	row            *md.Row
}

var longText = []string{"tinytext", "text", "mediumtext", "longtext", "blob", "json", "TINYTEXT", "TEXT", "MEDIUMTEXT", "LONGTEXT", "BLOB", "JSON"}

func createFieldType(r *md.Row) fieldType {
	return fieldType{
		row:            r,
		Name:           r.Name,
		Type:           r.Type.Name, //字段原长度
		Len:            r.Type.Len,
		DLen:           r.Type.DLen,  //字段实际长度
		Desc:           r.Desc.Raw,   //字段原说明信息
		Required:       !r.AllowNull, //空值默认值
		DefaultValue:   r.DefValue,   //默认值
		IsDate:         strings.EqualFold(r.Type.Name, "date") || strings.EqualFold(r.Type.Name, "datetime"),
		IsNumber:       strings.EqualFold(r.Type.Name, "number") || strings.EqualFold(r.Type.Name, "int"),
		IsPK:           md.HasConstraint(r.Constraints, "PK", "pk"),   //是否是主键
		IsSEQ:          md.HasConstraint(r.Constraints, "SEQ", "seq"), //是否自增字段
		IsExtFuncField: strings.HasPrefix(r.Raw, "^"),                 //扩展字段
		TMYSQL:         mySQLType(r.Type.Name, r.Type.Len, r.Type.DLen),
		VMYSQL:         f_mysql_get_def_value(r.DefValue),
		IsLongText:     r.Type.Len >= 64 || types.StringContains(longText, r.Type.Name),
	}
}

type color struct {
	Name      string
	FontColor string
	BgColor   string
}

// 显示组件
type displayCmpnt struct {
	Type   string //text,link,image,switch,pwd,mobile,progress,file,tag,daterange,date,multiselect,select
	Format string //字段显示格式，如:yyyy-MM-dd

}

func (s displayCmpnt) String() string {
	buff, _ := jsons.Marshal(s)
	return string(buff)
}

type ext struct {
	FormName string
}

type displayCmpnts map[string]displayCmpnt

func (d displayCmpnts) getCmpnt(r *md.Row, tps ...string) displayCmpnt {
	cmpnt := displayCmpnt{Type: "text"}
	var c displayCmpnt
	var tp string
	var ok bool
	for _, tp = range tps {
		if c, ok = d[strings.ToUpper(tp)]; ok {
			cmpnt = c
			break
		}
	}

	if !ok {
		if c, ok := d["*"]; ok {
			cmpnt = c
		} else if md.HasConstraint(r.Constraints, "sl", "SL") {
			cmpnt = displayCmpnt{Type: "select"}
		} else if strings.EqualFold(r.Type.Name, "date") ||
			strings.EqualFold(r.Type.Name, "datetime") {
			cmpnt = displayCmpnt{Type: "date"}
		}
	}
	if cmpnt.Format != "" {
		return cmpnt
	}
	cmpnt.Format = md.GetFormat(types.GetStringByIndex(tps, 0, tp), r.Constraints...)
	if cmpnt.Format != "" {
		return cmpnt
	}
	cmpnt.Format = md.GetFormat("f", r.Constraints...)
	return cmpnt
}

// 解析列表组件
func newLstCmpnt(r *md.Row) displayCmpnts {
	cmpnts := make(map[string]displayCmpnt)
	for _, k := range r.Constraints {
		tp, pns, f := md.GetConsByTagIgnorecase("tp", k)
		if tp == "" {
			continue
		}
		pni := strings.Split(pns, "-")
		for _, pn := range pni {
			cmpnts[types.GetString(strings.ToUpper(pn), "*")] = displayCmpnt{
				Type:   tp,
				Format: f,
			}
		}

	}
	return cmpnts
}

// 显示样式
type displayStyle struct {
	ListWidth    string //列表宽度
	Rows         int    //显示的行数，如textArea行数
	Position     string //位置 default,默认，换行newline
	HideOverflow bool   //超出隐藏
	CLR          color  //颜色信息
	Min          int    //最小值
	Max          int    //最大值
}

// 解析列表样式
func createStyle(r *md.Row) displayStyle {
	min, max := md.GetRanges(r.Constraints...)
	fc, bc, _ := md.GetConsByTagIgnorecase("color", r.Constraints...)
	bgcolor := md.HasConstraint(r.Constraints, "color")
	if bgcolor && bc == "" {
		bc = "colorful"
	}
	return displayStyle{
		ListWidth:    md.GetConsNameByTagIgnorecase("lw", r.Constraints...),
		Rows:         types.GetInt(md.GetConsNameByTagIgnorecase("row", r.Constraints...)),
		Position:     types.GetString(md.GetConsNameByTagIgnorecase("ps", r.Constraints...)),
		HideOverflow: md.HasConstraint(r.Constraints, "hof", "HOF"),
		Min:          types.GetInt(min),
		Max:          types.GetInt(max),
		CLR: color{
			Name:      "colorful",
			FontColor: fc,
			BgColor:   bc,
		},
	}
}

type enumType struct {
	IsEnum      bool   //是否是枚举类型
	EnumType    string //枚举名
	AssctColumn string //关联列
	PID         string //父级编号
	Group       string //分组
	IsDEColumn  bool
}

func createEnumType(r *md.Row) enumType {
	tp, pid, group := md.GetSelectName(r.Name, r.Constraints...)
	if pid == "-" {
		pid = ""
	}

	return enumType{
		IsEnum:      md.HasConstraint(r.Constraints, "sl", "SL"),
		EnumType:    tp,
		AssctColumn: types.DecodeString(strings.HasPrefix(pid, "@"), true, strings.Trim(pid, "@"), ""),
		PID:         types.DecodeString(strings.HasPrefix(pid, "@"), false, pid, ""),
		Group:       group,
		IsDEColumn:  md.HasConstraint(r.Constraints, "DE", "de"),
	}
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
func (c *Column) NIndex(i int) *Column {
	ncol := *c
	ncol.index = i
	return &ncol
}
