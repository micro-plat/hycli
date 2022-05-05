package api

import (
	"embed"
	"fmt"

	"github.com/micro-plat/hycli/md"
	"github.com/micro-plat/hycli/mgr/data"
	"github.com/micro-plat/hycli/output"
	"github.com/micro-plat/lib4go/types"
	"github.com/urfave/cli"
)

//CreateServiceByCtx 创建页面文件
func CreateServiceByCtx(c *cli.Context) (err error) {
	if len(c.Args()) == 0 {
		return fmt.Errorf("未指定md文件路径")
	}
	//获取输出路径
	outpath := types.GetString(c.Args().Get(1), "")
	return CreateService(c.Args().First(), c.String("table"), outpath)

}
func CreateService(mdpath string, tbs string, outpath string) error {
	//转换md文件
	otbls, err := md.Mds2Tables(mdpath)
	if err != nil {
		return err
	}

	//过滤表格
	tbls := otbls.Filter(tbs, true)

	//设置包名称
	tbls.JoinPkgName(outpath)

	//创建服务文件
	for _, tb := range tbls {
		if err = createService(outpath, tb); err != nil {
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
