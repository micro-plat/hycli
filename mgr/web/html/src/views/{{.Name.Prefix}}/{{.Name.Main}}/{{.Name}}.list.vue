<template tag="{-{.Marker}-}">
  {-{- $table := .}-}
  {-{- $opts :=$table.Optrs.ListOpts }-}
  {-{- $tbs := $table.Contact $table }-}
  {-{- $qcols :=  $table.Columns.GetFontQueryColumns}-}
  {-{- $ddmenuCols :=  $table.Columns.GetFontQueryColumns.GetEnumColumns.GetByCmpnt "ddmenu"}-}
  <div style="height: 100%">
    <div style="margin:0.8rem;"><h5 style="display:inline">{-{$table.Desc}-}</h5><span style="margin-left: 0.5rem; color:#999"> {-{$table.Conf.Get "desc"}-}</span>
      {-{- range $i,$c:= $ddmenuCols}-}
        <ddmenu :menuList="{-{.Name}-}List" @valueChanged="on{-{$c.Name}-}Change" v-model="form_{-{$table.UNQ}-}.{-{$c.Name}-}" menuType="{-{$c.Enum.EnumType}-}" ></ddmenu>
        {-{- end  }-}
    </div>
    <hr style="margin-top:0;color:#999"/>
    {-{- template "query.tmpl.html" $table }-}
    {-{- template "listbar.tmpl.html" $table }-}
    {-{- template "cmpnt.t.tmpl.html" $table }-}
    {-{- template "list.tmpl.html" $tbs }-}
  </div>
</template>
<script>
{-{- template "cmpnt.i.tmpl.js" $table }-}
export default {
  {-{- template "cmpnt.c.tmpl.js" $table }-}
data() {
  return {
    shortcuts: [
        {
          text: '今天',
          value: [new Date(),new Date()],
        },
        {
          text: '昨天',
          value: () => {
            const date = new Date()
            date.setTime(date.getTime() - 3600 * 1000 * 24)
            return [date,new Date()]
          },
        },
        {
          text: '本周',
          value: () => {
            const now = new Date()
            return [new Date(now.getFullYear(), now.getMonth(), now.getDate() - now.getDay()),new Date()]
          },
        },
        {
          text: '本月',
          value: () => {
            const now = new Date()
            return [new Date(now.getFullYear(), now.getMonth(), 1),new Date()]
          },
        },
      ],
    cmpnt_funcs:{},
    {-{- if eq $table.NeedBatchCheck true}-}
    bcheck:[],
    {-{- end}-}
    conf: {
      loading: false,
      progressColor: this.$theia.env.conf.progress || []
    },
    ganttIdx: -1,
      {-{- template "queryform.tmpl.js" $table }-}
  {-{- template "listbar.tmpl.js" $table }-}
};
  },
mounted() {
  this.loadEnums()
  {-{- range $i,$c:= $ddmenuCols}-}
  {-{- $memberClus := $table.Columns.GetFontQueryColumns.GetStaticColumns}-}
  {-{- if gt $memberClus.Len 0}-}
  {-{- range $i,$v := $memberClus}-}
  this.form_{-{$table.UNQ}-}.{-{$c.Name}-} = this.$theia.user.get("{-{$v.Ext.Name}-}")
  {-{- end}-}
  {-{- end}-}
  {-{- end}-}
  this.form_{-{$table.UNQ}-}.single_date_range_name = (this.multiQueryDateRange[0]||{}).value
  this.form_{-{$table.UNQ}-}.single_text_name = (this.multiQueryText[0]||{}).value
  {-{- range $x,$m:= $table.Optrs.All.GetCompnents}-}
  this.cmpnt_funcs["{-{$m.UNQ}-}"] = this.show_cmpnt_{-{$m.UNQ}-}
  {-{- end}-}

  this.form_{-{$table.UNQ}-} = Object.assign(this.form_{-{$table.UNQ}-},this.$route.params)
  {-{- $rpColumns := $table.Columns.GetColumns "rp"}-}
  {-{- if eq $rpColumns.Len 0}-}
  this.queryData_{-{ $table.UNQ }-} ()
  {-{- end}-}
    },
  {-{- if gt $rpColumns.Len 0}-}
 watch: {
    '$route' () {
      this.form_{-{$table.UNQ}-} = Object.assign(this.form_{-{$table.UNQ}-},this.$route.params)
      this.onQuery();//我的初始化方法
    }
  },
  {-{- end}-}
methods: {
  ganttChange(id) {
      this.ganttIdx = id
    },
  {-{- template "list.tmpl.js" $tbs }-}
  handleSizeChange(ps){
    this.form_{-{ $table.UNQ }-}.ps = ps
    this.onQuery(true)
  },
  {-{- range $i,$c:= $ddmenuCols}-}
        {-{- $memberClus := $table.Columns.GetFontQueryColumns.GetStaticColumns}-}
        {-{- if gt $memberClus.Len 0}-}
        {-{- range $i,$v := $memberClus}-}
  on{-{$c.Name}-}Change(v){
    this.$theia.user.set("{-{$v.Ext.Name}-}",v)
    this.loadEnums()
    this.onQuery(true)
  },
{-{- end}-}
{-{- end}-}
{-{- end}-}
  handleCurrentChange(pi){
    this.form_{-{ $table.UNQ }-}.pi = pi
    this.onQuery()
  },
  onQuery(refresh){
    if(refresh){
      this.form_{-{ $table.UNQ }-}.pi = 1
    }
    this.queryData_{-{ $table.UNQ }-} ()
  },
  colorful(r, name){
    if (this.$theia.env.conf.colorful[name]) {
      return this.$theia.env.conf.colorful[name][r] || ""
    }
    return this.$theia.env.conf.colorful[r] || ""
  },
  tagColor(r, name){
    if (this.$theia.env.conf.colorTag[name]) {
      return this.$theia.env.conf.colorTag[name][r] || ""
    }
    return this.$theia.env.conf.colorTag[r] || ""
  },
  goto(url, param){
    if (!url) {
      return;
    }
    if (!url.indexOf("://")) {
      this.$router.push({ path: url, query: param });
      return;
    }
    let p = this.$theia.url.encode(param)
    window.location = `${url}?${p}`
  },
  {-{- template "opts.tmpl.js" $tbs }-}
},
};
</script>
<style>
.el-dialog{
  border-radius:5px !important
}
</style>
<style scoped>
.el-form-item {
  margin-right: 10px !important;
}

.opts .el-button--small {
  margin-left: 0px;
  padding-left: 0px;
  margin-right: 4px;
}

.blist {
  margin-bottom: 8px;
  width: 100%;
}

.blist .el-radio-group {
  margin-right: 8px;
}

.ddl {
  margin-bottom: 8px;
  width: 100%;

}

.ddl .el-dropdown {
  margin-right: 8px;
}

.listbar {
  margin-bottom: 8px;
  text-align: center;
}
 .el-dropdown-link {
  cursor: pointer;
  color: var(--el-color-primary);
  display: flex;
  align-items: center;
  font-size: 0.65rem;
  margin-top: 7px;
}
/deep/.el-button--small{
  padding: 5px 4px;
}

.prepend /deep/.el-input__wrapper{
  border-top-right-radius:0;
  border-bottom-right-radius:0;
}
.prepend+/deep/.el-date-editor{
  border-top-left-radius:0;
  border-bottom-left-radius:0;
  border-left:0;
}

</style>