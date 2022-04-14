package {{.Name.Main}}

{{- $table := .|fltrCodeTable -}}
{{ $clen := (len $table.CRows)|minus}}
{{ $ulen := (len $table.URows)|minus}}

//get{{.Name.CName}}ListCount 获取{{.Desc}}列表条数
const get{{.Name.CName}}ListCount = `
select 
	count(1)
from {{.Name.Raw}} t
where
{{- range $i,$v := $table.QRows}}
{{- if eq "daterange" $v.RType.Type}}
	and t.{{$v.Name}} >=  if(@start_{{$v.Name}}='', t.{{$v.Name}},@start_{{$v.Name}})
	and t.{{$v.Name}} <  date_add(if(@end_{{$v.Name}}='',t.{{$v.Name}},@end_{{$v.Name}}), interval 1 day)
{{- else if eq "date" $v.RType.Type}}
	and t.{{$v.Name}} >=  if(@{{$v.Name}}='',t.{{$v.Name}},@{{$v.Name}})
	and t.{{$v.Name}} <  date_add(if(@{{$v.Name}}='',t.{{$v.Name}},@{{$v.Name}}),interval 1 day)	
{{- else if eq "input" $v.RType.Type}}
	?t.{{$v.Name}}
{{- else}}
	&t.{{$v.Name}}
{{- end -}}
{{end}}`


	
//get{{.Name.CName}}List 查询{{.Desc}}列表数据
const get{{.Name.CName}}List = `
select
	{{range $i,$v :=  $table.LRows -}}
	t.{{$v.Name}},
	{{end -}}
	1
from {{.Name.Raw}} t
where
	{{- range $i,$v := $table.QRows -}}
	{{- if eq "daterange" $v.RType.Type}}
	and t.{{$v.Name}} >=  if(@start_{{$v.Name}}='', t.{{$v.Name}},@start_{{$v.Name}})
	and t.{{$v.Name}} <  date_add(if(@end_{{$v.Name}}='',t.{{$v.Name}},@end_{{$v.Name}}), interval 1 day)
	{{- else if eq "date" $v.RType.Type}}
	and t.{{$v.Name}} >=  if(@{{$v.Name}}='',t.{{$v.Name}},@{{$v.Name}})
	and t.{{$v.Name}} <  date_add(if(@{{$v.Name}}='',t.{{$v.Name}},@{{$v.Name}}),interval 1 day)	
	{{- else if eq "input" $v.RType.Type}}
	?t.{{$v.Name}}
	{{- else}}
	&t.{{$v.Name}}
	{{- end -}}
	{{- end}}
order by {{ range $i,$v := $table.PKRows}}
	t.{{$v.Name}} desc
{{end -}}
limit @ps offset @offset`

// insert{{.Name.CName}} 保存{{.Desc}}数据
const insert{{.Name.CName}} = `
insert into {{.Name.Raw}}
(
	{{range $i,$v := $table.CRows -}}
	{{$v.Name}}{{if lt $i $clen }},{{end}}
	{{end -}}
)
values
(
	{{range $i,$v :=  $table.CRows -}}
	@{{$v.Name}}{{if lt $i $clen }},{{end}}
	{{end -}}
)`

//update{{.Name.CName}} 修改{{.Desc}}数据
const update{{.Name.CName}} = `
update {{.Name.Raw}} t set 
{{- range $i,$v := $table.URows}}
	t.{{$v.Name}} = @{{$v.Name}}{{if lt $i $ulen }},{{end}}
{{- end}}
where 
{{- range $i,$v := $table.PKRows}}
	&{{$v.Name}}
{{- end}}`

