{{$table :=.}}
{{ range $i,$c:= $table.LStatOpts }}
{{- $ucols := fltrColums $table $c.RwName -}}
stat_{{$c.UNQ}}:{
    {{range $j,$m:= $ucols -}}
    {{$m.Name}}:0,
    {{- end}}
},
{{end}}