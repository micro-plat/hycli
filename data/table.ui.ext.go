package data

import (
	"github.com/micro-plat/hycli/md"
	"github.com/micro-plat/lib4go/types"
)

type Ext struct {
	SLType    string
	ColorType string
	Format    string
	DefValue  string
	Valid     string
	FormName  string
}

func newExt(t string, r *md.Row, formName string) *Ext {
	ext := &Ext{FormName: types.GetString(formName, "form")}
	//处理sl参数
	if md.HasConstraint(r.Constraints, "sl") {
		ext.SLType = types.GetString(md.GetTPName("sl", r.Constraints...), r.Name)
	}

	//处理color参数
	if md.HasConstraint(r.Constraints, "color") {
		ext.ColorType = "colorful"
	}

	//处理输出格式
	if md.HasConstraint(r.Constraints, "f") {
		ext.Format = md.GetFormat(r.Constraints...)
	}
	//默认值
	ext.DefValue = md.GetDefValue(r.Constraints...)

	//验证方式
	ext.Valid = md.GetVdlValue(r.Constraints...)
	return ext

}
