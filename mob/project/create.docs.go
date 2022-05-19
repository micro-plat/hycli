package project

import (
	"embed"
	"fmt"

	"github.com/micro-plat/hycli/output"
	"github.com/micro-plat/lib4go/types"
	"github.com/urfave/cli"
)

//go:embed 0docs
var docsTmpls embed.FS
var docsTmplName = "0docs"

//CreateDocs 创建docs项目
func CreateDocs(name string, input interface{}) error {
	err := output.CreateFiles(docsTmpls,
		name,
		docsTmplName,
		docsTmplName,
		input)
	if err != nil {
		return err
	}
	return err
}

//CreateDocsByCtx 创建docs项目基础文件
func CreateDocsByCtx(c *cli.Context) (err error) {
	if len(c.Args()) == 0 {
		return fmt.Errorf("未指定项目名称")
	}
	name := types.GetString(c.Args().First(), "0docs")
	return CreateDocs(name, NewProject(name))
}
