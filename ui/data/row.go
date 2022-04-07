package data

import (
	"regexp"

	"github.com/micro-plat/hycli/md"
)

func toUIRow(r *md.Row) *UIRow {
	return New(r)
}

type UIType struct {
	Name      string
	Format    string
	Len       int
	InputType string
}

func NewUIType(name string, xlen int, cns []string) *UIType {
	uit := &UIType{Name: name, Len: xlen}
	reg := regexp.MustCompile(`f\(([^\(^\)]+)\)`)
	for _, c := range cns {
		names := reg.FindAllStringSubmatch(c, -1)
		if len(names) > 0 && len(names[0]) > 1 {
			uit.Format = names[0][1]
			goto OUT
		}
	}
OUT:
	uit.InputType = tpmap[name]
	return uit
}

type UIRow struct {
	*md.Row
	T *UIType
}

func New(r *md.Row) *UIRow {
	return &UIRow{Row: r, T: NewUIType(r.Type.Name, r.Type.Len, r.Constraints)}
}
