package md

import (
	"fmt"
	"strings"

	"github.com/micro-plat/lib4go/jsons"
	"github.com/micro-plat/lib4go/types"
)

// Tables markdown中的所有表信息
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
func (t Tables) Filter(n string, exclude bool) Tables {
	ns := strings.Split(n, ",")
	ntables := make(Tables, 0, 1)
	for _, tn := range t {
		if exclude && tn.Exclude {
			continue
		}
		for _, cc := range ns {
			if strings.Contains(tn.Name.Raw, cc) {
				ntables = append(ntables, tn)
				break
			}
		}
	}
	return ntables
}

func (t Tables) resetConf(cnf TableConfs) {

	cnfs := cnf.Map()
	if len(cnfs) == 0 {
		return
	}
	for _, tb := range t {

		//获取表的配置信息
		conf, ok := cnfs[tb.Name.RealName]
		if !ok {
			continue
		}
		tb.ExtInfo = types.GetString(conf.ExtInfo, tb.ExtInfo)
		tb.Desc = types.GetString(conf.Desc, tb.Desc)
		tb.Exclude = conf.Exclude || tb.Exclude

		//循环列，查找配置信息
		for _, row := range tb.Rows {
			rconf, ok := conf.Rows[row.Name]
			if !ok {
				continue
			}
			row.Constraints = append(row.Constraints, rconf.Constraints...)
			row.Desc = rconf.Desc
		}
	}
}
