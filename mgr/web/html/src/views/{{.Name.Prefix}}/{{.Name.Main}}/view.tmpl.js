{{- $table :=. -}}
{{- $pkRows:=$table.PKRows}}
{{- $vrow := $table.VRows -}}
{{- $viewOptRow :=$table.ViewOpts}}
 show(form) {
    this.conf.visible = true
    {{- range $i,$c:= $viewOptRow}}
      {{- if eq "TAB" $c.Name}}
      //{{ $c.Label}}查询
      {{- $ct:= fltrSearchUITable $c}}
      // {{if and (ne "" $c.RwName) (ne "" $c.FwName)}}
        this.form_{{$ct.UNQ}}.{{ $c.FwName}} = form.{{$c.RwName}}
      {{- else -}}
        {{- range $y,$x:= $pkRows}}
          {{- $currRow:= fltrForginKey $ct $table $x.Name -}}
          {{- if ne "" $currRow.UNQ}}
      // this.form_{{$ct.UNQ}}.{{$currRow.Name}} = form.{{$x.Name}}
          {{- end }}
        {{- end}}
      {{- end}}
      this.queryData_{{$ct.UNQ}}(form)
      {{- end}}
     {{- end}}
  
   let that = this;
    this.$theia.http
      .get("/{{.Name.MainPath|lower}}",form)
      .then((res) => {
            Object.assign(that.view, res)
        {{- range $i,$c := $vrow -}}
          {{- if eq true $c.RType.IsSelect}}
            that.view.{{$c.Name}}_label = that.$theia.enum.getNames("{{$c.Ext.SLType}}",res.{{$c.Name}})
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
  tagColor(r,name){
    if(this.$theia.env.conf.colorTag[name]){
          return this.$theia.env.conf.colorTag[name][r]||""
      }
     return this.$theia.env.conf.colorTag[r]||""
  },