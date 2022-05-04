package project

import (
	"embed"
	"fmt"

	"github.com/micro-plat/hycli/output"
	"github.com/urfave/cli"
)

//go:embed docs
var docsTmpls embed.FS
var docsTmplName = "docs"

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
	name := "docs"
	return CreateAPI(name, NewProject(name))
}
