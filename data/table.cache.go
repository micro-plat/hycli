package data

import "fmt"

var uiTableCaches = map[string]*UITable{}

func Cache(t *UITable) {
	if _, ok := uiTableCaches[t.Name.Raw]; !ok {
		uiTableCaches[t.Name.Raw] = t
	}
}
func Get(name string) *UITable {
	if name == "" {
		fmt.Printf("ui.table.name:%+v\n", uiTableCaches)
	}
	return uiTableCaches[name]
}
func Caches(t []*UITable) {
	for _, tb := range t {
		Cache(tb)
	}
}
