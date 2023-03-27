package data

import "github.com/micro-plat/lib4go/types"

type Tag struct {
	View string
	Add  string
	Edit string
}

func newTag(view bool, add bool, edit bool) *Tag {
	return &Tag{
		View: types.DecodeString(view, true, "", "~"),
		Add:  types.DecodeString(add, true, "", "~"),
		Edit: types.DecodeString(edit, true, "", "~"),
	}
}
