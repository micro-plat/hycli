package output

var funcs = map[string]interface{}{
	"in": in,
}

//数组中是否包含某个值
func in(list []string, f string) bool {
	for _, v := range list {
		if v == f {
			return true
		}
	}
	return false
}
