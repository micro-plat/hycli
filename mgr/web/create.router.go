package web

import (
	"embed"

	"github.com/micro-plat/hycli/data"
	"github.com/micro-plat/hycli/output"
)

//go:embed html/src/router
var routerTmpls embed.FS
var routerTmplName = "html/src/router"

//create 创建页面文件
func createRouter(root string, input interface{}) error {
	return output.CreateFiles(routerTmpls,
		root,
		prefix,
		routerTmplName,
		input,
		data.Funcs)
}
