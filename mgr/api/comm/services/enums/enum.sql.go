//go:build ignore

package enums

{{ $etable := .}}
var enumSQL = map[string]string{
	{{- range $i,$v:=$etable -}}
	{{- if and (ne "" $v.Enum.EnumType) (eq false $v.Enum.Multiple) }}
	"{{$v.Enum.EnumType}}":"select {{$v.Enum.Id}} value,{{if ne "" $v.Enum.PID}} {{$v.Enum.PID}} pid, {{end}}{{$v.Enum.Name}} name,{{- range $j,$v:=$v.Enum.DEColums -}}{{$v.Name}} {{$v.Desc}},{{end}}'{{$v.Enum.Type}}' type from {{$v.Name.Raw}} where 1=1  {{if ne "" $v.Enum.Status}} and {{$v.Enum.Status}} = 0{{end}} {{if ne "" $v.Enum.Expire}} and {{$v.Enum.Expire}} >= DATE_FORMAT(now(),'%Y-%m-%d'){{end}}",
	{{- end -}}
	{{- end}}
}

var unspecifiedEnum = []string{
	{{- range $i,$v:=$etable -}}
	{{- if and (ne "" $v.Enum.EnumType) (eq true $v.Enum.Multiple) }}
	"select {{$v.Enum.Id}} value,{{$v.Enum.Name}} name,{{- range $j,$v:=$v.Enum.DEColums -}}{{$v.Name}} {{$v.Desc}},{{end}}{{$v.Enum.Type}} type from {{$v.Name.Raw}} where {{$v.Enum.Type}}=if(@type='',{{$v.Enum.Type}},@type) {{if ne "" $v.Enum.Status}} and {{$v.Enum.Status}} = 0{{end}}  {{if ne "" $v.Enum.Expire}} and {{$v.Enum.Expire}} >= DATE_FORMAT(now(),'%Y-%m-%d'){{end}}",
	{{- end -}}
	{{- end}}
}
var sqls = make([]string, 0, 1)

func getSQLs(tp string) []string {
	//返回所有枚举
	if tp == "" || tp == "*" {
		return sqls
	}

	//根据枚举查询具体提供SQL
	vsqls, ok := enumSQL[tp]
	if ok {
		return []string{vsqls}
	}

	//指定未明确类型的枚举
	return unspecifiedEnum
}
func init() {
	for _, v := range enumSQL {
		sqls = append(sqls, v)
	}
	sqls = append(sqls, unspecifiedEnum...)
}
