//go:build db

package mysql

import (
	"embed"
	"fmt"
	"io/ioutil"

	"github.com/micro-plat/hydra"
)

func init() {
	hydra.OnReady(func() error {
		sqlList, err := readFiles(dirName)
		if err != nil {
			return err
		}
		hydra.Installer.DB.AddSQL(
			sqlList...,
		)
		return nil
	})
}

func readFiles(name string) ([]string, error) {
	entities, err := fs.ReadDir(name)
	if err != nil {
		return nil, err
	}
	content := make([]string, 0, 1)
	for _, entity := range entities {
		fpath := fmt.Sprintf("%s/%s", name, entity.Name())
		f, err := fs.Open(fpath)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		buff, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}
		content = append(content, string(buff))
	}
	return content, nil
}

//go:embed sqls
var fs embed.FS
var dirName = "sqls"
