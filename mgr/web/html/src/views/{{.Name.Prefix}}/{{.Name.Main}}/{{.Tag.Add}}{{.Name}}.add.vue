{-{- $table := .}-}
{-{- $cColumns :=  $table.GetColumnsByTPName "C"}-}
{-{- $acColumns :=  $table.GetColumnsByTPName "C-BC"}-}
{-{- $enumColumns :=$table.EnumColumns}-}

<template tag="{-{.Marker}-}">
  <el-dialog
    v-model="conf.visible"
    title="添加 {-{.Desc}-}"
   :width="conf.width"
    draggable
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
        visible: false,
        {-{- if gt (len $cColumns) 14}-}
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
    {-{- template "add.tmpl.js" $cColumns }-}
  },
  methods: {
    show(fm = {}) {
      this.conf.visible = true;
      this.loadEnums()
      let local = {}
      {-{- range $i,$c:= $table.BarOpts.GetByTag "ADD" }-}
        {-{- $read2window := $c.GetParams "read2window" -}-}
        {-{- if ne "" $read2window}-}
        
        //将数据保存到window缓存中
      local = window.{-{$read2window}-} || {}
        {-{- end -}-}
        {-{- end }-}
      let cache =  Object.assign(local,fm)
      this.form = Object.assign(cache,this.$route.params)

      {-{- range $i,$c := $acColumns.GetColumnsBy "alias"}-}
      {-{- $name := $c.GetOpt "alias" -}-}
      {-{- if ne "" $name}-}
      this.form.{-{$c.Name}-} = cache["{-{$name}-}"]
      {-{- end -}-}
      {-{- end -}-}

      {-{- range $i,$c := $cColumns}-}
      {-{- if eq true (f_string_start $c.Cmpnt.Type "multi")}-}
      this.form.{-{$c.Name}-} = this.form.{-{$c.Name}-}?this.form.{-{$c.Name}-}:[];
      {-{- end}-}
      {-{- end}-}

      {-{- range $j,$x:=$cColumns}-}
      {-{- range $i,$c:= $cColumns.GetColumnsByTriggerChangeEvent $x.Name}-}
      this.onChange_{-{$x.Name}-}(this.$route.params["{-{$x.Name}-}"]||this.form["{-{$x.Name}-}"],"{-{$x.Enum.EnumType}-}")
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
        {-{- range $i,$c:= $table.BarOpts.GetByTag "ADD" }-}
        {-{- $save2window := $c.GetParams "save2window" -}-}
        {-{- if ne "" $save2window}-}
        
        //将数据保存到window缓存中
        window.{-{$save2window}-} = postForm
        {-{- end -}-}
        {-{- end }-}


        {-{- $memberClus :=  $table.GetStaticColumns "fc" "#"}-}
      {-{- range $i,$v := $memberClus}-}
      {-{- $name := f_string_trim $i "#"}-}
        postForm.{-{$v.Name}-} = this.$theia.user.get("{-{$name}-}")
      {-{- end}-}

        //保存数据
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
          msg = code == 909? msg+"数据重复，请修改后重试":msg
          that.$notify.error({title: '失败',message:msg,duration:5000})
        })
    },
    hide() {
      this.conf.visible = false;
      this.$refs.form.resetFields();
    },
    {-{- range $i,$c:= $cColumns }-}
    {-{-  if (eq "rtext" $c.Cmpnt.Type )}-}
    onRtext{-{$c.Name}-}Change(text){
      this.{-{$c.Ext.FormName}-}.{-{$c.Name}-} = text
    },
    {-{- end}-}
    {-{- end}-}
    onUploadSuccess(response){
      {-{- range $i,$c:= $cColumns }-}
      {-{- if eq "file" $c.Cmpnt.Type  }-}
      {-{- if eq true (f_string_start $c.Cmpnt.Format "#")}-}
      this.form.{-{f_string_trim $c.Cmpnt.Format "#"}-} = response.path
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
      {-{- $group:= f_string_trim $c.Enum.Group "#"}-}
      this.{-{$c.Name}-}List = this.$theia.enum.get("{-{$c.Enum.EnumType}-}",val,{-{if f_string_start $c.Enum.Group "#"}-}this.$theia.user.get("{-{$group}-}"){-{else}-}"{-{$c.Enum.Group}-}" {-{end}-})
      this.form.{-{$c.Name}-} = null
      {-{- end}-}
     
      {-{- end}-}
      //添加联动
      {-{- $p := $x.GetParams "@change"}-}
      {-{- if gt (len $p) 0}-}
      let that = this
      if(tp){
        let lst= this.$theia.enum.get(tp)
        lst.forEach(el => {
          if(el.value == val)
          {-{- range $i,$c:= $p}-}
          that.form.{-{$i}-} = that.form.{-{$i}-}?that.form.{-{$i}-}:el.{-{$c}-}
          {-{- end}-}
        });
      }
      {-{- end}-}

    },
    {-{- end}-}
    {-{- end}-}
    {-{- template "enum.tmpl.js" $cColumns }-}
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