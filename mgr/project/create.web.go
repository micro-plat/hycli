package project

import (
	"embed"
	"fmt"

	"github.com/micro-plat/hycli/output"
	"github.com/micro-plat/lib4go/types"
	"github.com/urfave/cli"
)

//go:embed web
var tmpls embed.FS
var tmplName = "web"

// CreateWeb 创建web项目
func CreateWeb(name string, input interface{}, cover bool) error {
	return output.CreateFiles(tmpls,
		name,
		tmplName,
		tmplName,
		input, cover)
}

// CreateWebByCtx 创建web项目基础文件
func CreateWebByCtx(c *cli.Context) (err error) {
	if len(c.Args()) == 0 {
		return fmt.Errorf("未指定项目名称")
	}
	name := types.GetString(c.Args().First(), "web")
	return CreateWeb(name, NewProject(name), c.Bool("cover"))
}
