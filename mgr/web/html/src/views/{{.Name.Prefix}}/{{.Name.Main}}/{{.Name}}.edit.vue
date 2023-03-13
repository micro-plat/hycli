{{- $table := . }}
{{- $ucols := fltrColums $table "U"}}
{{- $enumColums :=$table.EnumColums}}
<template>
  <el-dialog
    v-model="conf.visible"
    title="修改 {{.Desc}}"
    width="60%"
    draggable
    :close-on-click-modal="false"
    :before-close="hide"
  >
<el-form :model="form" size="small" ref="form" :rules="rules">
{{template "add.tmpl.html" $ucols}}
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
      {{- template "add.tmpl.js" $ucols -}}
  },
  methods: {
    show(form) {
      this.conf.visible = true;
      this.get(form)
    },
      get(form){
        let that = this
        this.$theia.http.get("/{{.Name.MainPath|lower}}",form).then(res=>{
        {{- range $i,$c := $ucols -}}
        {{- if eq "switch" $c.ExCMPT.Type}}
          res.{{$c.Name}}_switch = res.{{$c.Name}} == 0
        {{- end}}
        {{- end}}
        Object.assign(that.form, res)
        }).catch(res=>{
          let code = res.response.status
          let msg = `{{.Desc}}查询失败(${code})`
          that.$notify.error({title: '失败',message:msg ,duration:5000})
        })
    },
    save(){
       {{range $i,$c:= $ucols -}}
         {{- if eq "switch" $c.ExCMPT.Type  -}}
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
        this.$theia.http.put("/{{.Name.MainPath|lower}}",this.form).then(res=>{
            that.$notify.success({title: '成功',message: '{{.Desc}}保存成功',duration:5000})
            that.$theia.enum.clear("{{.Name.Short}}")
            that.$emit("onsaved")
            that.hide()
        }).catch(res=>{
          let code = res.response.status
          let msg = `{{.Desc}}修改失败(${code})`
          that.$notify.error({title: '失败',message: msg,duration:5000})
        })
    },
    hide() {
      this.conf.visible = false;
      this.$refs.form.resetFields();
    },
     onUploadSuccess(response){
      {{range $i,$c:= $ucols -}}
      {{- if eq "file" $c.ExCMPT.Type  -}}
      this.form.{{$c.Name}} = response.path
      {{- end -}}
      {{- end}}
    },

 {{range $i,$c:= $enumColums -}}
    change{{$c.Name}}(value){
      this.{{$c.Name}}List =  this.$theia.enum.get("{{$c.Enum.EnumType}}",value)
    },
    {{- end}}
  },
};
</script>

<style>
</style>