package data

import "github.com/micro-plat/lib4go/types"

type Tag struct {
	View   string
	Add    string
	Edit   string
	Dialog string
	Cnfrm  string
}

func newTag(table *Table) *Tag {
	view := len(fltrColumns(table, VIEW_COLUMN)) > 0
	add := len(fltrColumns(table, ADD_COLUMN)) > 0
	edit := len(fltrColumns(table, UPDATE_COLUMN)) > 0
	tag := &Tag{
		View:   types.DecodeString(view, true, "", "~"),
		Add:    types.DecodeString(add, true, "", "~"),
		Edit:   types.DecodeString(edit, true, "", "~"),
		Dialog: "~",
		Cnfrm:  "~",
	}
	opts := table.ListOpts.Merge(table.BarOpts, table.ViewOpts, table.ViewExtCmptOpts.ALL)
	for _, opt := range opts {
		if opt.Tag == "DIALOG" {
			tag.Dialog = ""
			continue
		}
		if opt.Tag == "CNFRM" || opt.Tag == DELETE_TAG {
			tag.Cnfrm = ""
			continue
		}
	}
	return tag
}
