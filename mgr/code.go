package mgr

import (
	"fmt"

	"github.com/micro-plat/hycli/mgr/api"
	"github.com/micro-plat/hycli/mgr/web"
	"github.com/micro-plat/lib4go/types"
	"github.com/urfave/cli"
)

// CreateCodeByCtx 创建页面文件
func CreateCodeByCtx(c *cli.Context) (err error) {
	if len(c.Args()) == 0 {
		return fmt.Errorf("未指定md文件路径")
	}
	//获取输出路径
	apiOutpath := types.GetString(c.Args().Get(1), "./api")

	webOutput := types.GetString(c.Args().Get(1), "./web")

	cover := c.Bool("cover")

	//创建页面代码
	err = web.CreatePage(c.Args().First(), c.String("table"), webOutput, cover)
	if err != nil {
		return err
	}

	//创建服务代码
	err = api.CreateService(c.Args().First(), c.String("table"), apiOutpath, cover)
	if err != nil {
		return err
	}

	//创建枚举代码
	return api.CreateEnums(c.Args().First(), c.String("table"), apiOutpath, cover)
}
