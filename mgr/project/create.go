package project

import (
	"github.com/urfave/cli"
)

// CreateAPI 创建web项目
func Create(creatDocs bool, cover bool) error {

	//创建api项目

	if err := CreateAPI("", NewProject(""), cover); err != nil {
		return err
	}

	//创建web项目
	if err := CreateWeb("web", NewProject("web"), cover); err != nil {
		return err
	}

	//创建docs项目
	if creatDocs {
		if err := CreateDocs("0docs", NewProject("0docs"), cover); err != nil {
			return err
		}
	}

	return nil
}
func CreateByCtx(c *cli.Context) (err error) {
	return Create(c.Bool("docs"), c.Bool("cover"))
}
