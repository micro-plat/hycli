package md

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/micro-plat/lib4go/types"
)

type RName struct {
	Raw    string
	CName  string
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
	PkgName string
	Cache   interface{}
}

//NewTable 创建表
func NewTable(name, desc, extinfo string) *Table {
	fname := types.GetStringByIndex(getNames(name), 0)
	names := strings.Split(fname, "_")

	return &Table{
		Name: &RName{Raw: fname,
			Short:  strings.Join(names[1:], "_"),
			Prefix: names[0],
			CName:  ToCName(strings.Join(names[1:], "_")),
			Main:   strings.Join(names[1:2], "_")},
		PkgName: getPkgName(),
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
func getPkgName() string {
	gopath, _ := os.LookupEnv("GOPATH")
	if gopath == "" {
		return ""
	}
	currentPath, _ := os.Getwd()
	root := filepath.Join(gopath, "src")
	if strings.HasPrefix(strings.ToLower(currentPath), strings.ToLower(root)) {
		currentPath = currentPath[len(root)+1:]
	}
	return strings.Trim(strings.Replace(currentPath, "\\", "/", -1), "/")

}
