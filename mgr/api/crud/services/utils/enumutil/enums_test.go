package enumutil

import (
	"testing"

	"github.com/micro-plat/lib4go/assert"
	"github.com/micro-plat/lib4go/types"
)

func TestV(t *testing.T) {
	nmp := types.NewXMaps()
	nmp.Append(types.XMap{
		"name":  "保险业务商城支持功能",
		"type":  "dev_plan",
		"value": "175",
	})
	nmp.Append(types.XMap{
		"name":  "通用商城迭代v2.22",
		"type":  "dev_plan",
		"value": "176",
	})
	nmp.Append(types.XMap{
		"name":  "v2.22",
		"type":  "plan",
		"value": "174",
	})
	Append(nmp)
	assert.Equal(t, 2, len(enumsMap))
	assert.Equal(t, "175", GetValue("dev_plan", "保险业务商城支持功能"))
}
