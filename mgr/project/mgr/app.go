//go:build ignore

package main

import (
	"gitee.com/micro-plat/sso/sso"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/components"
	"github.com/micro-plat/hydra/conf/app"
	"github.com/micro-plat/hydra/hydra/servers/http"
	_ "{{.PkgName}}/api/services"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func init() {
	hydra.OnReady(func() error {
		if err := sso.Bind(hydra.S.Web, "http://localhost:6689", "sso", "B128F779D5741E701923346F7FA9F95C"); err != nil {
			return err
		}
		return nil
	})
	//启动时参数配置检查
	App.OnStarting(func(conf app.IAPPConf) error {
		//检查DB
		if _, err := components.Def.DB().GetDB(); err != nil {
			return err
		}

		//检查Cache
		if _, err := components.Def.Cache().GetCache(); err != nil {
			return err
		}
		return nil
	}, http.API)
}
