package data

import (
	"strings"

	"github.com/micro-plat/hycli/md"
	"github.com/micro-plat/lib4go/types"
)

var tpmap = map[string]string{
	"varchar2":      "input",
	"varchar":       "input",
	"string":        "input",
	"password":      "password",
	"pwd":           "password",
	"number":        "number",
	"decimal":       "number",
	"int":           "number",
	"tinyint":       "number",
	"bigint":        "number",
	"date":          "date",
	"time":          "time",
	"daterange":     "daterange",
	"datetime":      "daterange",
	"datetimerange": "daterange",
	"sl":            "select",
	"multiselect":   "multiselect",
	"select":        "select",
	"radio":         "select",
	"bool":          "select",
	"bl":            "select",
	"upload":        "file",
	"file":          "file",
	"link":          "link",
	"lnk":           "link",
	"tag":           "tag",
	"progress":      "progress",
	"switch":        "switch",
}

func getTP(vs ...string) string {
	for _, v := range vs {
		if v, ok := tpmap[v]; ok {
			return v
		}
	}

	return "input"
}

type UIType struct {
	Name      string
	Format    string
	Len       int
	IsDecimal bool
	IsNumber  bool
	Type      string
	Min       string
	Max       string
}

func NewUIType(t string, r *md.Row) *UIType {

	//1.未找到则查询tp对应的组件类型[如:tp(swith)]
	tpName := md.GetTPName("tp", r.Constraints...)

	//2.查找是否包含'sl'配置，[如：sl]
	selectName := md.GetSelectName(r.Name, r.Constraints...)

	//3. 没有则使用当前字段对应的类型
	tpName = types.GetString(tpName, types.DecodeString(selectName, "", r.Type.Name, "select"))

	uit := &UIType{
		Name:      r.Type.Name,
		Len:       r.Type.Len,
		IsDecimal: strings.EqualFold(getTP(tpName), "NUMBER") && r.Type.DLen > 0,
		IsNumber:  strings.EqualFold(getTP(tpName), "NUMBER"),
		Format:    md.GetFormat(r.Constraints...),
		Type:      getTP(tpName + md.GetRangeName()),
	}

	uit.Min, uit.Max = md.GetRanges(r.Constraints...)
	return uit
}
