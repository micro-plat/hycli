//go:build ignore
//go:build prod

package main

import (
	"embed"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/conf/server/api"
	"github.com/micro-plat/hydra/conf/server/auth/jwt"
	"github.com/micro-plat/hydra/conf/server/header"
	"github.com/micro-plat/hydra/conf/server/processor"
	"github.com/micro-plat/hydra/conf/server/static"
)

//go:embed web/dist/static
var staticFs embed.FS
var archive = "web/dist/static"

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func init() {
	hydra.OnReady(func() {

		hydra.Conf.Vars().DB().MySQLByConnStr("db", "{db.user_name}:{db.pwd}@tcp({db.ip}:3306)/{db.name}?charset=utf8")

		hydra.Conf.Web("6688", api.WithTimeout(300, 300),
			api.WithDNS("sso.hydra-cloud.com")).
			Static(static.WithAutoRewrite(), static.WithEmbed(archive, staticFs)).
			Processor(processor.WithServicePrefix("/api")).
			Header(header.WithCrossDomain()).
			Jwt(jwt.WithMode("HS512"),
				jwt.WithSecret("f0abd74b09bcc61449d66ae5d8128c18"),
				jwt.WithExpireAt(36000),
				jwt.WithAuthURL("/"),
				jwt.WithHeader(),
			)

	})
}