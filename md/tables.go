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
func (t Tables) JoinPkgName(tx string) {
	vtx := strings.Trim(strings.Trim(strings.Replace(tx, "\\", "/", -1), "."), "/")
	for _, tb := range t {
		tb.PkgName = fmt.Sprintf("%s/%s", tb.PkgName, vtx)
	}
}
func (t Tables) Filter(n string) Tables {
	if n == "" {
		return t
	}
	ns := strings.Split(n, ",")
	ntables := make(Tables, 0, 1)
	for _, tn := range t {
		for _, cc := range ns {
			if strings.Contains(tn.Name.Raw, cc) {
				ntables = append(ntables, tn)
				break
			}
		}

	}
	return ntables
}
