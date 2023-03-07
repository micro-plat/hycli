//go:build ignore

package enums

{{ $etable := .|fltrCodeTables}}
var enumSQL = map[string]string{
	{{- range $i,$v:=$etable -}}
	{{- if and ($v.AsEnum) (eq false $v.EType.Multiple) }}
	"{{$v.EType.Type}}":"select {{$v.EType.Id}} value,{{if ne "" $v.EType.PID}} {{$v.EType.PID}} pid, {{end}}{{$v.EType.Name}} name,{{- range $j,$v:=$v.EType.DERows -}}{{$v.Name}} {{$v.Desc}},{{end}}'{{$v.EType.Type}}' type from {{$v.Name.Raw}} where 1=1  {{if ne "" $v.EType.Status}} and {{$v.EType.Status}} = 0{{end}} {{if ne "" $v.EType.Expire}} and {{$v.EType.Expire}} >= DATE_FORMAT(now(),'%Y-%m-%d'){{end}}",
	{{- end -}}
	{{- end}}
}

var unspecifiedEnum = []string{
	{{- range $i,$v:=$etable -}}
	{{- if and ($v.AsEnum) (eq true $v.EType.Multiple) }}
	"select {{$v.EType.Id}} value,{{$v.EType.Name}} name,{{- range $j,$v:=$v.EType.DERows -}}{{$v.Name}} {{$v.Desc}},{{end}}{{$v.EType.Type}} type from {{$v.Name.Raw}} where {{$v.EType.Type}}=if(@type='',{{$v.EType.Type}},@type) {{if ne "" $v.EType.Status}} and {{$v.EType.Status}} = 0{{end}}  {{if ne "" $v.EType.Expire}} and {{$v.EType.Expire}} >= DATE_FORMAT(now(),'%Y-%m-%d'){{end}}",
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
