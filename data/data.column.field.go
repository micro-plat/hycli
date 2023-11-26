package data

import (
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

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

func (c *fieldType) LikeQuery() bool {
	return md.HasConstraint(c.row.Constraints, "#like")
}
