package data

import (
	"strings"

	"github.com/micro-plat/lib4go/types"
)

// config  配置信息
type config struct {
	KV map[string]string
}

// newConfig 构建表格配置信息
func newConfig(content string) *config {
	params := map[string]string{}
	ps := strings.Split(content, ";")
	for _, p := range ps {
		vs := strings.Split(p, ":")
		if key := types.GetStringByIndex(vs, 0); key != "" {
			params[key] = types.GetStringByIndex(vs, 1)
		}
	}
	return &config{KV: params}
}
func (c *config) IsStatic(v string) bool {
	return strings.HasPrefix(v, "@")
}
func (c *config) IsMember(v string) bool {
	return strings.HasPrefix(v, "#")
}
func (c *config) Type() string {
	return c.KV["@table"]
}
func (c *config) WriteLog() bool {
	return c.KV["@log"] == "true"
}
func (c *config) Get(v string) string {
	return c.KV[v]
}
