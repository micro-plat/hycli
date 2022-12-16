package data

var uiTableCaches = map[string]*UITable{}

func Cache(t *UITable) {
	uiTableCaches[t.Name.Raw] = t
	uiTableCaches[t.Name.Short] = t
}
func Get(name string) *UITable {
	if v, ok := uiTableCaches[name]; ok {
		return v
	}
	return &UITable{}
}
func Caches(t []*UITable) {
	for _, tb := range t {
		Cache(tb)
	}
}
