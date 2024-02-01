package api

import (
	"fmt"

	logs "github.com/lib4dev/cli/logger"
	"github.com/micro-plat/hycli/data"
	"github.com/micro-plat/hycli/output"
	"github.com/micro-plat/lib4go/types"
	"github.com/urfave/cli"
)

var curdTmplName = "crud/modules/const/db/curd/curd"
var curdPrefix = "crud/modules/const/db/curd"

// createCRUD 创建服务文件
func createCRUD(root string, input interface{}, cover bool) error {
	return output.CreateFiles(moduleTmpls,
		root,
		curdPrefix,
		curdTmplName,
		input,
		cover,
		data.Funcs)
}

// CreateSCURDcriptByCtx 创建CRUD脚本
func CreateSCURDcriptByCtx(c *cli.Context) (err error) {
	if len(c.Args()) == 0 {
		return fmt.Errorf("未指定md文件路径")
	}
	//获取输出路径
	outpath := types.GetString(c.Args().Get(1), "")
	cover := c.Bool("cover")
	return CreateCRUDScript(c.Args().First(), c.String("table"), outpath, cover)

}
func CreateCRUDScript(mdpath string, tbs string, outpath string, cover bool) error {

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

	data.Caches(ntbs)
	//创建服务文件
	for _, tb := range ntbs {
		if err = createCRUD(outpath, tb, cover); err != nil {
			return err
		}
	}
	return nil
}
