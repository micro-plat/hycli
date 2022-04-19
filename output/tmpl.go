package output

import (
	"bytes"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/micro-plat/lib4go/types"
)

//Translate 翻译模板
func TranslateFs(fs fs.FS, indexTmplName string, input interface{}, xfuncs ...map[string]interface{}) (t string, err error) {
	nfuncs := types.NewXMap()
	nfuncs.MergeMap(funcs)
	if len(xfuncs) > 0 {
		nfuncs.MergeMap(xfuncs[0])
	}
	var tmpl = template.New(indexTmplName).Funcs(nfuncs.ToMap())
	partten := strings.Replace(filepath.Join(filepath.Dir(indexTmplName), "*.tmpl.*"), "\\", "/", -1)
	tmpl, err = tmpl.ParseFS(fs, indexTmplName, partten)
	if err != nil {
		return "", fmt.Errorf("转换模板失败:%s %s %w", indexTmplName, partten, err)
	}
	buff := bytes.NewBufferString("")
	err = tmpl.ExecuteTemplate(buff, filepath.Base(indexTmplName), input)
	if err != nil {
		return "", fmt.Errorf("执行模板失败:%s\nerr: %w", indexTmplName, err)
	}
	return buff.String(), nil
}

//Translate 翻译模板
func TranslateContent(name string, content string, input interface{}, xfuncs ...map[string]interface{}) (string, error) {
	nfuncs := types.NewXMap()
	nfuncs.MergeMap(funcs)
	if len(xfuncs) > 0 {
		nfuncs.MergeMap(xfuncs[0])
	}
	var tmpl = template.New(name).Funcs(nfuncs.ToMap())
	np, err := tmpl.Parse(content)
	if err != nil {
		return "", err
	}
	buff := bytes.NewBufferString("")
	if err := np.Execute(buff, input); err != nil {
		return "", err
	}
	return buff.String(), nil
}
