//go:build dev

package main

import (
	"github.com/micro-plat/hydra"
)

// 初始化配置
func init() {
	hydra.OnReady(func() {
		//配置共有配置
		hydra.Conf.Vars().DB().MySQLByConnStr("db",
			"{db.user_name}:{db.pwd}@tcp({db.ip}:3306)/{db.name}?charset=utf8")

	})
}
