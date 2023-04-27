package member
{-{- $ft := .}-}
import (
	"github.com/micro-plat/hydra"
	"{-{$ft.PkgName}-}/modules/biz/system"
)

// getSystemHandle  获取系统信息
func getSystemHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取系统数据--------")

	ctx.Log().Info("2.获取系统信息")
	system, err := system.GetSystemInfo("-")
	if err != nil {
		return err
	}
	return system
}
