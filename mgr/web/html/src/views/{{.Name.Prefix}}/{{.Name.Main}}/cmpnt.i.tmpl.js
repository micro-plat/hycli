{{ $table := .}}

{{- if gt (len (fltrColums $table "c")) 0}}
import CAdd from "./{{.Name}}.add"
{{- end}}
{{- if gt (len (fltrColums $table "u")) 0}}
import CEdit from "./{{.Name}}.edit"
{{- end}}
{{- if gt (len (fltrColums $table "v")) 0}}
import CView from "./{{.Name}}.view"
{{- end}}
{{- if gt (len (fltrOptrs $table.ListOpts "dialog"))  0}}
import DLGOpts from "./{{.Name}}.dialog.vue"
{{- end}}
{{- if gt (len (fltrOptrs $table.ListOpts "confirm")) 0}}
import DLGCnfrm from "./{{.Name}}.cnfrm.vue"
{{- end}}
{{- if gt (len (fltrOptrs $table.ChartOpts "line")) 0}}
import ChartLine from "./{{.Name}}.chart_line.vue"
{{- end}}
{{- range $x,$m:=$table.ListOpts -}}
 {{- if eq "CMPNT" $m.Name  -}}
import {{$m.UNQ}} from "{{$m.URL}}"
{{- end}}
{{- end}}