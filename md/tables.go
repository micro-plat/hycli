package md

import (
	"fmt"
	"strings"

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
func (t Tables) Filter(n string) Tables {
	if n == "" {
		return t
	}
	ntables := make(Tables, 0, 1)
	for _, tn := range t {
		if strings.Contains(tn.Name.Raw, n) {
			ntables = append(ntables, tn)
		}
	}
	return ntables
}
