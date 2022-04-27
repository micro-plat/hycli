package output

import (
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/micro-plat/lib4go/types"
)

//Translate 翻译模板
func TranslateFs(fs fs.FS, indexTmplName string, hasParseFs bool, input interface{}, xfuncs ...map[string]interface{}) (t []byte, err error) {
	if isImage(indexTmplName) {
		f, err := fs.Open(indexTmplName)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		buff, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}
		return buff, nil
	}

	nfuncs := types.NewXMap()
	nfuncs.MergeMap(funcs)
	if len(xfuncs) > 0 {
		nfuncs.MergeMap(xfuncs[0])
	}
	var tmpl = template.New(indexTmplName).Funcs(nfuncs.ToMap())
	if hasParseFs {
		partten := strings.Replace(filepath.Join(filepath.Dir(indexTmplName), "*.tmpl.*"),
			"\\", "/", -1)
		tmpl, err = tmpl.ParseFS(fs, indexTmplName, partten)
		if err != nil {
			return nil, fmt.Errorf("转换模板失败:%s %s %w", indexTmplName, partten, err)
		}
	} else {
		tmpl, err = tmpl.ParseFS(fs, indexTmplName)
		if err != nil {
			return nil, fmt.Errorf("转换模板失败:%s %w", indexTmplName, err)
		}
	}

	buff := bytes.NewBufferString("")
	err = tmpl.ExecuteTemplate(buff, filepath.Base(indexTmplName), input)
	if err != nil {
		return nil, fmt.Errorf("执行模板失败:%s\nerr: %w", indexTmplName, err)
	}
	return trimIgnoreTag(buff.Bytes()), nil

}

//Translate 翻译模板
func TranslateContent(name string, content string, input interface{}, xfuncs ...map[string]interface{}) ([]byte, error) {

	nfuncs := types.NewXMap()
	nfuncs.MergeMap(funcs)
	if len(xfuncs) > 0 {
		nfuncs.MergeMap(xfuncs[0])
	}
	var tmpl = template.New(name).Funcs(nfuncs.ToMap())
	np, err := tmpl.Parse(content)
	if err != nil {
		return nil, err
	}
	buff := bytes.NewBufferString("")
	if err := np.Execute(buff, input); err != nil {
		return nil, err
	}
	return trimIgnoreTag(buff.Bytes()), nil
}
func isImage(p string) bool {
	np := strings.ToLower(p)
	return strings.HasSuffix(np, ".jpg") ||
		strings.HasSuffix(np, ".png") ||
		strings.HasSuffix(np, ".gif")
}
func trimIgnoreTag(buff []byte) []byte {
	prefix := []byte(`//go:build ignore`)
	if bytes.HasPrefix(buff, prefix) {
		return buff[len(prefix):]
	}
	return buff
}
