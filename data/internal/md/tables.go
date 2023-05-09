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

func (t Tables) resetConf(cnf TableConfs) Tables {

	cnfs := cnf.Map()
	if len(cnfs) == 0 {
		return t
	}
	tmpTable := make(map[string]bool)
	for _, tb := range t {
		tmpTable[tb.Name.Raw] = true
		//获取表的配置信息
		conf, ok := cnfs[tb.Name.RealName]
		if !ok {
			continue
		}
		tb.ExtInfo = types.GetString(conf.ExtInfo, tb.ExtInfo)
		tb.Settings = types.GetString(conf.Settings, tb.Settings)
		tb.Desc = types.GetString(conf.Desc, tb.Desc)
		tb.Exclude = conf.Exclude || tb.Exclude

		//循环列，查找配置信息
		for _, row := range tb.Rows {
			rconf, ok := conf.Rows[row.Name]
			if !ok {
				rconf, ok = conf.Rows[row.Raw]
			}
			if ok {
				row.Constraints = append(row.Constraints, rconf.Constraints...)
				row.Desc = rconf.Desc
				rconf.hasRow = true
			}

		}

		//查找原表未包含的列
		for _, r := range conf.Rows {
			if !r.hasRow {
				name := types.DecodeString(strings.HasPrefix(r.Name, "^"), true, r.Name, "^"+r.Name)
				row, _ := Line2TableRow(NewLine(r.Index, name, "number(20)", r.Desc.Raw))
				row.Constraints = append(row.Constraints, r.Constraints...)
				tb.AddRow(row)
			}

		}
	}

	//添加扩展表
	for _, c := range cnf {
		if _, ok := tmpTable[c.Name]; ok {
			continue
		}
		tb := NewTable(types.DecodeString(strings.HasPrefix(c.Name, "^"), true, c.Name, "^"+c.Name), c.Desc, c.ExtInfo, c.Settings)
		//查找原表未包含的列
		for _, r := range c.Rows {
			name := types.DecodeString(strings.HasPrefix(r.Name, "^"), true, r.Name, "^"+r.Name)
			row, _ := Line2TableRow(NewLine(r.Index, name, "number(20)", r.Desc.Raw))
			row.Constraints = append(row.Constraints, r.Constraints...)
			tb.AddRow(row)

		}
		t = append(t, tb)
	}
	return t

}
