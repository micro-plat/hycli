package page

import (
	"embed"

	"github.com/micro-plat/hycli/output"
	"github.com/micro-plat/hycli/ui/data"
)

//go:embed html/src/views
var viewTmpls embed.FS
var viewTmplName = "html/src/views"
var prefix = "html"

//create 创建页面文件
func createViews(root string, input interface{}) error {
	return output.CreateFiles(viewTmpls,
		root,
		prefix,
		viewTmplName,
		input,
		data.Funcs)
}
