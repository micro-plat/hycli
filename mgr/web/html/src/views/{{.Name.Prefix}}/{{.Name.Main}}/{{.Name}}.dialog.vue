{{- $table := .}}
{{- $optCols :=$table.ListOpts}}
<template>
  <div>
    {{ range $x,$m:=$optCols -}}
    {{- if eq "DIALOG" $m.Name -}}
    <el-dialog
      v-model="conf.{{$m.UNQ}}_visible"
      title="{{.Label}}"
      :close-on-click-modal="false"
      width="60%"
      draggable
      :before-close="hide_{{$m.UNQ}}"
    >
      {{- $rows:= fltrColums $table $m.RwName (sjoin "form_" $m.UNQ)}}
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
        {{ range $x,$m:=$optCols -}}
        {{- if eq "DIALOG" $m.Name -}}
        {{$m.UNQ}}_visible:false,
        {{- end}}
        {{- end}}
      },
  {{ range $x,$m:=$optCols -}}
     {{- if eq "DIALOG" $m.Name}}
      //{{$m.Label}} form by  {{$m.RwName}}
      {{- $rows:= fltrColums $table $m.RwName -}}
        {{range $i,$c:=$rows -}} 
        {{- if or (eq true $c.Enum.IsEnum) (eq "multiselect" $c.Cmpnt.Type)}}
    {{.Name}}List:this.$theia.enum.get("{{$c.Enum.EnumType}}","{{$c.Enum.PID}}","{{$c.Enum.Group}}"),
         {{- else}}
    {{$c.Name}}:"",
         {{- end -}}
        {{end}}
    form_{{$m.UNQ}}:{},
   {{- end}}
   {{- end}}
    {{ range $x,$m:=$optCols -}}
    {{- if eq "DIALOG" $m.Name}}
    rules_{{$m.UNQ}}:{
        {{- $rows:= fltrColums $table $m.RwName -}}
         {{- range $i,$c:=$rows}} 
          {{$c.Name}}:[{required:{{$c.Field.Required}}, message:"请输入{{$c.Label}}", trigger: 'blur'}],
          {{- end}}
        },
      {{- end}}
      {{- end}}


    }
  },
  methods:{
     {{- range $x,$m:=$optCols -}}
     {{- if eq "DIALOG" $m.Name}}
     //--------------------{{$m.Label}}---------------------------------
      //显示 {{$m.Label}} 弹出框 {{$m}}
      show_{{$m.UNQ}}(fm){
        {{- $ct:= fltrSearchUITable  $m -}}
        {{- $tbs := contactTBS  $table $ct -}}
        {{- $ctable := $tbs.Current -}}
        {{- $mtable := $tbs.Main -}}
        {{- $MLLERows:= fltrColums $mtable "l-le"}}
          
        //处理关联表{{$m.URL}}
        let currentForm = {}
        {{- range $i,$c := $MLLERows -}}
          {{- if eq true $c.Enum.IsEnum -}}
            {{- if eq $ctable.Enum.EnumType $c.Enum.EnumType}}
        currentForm.{{$table.Enum.Id}} = fm.{{$c.Name}}
            {{- end -}}
        {{- end -}}
        {{- end}}

         Object.assign(this.form_{{$m.UNQ}},currentForm)
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
        this.$refs.fm_{{$m.UNQ}}.validate((v=>{
            if(!v){
                return
            }
        let post_form_{{$m.UNQ}}=this.form_{{$m.UNQ}}
        {{- $rows:= fltrColums $table $m.RwName (sjoin "form_" $m.UNQ)}}
        {{range $i,$c:= $rows -}}
       
         {{- if eq "password" $c.Cmpnt.Type  -}}
        post_form_{{$m.UNQ}}.{{$c.Name}} = this.$theia.crypto.md5(this.form_{{$m.UNQ}}.{{$c.Name}})
         {{- end -}}
          {{end}}
        this.$theia.http.post("{{$m.URL}}",post_form_{{$m.UNQ}}).then(res=>{
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