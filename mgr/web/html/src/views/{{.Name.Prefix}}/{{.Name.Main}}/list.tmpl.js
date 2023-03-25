{{- $xtable := . -}}
{{- $table := $xtable.Current -}}
{{- $mtable := $xtable.Main -}}
{{- $qrow := fltrColums $table "q" -}}
{{- $lstRow := fltrColums $table "l" -}}
{{- $leRow := fltrColums $table "le" -}}
{{- $LLERows:= fltrColums $table "l-le" }}
{{- $MLLERows:= fltrColums $mtable "l-le" }}

    queryData_{{$table.UNQ}}(mform = {}){
    //构建查询参数
     {{- range $i,$c:= $qrow -}}
      {{- if eq "daterange" $c.Cmpnt.Type -}}
    this.form_{{$table.UNQ}}.start_{{$c.Name}} = null
    this.form_{{$table.UNQ}}.end_{{$c.Name}} = null
    if(this.form_{{$table.UNQ}}.{{- $c.Name}} && this.form_{{$table.UNQ}}.{{$c.Name}}.length > 1){
        this.form_{{$table.UNQ}}.start_{{$c.Name}} = this.form_{{$table.UNQ}}.{{$c.Name}}[0]
        this.form_{{$table.UNQ}}.end_{{$c.Name}} = this.form_{{$table.UNQ}}.{{$c.Name}}[1]
    }
      {{- end -}}
      {{- end}}
 
    //处理关联表{{$table.Name}} {{$xtable.Main.Name}} {{$table.Enum.EnumType}}
    {{- $exit := false -}}
    {{- range $x,$v := $xtable.Main.ViewOpts -}}
    {{- if eq $v.URL $table.Name.Raw -}}
    {{- if and (ne "" $v.RwName) (ne "" $v.FwName) -}}
    {{- $exit = true}}
   this.form_{{$table.UNQ}}.{{$v.FwName}} = mform.{{$v.RwName}}   
  {{end -}}
    {{- end -}}
{{- end -}}

    {{- if eq false $exit -}}
    {{- range $i,$c := $MLLERows -}}
      {{- if eq true $c.Enum.IsEnum -}}
        {{- if eq $table.Enum.EnumType $c.Enum.EnumType}}
    this.form_{{$table.UNQ}}.{{$table.Enum.Id}} = mform.{{$c.Name}}
        {{end -}}
    {{- end -}}
    {{- end -}}
    {{- end}}
   
    //发送查询请求

    let that = this
    that.conf.loading = true
    
    //构建统计查询
    {{- range $i,$v:= $table.LStatOpts}}
      this.$theia.http.get("{{$v.URL|lower}}",this.form_{{$table.UNQ}}).then(res=>{
        let item = res||{}
        {{$uxcols := fltrColums $table $v.RwName -}}
        {{- range $i,$c := $uxcols -}}
        {{- if and (ne "" $c.Cmpnt.Format) (eq true $c.Field.IsNumber) -}}
          that.stat_{{$v.UNQ}}.{{$c.Name}} = that.$theia.str.numberFormat(item.{{$c.Name}},'{{$c.Cmpnt.Format}}')
       {{- else -}}
         that.stat_{{$v.UNQ}}.{{$c.Name}} = item.{{$c.Name}}
        {{- end}}
        {{end}}
      });
    {{end}}
    
    //数据查询
    this.$theia.http.get("/{{$table.Name.MainPath|lower}}/query",this.form_{{$table.UNQ}}).then(res=>{
        that.conf.loading = false
        that.dataList_{{$table.UNQ}} = res.items||[]
        that.total_{{$table.UNQ}} = res.count
        that.dataList_{{$table.UNQ}}.forEach(item => {
            
          {{- range $i,$c := $LLERows -}}
        {{- if eq true $c.Enum.IsEnum}}
        item.{{$c.Name}}_label = that.$theia.enum.getNames("{{$c.Enum.EnumType}}",item.{{$c.Name}})
          {{- end -}}
          {{- if eq "switch" $c.Cmpnt.Type}}
        item.{{$c.Name}}_switch = item.{{$c.Name}} == 0
          {{- end}}
        {{- end}}

        {{- range $i,$c := $leRow}}
          {{- if and (ne "" $c.Cmpnt.Format) (eq true $c.Field.IsDate)}}
        item.le_{{$c.Name}} = that.$theia.str.dateFormat(item.{{$c.Name}},'{{$c.Cmpnt.Format}}')
        {{- else if and (ne "" $c.Cmpnt.Format) (eq true $c.Field.IsNumber)}}
        item.le_{{$c.Name}} = that.$theia.str.numberFormat(item.{{$c.Name}},'{{$c.Cmpnt.Format}}')
        {{- else if eq "mobile" $c.Cmpnt.Type}}
        item.le_{{$c.Name}} = that.$theia.str.phoneFormat(item.{{$c.Name}})
        {{- else if eq "cut" $c.Cmpnt.Type }}
        item.le_{{$c.Name}} = that.$theia.str.cut(item.{{$c.Name}},{{$c.Cmpnt.Format}})
          {{else}}
        item.le_{{$c.Name}} = item.{{$c.Name}}
        {{- end}}
        {{- end }}

        {{- range $i,$c := $lstRow}}
            {{- if and (ne "" $c.Cmpnt.Format) (eq true $c.Field.IsDate)}}
        item.{{$c.Name}} = that.$theia.str.dateFormat(item.{{$c.Name}},'{{$c.Cmpnt.Format}}')
          {{- else if and (ne "" $c.Cmpnt.Format) (eq true $c.Field.IsNumber)}}
        item.{{$c.Name}} = that.$theia.str.numberFormat(item.{{$c.Name}},'{{$c.Cmpnt.Format}}')
          {{- else if eq "mobile" $c.Cmpnt.Type}}
        item.{{$c.Name}} = that.$theia.str.phoneFormat(item.{{$c.Name}})
          {{- else if eq "cut" $c.Cmpnt.Type }}
        item.{{$c.Name}} = that.$theia.str.cut(item.{{$c.Name}},{{$c.Cmpnt.Format}})
          {{- end}}
          {{- end }}
      });
    })
  },

  {{ range $i,$c:= $qrow }}
  {{- if eq "ddl" $c.Cmpnt.Type -}}
  on{{.Name}}dropdownClick(f, x) {
    let mf = f || {}
    x.{{.Name}}_label = mf.name
    x.{{.Name}} = mf.value
    this.onQuery()
  },
{{- end }}{{ end }}


{{ range $i,$c:= $lstRow }}
{{- if eq "switch" $c.Cmpnt.Type -}}
on{{$c.Name}}SwitchChange(val) {
  let form = {}
  form.{{$c.Name}} = val==true?0:1;
  {{- range $i,$v :=  $table.PKColums}}
	form.{{$v.Name}} = this.form_{{$table.UNQ}}.{{$v.Name}}
  {{- end}}
  let that = this
  this.$theia.http.post("/{{$table.Name.MainPath|lower}}/switch",form).then(res=>{
    that.$notify.success({title: '成功',message: '修改{{$c.Label}}成功',duration:5000})
  }).catch(res=>{
    let code = ((res||{}).response||{}).status||0
    let msg = `修改{{$c.Label}}失败(${code})`
    that.$notify.error({title: '失败',message: msg,duration:5000})
  });
},
{{- end }}{{ end }}

