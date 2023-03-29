package api

import (
	"embed"

	"github.com/micro-plat/hycli/data"
	"github.com/micro-plat/hycli/output"
)

//go:embed crud/modules
var moduleTmpls embed.FS
var moduleTmplName = "crud/modules"

// createModules 创建服务文件
func createModules(root string, input interface{}, cover bool) error {
	return output.CreateFiles(moduleTmpls,
		root,
		prefix,
		moduleTmplName,
		input,
		cover,
		data.Funcs)
}
