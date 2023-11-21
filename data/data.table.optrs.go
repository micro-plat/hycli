package data

import "github.com/micro-plat/lib4go/types"

type tableOptrs struct {

	//ViewOpts 预览操作
	ViewOpts viewOptrs

	//ViewExtCmptOpts 预览页扩展组件
	ViewExtCmptOpts *viewExtOptrsMap

	//ListOpts 列表组件(显示在每行数据操作栏)
	ListOpts lstOptrs

	//LStatOpts 统计组件(显示在查询条件与表格之间)
	LStatOpts lstatOptrs

	//ChartOpts 图表组件(显示到统计组件下方)
	ChartOpts chartOptrs

	//ExtOpts 扩展组件(显示到图标组件的下方)
	ExtOpts extOptrs

	//图表组件,包含LStatOpts，ChartOpts，ExtOpts
	GraphOpts optrslst

	//BarOpts 工具栏操作(显示表格上方)
	BarOpts barOptrs

	//QueryOptrs 查询操作组件(显示在查询条件最后)
	QueryOptrs queryOptrs

	All optrslst

	OptrtMap map[string]*optrs
}

func NewTableOptrs() *tableOptrs {
	return &tableOptrs{
		OptrtMap: make(map[string]*optrs),
		All:      make(optrslst, 0),
	}
}
func (t *tableOptrs) Sort() {
	// t.ListOpts.Merge(t.BarOpts)
	t.BarOpts.Sort()
	t.ViewOpts.Sort()
	t.ListOpts.Sort()
	t.LStatOpts.Sort()
	t.ChartOpts.Sort()
	t.ExtOpts.Sort()
	t.QueryOptrs.Sort()
	t.ViewExtCmptOpts.ALL.Sort()

	all := make(optrslst, 0)
	all = append(all, t.BarOpts...)
	all = append(all, t.ViewOpts...)
	all = append(all, t.ListOpts...)
	all = append(all, t.LStatOpts...)
	all = append(all, t.ChartOpts...)
	all = append(all, t.ExtOpts...)
	all = append(all, t.QueryOptrs...)
	all = append(all, t.ViewExtCmptOpts.ALL...)
	for _, v := range all {
		if _, ok := t.OptrtMap[v.UNQ]; !ok {
			t.OptrtMap[v.UNQ] = v
			t.All = append(t.All, v)
		}
	}
	t.GraphOpts = append(t.GraphOpts, t.LStatOpts...)
	t.GraphOpts = append(t.GraphOpts, t.ChartOpts...)
	t.GraphOpts = append(t.GraphOpts, t.ExtOpts...)
	t.GraphOpts.Sort()
}
func (s optrslst) GetCompnents() optrslst {
	return s.GetByName("CMPNT")
}
func (s optrslst) GetConfirm() optrslst {
	return s.GetByTag("CNFRM")
}
func (s optrslst) GetDialog() optrslst {
	return s.GetByTag("DIALOG")
}
func (opts optrslst) GetAddOpt(tb ...string) *optrs {
	for _, v := range opts {
		if v.Tag == ADD_TAG {
			if len(tb) == 0 {
				return v
			}
			if types.StringContains(tb, v.Table) {
				return v
			}
		}
	}
	return &optrs{}
}
func (opts optrslst) GetEditOpt(tb ...string) *optrs {
	for _, v := range opts {
		if v.Tag == UPDATE_TAG {
			if len(tb) == 0 {
				return v
			}
			if types.StringContains(tb, v.Table) {
				return v
			}
		}
	}
	return &optrs{}
}
func (t *tableOptrs) SetViewExtCmptOpt(v *viewExtOptrsMap) {
	v.ALL.Sort()
	t.ViewExtCmptOpts = v
	for _, v := range t.ViewExtCmptOpts.ALL {
		if _, ok := t.OptrtMap[v.UNQ]; !ok {
			t.OptrtMap[v.UNQ] = v
			t.All = append(t.All, v)
		}
	}

}
