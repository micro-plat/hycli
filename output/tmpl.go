package output

import (
	"bytes"
	"html/template"

	"github.com/micro-plat/lib4go/types"
)

//Translate 翻译模板
func Translate(content string, input interface{}, xfuncs ...map[string]interface{}) (string, error) {
	nfuncs := types.NewXMap()
	nfuncs.MergeMap(funcs)
	if len(xfuncs) > 0 {
		nfuncs.MergeMap(xfuncs[0])
	}
	var tmpl = template.New("table").Funcs(nfuncs.ToMap())
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
