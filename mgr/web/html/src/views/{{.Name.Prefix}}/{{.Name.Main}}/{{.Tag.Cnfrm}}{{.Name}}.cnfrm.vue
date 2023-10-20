{-{- $table := .}-}
{-{- $opts := $table.ListOpts.Merge $table.BarOpts}-}
<template tag="{-{.Marker}-}">
  <div>
    {-{- range $x,$m:=$opts }-}
    {-{- if eq "CNFRM" $m.Tag }-}
    <el-dialog
      v-model="conf.{-{$m.UNQ}-}_visible"
      title="{-{$m.Label}-}"
      width="30%"
      draggable
      :close-on-click-modal="false"
      :before-close="hide_{-{$m.UNQ}-}">
    {-{- $label:= $m.GetParam "label" ""}-}
      <span 
        >确认{-{$m.Label}-}{-{$label}-}吗?</span
      >
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="hide_{-{$m.UNQ}-}" icon="close">取消</el-button>
          <el-button type="primary" @click="save_{-{$m.UNQ}-}" icon="select">确定</el-button>
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
        {-{- range $x,$m:=$opts }-}
        {-{- if eq "CNFRM" $m.Tag }-}
        {-{$m.UNQ}-}_visible:false,
        {-{- end}-}
        {-{- end}-}
      },
  {-{- range $x,$m:=$opts }-}
     {-{- if eq "CNFRM" $m.Tag }-}
      //{-{$m.Label}-} form by  [{-{$m.RwName}-}]
      form_{-{$m.UNQ}-}:{
       {-{- $cols:=  $table.GetColumnsByTPName $m.RwName }-}
        {-{- range $i,$c:=$cols }-} 
        {-{$c.Name}-}:"",
        {-{- end}-}
         },
   {-{- end}-}
   {-{- end}-}
    }
  },
  methods:{
     {-{- range $x,$m:=$opts }-}
     {-{- if eq "CNFRM" $m.Tag }-}
     //--------------------{-{$m.Label}-}---------------------------------
      //显示 {-{$m.Label}-} 弹出框
      show_{-{$m.UNQ}-}(fm){
         Object.assign(this.form_{-{$m.UNQ}-},fm)
         this.conf.{-{$m.UNQ}-}_visible = true;
      },
      //隐藏 {-{$m.Label}-} 弹出框
      hide_{-{$m.UNQ}-}(){
        this.conf.{-{$m.UNQ}-}_visible = false;
        // this.$refs.fm_{-{$m.UNQ}-}.resetFields();
      },
      //保存 {-{$m.Label}-} 弹出框数据
      save_{-{$m.UNQ}-}(){
        let that = this
        {-{- $url := f_string_translate $m.ReqURL (f_table_find_by_name $m.Table)}-}
        {-{- if eq false (f_string_contains $url "/") }-}
        {-{- $url = f_string_contact "/" (lower $table.Name.MainPath) "/" (f_string_translate $m.ReqURL (f_table_find_by_name $m.Table)) }-}
        {-{- end}-}
        this.$theia.http.post("{-{$url}-}",this.form_{-{$m.UNQ}-}).then(res=>{
          that.conf.confirmVisible = false 
          that.$notify.success({title: '成功',message: '{-{$m.Label}-}成功',duration:5000})
          that.$emit("onsaved")
          that.hide_{-{$m.UNQ}-}()
       }).catch(err=>{
          that.conf.confirmVisible = false
          let code = err.response.status
          let msg= `{-{$m.Label}-}失败(${code})`
          that.$notify.error({title: '失败',message:msg,duration:5000})
       })
    },
    //-----------------------------------------------------------
    {-{- end }-}
      {-{- end}-}
  },
}
</script>