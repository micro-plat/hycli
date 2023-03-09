package data

import (
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

// 操作信息
type optrs struct {
	Name   string //link,dialog,cmpnt,tab,cnfrm
	Label  string //修改,预览，删除
	URL    string
	RwName string
	FwName string
	UNQ    string
}
type lstOptrs []*optrs
type viewOptrs []*optrs

var detailOpts = &optrs{Name: "VIEW", Label: "详情", RwName: "V"}
var updateOpts = &optrs{Name: "UPDATE", Label: "修改", RwName: "U"}
var delOpts = &optrs{Name: "DEL", Label: "删除", RwName: "D"}

func createLstOptrs(t string) lstOptrs {
	optrs := make([]*optrs, 0, 1)
	optrs = append(optrs, createOptrs(t, "lst")...)
	optrs = append(optrs, createOptrs(t, "LST")...)
	return optrs
}
func createViewOptrs(t string) viewOptrs {
	optrs := make([]*optrs, 0, 1)
	optrs = append(optrs, createOptrs(t, "view")...)
	optrs = append(optrs, createOptrs(t, "VIEW")...)
	return optrs
}

func createOptrs(t string, tag string) []*optrs {
	list := md.GetExtOpt(t, tag)
	opts := make([]*optrs, 0, len(list))
	for _, lst := range list {
		opt := &optrs{
			Name:   strings.ToUpper(types.GetStringByIndex(lst, 1)),
			Label:  lst[0],
			URL:    types.GetStringByIndex(lst, 2),
			RwName: types.GetStringByIndex(lst, 3),
			FwName: types.GetStringByIndex(lst, 4),
			UNQ:    defFids.Next(),
		}
		opts = append(opts, opt)
	}
	return opts
}
