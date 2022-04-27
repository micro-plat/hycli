//go:build ignore
//go:build dev

package main

import (
	"github.com/micro-plat/hydra"

	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/conf/server/auth/jwt"
	"github.com/micro-plat/hydra/conf/server/header"
	"github.com/micro-plat/hydra/conf/server/processor"
)

//初始化配置
func init() {
	hydra.OnReady(func() {

		//配置共有配置
		hydra.Conf.Vars().DB().MySQLByConnStr("db", "{db.user_name}:{db.pwd}@tcp({db.ip}:3306)/{db.name}?charset=utf8")

		//登录的界面配置
		hydra.Conf.Web("8089", api.WithTimeout(300, 300)).
			Processor(processor.WithServicePrefix("/api")).
			Header(header.WithCrossDomain()).
			Jwt(jwt.WithMode("HS512"),
				jwt.WithSecret("bf8f3171946d8d5a13cca23aa6080c8e"),
				jwt.WithExpireAt(36000),
				jwt.WithAuthURL("/login"),
				jwt.WithHeader(),
				jwt.WithExcludes("/**/member/*", "/**/login/*"),
			)

	})
}
