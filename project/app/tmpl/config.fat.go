//go:build fat

package main

import (
	"github.com/micro-plat/hydra"
)

// bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func init() {
	hydra.OnReady(func() {
		hydra.Conf.Vars().DB().MySQLByConnStr("db",
			"{db.user_name}:{db.pwd}@tcp({db.ip}:3306)/{db.name}?charset=utf8")

	})
}
