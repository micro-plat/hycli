<template tag="{-{.Marker}-}">
  {-{- $table := .}-}
  {-{- $opts :=$table.ListOpts }-}
  {-{- $tbs := $table.Contact $table }-}
  {-{- $qcols :=  $table.GetColumnsByTPName "q"}-}
  <div style="height: 100%">
    <div style="margin:0.8rem;"><h5 style="display:inline">{-{$table.Desc}-}</h5><span style="margin-left: 0.5rem; color:#999"> {-{$table.Conf.Get "desc"}-}</span>
      {-{- range $i,$c:= $qcols}-}
        {-{- if and (eq "ddmenu" $c.Cmpnt.Type) (eq true $c.Enum.IsEnum)}-}
        <ddmenu :menuList="{-{.Name}-}List" @valueChanged="on{-{$c.Name}-}Change" v-model="form_{-{$table.UNQ}-}.{-{$c.Name}-}" menuType="{-{$c.Enum.EnumType}-}" ></ddmenu>
        {-{- end  }-}
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
  {-{- range $i,$c:= $qcols}-}
        {-{- if and (eq "ddmenu" $c.Cmpnt.Type) (eq true $c.Enum.IsEnum)}-}
        {-{- $memberClus :=  $table.GetStaticColumns "q" "#"}-}
        {-{- if gt (len $memberClus) 0}-}
        {-{- range $i,$v := $memberClus}-}
      {-{- $name := f_string_trim $i "#"}-}
      this.form_{-{$table.UNQ}-}.{-{$c.Name}-}=this.$theia.user.get("{-{$c.Name}-}")
      {-{- end}-}
       {-{- end}-}
        {-{- end}-}
    {-{- end}-}


  this.form_{-{$table.UNQ}-}.single_date_range_name = (this.multiQueryDateRange[0]||{}).value
  this.form_{-{$table.UNQ}-}.single_text_name = (this.multiQueryText[0]||{}).value
  {-{- $optRow:= $table.ListOpts.Merge $table.BarOpts}-}
  {-{- $cmpnts:=  $optRow.GetByName "CMPNT"}-}
  {-{- range $x,$m:= $cmpnts}-}
    this.cmpnt_funcs["{-{$m.UNQ}-}"] = this.show_cmpnt_{-{$m.UNQ}-}
  {-{- end}-}

  this.form_{-{$table.UNQ}-} = Object.assign(this.form_{-{$table.UNQ}-},this.$route.params)
  this.queryData_{-{ $table.UNQ }-} ()
      {-{- range $i, $c:=  $table.ChartOpts }-}
  this.$refs.chart_{-{ $c.UNQ }-}.show(this.form_{-{ $table.UNQ }-})
{-{- end }-}
{-{- range $i, $c:=  $table.ExtOpts }-}
  this.$refs.ext_{-{ $c.UNQ }-}.show(this.form_{-{ $table.UNQ }-})
{-{- end }-}

    },
  {-{- if ne "" ( $table.JoinNames  "rp" false "")}-}
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
    this.queryData_{-{ $table.UNQ }-} ()
  },

  {-{- range $i,$c:= $qcols}-}
        {-{- if and (eq "ddmenu" $c.Cmpnt.Type) (eq true $c.Enum.IsEnum)}-}
        {-{- $memberClus :=  $table.GetStaticColumns "q" "#"}-}
        {-{- if gt (len $memberClus) 0}-}
        {-{- range $i,$v := $memberClus}-}
      {-{- $name := f_string_trim $i "#"}-}
  on{-{$c.Name}-}Change(v){
    this.$theia.user.set("{-{$name}-}",v)
    this.onQuery(true)
  },
      {-{- end}-}
       {-{- end}-}
        {-{- end}-}
    {-{- end}-}

  handleCurrentChange(pi){
    this.form_{-{ $table.UNQ }-}.pi = pi
    this.queryData_{-{ $table.UNQ }-} ()
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