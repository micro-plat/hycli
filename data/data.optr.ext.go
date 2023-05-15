package data

import (
	"strings"

	"github.com/micro-plat/lib4go/types"
)

type viewExtOptrs struct {
	BarOptrs optrslst
	LstOptrs optrslst
}
type viewExtOptrsMap struct {
	ALL optrslst
	Map map[string]*viewExtOptrs
}

func (v viewExtOptrsMap) GetBarOptrs(tb string, parentUNQID string) optrslst {
	if p, ok := v.Map[tb]; ok {
		optlst := make(optrslst, 0, 1)
		for _, v := range p.BarOptrs {
			if v.ParentUNQ == parentUNQID {
				optlst = append(optlst, v)
			}
		}
		return optlst
	}
	return optrslst{}
}
func (v viewExtOptrsMap) GetLstOptrs(tb string) optrslst {
	if p, ok := v.Map[tb]; ok {
		return p.LstOptrs
	}
	return optrslst{}
}
func getViewExtCmptOptsByTable(t *Table) *viewExtOptrsMap {
	xtp := &viewExtOptrsMap{ALL: make(optrslst, 0, 1), Map: make(map[string]*viewExtOptrs)}
	for _, v := range t.ViewOpts {
		if strings.EqualFold(t.Name.Raw, v.Table) {
			continue
		}
		if _, ok := xtp.Map[v.Table]; !ok {
			xtp.Map[v.Table] = &viewExtOptrs{BarOptrs: make(optrslst, 0, 1), LstOptrs: make(optrslst, 0, 1)}
		}

		opts := getViewExtCmptOpts(v.Table, v.UNQ, v.Params.GetStatic())
		xtp.ALL = append(xtp.ALL, opts.BarOptrs...)
		xtp.ALL = append(xtp.ALL, opts.LstOptrs...)
		xtp.Map[v.Table].BarOptrs = append(xtp.Map[v.Table].BarOptrs, opts.BarOptrs...)
		xtp.Map[v.Table].LstOptrs = opts.LstOptrs
	}
	return xtp
}

func getViewExtCmptOpts(table string, parentUNQID string, params map[string]string) *viewExtOptrs {
	viewExtOpts := &viewExtOptrs{BarOptrs: make([]*optrs, 0, 1), LstOptrs: make([]*optrs, 0, 1)}
	tb := findTable(table)
	if tb.Table == nil {
		return viewExtOpts
	}
	for _, v := range tb.BarOpts {
		if strings.EqualFold(v.Params["@showInView"], "true") {
			nv := &optrs{}
			types.DeepCopyByGob(nv, v)
			nv.ParentUNQ = parentUNQID
			nv.UNQ = defFids.Next()
			nv.Params.Append(params)
			viewExtOpts.BarOptrs = append(viewExtOpts.BarOptrs, nv)
		}
	}
	for _, v := range tb.ListOpts {
		if strings.EqualFold(v.Params["@showInView"], "true") {
			nv := &optrs{}
			types.DeepCopyByGob(nv, v)
			nv.UNQ = defFids.Next()
			nv.Params.Append(params)
			viewExtOpts.LstOptrs = append(viewExtOpts.LstOptrs, nv)
		}
	}
	return viewExtOpts
}
