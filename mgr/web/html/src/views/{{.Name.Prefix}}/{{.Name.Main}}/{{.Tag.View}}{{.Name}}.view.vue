{-{- $table := . }-}
{-{- $vcols := fltrColumns $table "v" }-}
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
  {-{- range $i,$c:=fltrOptrs $viewOpts "step"}-}
  <el-steps class="steps" :active="conf.stepActive" align-center size="small" finish-status="success">
    {-{ range $j,$s:= fltrColumns $table $c.RwName}-}
    <el-step title="{-{$s.Label}-}" :description="view.{-{$s.Name}-}||'未设置'" />
    {-{- end }-}
  </el-steps>
  {-{- end}-}

    <el-tabs v-model="conf.selected">
      <el-tab-pane label="详情" name="first">
        {-{- template "view.tmpl.html" $vcols}-}
      </el-tab-pane>
      {-{- range $i,$c:= $viewOpts}-}
      {-{- if eq "CMPNT" $c.Name}-}
      <el-tab-pane label="{-{$c.Label}-}" name="{-{$c.Label}-}"  @tab-click="show_view_{-{$c.UNQ}-}">
         <{-{$c.UNQ}-} ref="view_{-{$c.UNQ}-}"></{-{$c.UNQ}-}>
      </el-tab-pane>
       {-{- else if eq "TAB" $c.Name}-}
        <el-tab-pane label="{-{$c.Label}-}" name="{-{$c.Label}-}"  @tab-click="show_view_{-{$c.UNQ}-}">
      
          {-{- $tlbar :=fltrOptrsByPUNQ $table.ViewExtCmptOpts  $c.UNQ }-}
           <!-- {-{$c.Table  }-},{-{ len $tlbar }-}-->
          {-{- if gt (len $tlbar) 0}-}
          <el-row>
            <el-col :span="24" class="text-right">
          {-{- range $i,$x:= $tlbar}-}
          <el-button type="success" icon="Plus"
           round size="small"
            @click="show_cmpnt_{-{$x.UNQ}-}">{-{ $x.Label}-}</el-button>
         {-{- end}-}
        </el-col>
          </el-row>
         {-{- end}-}
         {-{- $ct:= fltrSearchUITableAndResetUNQ $c}-}
        {-{- template "list.tmpl.html" $ct}-}
         </el-tab-pane>
      {-{- end}-}
      {-{- end }-}
    </el-tabs>
    <template #footer>
      <span style="height: 60px"> </span>
    </template>
    {-{- range $x,$m:= $table.ViewExtCmptOpts}-}
    {-{- if eq "CMPNT" $m.Name  }-}
    <{-{$m.UNQ}-} ref="cmpnt_{-{$m.UNQ}-}" @onsaved="show(form)"></{-{$m.UNQ}-}>
    {-{- end}-}
    {-{- end}-}
 
  </el-dialog>
</div>
</template>
<script>
 {-{- range $x,$m:=fltrOptrs (mergeOptrs $viewOpts $table.ViewExtCmptOpts) "CMPNT"}-}
import {-{$m.UNQ}-} from "{-{fltrTranslate $m.URL (fltrFindTable $m.Table)}-}"
{-{- end}-}

export default {
   components: {
    {-{- range $x,$m:=fltrOptrs (mergeOptrs $viewOpts $table.ViewExtCmptOpts) "CMPNT"}-}
    {-{$m.UNQ}-},
    {-{- end }-}
  },
  data() {
    return {
        conf:{
        visible:false,
        selected:"first",
        stepActive:0,
      },
      form:{},
       {-{- range $i,$c:=  $viewOpts }-}
       {-{- if eq "TAB" $c.Name}-}
          {-{- $ct:= fltrSearchUITableAndResetUNQ  $c}-}
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
   {-{- range $x,$m:=fltrOptrs $table.ViewExtCmptOpts "CMPNT"}-}
  show_cmpnt_{-{$m.UNQ}-}(){
    let form = Object.assign({},this.form)
    {-{- $c := fltrOptrsByStatic $m }-}
    {-{- range $k,$v := $c}-}
    form.{-{$k}-}="{-{$v}-}"
    {-{- end}-}
    this.$refs.cmpnt_{-{$m.UNQ}-}.show(form)
  },
{-{- end}-}
    {-{- template "view.tmpl.js" $table}-}
     {-{- range $i,$c:= $viewOpts }-}
       {-{- if eq "TAB" $c.Name }-}
          {-{- $ct:= fltrSearchUITableAndResetUNQ  $c }-}
          {-{- $tbs := contactTBS  $ct $table }-}
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