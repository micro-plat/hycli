{-{- $table := .}-}
{-{- $ucols :=  $table.Columns.GetFontUpdateColumns}-}
{-{- $enumColumns :=$table.Columns.GetEnumColumns}-}
{-{- $editOpts :=$table.Optrs.ListOpts.GetEditOpt }-}
{-{- $title := f_string_contact "修改" $table.Desc}-}
{-{- $title = $editOpts.GetParam "title" $title}-}

<template tag="{-{.Marker}-}">
  <el-dialog
    v-model="conf.visible"
    {-{ if $table.Enum.IsEnum  }-}v-if="conf.visible" {-{ end }-}
    :title="title"
    :width="conf.width"
    draggable
    align-center="true"
    :close-on-click-modal="false"
    :before-close="hide"
  >
<el-form :model="form"  ref="form" :rules="rules_{-{$table.UNQ}-}">
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
        {-{- $width = $editOpts.GetParam "width" "60%" }-}
        width:"{-{$width}-}",
        {-{- end  }-}
        uploadPath:this.$theia.env.join("/file/upload"),
      },
      title:"修改{-{$table.Desc}-}",
    {-{- $st := $table.NewScene $ucols}-}
    {-{- template "rules.tmpl.js" $st }-}
    form:{
        {-{- range $i,$c := $ucols}-}
        {-{- if $c.Cmpnt.Equal "switch"}-}
        {-{$c.Name}-}_switch:false,
        {-{- else if $c.Cmpnt.StartWith "multi"}-}
        {-{$c.Name}-}:[],
        {-{- else}-}
        {-{$c.Name}-}:"",
        {-{- end}-}
        {-{- end}-}
    },
    {-{- range $i,$c := $ucols.GetEnumColumns}-}
    {-{.Name}-}List:[],
    {-{- end}-}
    }
  },
  methods: {
    show(form) {
    {-{- $titles := $table.Columns.GetColumns "@title" }-}
    {-{- if gt $titles.Len 0 }-}
    this.title = form.{-{$titles.First.Name}-}||this.title
    {-{- end}-}
      this.conf.visible = true;
      this.loadEnums_{-{$table.UNQ}-}()
      this.form = Object.assign(form,this.$route.params)
      this.get(form)
    },
      get(form){
        let that = this
        this.conf.loading = true
        let postForm = {}
        {-{- range $i,$f := $table.Columns.GetPKColumns}-}
        postForm.{-{$f.Name}-} = form.{-{$f.Name}-}
        {-{- end}-}
        this.$theia.http.get("/{-{.Name.MainPath|lower}-}",postForm).then(res=>{
          {-{- range $i,$c := $ucols}-}
          {-{- if $c.Cmpnt.Equal "switch"}-}
          res.{-{$c.Name}-}_switch = res.{-{$c.Name}-} == 0
          {-{- else if $c.Cmpnt.Equal "tree" }-}
          that.$refs.tree_{-{$c.UNQ}-}.setCheckedKeys(res.{-{$c.Name}-}.split(","))
          {-{- else if $c.Cmpnt.StartWith "multi"}-}
          res.{-{$c.Name}-} = (res.{-{$c.Name}-}+"").split(",")
          {-{- end}-}
          {-{- end}-}
          that.form = Object.assign(that.form, res)
          {-{- range $j,$x:=$ucols}-}
          {-{- if and (eq true $x.Enum.IsEnum) ($x.Cmpnt.StartWith "tree")}-}
          let {-{$x.Name}-} = that.form.{-{$x.Name}-}.split(",")
          that.form.{-{$x.Name}-} = {-{$x.Name}-}
          {-{- else if and (eq true $x.Enum.IsEnum) ($x.Cmpnt.StartWith "cascader")}-}
          {-{- if eq $x.Cmpnt.Format "multiple"}-}
          let {-{$x.Name}-} = that.form.{-{$x.Name}-}.split(",")
          that.form.{-{$x.Name}-} = {-{$x.Name}-}
          {-{- else}-}
          that.form.{-{$x.Name}-} = that.form.{-{$x.Name}-}
          {-{- end}-}
          {-{- end}-}
          {-{- end}-}
        //处理枚举重新绑定
          this.loadEnums_{-{$table.UNQ}-}()
          that.conf.loading = false
        }).catch(res=>{
          let code = ((res||{}).response||{}).status||0
          let msg = `{-{.Desc}-}查询失败(${code})`
          that.$notify.error({title: '失败',message:msg ,duration:5000})
        })
    },
    save(){
       {-{- range $i,$c:= $ucols}-}
         {-{- if $c.Cmpnt.Equal "switch"}-}
        this.form.{-{$c.Name}-} = this.form.{-{$c.Name}-}_switch?0:1;
         {-{- end}-}
        {-{- end}-}
        this.$refs.form.validate((v=>{
            if(v) this.onSave()
        }))
    },
    onSave(){
      let that = this
        let postForm = Object.assign({},this.form)
        {-{- range $i,$c:= $ucols }-}
        {-{- if $c.Cmpnt.Equal "password"}-}
        postForm.{-{$c.Name}-} = this.$theia.crypto.md5(this.form.{-{$c.Name}-})
        {-{- else if  $c.Cmpnt.Equal "tree"}-}
        let checkKeys = this.$refs.tree_{-{$c.UNQ}-}.getCheckedKeys()||[]
        let halfKeys = this.$refs.tree_{-{$c.UNQ}-}.getHalfCheckedKeys()||[]
        postForm.{-{$c.Name}-} = checkKeys.concat(halfKeys).join(",")
        {-{- else if $c.Cmpnt.Equal "cascader"}-}
        let {-{$c.Name}-} = []
        let {-{$c.Name}-}Nodes = this.$refs.cascader_{-{$c.UNQ}-}.getCheckedNodes() || []
        {-{$c.Name}-}Nodes.forEach(v => {
          {-{$c.Name}-}.push(v.value)
        })
        postForm.{-{$c.Name}-} = {-{$c.Name}-}.join(",")
        {-{- else if $c.Cmpnt.StartWith "multi"}-}
        postForm.{-{$c.Name}-} = (postForm.{-{$c.Name}-}||[]).join(",")
        {-{- end }-}
        {-{- end}-}
        {-{- if ne "" $editOpts.Tag}-}
        {-{- $save2window := $editOpts.GetParam "save2window" -}-}
        {-{- if ne "" $save2window}-}
        //将数据保存到window缓存中
        window.{-{$save2window}-} = postForm
        {-{- range $i,$v := $table.Columns.GetPKColumns.GetValidColumns}-}
        window.{-{$v.Name}-} = null
        {-{- end}-}
        {-{- end}-}
        {-{- end}-}
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
    {-{-  if $c.Cmpnt.Equal "rtext"}-}
    onRtext{-{$c.Name}-}Change(text,urls){
      this.{-{$c.Ext.FormName}-}.{-{$c.Name}-} = text
      this.{-{$c.Ext.FormName}-}.{-{$c.Name}-}_urls = (urls||[]).join(",")
    },
    {-{- end}-}
    {-{- end}-}
     onUploadSuccess(response){
      {-{- range $i,$c:= $ucols}-}
      {-{- if $c.Cmpnt.Equal "file"}-}
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
      if (!val) {
        this.{-{$c.Name}-}List = []
        return
      }
      this.{-{$c.Name}-}List = this.$theia.enum.get("{-{$c.Enum.EnumType}-}",val,{-{if $c.Enum.GroupIsStatic}-}this.$theia.user.get("{-{$c.Enum.GetGroupValue}-}"){-{else}-}"{-{$c.Enum.Group}-}" {-{end}-})||[]
      {-{- $cxtriger := $ucols.GetColumnsByTriggerChangeEvent $c.Name}-}
      {-{- if gt (len $cxtriger) 0}-}
      {-{- range $mx,$cx := $cxtriger}-}
      this.onChange_{-{$c.Name}-}(this.form.{-{$c.Name}-},"{-{$cx.Enum.EnumType}-}")
      {-{- end}-}
      {-{- end}-}
      {-{- end}-}
      {-{- end}-}
      //添加联动
      {-{- $p := $x.GetParamMap "@change"}-}
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