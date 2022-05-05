{{- $table :=. -}}
{{- $vrow := $table.VRows -}}
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
   colorful(r){
     return this.$theia.env.conf.colorful[r]
  },