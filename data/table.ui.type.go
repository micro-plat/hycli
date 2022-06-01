package data

import (
	"strings"

	"github.com/micro-plat/hycli/md"
	"github.com/micro-plat/lib4go/types"
)

var aMap = map[string]string{
	"varchar2":      "input",
	"varchar":       "input",
	"string":        "input",
	"link":          "input",
	"lnk":           "input",
	"tag":           "input",
	"number":        "number",
	"decimal":       "number",
	"int":           "number",
	"tinyint":       "number",
	"bigint":        "number",
	"progress":      "number",
	"switch":        "switch",
	"swth":          "switch",
	"date":          "date",
	"time":          "time",
	"daterange":     "daterange",
	"datetime":      "daterange",
	"datetimerange": "daterange",
	"dtr":           "daterange",
	"sl":            "select",
	"select":        "select",
	"radio":         "select",
	"bool":          "select",
	"bl":            "select",
	"multiselect":   "multiselect",
	"file":          "file",
	"upload":        "file",
	"password":      "password",
	"pwd":           "password",
}
var vmap = map[string]string{
	"image":    "img",
	"img":      "img",
	"link":     "link",
	"lnk":      "link",
	"tag":      "tag",
	"progress": "progress",
	"switch":   "switch",
	"swth":     "switch",
}
var tpmap = map[string]map[string]string{
	"Q":  aMap,
	"C":  aMap,
	"U":  aMap,
	"L":  vmap,
	"LE": vmap,
	"V":  vmap,
}

func getShowTypeName(t string, vs ...string) string {
	for _, v := range vs {
		if _, ok := tpmap[t]; !ok {
			continue
		}
		if m, ok := tpmap[t][v]; ok {
			return m
		}
	}
	return ""
}
func getTypeName(vs ...string) string {
	for _, k := range vs {
		if v, ok := aMap[k]; ok {
			return v
		}
		if v, ok := vmap[k]; ok {
			return v
		}

	}
	return ""
}

type UIType struct {
	Name       string
	Format     string
	Len        int
	IsDecimal  bool
	IsNumber   bool
	IsDate     bool
	IsMobile   bool
	IsSelect   bool
	SelectName string
	Type       string
	Min        string
	Max        string
	Ext        string
}

func NewUIType(t string, r *md.Row) *UIType {

	// tname := md.GetTPName(t, r.Constraints...)

	//1.未找到则查询tp对应的组件类型[如:tp(swith)]
	tpName := md.GetTPName("tp", r.Constraints...)

	//2.查找是否包含'sl'配置，[如：sl]
	selectName := md.GetSelectName(r.Name, r.Constraints...)

	defName := types.DecodeString(selectName, "", r.Type.Name, "select")

	//3. 没有则使用当前字段对应的类型
	tpName = types.GetString(tpName, defName)

	uit := &UIType{
		Name:      r.Type.Name,
		Len:       r.Type.Len,
		IsDecimal: strings.EqualFold(getShowTypeName(t, tpName), "NUMBER") && r.Type.DLen > 0,
		IsNumber:  strings.EqualFold(getTypeName(r.Type.Name), "NUMBER"),
		IsDate: strings.EqualFold(getTypeName(r.Type.Name), "date") ||
			strings.EqualFold(getTypeName(r.Type.Name), "daterange"),
		IsMobile:   strings.EqualFold(tpName, "MOBILE"),
		Format:     md.GetFormat(r.Constraints...),
		Type:       getShowTypeName(t, tpName+md.GetRangeName()),
		IsSelect:   md.HasConstraint(r.Constraints, "sl"),
		SelectName: selectName,
	}

	uit.Min, uit.Max = md.GetRanges(r.Constraints...)
	return uit
}
