package data

import (
	"strings"

	"github.com/micro-plat/hycli/md"
	"github.com/micro-plat/lib4go/jsons"
	"github.com/micro-plat/lib4go/types"
)

const (
	UPDATE = "UPDATE"
	VIEW   = "VIEW"
	DEL    = "DEL"
)

var optLabels = map[string]string{
	UPDATE: "修改",
	VIEW:   "预览",
	DEL:    "删除",
}
var RwNames = map[string]string{
	UPDATE: "U",
	VIEW:   "V",
	DEL:    "D",
}

type Operation struct {
	Name   string //update,view,del,link,dialog,cmpnt,cnfrm
	Label  string //修改,预览，删除
	URL    string
	RwName string
	UNQ    string
}

func (o *Operation) String() string {
	r, _ := jsons.Marshal(o)
	return string(r)
}

func newOperation(name string, tname ...string) *Operation {
	return &Operation{
		Name:   name,
		Label:  types.GetString(optLabels[name], types.GetStringByIndex(tname, 0)),
		UNQ:    defFids.Next(),
		RwName: RwNames[name],
	}
}

//lst(授权,link->/a/b,m)
func newOperations(t string, tag string) []*Operation {
	list := md.GetExtOpt(t, tag)
	opts := make([]*Operation, 0, len(list))
	for _, lst := range list {
		opts = append(opts, &Operation{
			Name:   strings.ToUpper(types.GetStringByIndex(lst, 1)),
			Label:  lst[0],
			URL:    types.GetStringByIndex(lst, 2),
			RwName: types.GetStringByIndex(lst, 3),
			UNQ:    defFids.Next(),
		})
	}
	return opts
}
