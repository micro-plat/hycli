package data

import (
	"github.com/micro-plat/hycli/md"
)

type UITable struct {
	*md.Table
	QRows     []*UIRow //查询行
	LRows     []*UIRow //列表行
	LERows    []*UIRow //列表扩展行
	CRows     []*UIRow //创建行
	URows     []*UIRow //更新行
	VRows     []*UIRow //预览行
	DRows     []*UIRow //删除行
	PKRows    []*UIRow //主键行
	Optrs     []*Operation
	btExtOpt  []*Operation
	extExtOpt []*Operation
}

func fltrUITables(t md.Tables) []*UITable {
	tbs := make([]*UITable, 0, 1)
	for _, v := range t {
		tbs = append(tbs, fltrUITable(v))
	}
	return tbs
}
func fltrUITable(t *md.Table) *UITable {
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
		btExtOpt:  make([]*Operation, 0, 1),
		extExtOpt: make([]*Operation, 0, 1),
	}
	table.Optrs = newOperations(t.ExtInfo, "lst")
	table.btExtOpt = newOperations(t.ExtInfo, "bt")
	table.extExtOpt = newOperations(t.ExtInfo, "ext")
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
			if updateOptr == false {
				updateOptr = true
				table.Optrs = append(table.Optrs, newOperation(UPDATE))
			}
			table.URows = append(table.URows, NewUIRow("U", r))
		}
		if md.HasConstraint(r.Constraints, "D") && delOptr == false {
			delOptr = true
			table.Optrs = append(table.Optrs, newOperation(DEL))
			table.DRows = append(table.DRows, NewUIRow("D", r))
		}
		if md.HasConstraint(r.Constraints, "V") {
			if viewOptr == false {
				viewOptr = true
				table.Optrs = append(table.Optrs, newOperation(VIEW))
			}
			table.VRows = append(table.VRows, NewUIRow("V", r))
		}
	}

	return table
}
