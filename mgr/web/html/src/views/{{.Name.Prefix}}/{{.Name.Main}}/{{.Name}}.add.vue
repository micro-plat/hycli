{{- $table := .}}
{{- $ccolums := fltrColums $table "C"}}
{{- $enumColums :=$table.EnumColums}}
<template>
  <el-dialog
    v-model="conf.visible"
    title="添加 {{.Desc}}"
    width="60%"
    draggable
    :close-on-click-modal="false"
    :before-close="hide"
  >
  <el-form :model="form" size="small" ref="form" :rules="rules">
    {{- template "add.tmpl.html" $ccolums}}
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="hide">取消</el-button>
        <el-button type="primary" @click="save">提交</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script>
export default {
  data() {
    return {
      conf: {
        visible: false,
        uploadPath:this.$theia.env.join("/file/upload"),
      },
    {{- template "add.tmpl.js" $ccolums -}}
  },
  methods: {
    show(fm={}) {
      this.conf.visible = true;
      this.form = fm
    },
    save(){
       {{range $i,$c:= $ccolums -}}
         {{- if eq "switch" $c.Cmpnt.Type  -}}
        this.form.{{$c.Name}} = this.form.{{$c.Name}}_switch?0:1;
         {{- end -}}
          {{end}}
        this.$refs.form.validate((v=>{
            if(v){
                this.onSave()
               
            }
        }))
    },
    onSave(){
        let that = this
        let postForm=this.form
        {{range $i,$c:= $ccolums -}}
         {{- if eq "password" $c.Cmpnt.Type  -}}
        postForm.{{$c.Name}} = this.$theia.crypto.md5(this.form.{{$c.Name}})
         {{- end -}}
          {{end}}
        this.$theia.http.post("/{{.Name.MainPath|lower}}",postForm).then(res=>{
            that.$notify.success({title: '成功',message: '{{.Desc}}保存成功',duration:5000})
            that.$theia.enum.clear("{{.Name.Short}}")
           that.$emit("onsaved")
            that.hide()
        }).catch(res=>{
          let code = ((res||{}).response||{}).status||0
          let msg= `{{.Desc}}保存失败(${code})`
          that.$notify.error({title: '失败',message:msg,duration:5000})
          that.$emit("onsaved")
            that.hide()
        })
    },
    hide() {
      this.conf.visible = false;
      this.$refs.form.resetFields();
    },
    onUploadSuccess(response){
      {{range $i,$c:= $ccolums -}}
      {{- if eq "file" $c.Cmpnt.Type  -}}
      this.form.{{$c.Name}} = response.path
      {{- end -}}
      {{- end}}
    },
    {{- range $i,$c:= $ccolums -}}
    {{- $acolums := fltrAssctColums $ccolums $c.Name }}
    {{- if gt (len $acolums) 0}} 
    onChange_{{$c.Name}}(val){
      {{- range $j,$x:= $acolums  }}
      this.{{$x.Name}}List = this.$theia.enum.get("{{$x.Enum.EnumType}}",val)
      this.form.{{$x.Name}} = null
      {{- end}}
    },
    {{- end}}
    {{- end}}
},
}
</script>

<style scoped>
.avatar-uploader .avatar {
  width: 80px;
  height: 80px;
  display: block;
}
</style>