package data

import (
	
	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/ib4go/types"
)

type BaseColumn struct {
	ame   string //字段英文名
	Label string //字段中文名
	UNQ   string //字段唯一标识
}
type fieldType struct {
	Name              string //字段原名称
	Type              string //string,um,decimal,date
	SLen              string //字段原长度
	DefaultValue      string //默认值
	IsPK              bool   //是否是主键
	IsSEQ             bool   //是否自增字段
	IsCreateField     bool   //是否是创建字段
	IsUpdateField     bool   //是否是修改字段
	IsSelectField     bool   //是否是查询字段
	IsSelectViewField bool   //是否是查询行的预览
	IsExtFuncField    bool   //扩展功能字段
	IsFrontQueryField bool   //是否是前端查询字段
	IsBackQueryField  bool   //是否是后端查询字段
	IsDeleteField     bool   //是否是删除字段
	Idxs              []idx  //字段索引
}

func createFieldType(r *md.Row) fieldType {
	return fieldType{
		Name:              r.Name,
		Type:              r.Type.Name,
		SLen:              "",                                            //字段原长度
		Len:               0,                                             //字段实际长度
		Desc:              r.Desc.Raw,                                    //字段原说明信息
		AllowNull:         r.AllowNull,                                   //空值默认值
		DefaultValue:      r.DefValue,                                    //默认值
		IsPK:              md.HasConstraint(r.Constraints, "PK", "pk"),   //是否是主键
		IsSEQ:             md.HasConstraint(r.Constraints, "SEQ", "seq"), //是否自增字段
		IsCreateField:     md.HasConstraint(r.Constraints, "C", "c"),     //是否是创建字段
		IsUpdateField:     md.HasConstraint(r.Constraints, "U", "c"),     //是否是
		IsSelectField:     md.HasConstraint(r.Constraints, "L", "l"),     //是否是查询
		IsSelectViewField: md.HasConstraint(r.Constraints, "LE", "le"),   //是否是查询
		IsExtFuncField:    false,                                         //扩展功能
		IsFrontQueryField: md.HasConstraint(r.Constraints, "Q", "q"),     //是否是查询
		IsBackQueryField:  md.HasConstraint(r.Constraints, "BQ", "bq"),   //是否是后询
		IsDeleteField:     md.HasConstraint(r.Constraints, "D", "d"),
	}
}

type idx struct {
	Type  string
	Name  string
	Index int
}

type lstCmpt struct {
	Type      string //text,link,image,switch,pwd,mobile,progress,file,tag,daterange,date
	// ListWidth string //列表宽度
	// Rows      int    //显示的行数，如textArea行数
	// Position  string //位置 default,默认，换行
	Page    string //页面类型，字段显示格式，如:yyyy-MM-dd
	// CLR       color  //颜色信息
	// Min       int    //最小值
	// Max       int    //最大值
}

func createLstCMPT(r *md.Row) map[string]*lstCmpt {
	cmpts := make(map[string]lstCmpt)
	for _,k:=range r.Constraints{
		tp, pn := md.GetConsByTag("tp", k)
		if tp ==""{
			continue
		}
		cmpts[types.GetString(pn,"*")] = &lstCmpt{
			Type:tp,
			Page:pn,
		}
	}
	return cmpts
	
}

type color struct {
	Name      string
	FontColor string
	BgColor   string
}

type enumType struct {
	IsEnum   bool   //是否是枚举类型
	EnumName string //枚举名
	Id       string
	Name     string
	Type     string
	PID      string
	Status   string //枚举状态字段
	Expire   string //日期过期字段
	Multiple bool
}

type Column struct {
}

type Column struct {
	aseColumn           //字段基础息
	Field     fieldType //字段类型
	CMPT      lstCmpt   //列表组件
}

func newColum(r *d.Row) *Column {
	return &Column{
		BaseColumn: BaseColumn{
			Nme:   r.Name,
			Label: r.Desc.Name,
			UNQ:   defFids.Next(),
		},
		Field: createFieldType(r),
	}
}
