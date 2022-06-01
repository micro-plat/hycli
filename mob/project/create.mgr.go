package project

import (
	"embed"
	"fmt"

	"github.com/micro-plat/hycli/output"
	"github.com/micro-plat/lib4go/types"
	"github.com/urfave/cli"
)

//go:embed server
var apiTmpls embed.FS
var apiTmplName = "server"

//CreateAPI 创建web项目
func CreateAPI(name string, input interface{}) error {
	err := output.CreateFiles(apiTmpls,
		name,
		apiTmplName,
		apiTmplName,
		input)
	if err != nil {
		return err
	}
	return err
}

//CreateAPIByCtx 创建API项目基础文件
func CreateAPIByCtx(c *cli.Context) (err error) {
	if len(c.Args()) == 0 {
		return fmt.Errorf("未指定项目名称")
	}
	name := types.GetString(c.Args().First(), "mgr")
	return CreateAPI(name, NewProject(name))
}
