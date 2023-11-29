package data

import (
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

type EnumType struct {
	EnumType      string
	Id            string
	Name          string
	ExtName       string
	Type          string
	PID           string
	Group         string
	Status        string //枚举状态字段
	Expire        string //日期过期字段
	SortName      string //排序字段
	SortType      string //排序方式
	Multiple      bool
	DEColumns     Columns
	FilterColumns Columns
}

func newEnumType(enumType string, rs Columns) *EnumType {
	tp := &EnumType{EnumType: enumType, DEColumns: rs.GetEnumDelColumns()}
	tp.FilterColumns = rs.GetByCmpnt("DV")
	for _, r := range rs {
		if md.HasConstraint(r.Constraints, "DI") {
			tp.Id = types.GetString(md.GetConsNameByTagIgnorecase("DI", r.Constraints...), r.Name)
		}
		if md.HasConstraint(r.Constraints, "DN") {
			tp.Name = types.GetString(md.GetConsNameByTagIgnorecase("DN", r.Constraints...), r.Name)
		}
		if md.HasConstraint(r.Constraints, "DNE") {
			tp.ExtName = types.GetString(md.GetConsNameByTagIgnorecase("DNE", r.Constraints...), r.Name)
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
		if md.HasConstraint(r.Constraints, "DG", "dg") {
			tp.Group = r.Name
		}
		if md.HasConstraint(r.Constraints, "DSN", "dsn") {
			tp.SortName = r.Name
			tp.SortType = "asc"
		}
		if md.HasConstraint(r.Constraints, "DSND", "dsnd") {
			tp.SortName = r.Name
			tp.SortType = "desc"
		}
		if md.HasConstraint(r.Constraints, "expire") {
			tp.Expire = r.Name
		}
		if tp.SortName == "" {
			tp.SortName = tp.Id
			tp.SortType = "desc"
		}

	}
	if tp.Id == "" || tp.Name == "" {
		tp.EnumType = ""
	}
	return tp
}
func (t *EnumType) IsEnum() bool {
	return t.EnumType != ""
}
func (t *EnumType) IsEnumAndMuti() bool {
	return t.EnumType != "" && t.Multiple
}
func (t *EnumType) GroupIsStatic() bool {
	return strings.HasPrefix(t.Group, "#")
}

func (t *EnumType) GetGroupValue() string {
	return strings.Trim(t.Group, "#")
}