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

// CreateEnumsByCtx 创建页面文件
func CreateEnumsByCtx(c *cli.Context) (err error) {
	if len(c.Args()) == 0 {
		return fmt.Errorf("未指定md文件路径")
	}
	//获取输出路径
	outpath := types.GetString(c.Args().Get(1), "./")
	cover := c.Bool("cover")
	return CreateEnums(c.Args().First(), c.String("table"), outpath, cover)
}

// CreateEnums 创建页面文件
func CreateEnums(mdpath string, tbs string, outpath string, cover bool) error {

	logs.Log.Info("------------------生成枚举文件------------------")
	//转换md文件
	otbls, err := data.Mds2Tables(mdpath)
	if err != nil {
		return err
	}

	//过滤表格
	tbls := otbls.Filter(tbs, false)

	//设置包名称
	tbls.JoinPkgName(outpath)

	ntbs := data.NewTables(tbls)
	return createEnums(outpath, ntbs, cover)
}

//go:embed comm/services
var enumTmpls embed.FS
var enumTmplName = "comm/services"
var enumPrefix = "comm"

// createEnums 创建服务文件
func createEnums(root string, input interface{}, cover bool) error {
	return output.CreateFiles(enumTmpls,
		root,
		enumPrefix,
		enumTmplName,
		input,
		cover,
		data.Funcs)
}
