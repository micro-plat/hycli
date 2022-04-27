//go:build ignore

package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var App = hydra.NewApp(
	hydra.WithPlatName("mgrserver", "后台管理系统"),
	hydra.WithSystemName("mgrserver", "后台管理系统"),
	hydra.WithServerTypes(http.Web))

func main() {
	App.Start()
}
