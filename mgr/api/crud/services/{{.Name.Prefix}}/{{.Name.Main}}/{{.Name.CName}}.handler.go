//go:build ignore

package {{.Name.Main}}

import (
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)
{{- $table := . -}}

{{- $ccols := fltrNotNullCols (fltrColums $table "c") -}}
{{- $clen := (len $ccols)|minus}}
{{- $ucols := fltrNotNullCols  (fltrColums $table "u") -}}
{{- $ulen := (len $ucols)|minus -}}
{{ $pklen := (len $table.PKColums)|minus}}
//AuditInfoHandler 获取{{.Desc}}处理服务
type {{.Name.CName}}Handler struct {
	insertRequiredFields []string
	updateRequiredFields []string
	pkRequiredFields []string
}

func New{{.Name.CName}}Handler() *{{.Name.CName}}Handler {
	return &{{.Name.CName}}Handler{
		insertRequiredFields:[]string{ {{- range $i,$v := $ccols -}}
			"{{$v.Name}}"{{if lt $i $clen }},{{end}}{{- end -}}},
		updateRequiredFields:[]string{ {{- range $i,$v :=  $ucols -}}
			"{{$v.Name}}"{{if lt $i $ulen }},{{end}}{{- end -}}},
		pkRequiredFields:[]string{ {{- range $i,$v :=  $table.PKColums -}}
			"{{$v.Name}}"{{if lt $i $pklen }},{{end}}{{- end -}}},
	}
}

//DeleteHandle  删除{{.Desc}}数据
func (u *{{.Name.CName}}Handler) DeleteHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------删除{{.Desc}}数据--------")

	ctx.Log().Info("1.检查必须字段")
	if err := ctx.Request().Check(u.pkRequiredFields...); err != nil {
		return err
	}
	
	ctx.Log().Info("2.查询数据")
	i, err := hydra.C.DB().GetRegularDB().Execute(delete{{.Name.CName}}, ctx.Request().GetMap())
	if err != nil||i <= 0 {
		return errs.NewErrorf(http.StatusNotExtended, "删除数据出错:%+v", err)
	}
	return
}

//GetHandle  修改{{.Desc}}数据
func (u *{{.Name.CName}}Handler) GetHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------查询{{.Desc}}数据--------")

	ctx.Log().Info("1.检查必须字段")
	if err := ctx.Request().Check(u.pkRequiredFields...); err != nil {
		return err
	}
	
	ctx.Log().Info("2.查询数据")
	items, err := hydra.C.DB().GetRegularDB().Query(get{{.Name.CName}}, ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "保存数据出错:%+v", err)
	}
	return items.Get(0)
}

//PutHandle  修改{{.Desc}}数据
func (u *{{.Name.CName}}Handler) PutHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------保存{{.Desc}}数据--------")

	ctx.Log().Info("1.检查必须字段")
	if err := ctx.Request().Check(u.updateRequiredFields...); err != nil {
		return err
	}
	
	ctx.Log().Info("2.修改数据")
	rx, err := hydra.C.DB().GetRegularDB().Execute(update{{.Name.CName}}, ctx.Request().GetMap())
	if err != nil || rx == 0 {
		return errs.NewErrorf(http.StatusNotExtended, "保存数据出错:%+v,row:%d", err,rx)
	}
	return
}

//PostHandle  保存{{.Desc}}数据
func (u *{{.Name.CName}}Handler) PostHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------保存{{.Desc}}数据--------")

	ctx.Log().Info("1.检查必须字段")
	if err := ctx.Request().Check(u.insertRequiredFields...); err != nil {
		return err
	}
	
	ctx.Log().Info("2.添加新数据")
	rx, err := hydra.C.DB().GetRegularDB().Execute(insert{{.Name.CName}}, ctx.Request().GetMap())
	if err != nil || rx == 0 {
		return errs.NewErrorf(http.StatusNotExtended, "数据出错:%+v,row:%d", err,rx)
	}
	return
}

//QueryHandle  获取{{.Desc}}列表数据
func (u *{{.Name.CName}}Handler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取{{.Desc}}数据列表--------")

	ctx.Log().Info("1.查询数据条数")
	m := ctx.Request().GetMap()
	m["ps"] = ctx.Request().GetInt("ps", 10)
	m["offset"] = (ctx.Request().GetInt("pi", 1) - 1) * ctx.Request().GetInt("ps", 10)

	
	count, err := hydra.C.DB().GetRegularDB().Scalar(get{{.Name.CName}}ListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}

	rcount:=types.GetInt(count)
	if  rcount== 0 {
		return map[string]interface{}{
			"count":0,
		}
	}

	ctx.Log().Info("2.查询数据列表")
	items, err := hydra.C.DB().GetRegularDB().Query(get{{.Name.CName}}List, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据出错:%+v", err)
	}


	ctx.Log().Info("3.返回结果")
	return map[string]interface{}{
		"items": items,
		"count":rcount,
	}
}
