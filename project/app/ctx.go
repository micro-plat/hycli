package app

import (
	"embed"

	"github.com/micro-plat/hycli/data"
	"github.com/micro-plat/hycli/output"
	"github.com/micro-plat/lib4go/types"
	"github.com/urfave/cli"
)

// CreateByCtx 创建页面文件
func CreateByCtx(c *cli.Context) (err error) {
	//获取输出路径
	outpath := types.GetString(c.Args().Get(1), "")
	cover := c.Bool("cover")

	data := map[string]interface{}{
		"platName": types.GetString(c.String("platName"), ""),
		"sysName":  types.GetString(c.String("sysName"), ""),
		"PkgName":  data.GetPkgName(),
	}
	return createFiles(outpath, data, cover)
}

//go:embed tmpl
var srvsTmpls embed.FS
var srvsTmplName = "tmpl"
var prefix = "tmpl"

// createService 创建服务文件
func createFiles(root string, input interface{}, cover bool) error {
	return output.CreateFiles(srvsTmpls,
		root,
		prefix,
		srvsTmplName,
		input,
		cover,
		data.Funcs)
}
