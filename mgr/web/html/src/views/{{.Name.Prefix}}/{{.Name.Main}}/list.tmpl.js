{{- $xtable := . -}}
{{- $table := $xtable.Current -}}
{{- $mtable := $xtable.Main -}}
{{- $qrow := $table.QRows -}}
{{- $lstRow := $table.LRows -}}
{{- $leRow := $table.LERows -}}
{{- $LLERows:= mergeUIRow $lstRow $leRow}}
{{- $MLLERows:= mergeUIRow $mtable.LRows $mtable.LERows}}

    queryData_{{$table.UNQ}}(mform = {}){
    //构建查询参数
     {{- range $i,$c:= $qrow -}}
      {{- if eq "daterange" $c.RType.Type}}
    this.form_{{$table.UNQ}}.start_{{$c.Name}} = null
    this.form_{{$table.UNQ}}.end_{{$c.Name}} = null
    if(this.form_{{$table.UNQ}}.{{- $c.Name}} && this.form_{{$table.UNQ}}.{{$c.Name}}.length > 1){
        this.form_{{$table.UNQ}}.start_{{$c.Name}} = this.form_{{$table.UNQ}}.{{$c.Name}}[0]
        this.form_{{$table.UNQ}}.end_{{$c.Name}} = this.form_{{$table.UNQ}}.{{$c.Name}}[1]
    }
      {{end -}}
      {{- end}}
 
    //处理关联表{{$table.Name}} {{$xtable.Main.Name}}
    {{- range $i,$c := $MLLERows -}}
      {{- if eq true $c.RType.IsSelect -}}
        {{- if eq $table.Name.Short $c.Ext.SLType}}
        {{- if ne "" $c.RType.FKName}}
    this.form_{{$table.UNQ}}.{{$c.Name}} = mform.{{$c.RType.FKName}}
        {{- end -}}
        {{- end -}}
    {{- end -}}
    {{- end}}

    {{- range $x,$v := $xtable.Main.ViewOpts -}}
          {{if eq $v.URL $table.Name.Raw}}
          {{if and (ne "" $v.RwName) (ne "" $v.FwName)}}
     this.form_{{$table.UNQ}}.{{$v.FwName}} = mform.{{$v.RwName}}   
        {{end}}
          {{- end -}}
    {{end}}
     
    //发送查询请求
    let that = this
    that.conf.loading = true
    this.$theia.http.get("/{{$table.Name.MainPath|lower}}/query",this.form_{{$table.UNQ}}).then(res=>{
        that.conf.loading = false
        that.dataList_{{$table.UNQ}} = res.items||[]
        that.total_{{$table.UNQ}} = res.count
        that.dataList_{{$table.UNQ}}.forEach(item => {
            {{- range $i,$c := $LLERows}}
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
            item.{{$c.Name}}_label = that.$theia.enum.getNames("{{$c.Ext.SLType}}",item.{{$c.Name}})
          {{- end -}}
          {{- if eq "switch" $c.RType.Type}}
            item.{{$c.Name}}_switch = item.{{$c.Name}} == 0
          {{- end}}
      {{- end}}
      });
    })
  },

  {{ range $i,$c:= $qrow }}
  {{- if eq "ddl" $c.RType.Type -}}
  on{{.Name}}dropdownClick(f, x) {
    let mf = f || {}
    x.{{.Name}}_label = mf.name
    x.{{.Name}} = mf.value
    this.onQuery()
  },
{{- end }}{{ end }}