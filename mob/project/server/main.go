//go:build ignore

package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var App = hydra.NewApp(
	hydra.WithPlatName("mobserver", "移动端服务系统"),
	hydra.WithSystemName("mobserver", "移动端服务系统"),
	hydra.WithServerTypes(http.Web))

func main() {
	App.Start()
}
