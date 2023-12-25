{-{- $table := .}-}
{-{- $optCols:=  $table.Optrs.All.GetDialog}-}
<template tag="{-{.Marker}-}">
  <div>
    {-{- range $x,$m:=$optCols }-}
    {-{- $width:= $m.GetParam "width" "60%" }-}
    {-{- $rows:=  $table.Columns.GetColumns $m.RwName (f_string_contact "form_" $m.UNQ)}-}
   
    <!-- {-{$m.Label}-} -->
    <el-dialog v-model="conf.{-{$m.UNQ}-}_visible" 
      {-{ if $m.HasParam "#clearEnum"  }-}v-if="conf.{-{$m.UNQ}-}_visible" {-{ end }-}  draggable
      title="{-{$m.Label}-}" :close-on-click-modal="false"  width="{-{$width}-}"
      {-{ if gt (len $rows) 9 }-}  align-center="true" {-{ else}-}top="10vh" {-{ end }-}
      :before-close="hide_{-{$m.UNQ}-}">
    <el-form  ref="fm_{-{$m.UNQ}-}" :model="form_{-{$m.UNQ}-}" :rules="rules_{-{$m.UNQ}-}">
      {-{- template "add.tmpl.html" $rows}-}
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="hide_{-{$m.UNQ}-}" icon="close">取消</el-button>
        <el-button type="primary" @click="save_{-{$m.UNQ}-}" icon="select">提交</el-button>
      </span>
    </template>
    </el-dialog>
    {-{- end }-}
  </div>
</template>

