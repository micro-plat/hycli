package data

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/jsons"
	"github.com/micro-plat/lib4go/types"
)

type optParams map[string]string

func (o optParams) Get(n string) string {
	for k, v := range o {
		if k == n || "@"+n == k {
			return v
		}
	}
	return ""
}
func (o optParams) GetStatic() map[string]string {
	lst := make(map[string]string)
	for k, v := range o {
		if strings.HasPrefix(k, "@") {
			lst[k] = v
		}
	}
	return lst
}
func (o optParams) Append(p map[string]string) {
	for k, v := range p {
		o[k] = v
	}
}
func (o optParams) IsBatchCheck(k string) bool {
	v := o.Get(k)
	return types.StringContains(batchCheck, v)
}

// 操作信息
type optrs struct {
	Tag       string //
	Name      string //link,dialog,cmpnt,tab,cnfrm,form,view
	Cmd       string // batupdator lstbar chart
	Label     string //修改,预览，删除
	Desc      string //完整说明
	ICON      string //图标
	URL       string //组件URL
	RwName    string //当前表字段标记
	FwName    string //外部表字段标记
	ReqURL    string //服务请求URL
	IsMux     bool   //是否是复用组件
	index     int    //顺序编号
	Table     string
	Params    optParams
	ParentUNQ string //父组件编号
	UNQ       string
}

type optrslst []*optrs

type lstOptrs = optrslst
type viewOptrs = optrslst
type lstatOptrs = optrslst
type chartOptrs = optrslst
type barOptrs = optrslst
type viewExtCmptOpts = optrslst
type queryOptrs = optrslst
type extOptrs = optrslst

var viewOptCmd = []string{"view"}
var lstatOptCmd = []string{"lstat"}
var lstBarOptCmd = []string{"lst", "lstupdator"}
var batchCheck = []string{"bcheck", "@bcheck", "&bcheck"}
var qBarOptrCmd = []string{"export", "import", "lstbar", "batupdator", "barinsert", "batinsert", "list"}
var charOptrCmd = []string{"chart"}
var extCmptParam = []string{"add", "update", "view"}
var extOptrsCmd = []string{"tskbar", "xbar"}
var queryOptrCmd = []string{"qbar"}

const (
	PSN_QBAR   = "qbar"
	PSN_LSTBAR = "lstbar"
	PSN_QFORM  = "qform"
)

var addOpts = &optrs{Tag: ADD_TAG, URL: "@/views/{@prefix}/{@main}/{@name}.add", Name: "CMPNT", ICON: "Plus", Label: "添加", Desc: "添加", RwName: "C", UNQ: defFids.Next(), index: 1}
var queryOpts = &optrs{Tag: QUERY_TAG, URL: "", Name: "FORM", ICON: "Search", RwName: QUERY_COLUMN, Label: "查询", Desc: "查询", UNQ: defFids.Next(), index: 1}
var detailOpts = &optrs{Tag: VIEW_TAG, URL: "@/views/{@prefix}/{@main}/{@name}.view", Name: "CMPNT", Label: "详情", Desc: "详情", RwName: "V", UNQ: defFids.Next(), index: 1}
var updateOpts = &optrs{Tag: UPDATE_TAG, URL: "@/views/{@prefix}/{@main}/{@name}.edit", Name: "CMPNT", Label: "修改", Desc: "修改", RwName: "U", UNQ: defFids.Next(), index: 2}
var delOpts = &optrs{Tag: DELETE_TAG, URL: "@/views/{@prefix}/{@main}/{@name}.cnfrm", ReqURL: "/{@mainPath}/del", Name: "CMPNT", Label: "删除", Desc: "删除", RwName: "D", IsMux: true, index: 99}
var dialogOpts = &optrs{Tag: "DIALOG", URL: "@/views/{@prefix}/{@main}/{@name}.dialog", Name: "CMPNT", IsMux: true, index: 99}
var cnfrmOpts = &optrs{Tag: "CNFRM", URL: "@/views/{@prefix}/{@main}/{@name}.cnfrm", Name: "CMPNT", IsMux: true, index: 99}
var chartLinePieBarOpts = &optrs{Tag: "CHART", URL: "@/views/cmpnts/chart.base.vue", Name: "CMPNT"}
var taskBarOpts = &optrs{Tag: "TSKBAR", URL: "@/views/cmpnts/task.bar.vue", Name: "CMPNT"}
var xBarOpts = &optrs{Tag: "XBAR", URL: "@/views/cmpnts/xbar.vue", Name: "CMPNT"}

func (s *optrs) Get(tableName string) *optrs {
	nopts := *s
	nopts.Table = tableName
	// nopts.UNQ = defFids.Next()  11/18
	nopts.UNQ = defFids.Get(&nopts)
	return &nopts
}
func (s *optrs) GetAndSetTag(tableName string, tagName string) *optrs {
	nopts := *s
	nopts.Table = tableName
	nopts.Tag = tagName
	// nopts.UNQ = defFids.Next() 11/18
	nopts.UNQ = defFids.Get((&nopts))
	return &nopts
}
func (s *optrs) NeedBatchCheck() bool {
	for _, v := range s.Params {
		if types.StringContains(batchCheck, v) {
			return true
		}
	}
	return false
}
func (s *optrs) GetParams(p string) string {
	return s.Params.Get(p)
}

