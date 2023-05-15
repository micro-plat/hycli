package data

var uiTableCaches = map[string]*Table{}

func Cache(t *Table) {
	uiTableCaches[t.Name.Raw] = t
	uiTableCaches[t.Name.Short] = t
}
func Get(name string) *Table {
	if v, ok := uiTableCaches[name]; ok {
		return v
	}
	return &Table{}
}
func Caches(t Tables) {
	for _, tb := range t {
		Cache(tb)
	}

	//全局处理扩展选项
	for _, tx := range t {
		tx.LoadExtOptrs()
	}
}
