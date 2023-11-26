package data

import (
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/jsons"
	"github.com/micro-plat/lib4go/types"
)

// 显示组件
type displayCmpnt struct {
	Type   string //text,link,image,switch,pwd,mobile,progress,file,tag,daterange,date,multiselect,select
	Format string //字段显示格式，如:yyyy-MM-dd

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

func (s displayCmpnt) String() string {
	buff, _ := jsons.Marshal(s)
	return string(buff)
}
func (s displayCmpnt) IsStatic() bool {
	return strings.HasPrefix(s.Format, "#")
}
func (s displayCmpnt) StaticValue() string {
	return strings.Trim(s.Format, "#")
}
func (s displayCmpnt) StartWith(sx string) bool {
	lst := strings.Split(sx, "|")
	for _, v := range lst {
		if strings.HasPrefix(s.Type, v) {
			return true
		}
	}
	return false
}
func (s displayCmpnt) Equal(vs string) bool {
	lst := strings.Split(vs, "|")
	for _, v := range lst {
		if strings.EqualFold(s.Type, v) {
			return true
		}
	}
	return false

}
func (s displayCmpnt) HasFormat() bool {
	return s.Format != ""
}

type ext struct {
	FormName string
	TP       []string
	Name     string
}
