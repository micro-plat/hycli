package data

import (
	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

type EnumType struct {
	EnumType  string
	Id        string
	Name      string
	Type      string
	PID       string
	Status    string //枚举状态字段
	Expire    string //日期过期字段
	SortName  string //排序字段
	Multiple  bool
	DEColumns []*Column
}

func newEnumType(enumType string, rs []*md.Row, delColumn Columns) *EnumType {
	tp := &EnumType{EnumType: enumType, DEColumns: delColumn}
	for _, r := range rs {
		if md.HasConstraint(r.Constraints, "DI") {
			tp.Id = types.GetString(md.GetConsNameByTagIgnorecase("DI", r.Constraints...), r.Name)
		}
		if md.HasConstraint(r.Constraints, "DN") {
			tp.Name = types.GetString(md.GetConsNameByTagIgnorecase("DN", r.Constraints...), r.Name)
		}
		if md.HasConstraint(r.Constraints, "DT") {
			tp.Type = types.GetString(md.GetConsNameByTagIgnorecase("DT", r.Constraints...), r.Name)
			tp.Multiple = true
		}
		if md.HasConstraint(r.Constraints, "DS", "ds") {
			tp.Status = r.Name
		}
		if md.HasConstraint(r.Constraints, "DP", "dp") {
			tp.PID = r.Name
		}
		if md.HasConstraint(r.Constraints, "DSN", "dsn") {
			tp.SortName = r.Name
		}
		if md.HasConstraint(r.Constraints, "expire") {
			tp.Expire = r.Name
		}

	}
	if tp.Id == "" || tp.Name == "" {
		tp.EnumType = ""
	}
	return tp
}
