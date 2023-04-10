<template tag="{-{.Marker}-}">
  {-{- $table := .}-}
  {-{- $opts :=$table.ListOpts }-}
  <div style="height: 100%">
    {-{- template "query.tmpl.html" $table }-}
    {-{- template "listbar.tmpl.html" $table }-}
    {-{- template "cmpnt.t.tmpl.html" $table }-}
    {-{- template "list.tmpl.html" $table }-}
  </div>
</template>
<script>
{-{- template "cmpnt.i.tmpl.js" $table }-}
export default {
  {-{- template "cmpnt.c.tmpl.js" $table }-}
data() {
  return {
    conf: {
      loading: false,
      progressColor: this.$theia.env.conf.progress || []
    },
      {-{- template "queryform.tmpl.js" $table }-}
  {-{- template "listbar.tmpl.js" $table }-}
};
  },
mounted() {
  this.queryData_{-{ $table.UNQ }-} ()
      {-{- range $i, $c:=  $table.ChartOpts }-}
  this.$refs.chart_{-{ $c.UNQ }-}.show(this.form_{-{ $table.UNQ }-})
{-{- end }-}
    },
methods: {
  {-{- $tbs := contactTBS  $table $table }-}
  {-{- template "list.tmpl.js" $tbs }-}
  handleSizeChange(ps){
    this.form_{-{ $table.UNQ }-}.ps = ps
    this.queryData_{-{ $table.UNQ }-} ()
  },
  handleCurrentChange(pi){
    this.form_{-{ $table.UNQ }-}.pi = pi
    this.queryData_{-{ $table.UNQ }-} ()
  },
  onQuery(){
    this.form_{-{ $table.UNQ }-}.pi = 1
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
  {-{- template "opts.tmpl.js" $table }-}
},
};
</script>
<style>
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
}</style>