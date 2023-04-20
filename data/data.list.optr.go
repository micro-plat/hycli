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
	Tag       string
	Name      string //link,dialog,cmpnt,tab,cnfrm
	Cmd       string // lstupdator lstbar chart
	Label     string //修改,预览，删除
	ICON      string //图标
	URL       string //组件URL
	RwName    string //当前表字段标记
	FwName    string //外部表字段标记
	ReqURL    string //服务请求URL
	IsMux     bool   //是否是复用组件
	index     int    //顺序编号
	Table     string
	Params    map[string]string
	ParentUNQ string //父组件编号
	UNQ       string
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
var batchCheck = []string{"bcheck", "@bcheck", "&bcheck"}
var barOptrCmd = []string{"export", "import", "lstbar", "lstupdator"}
var charOptrCmd = []string{"chart"}
var extCmptParam = []string{"add"}
var extOptrsCmd = []string{"tskbar"}

var addOpts = &optrs{Tag: ADD_TAG, URL: "@/views/{@prefix}/{@main}/{@name}.add", Name: "CMPNT", ICON: "Plus", Label: "添加", RwName: "C", UNQ: defFids.Next(), index: 1}
var detailOpts = &optrs{Tag: VIEW_TAG, URL: "@/views/{@prefix}/{@main}/{@name}.view", Name: "CMPNT", Label: "详情", RwName: "V", UNQ: defFids.Next(), index: 1}
var updateOpts = &optrs{Tag: UPDATE_TAG, URL: "@/views/{@prefix}/{@main}/{@name}.edit", Name: "CMPNT", Label: "修改", RwName: "U", UNQ: defFids.Next(), index: 2}
var delOpts = &optrs{Tag: DELETE_TAG, URL: "@/views/{@prefix}/{@main}/{@name}.cnfrm", ReqURL: "/{@mainPath}/del", Name: "CMPNT", Label: "删除", RwName: "D", UNQ: defFids.Next(), IsMux: true, index: 99}
var dialogOpts = &optrs{Tag: "DIALOG", URL: "@/views/{@prefix}/{@main}/{@name}.dialog", Name: "CMPNT", IsMux: true, index: 99}
var cnfrmOpts = &optrs{Tag: "CNFRM", URL: "@/views/{@prefix}/{@main}/{@name}.cnfrm", Name: "CMPNT", IsMux: true, index: 99}
var chartLinePieBarOpts = &optrs{Tag: "CHART", URL: "@/views/cmpnts/chart.base.vue", Name: "CMPNT"}
var taskBarOpts = &optrs{Tag: "TSKBAR", URL: "@/views/cmpnts/task.bar.vue", Name: "CMPNT"}

func (s *optrs) Get(tableName string) *optrs {
	nopts := *s
	nopts.Table = tableName
	return &nopts
}
func (s *optrs) String() string {
	buff, _ := jsons.Marshal(s)
	return string(buff)
}
func (s optrslst) Len() int           { return len(s) }
func (s optrslst) Less(i, j int) bool { return s[i].index < s[j].index }
func (s optrslst) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s optrslst) Find(tag string) bool {
	for _, v := range s {
		if v.Tag == tag {
			return true
		}
	}
	return false
}

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

func (b barOptrs) NeedCheck(tb string) bool {
	for _, v := range b {
		for p, m := range v.Params {
			if types.StringContains(batchCheck, p) ||
				types.StringContains(batchCheck, m) {
				return true
			}
		}

	}
	return false
}

