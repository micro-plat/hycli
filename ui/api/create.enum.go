package api

import (
	"embed"
	"fmt"

	"github.com/micro-plat/hycli/md"
	"github.com/micro-plat/hycli/output"
	"github.com/micro-plat/hycli/ui/data"
	"github.com/urfave/cli"
)

//CreateEnums 创建页面文件
func CreateEnums(c *cli.Context) (err error) {
	if len(c.Args()) == 0 {
		return fmt.Errorf("未指定md文件路径")
	}

	//转换md文件
	otbls, err := md.Mds2Tables(c.Args().First())
	if err != nil {
		return err
	}

	//过滤表格
	tbls := otbls.Filter(c.String("table"))

	//获取输出路径
	root := "./"
	if len(c.Args()) > 1 {
		root = c.Args().Get(1)
	}

	return createEnums(root, tbls)
}

//go:embed comm/services
var enumTmpls embed.FS
var enumTmplName = "comm/services"
var enumPrefix = "comm"

//createEnums 创建服务文件
func createEnums(root string, input interface{}) error {
	return output.CreateFiles(enumTmpls,
		root,
		enumPrefix,
		enumTmplName,
		input,
		data.Funcs)
}
