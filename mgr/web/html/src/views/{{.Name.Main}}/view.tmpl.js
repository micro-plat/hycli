{{- $table :=. -}}
{{- $vrow := $table.VRows -}}
{{- $viewOptRow :=$table.ViewOpts}}
 show(form) {
    this.conf.visible = true
    let that = this;
    this.$theia.http
      .get("/{{.Name.CName|lower}}",form)
      .then((res) => {
        Object.assign(that.view, res)
    {{- range $i,$c := $vrow -}}
      {{- if eq true $c.RType.IsSelect}}
        that.view.{{$c.Name}}_label = that.$theia.enum.getName("{{$c.Ext.SLType}}",res.{{$c.Name}})
     {{- end -}}
      {{- end}}
      })
      .catch((res) => {
        let code = res.response.status;
        let msg = `{{.Desc}}查询失败(${code})`;
        that.$notify.error({ title: "失败", message: msg, duration: 5000 });
      });
  },
  {{ range $x,$m:=$viewOptRow -}}
  {{- if eq "CMPNT" $m.Name -}}
  show_view_{{$m.UNQ}}(){
    let that = this;
    let query={}
    {{- $rows:= fltrUIRows $table $m.RwName -}}
    {{range $i,$c:=$rows}} 
    query.{{$c.Name}} = that.view.{{$c.Name}}
    {{- end}}
    that.$refs.view_{{$m.UNQ}}.show(query)
  },
  {{- end}}
 {{- end}}
   colorful(r){
     return this.$theia.env.conf.colorful[r]
  },