package api

import (
	"embed"
	"fmt"

	"github.com/micro-plat/hycli/md"
	"github.com/micro-plat/hycli/output"
	"github.com/micro-plat/hycli/ui/data"
	"github.com/urfave/cli"
)

//CreateService 创建页面文件
func CreateService(c *cli.Context) (err error) {
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

	//创建服务文件
	for _, tb := range tbls {
		if err = createService(root, tb); err != nil {
			return err
		}
	}

	return nil

}

//go:embed crud/services
var srvsTmpls embed.FS
var srvsTmplName = "crud/services"
var prefix = "crud"

//createService 创建服务文件
func createService(root string, input interface{}) error {
	return output.CreateFiles(srvsTmpls,
		root,
		prefix,
		srvsTmplName,
		input,
		data.Funcs)
}
