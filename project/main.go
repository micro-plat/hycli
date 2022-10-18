package project

import (
	"github.com/lib4dev/cli/cmds"
	"github.com/micro-plat/hycli/project/app"
	"github.com/urfave/cli"
)

func init() {
	cmds.Register(
		cli.Command{
			Name:   "create",
			Usage:  "创建项目",
			Action: app.CreateByCtx,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "platName,p", Usage: `-平台名称`},
				cli.StringFlag{Name: "sysName,s", Usage: `-系统名称`},
			},
		},
	)
}
