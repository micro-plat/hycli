{-{- $table := . }-}
{-{- $vcols :=  $table.GetColumnsByName "v" }-}
{-{- $viewOpts :=$table.ViewOpts}-}
<template tag="{-{.Marker}-}">
  <div>
  <el-dialog
    v-model="conf.visible"
    title="{-{.Desc}-}详情"
    draggable
    width="68%"
    :close-on-click-modal="false"
    :before-close="hide"
  >
  {-{- range $i,$c:= $viewOpts.GetByName "step"}-}
  <el-steps class="steps" :active="conf.stepActive_{-{$c.UNQ}-}" align-center  finish-status="success">
    {-{- range $j,$s:=  $table.GetColumnsByName $c.RwName}-}
    <el-step title="{-{$s.Label}-}" :description="view.{-{$s.Name}-}||'未设置'" />
    {-{- end }-}
  </el-steps>
  {-{- end}-}
    <el-tabs v-model="conf.selected">
      {-{- range $i,$c:= $viewOpts}-}
      {-{- if eq "view" $c.Name}-}
      {-{- $currentTable := f_table_get_ttable  $c.Table }-}
      {-{- $vxcols :=  $currentTable.GetColumnsByName "v" }-}
      <el-tab-pane label="详情" name="{-{$c.UNQ}-}">
       {-{- template "view.tmpl.html" $vxcols}-}
      </el-tab-pane>
      {-{- else if eq "CMPNT" $c.Name}-}
      <el-tab-pane label="{-{$c.Label}-}" name="{-{$c.UNQ}-}"  @tab-click="show_view_{-{$c.UNQ}-}">
        <{-{$c.UNQ}-} ref="view_{-{$c.UNQ}-}"></{-{$c.UNQ}-}>
     </el-tab-pane>
       {-{- else if eq "TAB" $c.Name}-}
        <el-tab-pane label="{-{$c.Label}-}" name="{-{$c.UNQ}-}"  @tab-click="show_view_{-{$c.UNQ}-}">
          {-{- $tlbar := $table.ViewExtCmptOpts.GetBarOptrs  $c.Table $c.UNQ}-}
          {-{- if gt (len $tlbar) 0}-}
          <el-row>
            <el-col :span="24" class="text-right">
          {-{- range $i,$x:= $tlbar}-}
          <el-button type="success" icon="Plus"
           round  @click="show_cmpnt_{-{$x.UNQ}-}">{-{ $x.Label}-}</el-button>
         {-{- end}-}
        </el-col>
          </el-row>
         {-{- end}-}
         {-{- $ct:=  $c.GetAssociatedTable true}-}
         {-{- $tbs := $ct.Contact $table }-}
        {-{- template "list.tmpl.html" $tbs}-}
         </el-tab-pane>
      {-{- end}-}
      {-{- end }-}
    </el-tabs>
    <template #footer>
      <span style="height: 60px"> </span>
    </template>
    {-{- range $x,$m:= $table.ViewExtCmptOpts.ALL}-}
    {-{- if eq "CMPNT" $m.Name  }-}
    <{-{$m.UNQ}-} ref="cmpnt_{-{$m.UNQ}-}" @onsaved="show(form)"></{-{$m.UNQ}-}>
    {-{- end}-}
    {-{- end}-}
 
  </el-dialog>
</div>
</template>
<script>
{-{- $vaopts := $viewOpts.Merge $table.ViewExtCmptOpts.ALL}-}
{-{- $vaCmpntOpts :=  $vaopts.GetByName "CMPNT"}-}
 {-{- range $x,$m := $vaCmpntOpts}-}
import {-{$m.UNQ}-} from "{-{f_string_translate $m.URL (f_table_find_by_name $m.Table)}-}"
{-{- end}-}

export default {
   components: {

    {-{- range $x,$m:= $vaCmpntOpts}-}
    {-{$m.UNQ}-},
    {-{- end }-}
  },
  data() {
    return {
        conf:{
        visible:false,
        selected:"{-{($viewOpts.GetSelected).UNQ}-}",
        {-{- range $i,$c:= $viewOpts.GetByName "step"}-}
        stepActive_{-{$c.UNQ}-}:0,
        {-{- end}-}
      },
      form:{},
       {-{- range $i,$c:=  $viewOpts }-}
       {-{- if eq "TAB" $c.Name}-}
          {-{- $ct:=   $c.GetAssociatedTable true}-}
        {-{- template "queryform.tmpl.js" $ct}-}
      {-{- end}-}
      {-{- end }-}
      view: {
        {-{- range $i,$c := $vcols}-}
        {-{$c.Name}-}:"",
        {-{- end}-}
        },
      }
  },
  methods: {
   {-{- range $x,$m:= $table.ViewExtCmptOpts.ALL.GetByName "CMPNT"}-}
  show_cmpnt_{-{$m.UNQ}-}(fm = {}){
    let form = Object.assign({},this.form)
    form = Object.assign(form,fm)
 
    {-{- $c :=  $m.GetParamsByAtPrefix }-}
    {-{- range $k,$v := $c}-}
    {-{- if eq true (f_string_start $v "@")}-}
    form.{-{$k}-} = this.{-{f_string_trim $v "@"}-}
    {-{- else if eq true (f_string_start $v "&")}-}
    form.{-{$k}-} = (this.{-{f_string_trim $v "&"}-}||[]).join(",")
    {-{- else if eq true (f_string_start $v "#")}-}
    form.{-{$k}-} = fm.{-{f_string_trim $v "#"}-}
    {-{- else}-}
    form.{-{$k}-} = "{-{$v}-}"
    {-{- end}-}
    {-{- end}-}

    this.$refs.cmpnt_{-{$m.UNQ}-}.show(form)
  },
{-{- end}-}
    {-{- template "view.tmpl.js" $table}-}
     {-{- range $i,$c:= $viewOpts }-}
       {-{- if eq "TAB" $c.Name }-}
          {-{- $ct:=   $c.GetAssociatedTable true }-}
          {-{- $tbs := $ct.Contact $table }-}
        {-{- template "list.tmpl.js" $tbs}-}
      {-{- end }-}
      {-{- end }-}
  },
};
</script>
<style scoped>
/deep/.el-step__title{
    font-size: 0.8rem;
}
.steps{
  margin-top:8px;
  margin-bottom:  16px;
}
</style>