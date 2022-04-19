package output

import (
	"embed"
	"fmt"
	"path/filepath"
	"strings"

	logs "github.com/lib4dev/cli/logger"
)

//createFiles 创建文件
func CreateFiles(tmpls embed.FS, rootDir, prefix string, name string, input interface{}, xfuncs ...map[string]interface{}) error {
	entities, err := tmpls.ReadDir(name)
	if err != nil {
		return err
	}
	for _, entity := range entities {

		//处理目录
		fpath := fmt.Sprintf("%s/%s", name, entity.Name())
		if strings.Contains(fpath, ".tmpl") {
			continue
		}
		if entity.IsDir() {
			if err := CreateFiles(tmpls, rootDir, prefix, fpath, input, xfuncs...); err != nil {
				return err
			}
			continue
		}

		//翻译内容
		result, err := TranslateFs(tmpls, fpath, input, xfuncs...)
		if err != nil {
			return err
		}

		//构建路径
		npath := filepath.Join(".", rootDir, strings.TrimPrefix(fpath, prefix))
		npath, err = TranslateContent(npath, npath, input)
		if err != nil {
			return err
		}

		//创建文件
		fs, err := Create(npath, true)
		if err != nil {
			return err
		}

		//写入文件
		defer fs.Close()
		fs.WriteString(result)
		logs.Log.Info("生成文件:", npath)
	}
	return nil
}
