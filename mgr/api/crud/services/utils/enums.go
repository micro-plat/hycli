package utils

import "github.com/micro-plat/lib4go/types"

var enumsMap = map[string]map[string]types.XMap{}

func Append(xps types.XMaps) {
	for _, xp := range xps {
		tp := xp.GetString("type")
		if _, ok := enumsMap[tp]; !ok {
			enumsMap[tp] = map[string]types.XMap{}
			enumsMap[tp][xp.GetString("name")] = xp
			continue
		}
		enumsMap[tp][xp.GetString("name")] = xp
	}
}

func GetValue(tp string, name string) string {
	if v, ok := enumsMap[tp]; ok {
		if y, ok := v[name]; ok {
			return y.GetString("value")
		}
	}
	return ""
}
