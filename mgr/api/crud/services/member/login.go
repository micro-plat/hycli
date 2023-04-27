package member
{-{- $ft := .}-}
import (
	"github.com/micro-plat/hydra"
	"{-{$ft.PkgName}-}/modules/biz/member"
	"{-{$ft.PkgName}-}/modules/const/field"
)

var loginRequiredFields = []string{"user_name", "password"}

// LoginHandle  保存用户信息数据
func loginHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------用户登录数据--------")

	ctx.Log().Info("1.检查必须字段")
	if err := ctx.Request().Check(loginRequiredFields...); err != nil {
		return err
	}

	ctx.Log().Info("2.用户登录")
	state, err := member.Login(ctx.Request().GetString(field.UserName),
		ctx.Request().GetString(field.Password), "-")
	if err != nil {
		return err
	}
	ctx.User().Auth().Response(state)
	return state
}
