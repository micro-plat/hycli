{{- $table := . -}}
{{- $qrow := $table.QRows -}}
{{- $lstRow := $table.LRows -}}
{{- $leRow := $table.LERows -}}
{{- $LLERows:= mergeUIRow $lstRow $leRow}}

queryData_{{$table.UNQ}}(){
    //构建查询参数
    let that = this
     {{range $i,$c:= $qrow -}}
      {{- if eq "daterange" $c.RType.Type  -}}
   this.form_{{$table.UNQ}}.start_{{$c.Name}} = null
    this.form_{{$table.UNQ}}.end_{{$c.Name}} = null
    if(this.form_{{$table.UNQ}}.{{- $c.Name -}} && this.form_{{$table.UNQ}}.{{$c.Name}}.length>1){
      this.form_{{$table.UNQ}}.start_{{$c.Name}}=this.form_{{$table.UNQ}}.{{$c.Name}}[0]
      this.form_{{$table.UNQ}}.end_{{$c.Name}}=this.form_{{$table.UNQ}}.{{$c.Name}}[1]
    }
      {{end -}}
      {{- end -}}
    //发送查询请求
    this.conf.loading = true
    this.$theia.http.get("/{{$table.Name.CName|lower}}/query",this.form_{{$table.UNQ}}).then(res=>{
      that.conf.loading = false
      that.dataList_{{$table.UNQ}} = res.items||[]
      that.total_{{$table.UNQ}} = res.count
      that.dataList_{{$table.UNQ}}.forEach(item => {
        {{- range $i,$c := $LLERows -}}
        {{- if and (ne "" $c.RType.Format) (eq true $c.RType.IsDate)}}
         item.{{$c.Name}} = that.$theia.str.dateFormat(item.{{$c.Name}},'{{$c.RType.Format}}')
      {{- else if and (ne "" $c.RType.Format) (eq true $c.RType.IsNumber)}}
         item.{{$c.Name}} = that.$theia.str.numberFormat(item.{{$c.Name}},'{{$c.RType.Format}}')
      {{- else if eq true $c.RType.IsMobile}}
         item.{{$c.Name}} = that.$theia.str.phoneFormat(item.{{$c.Name}})
      {{- else if ne "" $c.RType.Format}}
         item.{{$c.Name}} = that.$theia.str.cut(item.{{$c.Name}},'{{$c.RType.Format}}')
      {{- end}}
      {{- end }}
      {{- range $i,$c := $LLERows -}}
     {{- if eq true $c.RType.IsSelect}}
        item.{{$c.Name}}_label = that.$theia.enum.getName("{{$c.Ext.SLType}}",item.{{$c.Name}})
      {{- end -}}
      {{- if eq "switch" $c.RType.Type}}
        item.{{$c.Name}}_switch = item.{{$c.Name}} == 0
      {{- end}}
      {{- end}}
      });
    })
  },