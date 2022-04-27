package project

import (
	"embed"
	"fmt"

	"github.com/micro-plat/hycli/output"
	"github.com/micro-plat/lib4go/types"
	"github.com/urfave/cli"
)

//go:embed mgr
var apiTmpls embed.FS
var apiTmplName = "mgr"

//CreateAPI 创建web项目
func CreateAPI(name string, input interface{}) error {
	err := output.CreateFiles(apiTmpls,
		name,
		apiTmplName,
		apiTmplName,
		input)
	if err != nil {
		return err
	}

	// //构建go.mod文件
	// session := sh.InteractiveSession()
	// session.SetDir(filepath.Join("./", name))
	// session.Command("go", "mod", "init")
	// _, err = session.Output()
	return err
}

//CreateAPIByCtx 创建API项目基础文件
func CreateAPIByCtx(c *cli.Context) (err error) {
	if len(c.Args()) == 0 {
		return fmt.Errorf("未指定项目名称")
	}
	name := types.GetString(c.Args().First(), "web")
	return CreateAPI(name, NewProject(name))
}
