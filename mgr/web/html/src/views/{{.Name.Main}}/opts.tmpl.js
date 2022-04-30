{{- $table := .}}
{{- $optRow :=$table.Optrs}}
{{ range $x,$m:=$optRow -}}
{{- if eq "DIALOG" $m.Name -}}
show_dialog_{{$m.UNQ}}(fm){
   this.$refs.dlgOpts.show_{{$m.UNQ}}(fm)
},
{{- else if eq "CMPNT" $m.Name -}}
show_cmpnt_{{$m.UNQ}}(fm){
  let query={}
  {{- $rows:= fltrUIRows $table $m.RwName -}}
    {{range $i,$c:=$rows}} 
  query.{{$c.Name}} = fm.{{$c.Name}}
    {{- end}}
  this.$refs.cmpnt_{{$m.UNQ}}.show(query)
},
 {{- else if or (eq "CNFRM" $m.Name) (eq "DEL" $m.Name) -}}
show_confirm_{{$m.UNQ}}(fm){
  this.$refs.dlgCnfrm.show_{{$m.UNQ}}(fm)
},
{{- end -}}
{{- end}}

{{ range $x,$m:=$optRow -}}
{{- if eq "LINK" $m.Name -}}
goto_{{$m.UNQ}}(fm){
  let query={}
  {{- $rows:= fltrUIRows $table $m.RwName -}}
    {{range $i,$c:=$rows}} 
  query.{{$c.Name}} = fm.{{$c.Name}}
    {{- end}}
  this.goto('{{$m.URL}}',query)
}
{{- end -}}
{{- end}}