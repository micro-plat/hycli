//go:build ignore

package enums

import (
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

//EnumsHandler 获取枚举处理服务
type EnumsHandler struct {
}

func NewEnumsHandler() *EnumsHandler {
	return &EnumsHandler{}
}

//EnumsHandler  获取枚举列表数据
func (u *EnumsHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取枚举数据列表--------")

	ctx.Log().Info("1.查询数据列表")
	sqls := getSQLs(ctx.Request().GetString("type"))
	list := make(types.XMaps, 0, 1)
	for _, sql := range sqls {
		items, err := hydra.C.DB().GetRegularDB().Query(sql, ctx.Request().GetMap())
		if err != nil {
			return errs.NewErrorf(http.StatusNotExtended, "查询数据出错:%+v", err)
		}
		list = append(list, items...)
	}
	ctx.Log().Info("2.返回结果")
	return list
}
