package {{.Name.Main}}

import (
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

{{- $table := .|fltrCodeTable -}}
{{ $clen := (len $table.CRows)|minus}}
{{ $ulen := (len $table.URows)|minus}}
//AuditInfoHandler 获取{{.Desc}}处理服务
type {{.Name.CName}}Handler struct {
	insertRequiredFields []string
	updateRequiredFields []string
}

func New{{.Name.CName}}Handler() *{{.Name.CName}}Handler {
	return &{{.Name.CName}}Handler{
		insertRequiredFields:[]string{ {{- range $i,$v :=  $table.CRows -}}
			"{{$v.Name}}"{{if lt $i $clen }},{{end}}{{- end -}}},
		updateRequiredFields:[]string{ {{- range $i,$v :=  $table.URows -}}
			"{{$v.Name}}"{{if lt $i $ulen }},{{end}}{{- end -}}},
	}
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
