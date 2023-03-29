{{- $table := .}}
{{- $opts:=mergeOptrs $table.ListOpts $table.BarOpts $table.ChartOpts}}
components: {
    {{- range $x,$m:=  $opts}}
    {{- if eq "CMPNT" $m.Name}}
    {{$m.UNQ}},
    {{- end}}
    {{- end}}
},