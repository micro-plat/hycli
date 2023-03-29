package md

import (
	"fmt"
	"strings"

	"github.com/micro-plat/lib4go/jsons"
	"github.com/micro-plat/lib4go/types"
)

type TableConfs []*TableConf

type TableConfsMap map[string]*TableConf

func (t TableConfs) String() string {
	buff, _ := jsons.Marshal(t)
	return string(buff)
}

func (t TableConfs) Map() TableConfsMap {
	m := TableConfsMap{}
	for _, tb := range t {
		m[tb.Name] = tb
	}
	return m
}

// Table 表名称
type TableConf struct {
	Name    string              `json:'name'` //表名
	Desc    string              `json:'desc'` //表描述
	Rows    map[string]*RowConf `json:'rows'` //原始行
	Exclude bool                `json:'exclude'`
	ExtInfo string              `json:'ext_info'` //扩展信息
}

// newTableConf 创建表
func newTableConf(name, desc, extinfo string) *TableConf {
	fname := types.GetStringByIndex(getNames(name), 0)
	rname := strings.Trim(fname, "^")

	return &TableConf{
		Name:    rname,
		Exclude: strings.HasPrefix(fname, "^"),
		Desc:    desc,
		Rows:    make(map[string]*RowConf),
		ExtInfo: extinfo,
	}
}
func (t TableConfs) Duplicate() error {
	names := make(map[string]bool)
	for _, v := range t {
		if _, ok := names[v.Name]; ok {
			return fmt.Errorf("存在重复的表名:%s", v.Name)
		}
		names[v.Name] = true
	}
	return nil
}

// Row 行信息
type RowConf struct {
	Index       int      `json:"index"`       //索引
	Name        string   `json:"name"`        //字段名
	Constraints []string `json:"constraints"` //约束
	Desc        *RDesc   `json:"desc"`        //描述
	hasRow      bool     //是否包含在原表中
}

// Lines2Conf 表数据行变为表
func Lines2Conf(lines Lines) (confs TableConfs, err error) {
	confs = TableConfs{}
	for _, tline := range lines.Lines {
		//markdown表格的表名，标题，标题数据区分行，共三行
		if len(tline) <= 3 {
			continue
		}
		var tb *TableConf
		for i, line := range tline {
			if i == 0 {
				//获取表名，描述名称
				name, err := getTableName(line)
				if err != nil {
					return nil, err
				}
				tb = newTableConf(name, getTableDesc(line), getTableExtInfo(line))
				continue
			}
			if i < 3 {
				continue
			}
			c, err := line2TableRowConf(line)
			if err != nil {
				return nil, err
			}
			//添加列
			tb.Rows[c.Name] = c
		}
		if tb != nil {
			confs = append(confs, tb)
		}
	}
	return confs, nil
}

func line2TableRowConf(line *Line) (*RowConf, error) {
	if strings.Count(line.Text, "|") != 4 {
		return nil, fmt.Errorf("配置表结构有误(行:%d)", line.LineID)
	}
	Columns := strings.Split(strings.Trim(strings.TrimSpace(line.Text), "|"), "|")
	if Columns[0] == "" {
		return nil, fmt.Errorf("字段名称不能为空 %s(行:%d)", line.Text, line.LineID)
	}
	desc := strings.TrimSpace(strings.Replace(Columns[2], "&#124;", "|", -1))
	c := &RowConf{
		Index:       line.LineID,
		Name:        strings.TrimSpace(strings.Replace(Columns[0], "&#124;", "|", -1)),
		Constraints: mergeConstraint(getNames(strings.TrimSpace(Columns[1]))),
		Desc:        &RDesc{Raw: desc, Name: getShortDesc(desc)},
	}
	// fmt.Println("---toconf:", Columns[0], ":", getNames(Columns[1]))

	return c, nil
}
