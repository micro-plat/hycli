package data

import (
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

type BaseColumn struct {
	Name      string //字段英文名
	Label     string //字段中文名
	UNQ       string //字段唯一标
	RawConsts []string
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
		Name:              strings.TrimLeft(r.Name, "_"),
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
		IsSelectField:     md.HasConstraint(r.Constraints, "LV", "l"),    //是否是查询
		IsSelectViewField: md.HasConstraint(r.Constraints, "LE", "le"),   //是否是查询
		IsViewField:       md.HasConstraint(r.Constraints, "V", "v"),     //是否是预览字段
		IsExtFuncField:    strings.HasPrefix(r.Name, "_"),                //扩展字段
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
	Page   string //页面类型

}
type ext struct {
	FormName string
}

type displayCmpts map[string]displayCmpt

func (d displayCmpts) getCmpt(tp string) displayCmpt {
	if c, ok := d[strings.ToUpper(tp)]; ok {
		return c
	}
	if c, ok := d["*"]; ok {
		return c
	}
	return displayCmpt{Type: "text"}
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
		tp, pn := md.GetConsByTagIgnorecase("tp", k)
		if tp == "" {
			continue
		}
		cmpts[types.GetString(strings.ToUpper(pn), "*")] = displayCmpt{
			Type: tp,
			Page: pn,
		}
	}
	return cmpts
}

// 解析列表样式
func createStyle(r *md.Row) displayStyle {
	min, max := md.GetRanges(r.Constraints...)
	fc, bc := md.GetConsByTagIgnorecase("color", r.Constraints...)
	return displayStyle{
		ListWidth:    md.GetConsNameByTagIgnorecase("lw", r.Constraints...),
		Rows:         types.GetInt(md.GetConsNameByTagIgnorecase("row", r.Constraints...)),
		Position:     types.GetString(md.GetConsNameByTagIgnorecase("ps")),
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
	JoinName string //关联字段名
}

func createEnumType(r *md.Row) enumType {
	sl, fname := md.GetSelectName(r.Name, r.Constraints...)
	return enumType{
		IsEnum:   sl != "",
		EnumType: sl,
		JoinName: fname,
	}
}

type Column struct { //字段基础息
	BaseColumn
	Field  fieldType    //字段类型
	Enum   enumType     //枚举类型
	QCMPT  displayCmpt  //查询显示组件
	LCMPT  displayCmpt  //列表显示组件
	LECMPT displayCmpt  //列表详情组件
	VCMPT  displayCmpt  //预览组件
	CCMPT  displayCmpt  //创建页面组件
	UCMPT  displayCmpt  //修改页面组件
	DCMPT  displayCmpt  //删除页面组件
	Style  displayStyle //显示样式
	Ext    ext
}

func newColum(r *md.Row) *Column {
	cmpts := newLstCMPT(r)
	return &Column{
		BaseColumn: BaseColumn{
			Name:      r.Name,
			Label:     r.Desc.Name,
			UNQ:       defFids.Next(),
			RawConsts: r.Constraints,
		},
		QCMPT:  cmpts.getCmpt("q"),
		LCMPT:  cmpts.getCmpt("l"),
		LECMPT: cmpts.getCmpt("le"),
		VCMPT:  cmpts.getCmpt("v"),
		CCMPT:  cmpts.getCmpt("c"),
		UCMPT:  cmpts.getCmpt("u"),
		DCMPT:  cmpts.getCmpt("d"),
		Enum:   createEnumType(r),
		Field:  createFieldType(r),
		Style:  createStyle(r),
	}
}
