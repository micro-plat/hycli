{{- $table := .|fltrUITable}}
{{- $optRow :=$table.Optrs}}
<template>
  <div>
    {{ range $x,$m:=$optRow -}}
    {{- if eq "DIALOG" $m.Name -}}
    <el-dialog
      v-model="conf.{{$m.UNQ}}_visible"
      title="{{.Label}}"
      width="30%"
      draggable
      :before-close="hide_{{$m.UNQ}}"
    >
      {{- $rows:= fltrUIRows $table $m.RwName (sjoin "form_" $m.UNQ)}}
      <el-form
        :model="form_{{$m.UNQ}}"
        size="small"
        ref="fm_{{$m.UNQ}}"
        :rules="rules_{{$m.UNQ}}"
      >
        {{- template "add.tmpl.html" $rows}}
      </el-form>
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
        {{ range $x,$m:=$optRow -}}
        {{- if eq "DIALOG" $m.Name -}}
        {{$m.UNQ}}_visible:false,
        {{- end}}
        {{- end}}
      },
  {{ range $x,$m:=$optRow -}}
     {{- if eq "DIALOG" $m.Name}}
      //{{$m.Label}} form by  {{$m.RwName}}
      form_{{$m.UNQ}}:{
       {{- $rows:= fltrUIRows $table $m.RwName -}}
        {{range $i,$c:=$rows -}} 
        {{$c.Name}}:"",
        {{- end}}
    },
   {{- end}}
   {{- end}}
    {{ range $x,$m:=$optRow -}}
    {{- if eq "DIALOG" $m.Name}}
    rules_{{$m.UNQ}}:{
        {{- $rows:= fltrUIRows $table $m.RwName -}}
         {{- range $i,$c:=$rows}} 
          {{$c.Name}}:[{required:{{$c.Required}}, message:"请输入{{$c.Desc}}", trigger: 'blur'}],
          {{- end}}
        },
      {{- end}}
      {{- end}}


    }
  },
  methods:{
     {{- range $x,$m:=$optRow -}}
     {{- if eq "DIALOG" $m.Name}}
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
        this.$refs.form_{{$m.UNQ}}.validate((v=>{
            if(!v){
                return
            }
        this.$theia.http.post("{{$m.URL}}",this.form_{{$m.UNQ}}).then(res=>{
            that.$notify.success({title: '成功',message: '提交成功',duration:5000})
            that.$emit("onsaved")
            that.hide_{{$m.UNQ}}()
            
        }).catch(res=>{
          let code = res.response.status
          let msg= `提交失败(${code})`
          that.$notify.error({title: '失败',message:msg,duration:5000})
        })
        }))
    },
    //-----------------------------------------------------------
    {{- end -}}
      {{- end}}
  },
}
</script>