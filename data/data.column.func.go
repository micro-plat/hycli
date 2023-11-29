package data

import (
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

func (c *Column) GetOpts(name string) (string, string, string) {
	f, s, t := md.GetConsByTagIgnorecase(name, c.RawConsts...)
	return f, s, t
}

func (c *Column) GetParamMap(name string) map[string]string {
	s := md.GetParamByTag(name, c.RawConsts...)
	item := strings.Split(s, ";")
	param := make(map[string]string)
	for _, v := range item {
		kv := strings.Split(v, ":")
		if len(kv) == 0 || kv[0] == "" {
			continue
		}
		if len(kv) == 2 {
			param[kv[0]] = kv[1]
			continue
		}
		param[kv[0]] = ""
	}
	return param

}
func (c *Column) GetOpt(name string) string {
	v, _, _ := c.GetOpts(name)
	return v
}
func (c *Column) GetOptInt(name string, i ...int) int {
	return types.GetInt(c.GetOpt(name), i...)
}

func (c *Column) HasCmpnt(p string) bool {
	pages := strings.Split(p, "-")
	for _, p := range pages {
		if md.HasConstraint(c.Constraints, strings.ToLower(p), strings.ToUpper(p)) {
			return true
		}
	}
	return false
}

// GetCmpntValue
func (c *Column) GetCmpntValue(name string, def string) string {
	v, page, _ := md.GetConsByTagIgnorecase(name, c.Constraints...)
	pages := strings.Split(page, "-")
	if len(pages) == 0 {
		return v
	}
	for _, p := range pages {
		for _, p1 := range c.Ext.TP {
			if strings.EqualFold(p1, p) {
				return v
			}
		}
	}
	return def

}
