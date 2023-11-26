package web

import (
	"embed"

	"github.com/micro-plat/hycli/data"
	"github.com/micro-plat/hycli/output"
)

//go:embed html/src/jspkg
var jspkgTmpls embed.FS
var jspkgTmplName = "html/src/jspkg"

// createjspkg 创建页面文件
func createjspkg(root string, input interface{}, cover bool) error {
	return output.CreateFiles(jspkgTmpls,
		root,
		prefix,
		jspkgTmplName,
		input,
		cover,
		data.Funcs)
}
