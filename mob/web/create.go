package web

import (
	"embed"
	"fmt"

	"github.com/micro-plat/hycli/data"
	"github.com/micro-plat/hycli/output"
	"github.com/micro-plat/lib4go/types"
	"github.com/urfave/cli"
)

// CreatePageByCtx 创建页面文件
func CreatePageByCtx(c *cli.Context) (err error) {
	if len(c.Args()) == 0 {
		return fmt.Errorf("未指定md文件路径")
	}
	//获取输出路径
	outpath := types.GetString(c.Args().Get(1), "./")
	cover := c.Bool("cover")
	return CreatePage(c.Args().First(), c.String("table"), outpath, cover)
}
func CreatePage(mdpath string, tbs string, outpath string, cover bool) error {
	//转换md文件
	otbls, err := data.Mds2Tables(mdpath)
	if err != nil {
		return err
	}
	//过滤表格
	tbls := otbls.Filter(tbs, true)

	ntbs := data.NewTables(tbls)
	data.Caches(ntbs)
	// data.ResetSelectRelation(ntbs)

	//创建页面文件
	for _, tb := range ntbs {
		if err = createViews(outpath, tb, cover); err != nil {
			return err
		}
	}

	// //创建路由文件

	if err = createRouter(outpath, tbls, cover); err != nil {
		return err
	}
	return nil
}

//go:embed html/src/views
var viewTmpls embed.FS
var viewTmplName = "html/src/views"
var prefix = "html"

// create 创建页面文件
func createViews(root string, input interface{}, cover bool) error {
	return output.CreateFiles(viewTmpls,
		root,
		prefix,
		viewTmplName,
		input,
		cover,
		data.Funcs)
}
