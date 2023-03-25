package data

import (
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

// 操作信息
type optrs struct {
	Tag    string
	Name   string //link,dialog,cmpnt,tab,cnfrm
	Label  string //修改,预览，删除
	URL    string
	RwName string
	FwName string
	Params map[string]string
	UNQ    string
}
type lstOptrs []*optrs
type viewOptrs []*optrs
type lstatOptrs []*optrs
type chartOptrs []*optrs
type barOptrs []*optrs
type ViewExtCmptOpts []*optrs

var batchCheck = []string{"bcheck"}
var barOptrCmd = []string{"export", "import", "bcheck"}
var charOptrCmd = []string{"chart"}
var extCmptParam = []string{"add"}

var addOpts = &optrs{Tag: "ADD", Name: "ADD", Label: "添加", RwName: "C", UNQ: defFids.Next()}
var detailOpts = &optrs{Tag: "VIEW", Name: "VIEW", Label: "详情", RwName: "V", UNQ: defFids.Next()}
var updateOpts = &optrs{Tag: "UPDATE", Name: "UPDATE", Label: "修改", RwName: "U", UNQ: defFids.Next()}
var delOpts = &optrs{Tag: "DEL", Name: "CNFRM", Label: "删除", RwName: "D", UNQ: defFids.Next()}

func (b barOptrs) NeedCheck(t string) bool {
	for _, v := range b {
		if types.StringContains(batchCheck, v.Tag) {
			return true
		}
	}
	return false
}

func createLstOptrs(t string) lstOptrs {
	optrs := make([]*optrs, 0, 1)
	optrs = append(optrs, createOptrs(t, "lst")...)
	optrs = append(optrs, createOptrs(t, "LST")...)
	return optrs
}
func createViewOptrs(t string) viewOptrs {
	optrs := make([]*optrs, 0, 1)
	optrs = append(optrs, createOptrs(t, "view")...)
	optrs = append(optrs, createOptrs(t, "VIEW")...)
	return optrs
}
func createLStatOptrs(t string) lstatOptrs {
	optrs := make([]*optrs, 0, 1)
	optrs = append(optrs, createOptrs(t, "lstat")...)
	optrs = append(optrs, createOptrs(t, "LSTAT")...)
	return optrs
}
func createChartOptrs(t string) chartOptrs {
	optrs := make([]*optrs, 0, 1)
	for _, v := range charOptrCmd {
		optrs = append(optrs, createOptrs(t, strings.ToLower(v))...)
		optrs = append(optrs, createOptrs(t, strings.ToUpper(v))...)
	}
	return optrs
}
func creatExtCmptOpts(opts ...[]*optrs) []*optrs {
	nopts := make([]*optrs, 0, 1)
	for _, opt := range opts {
		for _, view := range opt {
			for _, cmd := range extCmptParam {
				if v, ok := view.Params[cmd]; ok && types.GetBool(v) {
					xview := *view
					xview.Name = "CMPNT"
					xview.Tag = cmd
					xview.Label = view.Params["label"]
					nopts = append(nopts, &xview)
				}
			}
		}
	}
	return nopts

}
func createBarOptrs(t string) barOptrs {
	optrs := make([]*optrs, 0, 1)
	for _, v := range barOptrCmd {
		optrs = append(optrs, createOptrs(t, strings.ToLower(v))...)
		optrs = append(optrs, createOptrs(t, strings.ToUpper(v))...)
	}
	return optrs
}
func createOptrs(t string, tag string) []*optrs {
	list := md.GetExtOpt(t, tag)
	opts := make([]*optrs, 0, len(list))
	for _, lst := range list {
		opt := &optrs{
			Tag:    tag,
			Name:   strings.ToUpper(types.GetStringByIndex(lst, 1)),
			Label:  lst[0],
			URL:    types.GetStringByIndex(lst, 2),
			RwName: types.GetStringByIndex(lst, 3),
			FwName: types.GetStringByIndex(lst, 4),
			UNQ:    defFids.Next(),
			Params: make(map[string]string),
		}
		if strings.HasPrefix(opt.FwName, "{") && strings.HasSuffix(opt.FwName, "}") {
			content := strings.TrimRight(strings.TrimLeft(opt.FwName, "{"), "}")
			ps := strings.Split(content, ";")
			for _, p := range ps {
				vs := strings.Split(p, ":")
				if types.GetStringByIndex(vs, 1) != "" {
					opt.Params[types.GetStringByIndex(vs, 0)] = types.GetStringByIndex(vs, 1)
				}
			}
			if strings.Contains(opt.FwName, ":") {
				opt.RwName = opt.Params["rwName"]
				opt.FwName = opt.Params["fwName"]
			}
		}

		opts = append(opts, opt)
	}
	return opts
}
