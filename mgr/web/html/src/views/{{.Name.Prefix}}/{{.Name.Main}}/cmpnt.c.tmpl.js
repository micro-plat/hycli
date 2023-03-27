{{ $table := .}}
{{ $opts:=fltrMergeOptrs $table.ListOpts $table.BarOpts}}
components: {
    {{- if gt (len (fltrOptrs $table.ChartOpts "line-bar-pie")) 0}}
    ChartBase,
    {{- end}}
    {{- range $x,$m:= $opts}}
    {{- if eq "CMPNT" $m.Name}}
    {{$m.UNQ}},
    {{- end}}
    {{- end}}
},