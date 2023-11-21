package data

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

	//BarOpts 工具栏操作(显示表格上方)
	BarOpts barOptrs

	//QueryOptrs 查询操作组件(显示在查询条件最后)
	QueryOptrs queryOptrs

	All optrslst
}

func (t *tableOptrs) Sort() {
	t.ListOpts.Merge(t.BarOpts)
	t.BarOpts.Sort()
	t.ViewOpts.Sort()
	t.ListOpts.Sort()
	t.LStatOpts.Sort()
	t.ChartOpts.Sort()
	t.ExtOpts.Sort()
	t.QueryOptrs.Sort()
	t.ViewExtCmptOpts.ALL.Sort()

	t.All = make(optrslst, 0)
	t.All = append(t.All, t.BarOpts...)
	t.All = append(t.All, t.ViewOpts...)
	t.All = append(t.All, t.ListOpts...)
	t.All = append(t.All, t.LStatOpts...)
	t.All = append(t.All, t.ChartOpts...)
	t.All = append(t.All, t.ExtOpts...)
	t.All = append(t.All, t.QueryOptrs...)
	t.All = append(t.All, t.ViewExtCmptOpts.ALL...)
}