func (s *optrs) String() string {
	buff, _ := jsons.Marshal(s)
	return string(buff)
}
func (s *optrs) UNQKey() string {
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s",
		s.Table, s.ReqURL, s.URL,
		s.Name, s.RwName, s.FwName,
		s.Cmd)
}
func (s optrslst) Len() int           { return len(s) }
func (s optrslst) Less(i, j int) bool { return s[i].index < s[j].index }
func (s optrslst) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s optrslst) Find(tag string) bool {
	for _, v := range s {
		if strings.EqualFold(v.Tag, tag) {
			return true
		}
	}
	return false
}
func (s optrslst) Get(tag string) (*optrs, bool) {
	for _, v := range s {
		if strings.EqualFold(v.Tag, tag) {
			return v, true
		}
	}
	return nil, false
}
func (s optrslst) FindName(name string) bool {
	for _, v := range s {
		if strings.EqualFold(v.Name, name) {
			return true
		}
	}
	return false
}
func (s optrslst) String() string {
	buff, _ := jsons.Marshal(s)
	return string(buff)
}
func (s optrslst) Gets(tableName string) optrslst {
	optrslst := make(optrslst, 0, len(s))
	for _, v := range s {
		optrslst = append(optrslst, v.Get(tableName))
	}
	return optrslst
}

var extCmptOpts = map[string]*optrs{
	"add":    addOpts,
	"update": updateOpts,
	"view":   detailOpts,
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

func createLstBarOptrs(table *Table, t string) lstOptrs {
	optrs := createCmdsOptrs(table.Name.Raw, t, lstBarOptCmd)
	//构建操作
	if !optrslst(optrs).Find(VIEW_TAG) && len(table.GetColumnsByTPName(VIEW_COLUMN)) > 0 {
		optrs = append(optrs, detailOpts.Get(table.Name.Raw))
	}
	if !optrslst(optrs).Find(UPDATE_TAG) && len(table.GetColumnsByTPName(UPDATE_COLUMN)) > 0 {
		optrs = append(optrs, updateOpts.Get(table.Name.Raw))
	}

	if !optrslst(optrs).Find(DELETE_TAG) && len(table.GetColumnsByTPName(DELETE_COLUMN)) > 0 {
		optrs = append(optrs, delOpts.GetAndSetTag(table.Name.Raw, CNFRM))
	}
	if tag, ok := optrslst(optrs).Get(DELETE_TAG); ok {
		tag.Tag = CNFRM
	}

	return optrs
}
func createViewOptrs(table *Table, t string) (viewOptrs, viewExtCmptOpts) {
	viewOpts := createCmdsOptrs(table.Name.Raw, t, viewOptCmd)
	//构建操作
	if !optrslst(viewOpts).FindName(VIEW_TAG) && len(table.GetColumnsByTPName(VIEW_COLUMN)) > 0 {
		view := detailOpts.Get(table.Name.Raw)
		view.Name = "view"
		viewOpts = append(viewOpts, view)
	}
	extOpts := creatExtCmptOpts(viewOpts...)
	return viewOpts, extOpts
}
func createQueryOptrs(table *Table, t string) queryOptrs {
	opts := createCmdsOptrs(table.Name.Raw, t, queryOptrCmd)
	if !optrslst(opts).Find(QUERY_TAG) && len(table.GetColumnsByTPName(QUERY_COLUMN)) > 0 {
		opts = append(opts, queryOpts.Get(table.Name.Raw))
	}
	return opts
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
func createQBarOptrs(table *Table, t string) (barOptrs, bool) {
	opts := createCmdsOptrs(table.Name.Raw, t, qBarOptrCmd)
	if !optrslst(opts).Find(ADD_TAG) && len(table.GetColumnsByTPName(ADD_COLUMN)) > 0 {
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

func creatExtCmptOpts(opts ...*optrs) []*optrs {
	nopts := make([]*optrs, 0, 1)
	for _, view := range opts {
		for _, cmd := range extCmptParam {
			op, okOptr := extCmptOpts[cmd]
			v, ok := view.Params[cmd]
			if ok && okOptr {
				xview := *op
				// xview.UNQ = view.UNQ
				xview.Label = types.DecodeString(types.GetBool(v), true, view.Label, v)
				xview.Desc = xview.Label
				xview.Params = view.Params
				xview.Params["table"] = view.URL
				xview.Table = view.URL
				xview.ParentUNQ = view.UNQ
				// xview.UNQ = defFids.Next() 11/18
				xview.UNQ = defFids.Get(&xview)
				nopts = append(nopts, &xview)
			}
		}
	}
	return nopts
}

func createOptrs(tableName string, extInfo string, tag string) []*optrs {
	list := md.GetExtOpt(extInfo, tag)
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
		case "IMG":
			opt = *xBarOpts
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
		// opt.UNQ = defFids.Next() 11/18
		opt.UNQ = defFids.Get(&opt)
		opt.Params = make(map[string]string)
		rwName := types.GetStringByIndex(lst, 3)
		fwName := types.GetStringByIndex(lst, 4)
		//解析组件参数
		if strings.HasPrefix(fwName, "{") && strings.HasSuffix(fwName, "}") {
			content := strings.TrimRight(strings.TrimLeft(fwName, "{"), "}")
			ps := strings.Split(content, ";")
			for _, p := range ps {
				vs := strings.Split(p, ":")
				if key := types.GetStringByIndex(vs, 0); key != "" {
					opt.Params[key] = types.GetStringByIndex(vs, 1)
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
			if v, ok := opt.Params["idx"]; ok {
				opt.index = types.GetInt(v)
			}
			if v, ok := opt.Params["desc"]; ok {
				opt.Desc = v
			}
			if opt.Desc == "" {
				opt.Desc = opt.Label
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
		if opt.Desc == "" {
			opt.Desc = opt.Label
		}
		opts = append(opts, &opt)
	}
	return opts
}
