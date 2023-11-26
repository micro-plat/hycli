package data

import (
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

type enumType struct {
	IsEnum      bool   //是否是枚举类型
	EnumType    string //枚举名
	AssctColumn string //关联列
	PID         string //父级编号
	Group       string //分组
	IsDEColumn  bool
}

func (t *enumType) GetGroupValue() string {
	return strings.Trim(t.Group, "#")
}
func (t *enumType) GroupIsStatic() bool {
	return strings.HasPrefix(t.Group, "#")
}
func (t *enumType) HasAssctColumn() bool {
	return t.AssctColumn != ""
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
