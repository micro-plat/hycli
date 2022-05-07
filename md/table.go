package md

import (
	"strings"

	"github.com/micro-plat/lib4go/types"
)

type RName struct {
	Raw      string `json:'raw'`
	OName    string `json:'oname'` //原始名称
	CName    string `json:'cname'`
	Main     string `json:'main'`
	Short    string `json:'short'`
	Prefix   string `json:'prefix'`
	MainPath string `json:'main_path'`
}

func (r *RName) String() string {
	return r.Short
}

//Table 表名称
type Table struct {
	Name    *RName `json:'name'`     //表名
	Desc    string `json:'desc'`     //表描述
	Rows    []*Row `json:'rows'`     //原始行
	ExtInfo string `json:'ext_info'` //扩展信息

	Exclude bool        `json:'exclude'`
	PkgName string      `json:'pkg_name'`
	Cache   interface{} `json:'cache'`

	Tbs Tables `json:'tbs'`
}

//NewTable 创建表
func NewTable(name, desc, extinfo string) *Table {
	fname := types.GetStringByIndex(getNames(name), 0)
	rname := strings.Trim(fname, "^")
	names := strings.Split(rname, "_")

	return &Table{
		Name: &RName{
			Raw:      rname,
			Short:    strings.Join(names[1:], "_"),
			Prefix:   names[0],
			CName:    ToCName(strings.Join(names[1:], "_")),
			Main:     strings.Join(names[1:2], "_"),
			OName:    fname,
			MainPath: strings.Join(names[1:], "/"),
		},

		Exclude: strings.HasPrefix(fname, "^"),
		PkgName: GetPkgName(),
		Desc:    desc,
		Rows:    make([]*Row, 0, 1),
		ExtInfo: extinfo,
		Tbs:     Tables{},
	}
}

//AddRow 添加行信息
func (t *Table) AddRow(r *Row) error {
	cache := t.Cache
	t.Cache = nil
	t.Rows = append(t.Rows, r)
	t.Cache = cache
	return nil
}
