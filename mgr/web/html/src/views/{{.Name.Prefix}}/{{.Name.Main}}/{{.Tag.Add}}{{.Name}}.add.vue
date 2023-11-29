{-{- $table := .}-}
{-{- $cColumns :=  $table.Columns.GetFontCreateColumns}-}
{-{- $acColumns :=  $table.Columns.GetAllCreateColumns}-}
{-{- $enumColumns :=$table.Columns.GetEnumColumns}-}
{-{- $addOpts :=$table.Optrs.BarOpts.GetAddOpt }-}
{-{- $title := f_string_contact "添加" $table.Desc}-}
{-{- if ne "" $addOpts.Label}-}
{-{- $title = $addOpts.Label}-}
{-{- end}-}
<template tag="{-{.Marker}-}">
  <el-dialog
    v-model="conf.visible"
    :title="title"
   :width="conf.width"
    draggable
    {-{ if gt $cColumns.Len 9 }-}  align-center="true" {-{ else}-} top="10vh" {-{ end }-}
    :close-on-click-modal="false"
    :before-close="hide"
  >
  <el-form :model="form"  ref="form" :rules="rules">
    {-{- template "add.tmpl.html" $cColumns}-}
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
        title:"{-{$title}-}",
        visible: false,
        {-{- if gt $cColumns.Len 14}-}
        width:"70%",
        {-{- else }-}
        {-{- $width:= "60%"}-}
        {-{- $c :=$table.Optrs.BarOpts.GetAddOpt}-}
        {-{- $width = $c.GetParam "width" "60%" }-}
        width:"{-{$width}-}",
        {-{- end  }-}
        uploadPath:this.$theia.env.join("/file/upload"),
      },
    {-{- template "add.tmpl.js" $cColumns }-}
  },
  methods: {
    show(fm = {}) {
      this.title = fm.$title || this.title
      this.conf.visible = true;
      this.loadEnums_{-{$table.UNQ}-}()
      let local = {}
      {-{- $c :=$table.Optrs.BarOpts.GetAddOpt}-}
      {-{- $read2window := $c.GetParams "read2window" -}-}
      {-{- if ne "" $read2window}-}
        //将数据保存到window缓存中
      local = window.{-{$read2window}-} || {}
      {-{- end}-}
      let cache = Object.assign(local,fm)
      this.form = Object.assign(cache,this.$route.params)

      {-{- range $i,$c := $acColumns.GetColumnsBy "alias"}-}
      {-{- $name := $c.GetOpt "alias" -}-}
      {-{- if ne "" $name}-}
      this.form.{-{$c.Name}-} = cache["{-{$name}-}"]
      {-{- end -}-}
      {-{- end -}-}

      {-{- range $i,$c := $cColumns.GetStartColumns "multi"}-}
      this.form.{-{$c.Name}-} = this.form.{-{$c.Name}-}?this.form.{-{$c.Name}-}:[];
      {-{- end}-}

      {-{- range $j,$x:=$cColumns}-}
      {-{- range $i,$c:= $cColumns.GetColumnsByTriggerChangeEvent $x.Name}-}
      this.onChange_{-{$x.Name}-}(this.$route.params["{-{$x.Name}-}"]||this.form["{-{$x.Name}-}"],"{-{$x.Enum.EnumType}-}")
      {-{- end}-} 
      {-{- end}-}
    },
    save(){
      {-{- range $i,$c:= $cColumns.GetByCmpnt "switch" }-}
      this.form.{-{$c.Name}-} = this.form.{-{$c.Name}-}_switch?0:1;
      {-{- end }-}
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
       {-{- else if $c.Cmpnt.StartWith "multi"}-}
        postForm.{-{$c.Name}-} = (postForm.{-{$c.Name}-}||[]).join(",")
        {-{- end }-}
        {-{- end}-}
        {-{- $c:= $table.Optrs.BarOpts.GetAddOpt }-}
        {-{- $save2window := $c.GetParams "save2window" -}-}
        {-{- if ne "" $save2window}-}
        //将数据保存到window缓存中
        window.{-{$save2window}-} = postForm
        {-{- end}-}
        {-{- $memberClus := $table.Columns.GetFontBackCreateColumns.GetStaticColumns}-}
      {-{- range $i,$v := $memberClus}-}
        postForm.{-{$v.Name}-} = this.$theia.user.get("{-{$v.Ext.Name}-}")
      {-{- end}-}

        //保存数据
        this.$theia.http.post("/{-{.Name.MainPath|lower}-}",postForm).then(res=>{
            that.$notify.success({title: '成功',message: '{-{.Desc}-}保存成功',duration:5000})
            that.hide()
            {-{- if $table.Enum.IsEnumAndMuti}-}
            that.$theia.enum.clear()
            {-{- else}-}
            that.$theia.enum.clear("{-{$table.Enum.EnumType}-}")
            that.$theia.enum.get("{-{$table.Enum.EnumType}-}")
            {-{- end}-}
            
            that.$emit("onsaved")
        }).catch(res=>{
          let code = ((res||{}).response||{}).status||0
          let msg= `{-{.Desc}-}保存失败(${code})`
          msg = code == 909? msg+"数据重复，请修改后重试":msg
          that.$notify.error({title: '失败',message:msg,duration:5000})
        })
    },
    hide() {
      this.conf.visible = false;
      this.$refs.form.resetFields();
    },
    {-{- range $i,$c:= $cColumns.GetByCmpnt "rtext"}-}
    onRtext{-{$c.Name}-}Change(text){
      this.{-{$c.Ext.FormName}-}.{-{$c.Name}-} = text
    },
    {-{- end}-}
    onUploadSuccess(response){
      {-{- range $i,$c:= $cColumns }-}
      {-{- if eq "file" $c.Cmpnt.Type  }-}
      {-{- if eq true $c.Cmpnt.IsStatic}-}
      this.form.{-{$c.Cmpnt.StaticValue}-} = response.path
      {-{- else}-}
      this.form.{-{$c.Name}-} = response.path
      {-{- end}-}
      {-{- end }-}
      {-{- end}-}
    },
    {-{- range $j,$x:=$cColumns}-}
    {-{- $ctriger := $cColumns.GetColumnsByTriggerChangeEvent $x.Name}-}
     {-{- if gt (len $ctriger) 0}-}
    onChange_{-{$x.Name}-}(val,tp){
      {-{- range $i,$c:= $ctriger}-}
      {-{- if eq true ($c.IsAssctColumn $x.Name)}-}
      if (!val) {
        this.{-{$c.Name}-}List = []
        this.form.{-{$c.Name}-} = null
        return
      }
      this.{-{$c.Name}-}List = this.$theia.enum.get("{-{$c.Enum.EnumType}-}",val,{-{if $c.Enum.GroupIsStatic}-} this.$theia.user.get("{-{$c.Enum.GetGroupValue}-}"){-{else}-}"{-{$c.Enum.Group}-}" {-{end}-})||[]
      this.form.{-{$c.Name}-} =this.{-{$c.Name}-}List.length>0?(this.{-{$c.Name}-}List[0]||{}).value:null
      {-{- $cxtriger := $cColumns.GetColumnsByTriggerChangeEvent $c.Name}-}
      {-{- if gt (len $cxtriger) 0}-}
      //cxtriger:{-{ $cxtriger}-}
      {-{- range $mx,$cx := $cxtriger}-}
      this.onChange_{-{$c.Name}-}(this.form.{-{$c.Name}-},"{-{$cx.Enum.EnumType}-}")
      {-{- end}-}
      {-{- end}-}
      {-{- end}-}
      {-{- end}-}
      
      {-{- $p := $x.GetParamMap "@change"}-}
      {-{- if gt (len $p) 0}-}
      //添加联动
      let that = this
      if(tp){
        let lst= this.$theia.enum.get(tp)
        var hasValues = {}
        lst.forEach(el => {
          if(el.value == val)
          {-{- range $i,$c:= $p}-}
          if (!that.form.{-{$i}-} || that.form._{-{$i}-} == that.form.{-{$i}-}) {
            hasValues.{-{$i}-} = true
            that.form.{-{$i}-} = el.{-{$c}-}
            that.form._{-{$i}-} = el.{-{$c}-}
          }
          {-{- end}-}
        });
        {-{- range $i,$c:= $p}-}
        if(!hasValues.{-{$i}-}){
          that.form.{-{$i}-} = ""
          that.form._{-{$i}-} = ""
        }
        {-{- end}-}
      }
      {-{- end}-}
    },
    {-{- end}-}
    {-{- end}-}
    {-{- $st := $table.NewScene $cColumns}-}
    {-{- template "enum.tmpl.js" $st }-}
},
}
</script>

<style scoped>
.avatar-uploader .avatar {
  width: 80px;
  height: 80px;
  display: block;
}
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