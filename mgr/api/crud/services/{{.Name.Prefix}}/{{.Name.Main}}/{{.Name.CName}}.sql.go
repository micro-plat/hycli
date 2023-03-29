//go:build ignore

package {{.Name.Main}}

{{- $table := .}}
{{- $CColumns:= fltrColumnsExcludeExt (fltrColumns $table "c")}}
{{- $UColumns:= fltrColumnsExcludeExt (fltrColumns $table "u")}}
{{- $switchs := fltrCmpnt $table "switch" "l"}}
{{- $clen := (len $CColumns)|minus}}
{{- $ulen := (len $UColumns)|minus}}
{{- $slen := (len $switchs)|minus}}
{{- $sigleQueryCols:= fltrColumnsExcludeExt (fltrColumns $table "l-le")}}
{{- $totalQCols:= fltrColumnsExcludeExt (fltrColumns $table "q-bq")}}
{{- $vlen := (len $sigleQueryCols)|minus}}

{{- if gt (len (fltrColumns $table "q-bq")) 0}}

//get{{.Name.CName}}ListCount 获取{{.Desc}}列表条数
const get{{.Name.CName}}ListCount = `
select 
	count(1)
from {{.Name.Raw}} t
where
{{- range $i,$v := $totalQCols}}
{{- if eq "daterange" $v.Cmpnt.Type}}
	and t.{{$v.Name}} >=  if(@start_{{$v.Name}}='', t.{{$v.Name}},@start_{{$v.Name}})
	and t.{{$v.Name}} <  date_add(if(@end_{{$v.Name}}='',t.{{$v.Name}},@end_{{$v.Name}}), interval 1 day)
{{- else if eq "date" $v.Cmpnt.Type}}
	and t.{{$v.Name}} >=  if(@{{$v.Name}}='',t.{{$v.Name}},@{{$v.Name}})
	and t.{{$v.Name}} <  date_add(if(@{{$v.Name}}='',t.{{$v.Name}},@{{$v.Name}}),interval 1 day)	
{{- else if eq "input" $v.Cmpnt.Type}}
	?t.{{$v.Name}}
{{- else}}
	&t.{{$v.Name}}
{{- end}}
{{- end}}`

//get{{.Name.CName}}List 查询{{.Desc}}列表数据
const get{{.Name.CName}}List = `
select
	{{- range $i,$v :=  $sigleQueryCols}}
	t.{{$v.Name}},
	{{- end}}
	1
from {{.Name.Raw}} t
where
	{{- range $i,$v := $totalQCols}}
	
	{{- if eq "daterange" $v.Cmpnt.Type}}
	and t.{{$v.Name}} >=  if(@start_{{$v.Name}}='', t.{{$v.Name}},@start_{{$v.Name}})
	and t.{{$v.Name}} <  date_add(if(@end_{{$v.Name}}='',t.{{$v.Name}},@end_{{$v.Name}}), interval 1 day)
	{{- else if eq "date" $v.Cmpnt.Type}}
	and t.{{$v.Name}} >=  if(@{{$v.Name}}='',t.{{$v.Name}},@{{$v.Name}})
	and t.{{$v.Name}} <  date_add(if(@{{$v.Name}}='',t.{{$v.Name}},@{{$v.Name}}),interval 1 day)	
	{{- else if eq "input" $v.Cmpnt.Type}}
	?t.{{$v.Name}}
	{{- else}}
	&t.{{$v.Name}}
	{{- end}}
	{{- end}}
order by {{range $i,$v :=fltrColumnsExcludeExt $table.PKColumns}}
	t.{{$v.Name}} desc
{{- end}}
limit @ps offset @offset`
{{- end}}

{{- if gt (len (fltrColumns $table "c")) 0}}

// insert{{.Name.CName}} 保存{{.Desc}}数据
const insert{{.Name.CName}} = `
insert into {{.Name.Raw}}
(
	{{- range $i,$v :=fltrColumnsExcludeExt (fltrColumns $table "c")}}
	{{$v.Name}}{{if lt $i $clen}},{{end}}
	{{- end}}
)
values
(
	{{- range $i,$v := fltrColumnsExcludeExt (fltrColumns $table "c")}}
	@{{$v.Name}}{{if lt $i $clen}},{{end}}
	{{- end}}
)`
{{- end}}

{{- if gt (len (fltrColumnsExcludeExt ( fltrColumns $table "u"))) 0}}

//update{{.Name.CName}} 修改{{.Desc}}数据
const update{{.Name.CName}} = `
update {{.Name.Raw}} t set 
{{- range $i,$v :=fltrColumnsExcludeExt (fltrColumns $table "u")}}
	t.{{$v.Name}} = @{{$v.Name}}{{if lt $i $ulen}},{{end}}
{{- end}}
where 
{{- range $i,$v :=fltrColumnsExcludeExt $table.PKColumns}}
	&{{$v.Name}}
{{- end}}`
{{- end}}

{{- if gt (len (fltrColumnsExcludeExt (fltrColumns $table "l-le"))) 0}}

//get{{.Name.CName}} 查询单条{{.Desc}}数据
const get{{.Name.CName}} = `
select
{{- range $i,$v := $sigleQueryCols}}
	t.{{$v.Name}}{{if lt $i $vlen}},{{end}}
{{- end}}
from {{.Name.Raw}} t
where 
{{- range $i,$v :=fltrColumnsExcludeExt $table.PKColumns}}
	&{{$v.Name}}
{{- end}}`
{{- end}}

{{- if gt (len (fltrColumns $table "D")) 0}}

//delete{{.Name.CName}} 删除单条{{.Desc}}数据
const delete{{.Name.CName}} = `
delete from {{.Name.Raw}}
where 
{{- range $i,$v :=fltrColumnsExcludeExt $table.PKColumns}}
	&{{$v.Name}}
{{- end}}`
{{- end}}

{{- if gt (len $switchs) 0}}

//switch{{.Name.CName}} 删除单条{{.Desc}}数据
const switch{{.Name.CName}} = `
update {{.Name.Raw}} t set 
{{- range $i,$v := $switchs}}
	t.{{$v.Name}} = @{{$v.Name}}{{if lt $i $slen}},{{end}}
{{- end}}
where 
{{- range $i,$v :=fltrColumnsExcludeExt $table.PKColumns}}
	&{{$v.Name}}
{{- end}}`
{{- end}}