package data

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/jsons"
	"github.com/micro-plat/lib4go/types"
)

// 操作信息
type optrs struct {
	Tag      string
	Name     string //link,dialog,cmpnt,tab,cnfrm
	Label    string //修改,预览，删除
	ICON     string //图标
	URL      string //组件URL
	RwName   string //当前表字段标记
	FwName   string //外部表字段标记
	ReqURL   string //服务请求URL
	IsMux    bool   //是否是复用组件
	index    int    //顺序编号
	ExtTable string
	Params   map[string]string
	UNQ      string
}
type optrslst []*optrs
type formatOptrs optrslst
type lstOptrs optrslst
type viewOptrs optrslst
type lstatOptrs optrslst
type chartOptrs optrslst
type barOptrs optrslst
type viewExtCmptOpts optrslst
type extOptrs optrslst

var viewOptCmd = []string{"view"}
var lstatOptCmd = []string{"lstat"}
var lstOptCmd = []string{"lst"}
var batchCheck = []string{"bcheck"}
var barOptrCmd = []string{"export", "import", "bcheck"}
var charOptrCmd = []string{"chart"}
var extCmptParam = []string{"add"}
var extOptrsCmd = []string{"tskbar"}

var addOpts = &optrs{Tag: "ADD", URL: "@/views/{@prefix}/{@main}/{@name}.add", Name: "CMPNT", ICON: "Plus", Label: "添加", RwName: "C", UNQ: defFids.Next(), index: 1}
var detailOpts = &optrs{Tag: "VIEW", URL: "./{@name}.view", Name: "CMPNT", Label: "详情", RwName: "V", UNQ: defFids.Next(), index: 1}
var updateOpts = &optrs{Tag: "UPDATE", URL: "./{@name}.edit", Name: "CMPNT", Label: "修改", RwName: "U", UNQ: defFids.Next(), index: 2}
var delOpts = &optrs{Tag: "CNFRM", URL: "./{@name}.cnfrm", ReqURL: "/{@mainPath}/del", Name: "CMPNT", Label: "删除", RwName: "D", UNQ: defFids.Next(), IsMux: true, index: 99}
var dialogOpts = &optrs{Tag: "DIALOG", URL: "./{@name}.dialog", Name: "CMPNT", IsMux: true, index: 99}
var cnfrmOpts = &optrs{Tag: "CNFRM", URL: "./{@name}.cnfrm", Name: "CMPNT", IsMux: true, index: 99}
var chartLinePieBarOpts = &optrs{Tag: "CHART", URL: "@/views/cmpnts/chart.base.vue", Name: "CMPNT"}
var taskBarOpts = &optrs{Tag: "TSKBAR", URL: "@/views/cmpnts/task.bar.vue", Name: "CMPNT"}

func (s optrslst) Len() int           { return len(s) }
func (s optrslst) Less(i, j int) bool { return s[i].index < s[j].index }
func (s optrslst) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

var extCmptOpts = map[string]*optrs{
	"add": addOpts,
}

func (f formatOptrs) String() string {
	buff, err := jsons.Marshal(f)
	if err != nil {
		return fmt.Sprintf("err:%s", err.Error())
	}
	return string(buff)
}

func (b barOptrs) NeedCheck(t string) bool {
	for _, v := range b {
		if types.StringContains(batchCheck, v.Tag) {
			return true
		}
	}
	return false
}

func createLstOptrs(table *Table, t string) lstOptrs {
	optrs := createCmdsOptrs(t, lstOptCmd)
	//构建操作
	if len(fltrColumns(table, VIEW_COLUMN)) > 0 {
		optrs = append(optrs, detailOpts)
	}
	if len(fltrColumns(table, UPDATE_COLUMN)) > 0 {
		optrs = append(optrs, updateOpts)
	}
	if len(fltrColumns(table, DELETE_COLUMN)) > 0 {
		optrs = append(optrs, delOpts)
	}
	return optrs
}
func createViewOptrs(t string) (viewOptrs, viewExtCmptOpts) {
	viewOpts := createCmdsOptrs(t, viewOptCmd)
	extOpts := creatExtCmptOpts(viewOpts)
	return viewOpts, extOpts
}
func createLStatChartOptrs(t string) (lstatOptrs, chartOptrs) {
	optrs := createCmdsOptrs(t, lstatOptCmd)
	copts := createChartOptrs(t)
	return optrs, copts
}
func createChartOptrs(t string) chartOptrs {
	return createCmdsOptrs(t, charOptrCmd)
}
func createExtOptrs(t string) extOptrs {
	return createCmdsOptrs(t, extOptrsCmd)
}
func createBarOptrs(table *Table, t string) (barOptrs, bool) {
	opts := createCmdsOptrs(t, barOptrCmd)
	if len(fltrColumns(table, ADD_COLUMN)) > 0 {
		opts = append(opts, addOpts)
	}
	return opts, barOptrs(opts).NeedCheck(table.Name.Raw)
}
func createCmdsOptrs(t string, cmds []string) []*optrs {
	optrs := make([]*optrs, 0, 1)
	for _, v := range cmds {
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
				if v, ok := view.Params[cmd]; ok {
					if opts, ok := extCmptOpts[cmd]; ok {
						xview := *opts
						xview.UNQ = view.UNQ
						xview.Label = types.DecodeString(types.GetBool(v), true, view.Label, v)
						xview.Params = view.Params
						xview.Params["table"] = view.URL
						nopts = append(nopts, &xview)
					}
				}
			}
		}
	}
	return nopts
}

func createOptrs(t string, tag string) []*optrs {
	list := md.GetExtOpt(t, tag)
	opts := make([]*optrs, 0, len(list))
	if len(list) == 0 {
		return opts
	}
	for _, lst := range list {
		opt := optrs{}
		name := strings.ToUpper(types.GetStringByIndex(lst, 1))
		switch name {
		case "DIALOG":
			opt = *dialogOpts
			opt.ReqURL = types.GetStringByIndex(lst, 2)
		case "CNFRM":
			opt = *cnfrmOpts
			opt.ReqURL = types.GetStringByIndex(lst, 2)
		case "LINE", "PIE", "BAR":
			opt = *chartLinePieBarOpts
			opt.Tag = name
			opt.ReqURL = types.GetStringByIndex(lst, 2)
		case "TSKBAR":
			opt = *taskBarOpts
			opt.Tag = name
			opt.ReqURL = types.GetStringByIndex(lst, 2)
		default:
			opt = optrs{Tag: tag, Name: name, URL: types.GetStringByIndex(lst, 2), index: 99}
		}

		opt.Label = types.GetStringByIndex(lst, 0)
		opt.RwName = types.GetStringByIndex(lst, 3)
		opt.FwName = types.GetStringByIndex(lst, 4)
		opt.UNQ = defFids.Next()
		opt.Params = make(map[string]string)

		//解析组件参数
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
			opt.ExtTable = opt.Params["table"]
			opt.ICON = types.GetString(opt.Params["icon"], opt.ICON)
		}
		opts = append(opts, &opt)
	}
	return opts
}
