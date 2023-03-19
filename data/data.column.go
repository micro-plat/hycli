package data

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

type BaseColumn struct {
	Name      string //字段英文名
	Label     string //字段中文名
	UNQ       string //字段唯一标
	RawConsts []string
	rawRow    *md.Row
	index     int //顺序编号
}
type fieldType struct {
	Name              string //字段原名称
	Type              string //strin,um,decimal,date
	Desc              string
	Required          bool
	SLen              string //字段原长度
	Len               int    //字段原长度
	DLen              int    //？
	DefaultValue      string //默认值
	IsDate            bool
	IsNumber          bool
	IsPK              bool  //是否是主键
	IsSEQ             bool  //是否自增字段
	IsCreateField     bool  //是否是创建字段
	IsUpdateField     bool  //是否是修改字段
	IsSelectField     bool  //是否是查询字段
	IsSelectViewField bool  //是否是查询行的预览
	IsViewField       bool  //是否是预览字段
	IsExtFuncField    bool  //扩展功能字段
	IsFrontQueryField bool  //是否是前端查询字段
	IsBackQueryField  bool  //是否是后端查询字段
	IsDeleteField     bool  //是否是删除字段
	Idxs              []idx //字段索引
}

func createFieldType(r *md.Row) fieldType {
	return fieldType{
		Name:              r.Name,
		Type:              r.Type.Name, //字段原长度
		Len:               r.Type.Len,
		DLen:              r.Type.DLen,  //字段实际长度
		Desc:              r.Desc.Raw,   //字段原说明信息
		Required:          !r.AllowNull, //空值默认值
		DefaultValue:      r.DefValue,   //默认值
		IsDate:            strings.EqualFold(r.Type.Name, "date") || strings.EqualFold(r.Type.Name, "datetime"),
		IsNumber:          strings.EqualFold(r.Type.Name, "number") || strings.EqualFold(r.Type.Name, "int"),
		IsPK:              md.HasConstraint(r.Constraints, "PK", "pk"),   //是否是主键
		IsSEQ:             md.HasConstraint(r.Constraints, "SEQ", "seq"), //是否自增字段
		IsCreateField:     md.HasConstraint(r.Constraints, "C", "c"),     //是否是创建字段
		IsUpdateField:     md.HasConstraint(r.Constraints, "U", "u"),     //是否是
		IsSelectField:     md.HasConstraint(r.Constraints, "L", "l"),     //是否是查询
		IsSelectViewField: md.HasConstraint(r.Constraints, "LE", "le"),   //是否是查询
		IsViewField:       md.HasConstraint(r.Constraints, "V", "v"),     //是否是预览字段
		IsExtFuncField:    strings.HasPrefix(r.Raw, "^"),                 //扩展字段
		IsFrontQueryField: md.HasConstraint(r.Constraints, "Q", "q"),     //是否是查询
		IsBackQueryField:  md.HasConstraint(r.Constraints, "BQ", "bq"),   //是否是后询
		IsDeleteField:     md.HasConstraint(r.Constraints, "D", "d"),
	}
}

type color struct {
	Name      string
	FontColor string
	BgColor   string
}

type idx struct {
	Type  string
	Name  string
	Index int
}

// 显示组件
type displayCmpt struct {
	Type   string //text,link,image,switch,pwd,mobile,progress,file,tag,daterange,date,multiselect,select
	Format string //字段显示格式，如:yyyy-MM-dd
	// Page   string //页面类型

}
type ext struct {
	FormName string
}

type displayCmpts map[string]displayCmpt

