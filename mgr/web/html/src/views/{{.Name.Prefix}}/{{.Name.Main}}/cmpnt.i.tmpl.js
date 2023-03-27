{{- $table := .}}
{{- $opts:=fltrMergeOptrs $table.ListOpts $table.BarOpts}}
{{- if gt (len (fltrOptrs $table.ChartOpts "line-bar-pie")) 0}}
import ChartBase from "@/views/cmpnts/chart.base.vue"
{{- end}}
{{- range $x,$m:=$opts}}
 {{- if eq "CMPNT" $m.Name}}
import {{$m.UNQ}} from "{{fltrTranslate $m.URL $table}}"
{{- end}}
{{- end}}


