package member
{-{- $ft := .}-}
import (
	"github.com/micro-plat/hydra"
	"{-{$ft.PkgName}-}/modules/biz/member"
)

// getUserHandle  获取用户信息
func getUserHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取用户数据--------")

	ctx.Log().Info("1.获取用户信息")
	member, err := member.GetMemberState(ctx)
	if err != nil {
		return err
	}
	return member
}
