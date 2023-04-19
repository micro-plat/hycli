{-{- $table := .}-}
{-{- $cColumns := fltrColumns $table "C"}-}
{-{- $enumColumns :=$table.EnumColumns}-}
<template tag="{-{.Marker}-}">
  <el-dialog
    v-model="conf.visible"
    title="添加 {-{.Desc}-}"
    width="60%"
    draggable
    :close-on-click-modal="false"
    :before-close="hide"
  >
  <el-form :model="form" size="small" ref="form" :rules="rules">
    {-{- template "add.tmpl.html" $cColumns}-}
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
    {-{- template "add.tmpl.js" $cColumns }-}
  },
  methods: {
    show(fm = {}) {
      this.conf.visible = true;
      this.form = Object.assign(fm,this.$route.params)
      {-{- range $i,$c:= $cColumns }-}
      {-{- if and (gt (len (fltrAssctColumns $cColumns $c.Name)) 0) (eq true (fltrHasConst $c "rp"))}-} 
      this.onChange_{-{$c.Name}-}(this.$route.params["{-{$c.Name}-}"]||this.form["{-{$c.Name}-}"])
      {-{- end}-} 
      {-{- end}-}
    },
    save(){
       {-{- range $i,$c:= $cColumns }-}
         {-{- if eq "switch" $c.Cmpnt.Type  }-}
        this.form.{-{$c.Name}-} = this.form.{-{$c.Name}-}_switch?0:1;
         {-{- end }-}
          {-{- end}-}
        this.$refs.form.validate((v=>{
            if(v){
                this.onSave()
            }
        }))
    },
    onSave(){
        let that = this
        let postForm = Object.assign({},this.form)
        {-{- range $i,$c:= $cColumns }-}
        {-{- if eq "password" $c.Cmpnt.Type  }-}
        postForm.{-{$c.Name}-} = this.$theia.crypto.md5(this.form.{-{$c.Name}-})
        {-{- else if eq true (fltrStart $c.Cmpnt.Type "multi")}-}
        postForm.{-{$c.Name}-} = (postForm.{-{$c.Name}-}||[]).join(",")
        {-{- end }-}
        {-{- end}-}
        this.$theia.http.post("/{-{.Name.MainPath|lower}-}",postForm).then(res=>{
            that.$notify.success({title: '成功',message: '{-{.Desc}-}保存成功',duration:5000})
            that.hide()
            {-{- if ne "" $table.Enum.EnumType}-}
            {-{- if eq true $table.Enum.Multiple}-}
            that.$theia.enum.clear()
            {-{- else}-}
            that.$theia.enum.clear("{-{$table.Enum.EnumType}-}")
            that.$theia.enum.get("{-{$table.Enum.EnumType}-}")
            {-{- end}-}
            {-{- end}-}
            that.$emit("onsaved")
        }).catch(res=>{
          let code = ((res||{}).response||{}).status||0
          let msg= `{-{.Desc}-}保存失败(${code})`
          that.$notify.error({title: '失败',message:msg,duration:5000})
        })
    },
    hide() {
      this.conf.visible = false;
      this.$refs.form.resetFields();
    },
    onUploadSuccess(response){
      {-{- range $i,$c:= $cColumns }-}
      {-{- if eq "file" $c.Cmpnt.Type  }-}
      {-{- if eq true (fltrStart $c.Cmpnt.Format "#")}-}
      this.form.{-{fltrTrim $c.Cmpnt.Format "#"}-} = response.path
      {-{- else}-}
      this.form.{-{$c.Name}-} = response.path
      {-{- end}-}
      
      {-{- end }-}
      {-{- end}-}
    },
    {-{- range $i,$c:= $cColumns }-}
    {-{- $aColumns := fltrAssctColumns $cColumns $c.Name }-}
    {-{- if gt (len $aColumns) 0}-} 
    onChange_{-{$c.Name}-}(val){
      {-{- range $j,$x:= $aColumns  }-}
      this.{-{$x.Name}-}List = this.$theia.enum.get("{-{$x.Enum.EnumType}-}",val)
      this.form.{-{$x.Name}-} = null
      {-{- end}-}
    },
    {-{- end}-}
    {-{- end}-}
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