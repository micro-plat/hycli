{{- $xtable := . -}}
{{- $table := $xtable.Current -}}
{{- $mtable := $xtable.Main -}}
{{- $qrow := $table.QColums -}}
{{- $lstRow := $table.LColums -}}
{{- $leRow := $table.LEColums -}}
{{- $LLERows:= $table.LLEColums }}
{{- $MLLERows:= $mtable.LLEColums }}

    queryData_{{$table.UNQ}}(mform = {}){
    //构建查询参数
     {{- range $i,$c:= $qrow -}}
      {{- if eq "daterange" $c.CCMPT.Type}}
    this.form_{{$table.UNQ}}.start_{{$c.Name}} = null
    this.form_{{$table.UNQ}}.end_{{$c.Name}} = null
    if(this.form_{{$table.UNQ}}.{{- $c.Name}} && this.form_{{$table.UNQ}}.{{$c.Name}}.length > 1){
        this.form_{{$table.UNQ}}.start_{{$c.Name}} = this.form_{{$table.UNQ}}.{{$c.Name}}[0]
        this.form_{{$table.UNQ}}.end_{{$c.Name}} = this.form_{{$table.UNQ}}.{{$c.Name}}[1]
    }
      {{end -}}
      {{- end}}
 

    //处理关联表{{$table.Name}} {{$xtable.Main.Name}}  ：todo
    {{- range $i,$c := $MLLERows -}}
      {{- if eq true $c.Enum.IsEnum -}}
        {{- if eq $table.Name.Short $c.Enum.EnumType}}
        {{- if ne "" $mtable.PKColum.Name}}
    this.form_{{$table.UNQ}}.{{$c.Name}} = mform.{{$mtable.PKColum.Name}}
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
            
          {{- range $i,$c := $LLERows -}}
        {{- if eq true $c.Enum.IsEnum}}
            item.{{$c.Name}}_label = that.$theia.enum.getNames("{{$c.Enum.EnumType}}",item.{{$c.Name}})
          {{- end -}}
          {{- if eq "switch" $c.CCMPT.Type}}
            item.{{$c.Name}}_switch = item.{{$c.Name}} == 0
          {{- end}}
      {{- end}}
      {{- range $i,$c := $LLERows}}
            {{- if and (ne "" $c.LCMPT.Format) (eq true $c.Field.IsDate)}}
            item.{{$c.Name}} = that.$theia.str.dateFormat(item.{{$c.Name}},'{{$c.LCMPT.Format}}')
          {{- else if and (ne "" $c.LCMPT.Format) (eq true $c.Field.IsNumber)}}
            item.{{$c.Name}} = that.$theia.str.numberFormat(item.{{$c.Name}},'{{$c.LCMPT.Format}}')
          {{- else if eq "mobile" $c.LCMPT.Type}}
            item.{{$c.Name}} = that.$theia.str.phoneFormat(item.{{$c.Name}})
          {{- else if eq "text" $c.LCMPT.Type}}
            item.{{$c.Name}} = that.$theia.str.cut(item.{{$c.Name}},{{$c.Field.Len}})
          {{- end}}
          {{- end }}
      });
    })
  },

  {{ range $i,$c:= $qrow }}
  {{- if eq "ddl" $c.CCMPT.Type -}}
  on{{.Name}}dropdownClick(f, x) {
    let mf = f || {}
    x.{{.Name}}_label = mf.name
    x.{{.Name}} = mf.value
    this.onQuery()
  },
{{- end }}{{ end }}