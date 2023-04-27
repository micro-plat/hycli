package member
{-{- $ft := .}-}
import (
	"github.com/micro-plat/hydra"
	"{-{$ft.PkgName}-}/modules/biz/member"
	"{-{$ft.PkgName}-}/modules/biz/menu"
)

// getMenuHandle  获取菜单信息
func getMenuHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取菜单数据--------")

	ctx.Log().Info("1.获取用户信息")
	member, err := member.GetMemberState(ctx)
	if err != nil {
		return err
	}

	ctx.Log().Info("2.获取用户菜单", member.UserID, member.SysIdent)
	system, err := menu.GetMenusByUserID(member.UserID, member.SysIdent)
	if err != nil {
		return err
	}
	return system
}
