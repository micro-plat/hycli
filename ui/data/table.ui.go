package data

import (
	"github.com/micro-plat/hycli/md"
)

type UITable struct {
	*md.Table
	QRows  []*UIRow
	LRows  []*UIRow
	CRows  []*UIRow
	URows  []*UIRow
	PKRows []*UIRow
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
		Table:  t,
		QRows:  make([]*UIRow, 0, 1),
		LRows:  make([]*UIRow, 0, 1),
		CRows:  make([]*UIRow, 0, 1),
		URows:  make([]*UIRow, 0, 1),
		PKRows: make([]*UIRow, 0, 1),
	}
	for _, r := range t.Rows {
		if md.HasConstraint(r.Constraints, "Q") {
			table.QRows = append(table.QRows, NewUIRow("Q", r))
		}
		if md.HasConstraint(r.Constraints, "L") {
			table.LRows = append(table.LRows, NewUIRow("L", r))
		}
		if md.HasConstraint(r.Constraints, "C") {
			table.CRows = append(table.CRows, NewUIRow("C", r))
		}
		if md.HasConstraint(r.Constraints, "U") {
			table.URows = append(table.URows, NewUIRow("U", r))
		}
		if md.HasConstraint(r.Constraints, "PK") {
			table.PKRows = append(table.PKRows, NewUIRow("", r))
		}
	}

	return table
}