func (d displayCmpts) getCmpt(r *md.Row, tps ...string) displayCmpt {
	cmpt := displayCmpt{Type: "text"}
	var c displayCmpt
	var tp string
	var ok bool
	for _, tp = range tps {
		if c, ok = d[strings.ToUpper(tp)]; ok {
			cmpt = c
			break
		}
	}

	if !ok {
		if c, ok := d["*"]; ok {
			cmpt = c
		} else if md.HasConstraint(r.Constraints, "sl", "SL") {
			cmpt = displayCmpt{Type: "select"}
		} else if strings.EqualFold(r.Type.Name, "date") ||
			strings.EqualFold(r.Type.Name, "datetime") {
			cmpt = displayCmpt{Type: "date"}
		}
	}
	if cmpt.Format != "" {
		return cmpt
	}
	cmpt.Format = md.GetFormat(types.GetStringByIndex(tps, 0, tp), r.Constraints...)
	if cmpt.Format != "" {
		return cmpt
	}
	cmpt.Format = md.GetFormat("f", r.Constraints...)
	return cmpt
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

// 解析列表组件
func newLstCMPT(r *md.Row) displayCmpts {
	cmpts := make(map[string]displayCmpt)
	for _, k := range r.Constraints {
		tp, pns, f := md.GetConsByTagIgnorecase("tp", k)
		if tp == "" {
			continue
		}
		pni := strings.Split(pns, "-")
		for _, pn := range pni {
			cmpts[types.GetString(strings.ToUpper(pn), "*")] = displayCmpt{
				Type:   tp,
				Format: f,
			}
		}

	}
	return cmpts
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
	IsEnum   bool   //是否是枚举类型
	EnumType string //枚举名
	PID      string //父级编号
	Group    string //分组
	// JoinName   string //关联字段名
	IsDEColumn bool
}

func createEnumType(r *md.Row) enumType {
	tp, pid, group := md.GetSelectName(r.Name, r.Constraints...)

	return enumType{
		IsEnum:   md.HasConstraint(r.Constraints, "sl", "SL"),
		EnumType: tp,
		PID:      pid,
		Group:    group,
		// JoinName:   fname,
		IsDEColumn: md.HasConstraint(r.Constraints, "DE", "de"),
	}
}

type Column struct { //字段基础息
	BaseColumn
	Field   fieldType    //字段类型
	Enum    enumType     //枚举类型
	QCMPT   displayCmpt  //查询显示组件
	LCMPT   displayCmpt  //列表显示组件
	LECMPT  displayCmpt  //列表详情组件
	VCMPT   displayCmpt  //预览组件
	CCMPT   displayCmpt  //创建页面组件
	UCMPT   displayCmpt  //修改页面组件
	DCMPT   displayCmpt  //删除页面组件
	ExCMPT  displayCmpt  //扩展参数组件
	Style   displayStyle //显示样式
	allCMPT displayCmpts //所有组件
	Ext     ext
}

type Columns []*Column

func (s Columns) Len() int           { return len(s) }
func (s Columns) Less(i, j int) bool { return s[i].index < s[j].index }
func (s Columns) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func newColum(r *md.Row) *Column {
	cmpts := newLstCMPT(r)
	return &Column{
		BaseColumn: BaseColumn{
			Name:      r.Name,
			Label:     r.Desc.Name,
			UNQ:       defFids.Next(),
			RawConsts: r.Constraints,
			rawRow:    r,
			index:     99,
		},
		allCMPT: cmpts,
		QCMPT:   cmpts.getCmpt(r, "q"),
		LCMPT:   cmpts.getCmpt(r, "l", "le"),
		LECMPT:  cmpts.getCmpt(r, "l", "le"),
		VCMPT:   cmpts.getCmpt(r, "v"),
		CCMPT:   cmpts.getCmpt(r, "c", "u"),
		UCMPT:   cmpts.getCmpt(r, "c", "u"),
		DCMPT:   cmpts.getCmpt(r, "d"),
		Enum:    createEnumType(r),
		Field:   createFieldType(r),
		Style:   createStyle(r),
		Ext:     ext{FormName: "form"},
	}
}
func (c *Column) ResetExtCMPT(t string) {
	c.ExCMPT = c.allCMPT.getCmpt(c.rawRow, t)
}
func (c *Column) QIndex(i int) *Column {
	return c.resetIndex("q", i)
}
func (c *Column) LIndex(i int) *Column {
	return c.resetIndex("l", i)
}
func (c *Column) LeIndex(i int) *Column {
	return c.resetIndex("le", i)
}
func (c *Column) VIndex(i int) *Column {
	return c.resetIndex("v", i)
}
func (c *Column) CIndex(i int) *Column {
	return c.resetIndex("c", i)
}
func (c *Column) UIndex(i int) *Column {
	return c.resetIndex("u", i)
}
func (c *Column) DIndex(i int) *Column {
	return c.resetIndex("d", i)
}
func (c *Column) ExtIndex(t string, i int) *Column {
	return c.resetIndex(t, i)
}
func (c *Column) resetIndex(x string, i int) *Column {
	ncol := *c
	fc, _, _ := md.GetConsByTagIgnorecase(fmt.Sprintf("%sidx", x), c.RawConsts...)
	ncol.index = types.GetInt(fc, i+100+2)
	return &ncol
}
