package system

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

func GetSystemInfo(ident string) (types.IXMap, error) {
	db := hydra.C.DB().GetRegularDB()
	data, err := db.Query(getSystemInfo, map[string]interface{}{
		"ident": ident,
	})
	if err != nil {
		return nil, err
	}
	if data.IsEmpty() {
		return nil, errs.NewErrorf(901, "系统不存在%s", ident)
	}
	return data.Get(0), nil
}
