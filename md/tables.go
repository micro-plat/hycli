package md

import (
	"fmt"

	"github.com/micro-plat/lib4go/jsons"
)

//Tables markdown中的所有表信息
type Tables []*Table

func (r Tables) String() string {
	buff, _ := jsons.Marshal(r)
	return string(buff)
}
func (t Tables) Duplicate() error {
	names := make(map[string]bool)
	for _, v := range t {
		if _, ok := names[v.Name.Raw]; ok {
			return fmt.Errorf("存在重复的表名:%s", v.Name)
		}
		names[v.Name.Raw] = true
	}
	return nil
}
