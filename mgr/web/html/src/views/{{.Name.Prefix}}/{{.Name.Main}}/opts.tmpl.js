{{- $table := . -}}
{{- $optRow:= fltrMergeOptrs $table.ListOpts $table.BarOpts}}
{{- range $x,$m:=$optRow -}}
{{- if eq "DIALOG" $m.Name}}
  show_dialog_{{$m.UNQ}}(fm = {}){
    this.$refs.dlgOpts.show_{{$m.UNQ}}(fm)
  },
{{- else if eq "CMPNT" $m.Name}}
  show_cmpnt_{{$m.UNQ}}(fm = {}){
    let query = {}
    {{- $rows:= fltrColums $table $m.RwName -}}
      {{range $i,$c:=$rows}} 
    query.{{$c.Name}} = fm.{{$c.Name}}
      {{- end}}
    {{- $pkrows:=  $table.PKColums -}}
      {{range $i,$c:=$pkrows}} 
    query.{{$c.Name}} = fm.{{$c.Name}}
      {{- end}}
    {{- if eq "DEL" $m.Tag }}
    this.$refs.cmpnt_{{$m.UNQ}}.show_{{$m.UNQ}}(fm)
    {{else}}
    this.$refs.cmpnt_{{$m.UNQ}}.show(query)
    {{- end}}
  },
{{- end -}}
{{- end}}
{{ range $x,$m:=$optRow -}}
{{- if eq "LINK" $m.Name -}}
  goto_{{$m.UNQ}}(fm = {}){
    let query = {}
    {{- $rows:= fltrColums $table $m.RwName -}}
      {{range $i,$c:=$rows}} 
    query.{{$c.Name}} = fm.{{$c.Name}}
      {{- end}}
    this.goto('{{$m.URL}}',query)
  }
{{- end -}}
{{- end -}}


