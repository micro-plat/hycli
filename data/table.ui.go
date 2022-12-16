package data

import (
	"encoding/json"
	"strings"

	"github.com/micro-plat/hycli/md"
	"github.com/micro-plat/lib4go/jsons"
)

type TUITable struct {
	*UITable
	IsTmpl bool
}
type UITable struct {
	*md.Table
	QRows     []*UIRow `json:'qrows'`   //查询行
	LRows     []*UIRow `json:'lrows'`   //列表行
	LERows    []*UIRow `json:'lerows'`  //列表扩展行
	CRows     []*UIRow `json:'crows'`   //创建行
	URows     []*UIRow `json:'urows'`   //更新行
	VRows     []*UIRow `json:'vrows'`   //预览行
	DRows     []*UIRow `json:'drows'`   //删除行
	PKRows    []*UIRow `json:'pk_rows'` //主键行
	Optrs     []*Operation
	BatchOpts []*Operation
	QueryOpts []*Operation
	ViewOpts  []*Operation //预览
	UNQ       string       `json:'unq'`
}

func (u *UITable) String() string {
	return u.Name.String()
}
func (u *UITable) Copy() *UITable {
	tb := &UITable{}
	buff, err := jsons.Marshal(u)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(buff, &tb)
	if err != nil {
		panic(err)
	}
	return tb
}
func FltrUITables(t md.Tables) []*UITable {
	tbs := make([]*UITable, 0, 1)
	for _, v := range t {
		tbs = append(tbs, FltrUITable(v))
	}
	return tbs
}

// 全局查找指定表
func fltrSearchUITable(name *Operation) *TUITable {
	uname := name.URL
	if strings.Contains(name.URL, "/") {
		uname = strings.Replace(strings.Trim(name.URL, "/"), "/", "_", -1)
	}
	ut := Get(uname)
	return &TUITable{UITable: ut, IsTmpl: true}
}

func ResetSelectRelation(tbs []*UITable) {
	for _, tn := range tbs {
		lleRows := mergeUIRow(tn.LRows, tn.LERows)
		for _, row := range lleRows {
			if row.RType.IsSelect {
				pkRows := Get(row.RType.SelectName).PKRows
				if len(pkRows) > 0 {
					row.RType.FKName = pkRows[0].Name
				}

			}
		}
	}
}

func FltrUITable(t *md.Table) *UITable {
	if t.Cache != nil {
		return t.Cache.(*UITable)
	}
	table := &UITable{
		Table:     t,
		QRows:     make([]*UIRow, 0, 1),
		LRows:     make([]*UIRow, 0, 1),
		LERows:    make([]*UIRow, 0, 1),
		CRows:     make([]*UIRow, 0, 1),
		URows:     make([]*UIRow, 0, 1),
		VRows:     make([]*UIRow, 0, 1),
		DRows:     make([]*UIRow, 0, 1),
		PKRows:    make([]*UIRow, 0, 1),
		Optrs:     make([]*Operation, 0, 1),
		BatchOpts: make([]*Operation, 0, 1),
		QueryOpts: make([]*Operation, 0, 1),
		ViewOpts:  make([]*Operation, 0, 1),
		UNQ:       defFids.Next(),
	}
	t.Cache = table
	table.Optrs = newOperations("", t.ExtInfo, "lst")
	table.BatchOpts = newOperations("", t.ExtInfo, "blst")
	table.QueryOpts = newOperations("", t.ExtInfo, "qlst")
	table.ViewOpts = newOperations("", t.ExtInfo, "view")
	updateOptr := false
	viewOptr := false
	delOptr := false
	for _, r := range t.Rows {
		if md.HasConstraint(r.Constraints, "Q") {
			table.QRows = append(table.QRows, NewUIRow("Q", r))
		}
		if md.HasConstraint(r.Constraints, "L") {
			table.LRows = append(table.LRows, NewUIRow("L", r))
		}
		if md.HasConstraint(r.Constraints, "LE") {
			table.LERows = append(table.LERows, NewUIRow("LE", r))
		}
		if md.HasConstraint(r.Constraints, "C") {
			table.CRows = append(table.CRows, NewUIRow("C", r))
		}
		if md.HasConstraint(r.Constraints, "PK") {
			table.PKRows = append(table.PKRows, NewUIRow("", r))
		}
		if md.HasConstraint(r.Constraints, "U") {
			if !updateOptr {
				updateOptr = true
				table.Optrs = append(table.Optrs, newOperation(UPDATE))
			}
			table.URows = append(table.URows, NewUIRow("U", r))
		}
		if md.HasConstraint(r.Constraints, "D") && !delOptr {
			delOptr = true
			table.Optrs = append(table.Optrs, newOperation(DEL))
			table.DRows = append(table.DRows, NewUIRow("D", r))
		}
		if md.HasConstraint(r.Constraints, "V") {
			if !viewOptr {
				viewOptr = true
				table.Optrs = append(table.Optrs, newOperation(VIEW))
			}
			table.VRows = append(table.VRows, NewUIRow("V", r))
		}
	}

	return table
}
