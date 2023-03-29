package api

import (
	"embed"
	"fmt"

	logs "github.com/lib4dev/cli/logger"
	"github.com/micro-plat/hycli/data"
	"github.com/micro-plat/hycli/output"
	"github.com/micro-plat/lib4go/types"
	"github.com/urfave/cli"
)

// CreateServiceByCtx 创建页面文件
func CreateServiceByCtx(c *cli.Context) (err error) {
	if len(c.Args()) == 0 {
		return fmt.Errorf("未指定md文件路径")
	}
	//获取输出路径
	outpath := types.GetString(c.Args().Get(1), "")
	cover := c.Bool("cover")
	return CreateService(c.Args().First(), c.String("table"), outpath, cover)

}
func CreateService(mdpath string, tbs string, outpath string, cover bool) error {

	logs.Log.Info("------------------生成接口文件------------------")
	//转换md文件
	otbls, err := data.Mds2Tables(mdpath)
	if err != nil {
		return err
	}

	//过滤表格
	tbls := otbls.Filter(tbs, true)

	//设置包名称
	tbls.JoinPkgName(outpath)

	ntbs := data.NewTables(tbls)

	//创建服务文件
	for _, tb := range ntbs {
		if err = createService(outpath, tb, cover); err != nil {
			return err
		}
		if err = createModules(outpath, tb, cover); err != nil {
			return err
		}
	}
	return nil
}

//go:embed crud/services
var srvsTmpls embed.FS
var srvsTmplName = "crud/services"
var prefix = "crud"

// createService 创建服务文件
func createService(root string, input interface{}, cover bool) error {
	return output.CreateFiles(srvsTmpls,
		root,
		prefix,
		srvsTmplName,
		input,
		cover,
		data.Funcs)
}
