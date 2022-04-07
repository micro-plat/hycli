package md

import (
	"strings"

	"github.com/micro-plat/lib4go/types"
)

type RName struct {
	Raw    string
	Main   string
	Short  string
	Prefix string
}

func (r *RName) String() string {
	return r.Short
}

//Table 表名称
type Table struct {
	Name    *RName //表名
	Desc    string //表描述
	Rows    []*Row //原始行
	ExtInfo string //扩展信息
}

//NewTable 创建表
func NewTable(name, desc, extinfo string) *Table {
	fname := types.GetStringByIndex(getNames(name), 0)
	names := strings.Split(fname, "_")

	return &Table{
		Name: &RName{Raw: fname,
			Short:  strings.Join(names[1:], "_"),
			Prefix: names[0],
			Main:   strings.Join(names[1:2], "_")},
		Desc:    desc,
		Rows:    make([]*Row, 0, 1),
		ExtInfo: extinfo,
	}
}

//AddRow 添加行信息
func (t *Table) AddRow(r *Row) error {
	t.Rows = append(t.Rows, r)
	return nil
}
