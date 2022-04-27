package data

import (
	"github.com/micro-plat/hycli/md"
	"github.com/micro-plat/lib4go/jsons"
	"github.com/micro-plat/lib4go/types"
)

type EnumType struct {
	Id       string
	Name     string
	Type     string
	Multiple bool
	DERows   []*CodeRow
}

func (e *EnumType) String() string {
	buff, _ := jsons.Marshal(e)
	return string(buff)
}

func NewEnumType(r *md.Row) *EnumType {
	return &EnumType{
		Id:   types.GetString(md.GetTPName("DI", r.Constraints...), r.Name),
		Name: types.GetString(md.GetTPName("DN", r.Constraints...), r.Name),
		Type: types.GetString(md.GetTPName("DT", r.Constraints...), r.Name),
	}
}
func (d *EnumType) SetTypeIfAbsence(tp string) {
	if d.Type == "" {
		d.Type = tp
	}
}
func (d *EnumType) IsValid() bool {
	return d.Id != "" && d.Name != "" && d.Type != ""
}
