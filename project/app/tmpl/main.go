package main

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/http"
)

var App = hydra.NewApp(
	hydra.WithPlatName("{{.platName}}", "{{.platName}}"),
	hydra.WithSystemName("{{.sysName}}", "{{.sysName}}"),
	hydra.WithServerTypes(http.API))

func main() {
	App.Start()
}
