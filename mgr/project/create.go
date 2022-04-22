package project

import (
	"embed"
	"fmt"

	"github.com/micro-plat/hycli/output"
	"github.com/urfave/cli"
)

//go:embed html
var tmpls embed.FS
var tmplName = "html"

//CreateProject 创建web项目基础文件
func Create(c *cli.Context) (err error) {
	if len(c.Args()) == 0 {
		return fmt.Errorf("未指定项目名称")
	}
	return create(c.Args().First())
}

//create 创建web项目
func create(name string) error {
	return output.CreateFiles(tmpls,
		name,
		tmplName,
		tmplName,
		&Project{Name: name})
}
