//go:build ignore
{-{- $table := .|f_table_first}-}
package enums

import (
	"net/http"
	"sync"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
	"{-{$table.PkgName}-}/services/utils"
	"{-{$table.PkgName}-}/modules/biz/member"
)

// EnumsHandler 获取枚举处理服务
type EnumsHandler struct {
	once sync.Once
}

func NewEnumsHandler() *EnumsHandler {
	return &EnumsHandler{}
}

// EnumsHandler  获取枚举列表数据
func (u *EnumsHandler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取枚举数据列表--------")
	u.once.Do(u.load)
	ctx.Log().Info("1.查询数据列表",ctx.Request().GetString("type"))
	sqls := getSQLs(ctx.Request().GetString("type"))
	list := make(types.XMaps, 0, 1)
	member.GetSetMemberCtx(ctx)
	for _, sql := range sqls {
		items, err := hydra.C.DB().GetRegularDB().Query(sql, ctx.Request().GetMap())
		if err != nil {
			return errs.NewErrorf(http.StatusNotExtended, "查询数据出错:%+v", err)
		}
		list = append(list, items...)
	}
	utils.Append(list)
	ctx.Log().Info("2.返回结果")
	return list
}
func (u *EnumsHandler) load() {
	for _, v := range enumSQL {
		sqls = append(sqls, v)
	}
	sqls = append(sqls, unspecifiedEnum...)
	currentSQL := getSQLs("")
	list := make(types.XMaps, 0, 1)
	for _, sql := range currentSQL {
		items, _ := hydra.C.DB().GetRegularDB().Query(sql, map[string]interface{}{})
		list = append(list, items...)
	}
	utils.Append(list)
}
