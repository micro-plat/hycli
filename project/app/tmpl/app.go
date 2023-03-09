//go:build ignore

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/hydra/conf/app"
	_ "{{.PkgName}}/services"
)

// init 检查应用程序配置文件，并根据配置初始化服务
func init() {
	//启动时参数配置检查
	App.OnStarting(func(conf app.IAPPConf) error {
		//检查DB
		if _, err := components.Def.DB().GetDB(); err != nil {
			return err
		}
		return nil
	})
}
