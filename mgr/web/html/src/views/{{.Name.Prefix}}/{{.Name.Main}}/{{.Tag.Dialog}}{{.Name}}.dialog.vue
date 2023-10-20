{-{- $table := .}-}
{-{- $optCols:=  $table.ListOpts.Merge $table.BarOpts}-}
<template tag="{-{.Marker}-}">
  <div>
    {-{- range $x,$m:=$optCols }-}
    {-{- if eq "DIALOG" $m.Tag }-}
    <el-dialog
      v-model="conf.{-{$m.UNQ}-}_visible"
      title="{-{.Label}-}"
      :close-on-click-modal="false"
      width="60%"
      draggable
      :before-close="hide_{-{$m.UNQ}-}"
    >
      {-{- $rows:=  $table.GetColumnsByTPName $m.RwName (f_string_contact "form_" $m.UNQ)}-}
      <el-form
        :model="form_{-{$m.UNQ}-}"
        
        ref="fm_{-{$m.UNQ}-}"
        :rules="rules_{-{$m.UNQ}-}"
      >
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
    {-{- end }-}
  </div>
</template>

<script>
export default {
  data(){
    return{
      conf:{
        {-{- range $x,$m:=$optCols }-}
        {-{- if eq "DIALOG" $m.Tag }-}
        {-{$m.UNQ}-}_visible:false,
        {-{- end}-}
        {-{- end}-}
      },
  {-{- range $x,$m:=$optCols }-}
     {-{- if eq "DIALOG" $m.Tag}-}
      //{-{$m.Label}-} form by  {-{$m.RwName}-}
      {-{- $rows:=  $table.GetColumnsByTPName $m.RwName }-}
        {-{- range $i,$c:=$rows }-} 
        {-{- if or (eq true $c.Enum.IsEnum) (eq true (f_string_start $c.Cmpnt.Type "multi"))}-}
    {-{.Name}-}List:this.$theia.enum.get("{-{$c.Enum.EnumType}-}","{-{$c.Enum.PID}-}","{-{$c.Enum.Group}-}"),
         {-{- else}-}
    {-{$c.Name}-}:"",
         {-{- end }-}
        {-{- end}-}
    form_{-{$m.UNQ}-}:{},
   {-{- end}-}
   {-{- end}-}
    {-{- range $x,$m:=$optCols }-}
    {-{- if eq "DIALOG" $m.Tag}-}
    rules_{-{$m.UNQ}-}:{
        {-{- $rows:=  $table.GetColumnsByTPName $m.RwName }-}
         {-{- range $i,$c:=$rows}-} 
          {-{$c.Name}-}:[{required:{-{$c.Field.Required}-}, message:"请输入{-{$c.Label}-}", trigger: 'blur'}],
          {-{- end}-}
        },
      {-{- end}-}
      {-{- end}-}
    }
  },
  methods:{
     {-{- range $x,$m:=$optCols }-}
     {-{- if eq "DIALOG" $m.Tag}-}
     //--------------------{-{$m.Label}-}---------------------------------
      //显示 {-{$m.Label}-} 弹出框 {-{$m}-}
      show_{-{$m.UNQ}-}(fm){
        {-{- $ct:=   $m.GetAssociatedTable }-}
        {-{- $tbs := $table.Contact $ct }-}
        {-{- $ctable := $tbs.Current }-}
        {-{- $mtable := $tbs.Main }-}
        {-{- $MLLERows:=  $mtable.GetColumnsByTPName "l-le"}-}
          
        //处理关联表{-{$m.URL}-}
        let currentForm = {}
        {-{- range $i,$c := $MLLERows }-}
          {-{- if eq true $c.Enum.IsEnum }-}
            {-{- if eq $ctable.Enum.EnumType $c.Enum.EnumType}-}
        currentForm.{-{$table.Enum.Id}-} = fm.{-{$c.Name}-}
            {-{- end }-}
        {-{- end }-}
        {-{- end}-}

         Object.assign(this.form_{-{$m.UNQ}-},currentForm)
         Object.assign(this.form_{-{$m.UNQ}-},fm)
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
            if(!v){
                return
            }
        let post_form_{-{$m.UNQ}-} = this.form_{-{$m.UNQ}-}
        {-{- $rows:=  $table.GetColumnsByTPName $m.RwName (f_string_contact "form_" $m.UNQ)}-}
        {-{range $i,$c:= $rows }-}
       
         {-{- if eq "password" $c.Cmpnt.Type  }-}
        post_form_{-{$m.UNQ}-}.{-{$c.Name}-} = this.$theia.crypto.md5(this.form_{-{$m.UNQ}-}.{-{$c.Name}-})
         {-{- end }-}
          {-{end}-}
        this.$theia.http.post("/{-{$table.Name.MainPath|lower}-}/{-{f_string_translate $m.ReqURL (f_table_find_by_name $m.Table)}-}",post_form_{-{$m.UNQ}-}).then(res=>{
            that.$notify.success({title: '成功',message: '提交成功',duration:5000})
            that.$emit("onsaved")
            that.hide_{-{$m.UNQ}-}()
            
        }).catch(res=>{
          let code = ((res||{}).response||{}).status||0
          let msg= `提交失败(${code})`
          that.$notify.error({title: '失败',message:msg,duration:5000})
        })
        }))
    },
    //-----------------------------------------------------------
    {-{- end }-}
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