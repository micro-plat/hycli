{{- $table := .}}
{{- $opts :=$table.ListOpts}}
<template>
  <div>
    {{ range $x,$m:=$opts -}}
    {{- if or (eq "CNFRM" $m.Name) (eq "DEL" $m.Tag) -}}
    <el-dialog
      v-model="conf.{{$m.UNQ}}_visible"
      title="{{$m.Label}}"
      width="30%"
      draggable
      :close-on-click-modal="false"
      :before-close="hide_{{$m.UNQ}}"
    >
      {{- $cols:= fltrColums $table $m.RwName (sjoin "form_" $m.UNQ)}}
      <span
        >确认{{$m.Label}}({{range $i,$c:= $cols}} {{ $c.Label }}:<span
          v-text="form_{{$m.UNQ}}.{{$c.Name}}"
        ></span
        >{{ end }}) 吗?</span
      >
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="hide_{{$m.UNQ}}">取消</el-button>
          <el-button type="primary" @click="save_{{$m.UNQ}}">提交</el-button>
        </span>
      </template>
    </el-dialog>
    {{- end }}
    {{- end }}
  </div>
</template>

<script>
export default {
  data(){
    return{
      conf:{
        {{ range $x,$m:=$opts -}}
        {{- if or (eq "CNFRM" $m.Name) (eq "DEL" $m.Tag) -}}
        {{$m.UNQ}}_visible:false,
        {{- end}}
        {{- end}}
      },
  {{ range $x,$m:=$opts -}}
     {{- if or (eq "CNFRM" $m.Name) (eq "DEL" $m.Tag) -}}
      //{{$m.Label}} form by  [{{$m.RwName}}]
      form_{{$m.UNQ}}:{
       {{- $cols:= fltrColums $table $m.RwName -}}
        {{range $i,$c:=$cols -}} 
        {{$c.Name}}:"",
        {{- end}}
         },
   {{- end}}
   {{- end}}
    }
  },
  methods:{
     {{- range $x,$m:=$opts -}}
     {{- if or (eq "CNFRM" $m.Name) (eq "DEL" $m.Tag) -}}
     //--------------------{{$m.Label}}---------------------------------
      //显示 {{$m.Label}} 弹出框
      show_{{$m.UNQ}}(fm){
         Object.assign(this.form_{{$m.UNQ}},fm)
         this.conf.{{$m.UNQ}}_visible = true;
      },
      //隐藏 {{$m.Label}} 弹出框
      hide_{{$m.UNQ}}(){
        this.conf.{{$m.UNQ}}_visible = false;
        this.$refs.fm_{{$m.UNQ}}.resetFields();
      },
      //保存 {{$m.Label}} 弹出框数据
      save_{{$m.UNQ}}(){
        let that = this
        {{ if eq "DEL" $m.Tag -}}
        this.$theia.http.del("/{{$table.Name.MainPath|lower}}",this.form_{{$m.UNQ}}).then(res=>{
          that.conf.confirmVisible = false 
          that.$notify.success({title: '成功',message: '{{$m.Label}}成功',duration:5000})
          that.$emit("onsaved")
          that.hide_{{$m.UNQ}}()
       }).catch(err=>{
          that.conf.confirmVisible = false
          let code = res.response.status
          let msg= `{{$m.Label}}失败(${code})`
          that.$notify.error({title: '失败',message:msg,duration:5000})
       })
        {{- else}}
        this.$theia.http.post("{{$m.URL}}",this.form_{{$m.UNQ}}).then(res=>{
            that.$notify.success({title: '成功',message: '提交成功',duration:5000})
            that.$emit("onsaved")
            that.hide_{{$m.UNQ}}()
        }).catch(res=>{
          let code = res.response.status
          let msg= `提交失败(${code})`
          that.$notify.error({title: '失败',message:msg,duration:5000})
        })
        {{end}}
    },
    //-----------------------------------------------------------
    {{- end -}}
      {{- end}}
  },
  
}
</script>