package ui

import (
	"github.com/lib4dev/cli/cmds"
	"github.com/micro-plat/hycli/ui/page"
	"github.com/micro-plat/hycli/ui/project"
	"github.com/urfave/cli"
)

func init() {
	cmds.Register(
		cli.Command{
			Name:  "ui",
			Usage: "创建vue前端项目",
			Subcommands: cli.Commands{
				{
					Name:   "create",
					Usage:  "创建项目",
					Action: project.Create,
				},
				{
					Name:   "page",
					Usage:  "创建web页面",
					Action: page.Create,
					Flags: []cli.Flag{
						cli.StringFlag{Name: "table,t", Usage: `-表名称`},
						cli.StringFlag{Name: "kw,k", Usage: `-约束字段`},
						cli.BoolFlag{Name: "w2f,f", Usage: `-生成到文件`},
						cli.BoolFlag{Name: "cover,v", Usage: `-文件已存在时自动覆盖`},
					},
				},
			},
		})

}