<script>
export default {
  data(){
    return{
      conf:{
        {-{- range $x,$m:=$optCols }-}
        {-{$m.UNQ}-}_visible:false,//{-{$m.Label}-}
        {-{- end}-}
      },
      {-{- range $x,$m:=$optCols }-}
      {-{- $rows:=  $table.Columns.GetEnumColumns.GetColumns $m.RwName }-}
      {-{- range $i,$c:=$rows }-} 
      {-{- if $c.Cmpnt.StartWith "tree|cascader|multi"}-}
      {-{.Name}-}List:[],
      {-{$c.Name}-}:"",
      {-{- end }-}
      {-{- end}-}
      form_{-{$m.UNQ}-}:{},//{-{$m.Label}-}
      {-{- end}-}
      {-{- range $x,$m:=$optCols }-}
      //{-{$m.Label}-}
      rules_{-{$m.UNQ}-}:{
        {-{- $rows:=  $table.Columns.GetColumns $m.RwName }-}
        {-{- range $i,$c:=$rows}-} 
        {-{$c.Name}-}:[{required:{-{$c.Field.Required}-}, message:"请输入{-{$c.Label}-}", trigger: 'blur'}],
        {-{- end}-}
      },
      {-{- end}-}
    }
  },
  methods:{
     {-{- range $x,$m:=$optCols }-}
     //--------------------{-{$m.Label}-}---------------------------------
      //显示 {-{$m.Label}-} 弹出框
      show_{-{$m.UNQ}-}(fm){
        //处理输入参数
        {-{- $ct:= $m.GetAssociatedTable }-}
        {-{- $tbs := $table.Contact $ct }-}
        {-{- $ctable := $tbs.Current }-}
        {-{- $mtable := $tbs.Main }-}
        {-{- $MLLERows:=  $mtable.Columns.GetFontListColumns}-}
        let currentForm = {}
        {-{- range $i,$c := $MLLERows }-}
        {-{- if $c.Enum.IsEnum }-}
        {-{- if eq $ctable.Enum.EnumType $c.Enum.EnumType}-}
        currentForm.{-{$table.Enum.Id}-} = fm.{-{$c.Name}-}
        {-{- end }-}
        {-{- end }-}
        {-{- end}-}
        Object.assign(this.form_{-{$m.UNQ}-},currentForm)
        {-{- range $i,$c := $ct.Columns.GetColumns $m.RwName}-}
        this.form_{-{$m.UNQ}-}.{-{$c.Name}-} = fm.{-{$c.Name}-}
        {-{- end}-}
        {-{- range $i,$c := $ct.Columns.GetColumns $m.FwName}-}
        this.form_{-{$m.UNQ}-}.{-{$c.Name}-} = fm.{-{$c.Name}-}
        {-{- end}-}
        //特殊参数处理
        {-{- $c :=  $m.GetParamsByAtPrefix }-}
        {-{- range $k,$v := $c}-}
        {-{- if eq true (f_string_start $v "@")}-}
        this.form_{-{$m.UNQ}-}.{-{$k}-} = this.{-{f_string_trim $v "@"}-}
        {-{- else if eq true (f_string_start $v "&")}-}
        // this.form_{-{$m.UNQ}-}.{-{$k}-} = (this.{-{f_string_trim $v "&"}-}||[]).join(",")
        {-{- else if eq true (f_string_start $v "#")}-}
        this.form_{-{$m.UNQ}-}.{-{$k}-} = fm.{-{f_string_trim $v "#"}-}
        {-{- else if eq true (f_string_start $v "!")}-}
        {-{- $xv:=f_string_trim $v "!"}-}
        query.{-{$k}-} = this.$theia.user.get("{-{$xv}-}")
        {-{- else}-}
        this.form_{-{$m.UNQ}-}.{-{$k}-} = "{-{$v}-}"
        {-{- end}-}
        {-{- end}-}
        this.loadEnums_{-{$m.UNQ}-}()
        this.conf.{-{$m.UNQ}-}_visible = true;
      },

      //隐藏 {-{$m.Label}-} 弹出框
      hide_{-{$m.UNQ}-}(){
        this.conf.{-{$m.UNQ}-}_visible = false;
        this.$refs.fm_{-{$m.UNQ}-}.resetFields();
      },

      //保存 {-{$m.Label}-} 弹出框数据
      save_{-{$m.UNQ}-}(){
        let that = this
        this.$refs.fm_{-{$m.UNQ}-}.validate((v=>{
          if(!v) return
          let post_form_{-{$m.UNQ}-} = this.form_{-{$m.UNQ}-}
          {-{- $rows:=  $table.Columns.GetColumns $m.RwName (f_string_contact "form_" $m.UNQ)}-}
          {-{- range $i,$c:= $rows }-}
          {-{- if $c.Cmpnt.Equal "password" }-}
          post_form_{-{$m.UNQ}-}.{-{$c.Name}-} = this.$theia.crypto.md5(this.form_{-{$m.UNQ}-}.{-{$c.Name}-})
          {-{- else if $c.Cmpnt.Equal "tree" }-}
          post_form_{-{$m.UNQ}-}.{-{$c.Name}-} = this.$refs.tree_{-{$c.UNQ}-}.getCheckedKeys().join(",")
          {-{- else if $c.Cmpnt.Equal "cascader" }-}
          let {-{$c.Name}-} = []
          let {-{$c.Name}-}Nodes = this.$refs.cascader_{-{$c.UNQ}-}.getCheckedNodes() || []
          {-{$c.Name}-}Nodes.forEach(v => {
            {-{$c.Name}-}.push(v.value)
          })
          post_form_{-{$m.UNQ}-}.{-{$c.Name}-} = {-{$c.Name}-}.join(",")
          {-{- else if $c.Cmpnt.StartWith "multi"}-}
          post_form_{-{$m.UNQ}-}.{-{$c.Name}-} = (post_form_{-{$m.UNQ}-}.{-{$c.Name}-}||[]).join(",")
          {-{- end }-}
          {-{- end}-}
          {-{- $fxname := f_string_contact "f" $m.RwName}-}
          {-{- $memberClus :=  ($table.Columns.GetColumns $fxname).GetStaticColumns}-}
          {-{- range $i,$v := $memberClus}-}
          post_form_{-{$m.UNQ}-}.{-{$v.Name}-} = this.$theia.user.get("{-{$v.Ext.Name}-}")
          {-{- end}-}
          this.$theia.http.post("/{-{$table.Name.MainPath|lower}-}/{-{f_string_translate $m.ReqURL (f_table_find_by_name $m.Table)}-}",post_form_{-{$m.UNQ}-}).then(res=>{
            that.$notify.success({title: '成功',message: '提交成功',duration:5000})
            that.$emit("onsaved")
            that.hide_{-{$m.UNQ}-}()
            {-{- range $k,$v :=  $m.GetEventByAtPrefix}-}
            {-{- if eq $k "clearEnum"}-}
            that.$theia.enum.clear("{-{$v}-}")
            {-{- end}-}
            {-{- end}-}
          }).catch(res=>{
            let code = ((res||{}).response||{}).status||0
            let msg= `提交失败(${code})`
            that.$notify.error({title: '失败',message:msg,duration:5000})
          })
      }))
    },
    {-{- end }-}
    {-{- range $x,$m:=$optCols }-}
    //{-{$m.Label}-}
    {-{- $rows:=  $table.Columns.GetEnumColumns.GetColumns $m.RwName }-}
    {-{- $st := $m.NewScene $rows}-}
    {-{- template "enum.tmpl.js" $st }-}
    {-{- end}-}
  },
}
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