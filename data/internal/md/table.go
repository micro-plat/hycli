package md

import (
	"strings"

	"github.com/micro-plat/hycli/output"
	"github.com/micro-plat/lib4go/types"
)

type RName struct {
	Raw      string `json:'raw'`
	OName    string `json:'oname'`  //原始名称，包含~
	CName    string `json:'cname'`  //单词间首字母大写
	LOName   string `json:'loName'` //全小写，"."进行分隔
	Main     string `json:'main'`
	Short    string `json:'short'`
	RealName string `json:'realName'` ////原始名称去掉~等特殊字符
	Prefix   string `json:'prefix'`
	MainPath string `json:'main_path'`
}

func (r *RName) String() string {
	return r.Short
}

// Table 表名称
type Table struct {
	Name     *RName `json:'name'`     //表名
	Desc     string `json:'desc'`     //表描述
	Rows     []*Row `json:'rows'`     //原始行
	ExtInfo  string `json:'ext_info'` //扩展信息
	Settings string `json:'json'`

	Exclude bool        `json:'exclude'`
	PkgName string      `json:'pkg_name'`
	Cache   interface{} `json:'cache'`
	Marker  string
	Tbs     Tables `json:'tbs'`
}

// NewTable 创建表
func NewTable(name, desc, extinfo string, settings string) *Table {
	fname := types.GetStringByIndex(getNames(name), 0)
	rname := strings.Trim(fname, "^")
	names := strings.Split(rname, "_")

	return &Table{
		Name: &RName{
			Raw:      rname, //表名称
			RealName: rname,
			Short:    strings.Join(names[1:], "_"),          //以下划线连接的第二段后面的名称
			Prefix:   names[0],                              //前缀名
			CName:    ToCName(strings.Join(names[1:], "_")), //class name,去掉第一段，后面几段首字母大写
			Main:     strings.Join(names[1:2], "_"),         //第二段字符
			OName:    fname,                                 //原始名称
			MainPath: strings.Join(names, "/"),              //以斜杠分隔的路径
			LOName:   strings.ToLower(strings.Join(names[1:], ".")),
		},

		Exclude:  strings.HasPrefix(name, "^"),
		PkgName:  GetPkgName(),
		Desc:     desc,
		Rows:     make([]*Row, 0, 1),
		ExtInfo:  extinfo,
		Settings: settings,
		Tbs:      Tables{},
		Marker:   output.MARKER,
	}
}

// AddRow 添加行信息
func (t *Table) AddRow(r *Row) error {
	cache := t.Cache
	t.Cache = nil
	t.Rows = append(t.Rows, r)
	t.Cache = cache
	return nil
}
