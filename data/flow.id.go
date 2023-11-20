package data

import (
	"fmt"
	"strings"
)

const (
	chars = "abcdefghijklmnopqrstuvwxyz"
)

var defFids = newFids(3)
var defFidsmp = map[string]string{}

type fids struct {
	len       int
	charIndex []int
}

func newFids(l int) *fids {
	return &fids{
		len:       l,
		charIndex: make([]int, l),
	}
}
func (f *fids) Current() string {
	sb := strings.Builder{}
	for _, v := range f.charIndex {
		sb.WriteString(string(chars[v]))
	}
	return sb.String()
}
func (f *fids) Next() string {
	add(f.charIndex)
	return f.Current()
}
func (f *fids) Get(o *optrs) string {
	if strings.EqualFold(o.Tag, "DEL") || strings.EqualFold(o.Tag, "CNFRM") {
		var key = fmt.Sprintf("%s/%s/%s", o.Table, o.Label, o.URL)
		_, ok := defFidsmp[key]
		if !ok {
			defFidsmp[key] = f.Next()
		}
		return defFidsmp[key]
	}
	return f.Next()

}
func add(idxs []int) {
	lenx := len(idxs)
	if lenx == 0 {
		return
	}

	if idxs[lenx-1] < len(chars)-1 {
		idxs[lenx-1]++
	} else {
		idxs[lenx-1] = 0
		if lenx > 1 {
			add(idxs[0 : len(idxs)-1])
		}
	}
}
