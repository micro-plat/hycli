package project

import (
	"github.com/urfave/cli"
)

//CreateAPI 创建web项目
func Create() error {

	//创建api项目

	if err := CreateAPI("", NewProject("")); err != nil {
		return err
	}

	//创建web项目
	if err := CreateWeb("web", NewProject("web")); err != nil {
		return err
	}
	return nil
}
func CreateByCtx(c *cli.Context) (err error) {
	return Create()
}
