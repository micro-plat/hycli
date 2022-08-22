//go:build ignore
//go:build dev

package main

import (
	"github.com/micro-plat/hydra"

	"gitee.com/micro-plat/sso/ssov5"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/conf/server/auth"
	"github.com/micro-plat/hydra/conf/server/auth/jwt"
	"github.com/micro-plat/hydra/conf/server/header"
	"github.com/micro-plat/hydra/conf/server/processor"
)

// 初始化配置
func init() {
	hydra.OnReady(func() error {

		//配置共有配置
		hydra.Conf.Vars().DB().MySQLByConnStr("db",
			"{db.user_name}:{db.pwd}@tcp({db.ip}:3306)/{db.name}?charset=utf8")

		return nil
	})
}