func createLstOptrs(table *Table, t string) lstOptrs {
	optrs := createCmdsOptrs(table.Name.Raw, t, lstOptCmd)
	//构建操作
	if !optrslst(optrs).Find(VIEW_TAG) && len(fltrColumns(table, VIEW_COLUMN)) > 0 {
		optrs = append(optrs, detailOpts.Get(table.Name.Raw))
	}
	if !optrslst(optrs).Find(UPDATE_TAG) && len(fltrColumns(table, UPDATE_COLUMN)) > 0 {
		optrs = append(optrs, updateOpts.Get(table.Name.Raw))
	}
	if !optrslst(optrs).Find(DELETE_TAG) && len(fltrColumns(table, DELETE_COLUMN)) > 0 {
		optrs = append(optrs, delOpts.Get(table.Name.Raw))
	}
	return optrs
}
func createViewOptrs(tableName, t string) (viewOptrs, viewExtCmptOpts) {
	viewOpts := createCmdsOptrs(tableName, t, viewOptCmd)
	extOpts := creatExtCmptOpts(viewOpts)
	return viewOpts, extOpts
}
func createLStatChartOptrs(tableName, t string) (lstatOptrs, chartOptrs) {
	optrs := createCmdsOptrs(tableName, t, lstatOptCmd)
	copts := createChartOptrs(tableName, t)
	return optrs, copts
}
func createChartOptrs(tableName, t string) chartOptrs {
	return createCmdsOptrs(tableName, t, charOptrCmd)
}
func createExtOptrs(tableName string, t string) extOptrs {
	return createCmdsOptrs(tableName, t, extOptrsCmd)
}
func createBarOptrs(table *Table, t string) (barOptrs, bool) {
	opts := createCmdsOptrs(table.Name.Raw, t, barOptrCmd)
	if !optrslst(opts).Find(ADD_TAG) && len(fltrColumns(table, ADD_COLUMN)) > 0 {
		opts = append(opts, addOpts.Get(table.Name.Raw))
	}
	return opts, barOptrs(opts).NeedCheck(table.Name.Raw)
}
func createCmdsOptrs(tableName string, t string, cmds []string) []*optrs {
	optrs := make([]*optrs, 0, 1)
	for _, v := range cmds {
		optrs = append(optrs, createOptrs(tableName, t, strings.ToLower(v))...)
		optrs = append(optrs, createOptrs(tableName, t, strings.ToUpper(v))...)
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
						// xview.UNQ = view.UNQ
						xview.UNQ = defFids.Next()
						xview.Label = types.DecodeString(types.GetBool(v), true, view.Label, v)
						xview.Params = view.Params
						xview.Params["table"] = view.URL
						xview.Table = view.URL
						xview.ParentUNQ = view.UNQ
						nopts = append(nopts, &xview)
					}
				}
			}
		}
	}
	return nopts
}

func createOptrs(tableName string, t string, tag string) []*optrs {
	list := md.GetExtOpt(t, tag)
	opts := make([]*optrs, 0, len(list))
	if len(list) == 0 {
		return opts
	}
	for _, lst := range list {
		opt := optrs{}
		name := strings.ToUpper(types.GetStringByIndex(lst, 1))
		switch name {
		case VIEW_TAG:
			opt = *detailOpts
			opt.Table = types.GetStringByIndex(lst, 2, tableName)
		case UPDATE_TAG:
			opt = *updateOpts
			opt.Table = types.GetStringByIndex(lst, 2, tableName)
		case ADD_TAG:
			opt = *addOpts
			opt.Table = types.GetStringByIndex(lst, 2, tableName)
		case DELETE_TAG:
			opt = *delOpts
			opt.Table = types.GetStringByIndex(lst, 2, tableName)
		case "DIALOG":
			opt = *dialogOpts
			opt.ReqURL = types.GetStringByIndex(lst, 2)
			opt.Table = tableName
		case "CNFRM":
			opt = *cnfrmOpts
			opt.ReqURL = types.GetStringByIndex(lst, 2)
			opt.Table = tableName
		case "LINE", "PIE", "BAR":
			opt = *chartLinePieBarOpts
			opt.Tag = name
			opt.ReqURL = types.GetStringByIndex(lst, 2)
			opt.Table = tableName
		case "TSKBAR":
			opt = *taskBarOpts
			opt.Tag = name
			opt.ReqURL = types.GetStringByIndex(lst, 2)
			opt.Table = tableName
		case TAB_TAG:
			opt = optrs{
				Tag:   tag,
				Name:  name,
				URL:   types.GetStringByIndex(lst, 2),
				index: 99,
				Table: types.GetStringByIndex(lst, 2),
			}
		default:
			opt = optrs{
				Tag:   tag,
				Name:  name,
				URL:   types.GetStringByIndex(lst, 2),
				index: 99,
				Table: tableName,
			}
		}

		opt.Cmd = tag
		opt.Label = types.GetStringByIndex(lst, 0)
		opt.UNQ = defFids.Next()
		opt.Params = make(map[string]string)
		rwName := types.GetStringByIndex(lst, 3)
		fwName := types.GetStringByIndex(lst, 4)
		//解析组件参数
		if strings.HasPrefix(fwName, "{") && strings.HasSuffix(fwName, "}") {
			content := strings.TrimRight(strings.TrimLeft(fwName, "{"), "}")
			ps := strings.Split(content, ";")
			for _, p := range ps {
				vs := strings.Split(p, ":")
				if types.GetStringByIndex(vs, 1) != "" {
					opt.Params[types.GetStringByIndex(vs, 0)] = types.GetStringByIndex(vs, 1)
				}
			}
			if opt.Params["rwName"] != "" {
				opt.RwName = opt.Params["rwName"]
			}
			if opt.Params["fwName"] != "" {
				opt.FwName = opt.Params["fwName"]
			}
			if tb, ok := opt.Params["table"]; ok {
				opt.Table = tb
			}

			opt.ICON = types.GetString(opt.Params["icon"], opt.ICON)
			opts = append(opts, &opt)
			continue
		}
		if rwName != "" {
			opt.RwName = rwName
		}
		if fwName != "" {
			opt.FwName = fwName
		}
		opts = append(opts, &opt)
	}
	return opts
}
