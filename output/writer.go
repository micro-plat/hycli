package output

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	logs "github.com/lib4dev/cli/logger"
	"github.com/micro-plat/lib4go/errs"
)

// createFiles 创建文件
func CreateFiles(tmpls embed.FS, rootDir, prefix string, name string, input interface{}, cover bool, xfuncs ...map[string]interface{}) error {
	entities, err := tmpls.ReadDir(name)
	if err != nil {
		return err
	}
	hasTmplFile := hasTmpl(name, entities)
	for _, entity := range entities {

		//处理目录
		fpath := fmt.Sprintf("%s/%s", name, entity.Name())
		if strings.Contains(fpath, ".tmpl") {
			continue
		}
		if entity.IsDir() {
			if err := CreateFiles(tmpls, rootDir, prefix, fpath, input, cover, xfuncs...); err != nil {
				return err
			}
			continue
		}

		//翻译内容{-{ }-}
		result, err := TranslateFs(tmpls, fpath, hasTmplFile, input, xfuncs...)
		if err != nil {
			return err
		}

		//构建路径{{}}
		npath := filepath.Join(".", rootDir, strings.TrimPrefix(fpath, prefix))
		npathBytes, err := TranslatePath(npath, npath, input)
		if err != nil {
			return err
		}
		//路径中包含~符合则忽略此文件
		if bytes.Contains(npathBytes, []byte("~")) {
			continue
		}

		//创建文件
		fs, err := Create(string(npathBytes), cover)
		if err != nil {
			if errs.GetCode(err) == 204 {
				logs.Log.Warn(err)
				continue
			}
			return err
		}

		//写入文件
		defer fs.Close()
		_, err = fs.Write(result)
		if err != nil {
			return err
		}
		logs.Log.Info("生成文件:", string(npathBytes))
	}
	return nil
}
func hasTmpl(name string, es []fs.DirEntry) bool {
	for _, e := range es {
		fpath := fmt.Sprintf("%s/%s", name, e.Name())
		if strings.Contains(fpath, ".tmpl.") {
			return true
		}
	}
	return false
}
