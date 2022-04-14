package output

import "strings"

var funcs = map[string]interface{}{
	"in":    in,
	"lower": lower,
	"minus": minus,
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
func lower(f string) string {
	return strings.ToLower(f)
}
func minus(v int) int {
	return v - 1
}
