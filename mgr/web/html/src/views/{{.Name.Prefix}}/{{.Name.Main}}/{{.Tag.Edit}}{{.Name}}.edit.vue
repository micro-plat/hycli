{-{- $table := .}-}
{-{- $ucols :=  $table.GetColumnsByTPName "U"}-}
{-{- $enumColumns :=$table.EnumColumns}-}
{-{- $editOpts :=$table.BarOpts.GetEditOpt }-}
{-{- $title := f_string_contact "修改" $table.Desc}-}
{-{- if ne "" $editOpts.Label}-}
{-{- $title = $editOpts.Label}-}
{-{- end}-}
<template tag="{-{.Marker}-}">
  <el-dialog
    v-model="conf.visible"
    title="{-{$title}-}"
    :width="conf.width"
    draggable
    align-center="true"
    :close-on-click-modal="false"
    :before-close="hide"
  >
<el-form :model="form"  ref="form" :rules="rules">
{-{- template "add.tmpl.html" $ucols}-}
</el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="hide" icon="close">取消</el-button>
        <el-button type="primary" @click="save" icon="select">提交</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script>
import rtext from "@/views/cmpnts/rtext.vue"
export default {
  components: {
    rtext
},
  data() {
    return {
      conf: {
        visible: false,
        {-{- if gt (len $ucols) 14}-}
        width:"70%",
        {-{- else }-}
        {-{- $width:= "60%"}-}
        {-{- range $i,$c:= $table.BarOpts.GetByTag "ADD" }-}
        {-{- $width = $c.GetParam "width" "60%" }-}
        {-{- end}-}
        width:"{-{$width}-}",
      {-{- end  }-}
        uploadPath:this.$theia.env.join("/file/upload"),
      },
      {-{- template "add.tmpl.js" $ucols}-}
  },
  methods: {
    show(form) {
      this.conf.visible = true;
      this.loadEnums()
      this.form = Object.assign(form,this.$route.params)
      this.get(form)
    },
      get(form){
        let that = this
        this.conf.loading = true
        this.$theia.http.get("/{-{.Name.MainPath|lower}-}",form).then(res=>{
        {-{- range $i,$c := $ucols}-}
        {-{- if eq "switch" $c.Cmpnt.Type}-}
          res.{-{$c.Name}-}_switch = res.{-{$c.Name}-} == 0
        {-{- else if eq "tree" $c.Cmpnt.Type  }-}
        that.$refs.tree_{-{$c.UNQ}-}.setCheckedKeys(res.{-{$c.Name}-}.split(","))
        {-{- else if eq true (f_string_start $c.Cmpnt.Type "multi")}-}
          res.{-{$c.Name}-} = (res.{-{$c.Name}-}+"").split(",")
        {-{- end}-}
        {-{- end}-}
        that.form = Object.assign(that.form, res)
        {-{- range $j,$x:=$ucols}-}
        {-{- if and (eq true $x.Enum.IsEnum) (eq true (f_string_start $x.Cmpnt.Type "tree|cascader"))}-}
        let {-{$x.Name}-} = that.form.{-{$x.Name}-}.split(",")
        that.form.{-{$x.Name}-} = {-{$x.Name}-}
        {-{- end}-}
      {-{- end}-}
      //处理枚举重新绑定
      that.loadEnums()
      that.conf.loading = false
        }).catch(res=>{
          let code = ((res||{}).response||{}).status||0
          let msg = `{-{.Desc}-}查询失败(${code})`
          that.$notify.error({title: '失败',message:msg ,duration:5000})
        })
    },
    save(){
       {-{- range $i,$c:= $ucols}-}
         {-{- if eq "switch" $c.Cmpnt.Type }-}
        this.form.{-{$c.Name}-} = this.form.{-{$c.Name}-}_switch?0:1;
         {-{- end}-}
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
        {-{- range $i,$c:= $ucols }-}
        {-{- if eq "password" $c.Cmpnt.Type  }-}
        postForm.{-{$c.Name}-} = this.$theia.crypto.md5(this.form.{-{$c.Name}-})
        {-{- else if eq "tree" $c.Cmpnt.Type  }-}
        let checkKeys = this.$refs.tree_{-{$c.UNQ}-}.getCheckedKeys()||[]
        let halfKeys = this.$refs.tree_{-{$c.UNQ}-}.getHalfCheckedKeys()||[]
        postForm.{-{$c.Name}-} = checkKeys.concat(halfKeys).join(",")
        {-{- else if eq "cascader" $c.Cmpnt.Type  }-}
        let {-{$c.Name}-} = []
        let {-{$c.Name}-}Nodes = this.$refs.cascader_{-{$c.UNQ}-}.getCheckedNodes() || []
        {-{$c.Name}-}Nodes.forEach(v => {
          {-{$c.Name}-}.push(v.value)
        })
        postForm.{-{$c.Name}-} = {-{$c.Name}-}.join(",")
        {-{- else if eq true (f_string_start $c.Cmpnt.Type "multi")}-}
        postForm.{-{$c.Name}-} = (postForm.{-{$c.Name}-}||[]).join(",")
        {-{- end }-}
        {-{- end}-}
        {-{- range $i,$c:=  $table.ListOpts.GetByTag "UPDATE" }-}
        {-{- $save2window := $c.GetParams "save2window" -}-}
        {-{- if ne "" $save2window}-}
        
        //将数据保存到window缓存中
        window.{-{$save2window}-} = postForm
        {-{- range $i,$v := $table.PKColumns.GetValidColumns}-}
        window.{-{$v.Name}-} = null
      {-{- end}-}
        {-{- end -}-}
        {-{- end }-}

        this.$theia.http.put("/{-{.Name.MainPath|lower}-}",postForm).then(res=>{
            that.$notify.success({title: '成功',message: '{-{.Desc}-}保存成功',duration:5000})
            {-{- if ne "" $table.Enum.EnumType}-}
            that.$theia.enum.clear("{-{$table.Enum.EnumType}-}")
            that.$theia.enum.get("{-{$table.Enum.EnumType}-}")
            {-{- end}-}
            that.hide()
            that.$emit("onsaved")
        }).catch(res=>{
          let code = ((res||{}).response||{}).status||0
          let msg = `{-{.Desc}-}修改失败(${code})`
          that.$notify.error({title: '失败',message: msg,duration:5000})
        })
    },
    hide() {
      this.conf.visible = false;
      this.$refs.form.resetFields();
    },
    {-{- range $i,$c:= $ucols }-}
    {-{-  if (eq "rtext" $c.Cmpnt.Type )}-}
    onRtext{-{$c.Name}-}Change(text){
      this.{-{$c.Ext.FormName}-}.{-{$c.Name}-} = text
    },
    {-{- end}-}
    {-{- end}-}
     onUploadSuccess(response){
      {-{- range $i,$c:= $ucols}-}
      {-{- if eq "file" $c.Cmpnt.Type }-}
      this.form.{-{$c.Name}-} = response.path
      {-{- end}-}
      {-{- end}-}
    },

    {-{- range $j,$x:=$ucols}-}
    {-{- $ctriger := $ucols.GetColumnsByTriggerChangeEvent $x.Name}-}
     {-{- if gt (len $ctriger) 0}-}
    onChange_{-{$x.Name}-}(val,tp){
      if(this.conf.loading || !val||!this.form.{-{$x.Name}-}){
        return
      }
      {-{- range $i,$c:= $ctriger}-}
      {-{- if eq true ($c.IsAssctColumn $x.Name)}-}
      {-{- $group:= f_string_trim $c.Enum.Group "#"}-}
      if (!val) {
        this.{-{$c.Name}-}List = []
        // this.form.{-{$c.Name}-} = null
        return
      }
      this.{-{$c.Name}-}List = this.$theia.enum.get("{-{$c.Enum.EnumType}-}",val,{-{if f_string_start $c.Enum.Group "#"}-}this.$theia.user.get("{-{$group}-}"){-{else}-}"{-{$c.Enum.Group}-}" {-{end}-})||[]
      // this.form.{-{$c.Name}-} =this.{-{$c.Name}-}List.length>0?(this.{-{$c.Name}-}List[0]||{}).value:null
      {-{- $cxtriger := $ucols.GetColumnsByTriggerChangeEvent $c.Name}-}
      {-{- if gt (len $cxtriger) 0}-}
      {-{- range $mx,$cx := $cxtriger}-}
      this.onChange_{-{$c.Name}-}(this.form.{-{$c.Name}-},"{-{$cx.Enum.EnumType}-}")
      {-{- end}-}
      {-{- end}-}
      {-{- end}-}
      {-{- end}-}
      //添加联动
      {-{- $p := $x.GetParams "@change"}-}
      {-{- if gt (len $p) 0}-}
      let that = this
      if(tp){
        let lst= this.$theia.enum.get(tp)
        var hasValues = {}
        lst.forEach(el => {
          if(el.value == val)
          {-{- range $i,$c:= $p}-}
          if (!that.form.{-{$i}-} || !that.form._{-{$i}-} || that.form._{-{$i}-} == that.form.{-{$i}-}) {
            hasValues.{-{$i}-} = true
            that.form.{-{$i}-} = el.{-{$c}-}
            that.form._{-{$i}-} = el.{-{$c}-}
          }
          {-{- end}-}
        });
        {-{- range $i,$c:= $p}-}
        if(!hasValues.{-{$i}-}){
          that.form._{-{$i}-} = that.form.{-{$i}-}
        }
        {-{- end}-}
      }
      {-{- end}-}

    },
    {-{- end}-}
    {-{- end}-}
    {-{- $st := $table.NewScene $ucols}-}
    {-{- template "enum.tmpl.js" $st }-}
  },
};
</script>

<style scoped>
.form_item_info{
  margin-left: 4px;
  color:#999;
}
.form_item_info:hover{
  cursor: pointer;
  font-weight: 500;
  color:#333;
}
</style>