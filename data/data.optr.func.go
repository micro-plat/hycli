package data

import (
	"sort"
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

func (o optrslst) Merge(optsm ...optrslst) optrslst {
	lst := make(optrslst, 0, 1)
	v := map[string]bool{}
	xlst := make([]optrslst, 0, len(optsm)+1)
	xlst = append(xlst, o)
	xlst = append(xlst, optsm...)
	for _, opts := range xlst {
		for _, opt := range opts {
			if _, ok := v[opt.UNQ]; !ok {
				lst = append(lst, opt)
				v[opt.UNQ] = true
			}

		}
	}

	sort.Sort(lst)
	return lst

}
func (opts optrslst) GetByName(tps string) optrslst {
	nopts := make(optrslst, 0, 1)
	tpn := strings.Split(tps, "-")
	for _, v := range opts {
		for _, tp := range tpn {
			if strings.EqualFold(v.Name, tp) {
				nopts = append(nopts, v)
			}
		}
	}
	sort.Sort(nopts)
	return nopts
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
func (opts optrslst) GetByExternalTable(tb string) optrslst {
	nopts := make(optrslst, 0, 1)
	for _, v := range opts {
		if strings.EqualFold(v.Table, tb) {
			nopts = append(nopts, v)
		}
	}
	sort.Sort(nopts)
	return nopts
}

func (opts optrslst) GetSelected() *optrs {
	for _, optr := range opts {
		if v, ok := optr.Params["selected"]; ok && v == "true" {
			return optr
		}
	}
	if len(opts) > 0 {
		return opts[0]
	}
	return &optrs{}

}

func (opts optrslst) GetByCmptName(cmds string) optrslst {
	nopts := make(optrslst, 0, 1)
	cmd := strings.Split(cmds, "-")
	for _, v := range opts {
		for _, c := range cmd {
			if strings.EqualFold(v.Cmd, c) {
				nopts = append(nopts, v)
			}
		}
	}
	sort.Sort(nopts)
	return nopts
}

func (opts optrslst) GetByTag(tag string) optrslst {
	nopts := make(optrslst, 0, 1)
	tags := strings.Split(tag, "-")
	for _, v := range opts {
		for _, tp := range tags {
			if strings.EqualFold(v.Tag, tp) {
				nopts = append(nopts, v)
			}
		}
	}
	sort.Sort(nopts)
	return nopts
}

func (opts *optrs) GetParamsByAtPrefix() map[string]string {
	outs := make(map[string]string)
	for k, v := range opts.Params {
		if strings.HasPrefix(k, "@") {
			outs[strings.Trim(k, "@")] = v
		}
	}

	return outs
}
func (opt *optrs) GetParam(name string, def string) string {
	if v, ok := opt.Params[name]; ok {
		return v
	}
	return def
}
func (opt *optrs) GetExprs(name string, def string) []*expr {
	p := opt.GetParam(name, def)
	return str2Expr(p)
}
func str2Expr(f string) []*expr {
	prs := md.GetExprs(f)
	lst := make([]*expr, 0, 1)
	for _, pr := range prs {
		lst = append(lst, &expr{
			Name:   types.DecodeString(strings.HasPrefix(pr[0], "@"), true, "", pr[0]),
			Field:  types.DecodeString(strings.HasPrefix(pr[0], "@"), true, strings.Trim(pr[0], "@"), ""),
			Symbol: types.GetStringByIndex(pr, 1, ""),
			Value:  types.GetStringByIndex(pr, 2, ""),
		})
	}
	return lst
}
