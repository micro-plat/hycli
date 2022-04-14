package page

import (
	"fmt"

	"github.com/micro-plat/hycli/md"
	"github.com/urfave/cli"
)

//CreateProject 创建页面文件
func Create(c *cli.Context) (err error) {
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

	//创建页面文件
	for _, tb := range tbls {
		if err = createViews(root, tb); err != nil {
			return err
		}
	}

	// //创建路由文件
	// if err = createRouter(root, tbls); err != nil {
	// 	return err
	// }
	return nil

}
