{{ $table := .}}
components: {
    {{- if gt (len (fltrColums $table "c")) 0}}
    CAdd,
    {{- end}}
    {{- if gt (len (fltrColums $table "u")) 0}}
    CEdit,
    {{- end}}
    {{- if gt (len (fltrColums $table "v")) 0}}
    CView,
    {{- end}}
    {{- if gt (len (fltrOptrs $table.ListOpts "dialog"))  0}}
    DLGOpts,
    {{- end}}
    {{- if gt (len (fltrOptrs $table.ListOpts "CNFRM")) 0}}
    DLGCnfrm,
    {{- end}}
    {{- if gt (len (fltrOptrs $table.ChartOpts "line-bar-pie")) 0}}
    ChartBase,
    {{- end}}
    {{- range $x,$m:=$table.ListOpts -}}
    {{- if eq "CMPNT" $m.Name -}}
    {{$m.UNQ}},
    {{- end}}
    {{- end}}
},