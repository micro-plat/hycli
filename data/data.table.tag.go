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
	view := table.Columns.GetViewColumns().Len() > 0
	add := table.Columns.GetAllCreateColumns().Len() > 0
	edit := table.Columns.GetAllUpdateColumns().Len() > 0
	tag := &Tag{
		View:   types.DecodeString(view, true, "", "~"),
		Add:    types.DecodeString(add, true, "", "~"),
		Edit:   types.DecodeString(edit, true, "", "~"),
		Dialog: "~",
		Cnfrm:  "~",
	}
	opts := table.Optrs.ListOpts.Merge(table.Optrs.BarOpts,
		table.Optrs.ViewOpts,
		table.Optrs.ViewExtCmptOpts.ALL)
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
