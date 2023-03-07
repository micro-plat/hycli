package data

import (
	"github.com/micro-plat/hycli/md"
	"github.com/micro-plat/lib4go/types"
)

type CodeTable struct {
	*md.Table
	EType  *EnumType
	AsEnum bool
	QRows  []*CodeRow
	BQRows []*CodeRow
	LRows  []*CodeRow
	LERows []*CodeRow
	CRows  []*CodeRow
	URows  []*CodeRow
	VRows  []*CodeRow
	PKRows []*CodeRow
}

func flterMainTable(tbs []*CodeTable) []*CodeTable {
	cache := map[string]bool{}
	ntbs := make([]*CodeTable, 0, len(tbs))
	for _, t := range tbs {
		if _, ok := cache[t.Name.Main]; !ok {
			cache[t.Name.Main] = true
			ntbs = append(ntbs, t)
		}
	}
	return ntbs
}

// 过滤行
func fltrCodeTables(t md.Tables) []*CodeTable {
	tbs := make([]*CodeTable, 0, 1)
	for _, v := range t {
		tbs = append(tbs, fltrCodeTable(v))
	}
	return tbs
}

// 过滤行
func fltrCodeTable(t *md.Table) *CodeTable {

	table := &CodeTable{
		Table:  t,
		QRows:  make([]*CodeRow, 0, 1),
		BQRows: make([]*CodeRow, 0, 1),
		LRows:  make([]*CodeRow, 0, 1),
		LERows: make([]*CodeRow, 0, 1),
		CRows:  make([]*CodeRow, 0, 1),
		URows:  make([]*CodeRow, 0, 1),
		PKRows: make([]*CodeRow, 0, 1),
	}
	table.EType = &EnumType{DERows: make([]*CodeRow, 0, 1)}
	tp := &EnumType{}
	for _, r := range t.Rows {
		if md.HasConstraint(r.Constraints, "Q") {
			table.QRows = append(table.QRows, NewCodeRow("Q", r))
		}
		if md.HasConstraint(r.Constraints, "BQ") {
			table.BQRows = append(table.BQRows, NewCodeRow("BQ", r))
		}
		if md.HasConstraint(r.Constraints, "L") {
			table.LRows = append(table.LRows, NewCodeRow("L", r))
		}
		if md.HasConstraint(r.Constraints, "LE") {
			table.LERows = append(table.LERows, NewCodeRow("LE", r))
		}
		if md.HasConstraint(r.Constraints, "C") {
			table.CRows = append(table.CRows, NewCodeRow("C", r))
		}
		if md.HasConstraint(r.Constraints, "U") {
			table.URows = append(table.URows, NewCodeRow("U", r))
		}
		if md.HasConstraint(r.Constraints, "V") {
			table.VRows = append(table.VRows, NewCodeRow("V", r))
		}
		if md.HasConstraint(r.Constraints, "PK") {
			table.PKRows = append(table.PKRows, NewCodeRow("", r))
		}
		if md.HasConstraint(r.Constraints, "DI") {
			tp.Id = types.GetString(md.GetConsNameByTag("DI", r.Constraints...), r.Name)
		}
		if md.HasConstraint(r.Constraints, "DN") {
			tp.Name = types.GetString(md.GetConsNameByTag("DN", r.Constraints...), r.Name)
		}
		if md.HasConstraint(r.Constraints, "DT") {
			tp.Type = types.GetString(md.GetConsNameByTag("DT", r.Constraints...), r.Name)
			tp.Multiple = true
		}
		if md.HasConstraint(r.Constraints, "DE") {
			tp.DERows = append(tp.DERows, NewCodeRow("", r))
		}
		if md.HasConstraint(r.Constraints, "DS") {
			tp.Status = r.Name
		}
		if md.HasConstraint(r.Constraints, "DP") {
			tp.PID = r.Name
		}
		if md.HasConstraint(r.Constraints, "expire") {
			tp.Expire = r.Name
		}
	}
	tp.SetTypeIfAbsence(t.Name.Short)
	if tp.IsValid() {
		table.EType = tp
		table.AsEnum = true
	}
	return table
}
