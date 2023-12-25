package mgr

import (
	"github.com/lib4dev/cli/cmds"
	"github.com/micro-plat/hycli/mgr/api"
	"github.com/micro-plat/hycli/mgr/project"
	"github.com/micro-plat/hycli/mgr/web"
	"github.com/urfave/cli"
)

func init() {
	cmds.Register(
		cli.Command{
			Name:  "sql",
			Usage: "SQL语句",
			Subcommands: cli.Commands{
				{
					Name:   "create",
					Usage:  "创建表结构",
					Action: api.CreateSQLScriptByCtx,
					Flags: []cli.Flag{
						cli.StringFlag{Name: "table,t", Usage: `-表名称`},
						cli.BoolFlag{Name: "cover,v", Usage: `-文件已存在时自动覆盖`},
					},
				},
				{
					Name:   "crud",
					Usage:  "创建表结构",
					Action: api.CreateSCURDcriptByCtx,
					Flags: []cli.Flag{
						cli.StringFlag{Name: "table,t", Usage: `-表名称`},
						cli.BoolFlag{Name: "cover,v", Usage: `-文件已存在时自动覆盖`},
					},
				},
			},
		},
		cli.Command{
			Name:  "mgr",
			Usage: "创建vue前端项目",
			Subcommands: cli.Commands{
				{
					Name:   "create",
					Usage:  "创建项目",
					Action: project.CreateByCtx,
					Flags: []cli.Flag{
						cli.StringFlag{Name: "docs,c", Usage: `-创建docs目录`},
					},
				},
				{
					Name:   "code",
					Usage:  "创建web页面、后端服务代码",
					Action: CreateCodeByCtx,
					Flags: []cli.Flag{
						cli.StringFlag{Name: "table,t", Usage: `-表名称`},
						cli.BoolFlag{Name: "cover,v", Usage: `-文件已存在时自动覆盖`},
					},
				},
				{
					Name:   "web",
					Usage:  "创建web页面",
					Action: web.CreatePageByCtx,
					Flags: []cli.Flag{
						cli.StringFlag{Name: "table,t", Usage: `-表名称`},
						cli.BoolFlag{Name: "cover,v", Usage: `-文件已存在时自动覆盖`},
					},
				},
				{
					Name:   "api",
					Usage:  "创建api服务代码",
					Action: api.CreateServiceByCtx,
					Flags: []cli.Flag{
						cli.StringFlag{Name: "table,t", Usage: `-表名称`},
						cli.BoolFlag{Name: "cover,v", Usage: `-文件已存在时自动覆盖`},
					},
				},
				{
					Name:   "enums",
					Usage:  "创建enums服务代码",
					Action: api.CreateEnumsByCtx,
					Flags: []cli.Flag{
						cli.StringFlag{Name: "table,t", Usage: `-表名称`},
						cli.BoolFlag{Name: "cover,v", Usage: `-文件已存在时自动覆盖`},
					},
				},
			},
		})
}
