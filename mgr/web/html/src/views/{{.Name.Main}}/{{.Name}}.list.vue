<template>
  {{- $table := . -}}
  {{- $qrow := $table.QRows -}}
  {{- $lstRow := $table.LRows -}}
  {{- $leRow := $table.LERows -}}
  {{- $optRow :=$table.Optrs -}}
  {{- $delRow :=$table.DRows -}}
  {{- $LLERows:= mergeUIRow $lstRow $leRow}}
  <div style="height: 100%">
    <!-- 查询条件 -->
    {{template "query.tmpl.html" $table}}
    <!-- 列表 -->
    {{template "list.tmpl.html" $table}}
    <CAdd ref="cadd" @onsaved="onQuery"></CAdd>
    <CEdit ref="cedit" @onsaved="onQuery"></CEdit>
    <CView ref="cview" @onsaved="onQuery"></CView>
    <DLGOpts ref="dlgOpts" @onsaved="onQuery"></DLGOpts>
    <DLGCnfrm ref="dlgCnfrm" @onsaved="onQuery"></DLGCnfrm>
  </div>
</template>
<script>
import CAdd from "./{{.Name}}.add"
import CEdit from "./{{.Name}}.edit"
import CView from "./{{.Name}}.view"
import DLGOpts from "./{{.Name}}.dialog.vue"
import DLGCnfrm from "./{{.Name}}.cnfrm.vue"
{{range $x,$m:=$optRow -}}
 {{- if eq "CMPNT" $m.Name  -}}
import {{$m.UNQ}} from "{{$m.URL}}"
{{- end}}
{{- end}}
export default {
   components: {
		CAdd,
    CEdit,
    CView,
    DLGOpts,
    DLGCnfrm,
    {{ range $x,$m:=$optRow -}}
     {{- if eq "CMPNT" $m.Name -}}
    {{$m.UNQ}},
    {{- end -}}
    {{end}}
  },
  data() {
    return {
      conf:{
        loading:false,
        progressColor: this.$theia.env.conf.progress||[]
      },
      {{- template "queryform.tmpl.js" $table -}}
    };
  },
    created() {
      this.queryData_{{$table.UNQ}}()
    },
  methods:{
    {{- template "list.tmpl.js" $table -}}
    handleSizeChange(ps){
      this.form_{{$table.UNQ}}.ps = ps
     this.queryData_{{$table.UNQ}}()
    },
    handleCurrentChange(pi){
      this.form_{{$table.UNQ}}.pi = pi
     this.queryData_{{$table.UNQ}}()
    },
    onQuery(){
      this.form_{{$table.UNQ}}.pi = 1
      this.queryData_{{$table.UNQ}}()
    },
    showAdd(){
       this.$refs.cadd.show()
    },
     showView(q){ 
       this.$refs.cview.show(q)
    },
    showEdit(q){
        this.$refs.cedit.show(q)
      },
    colorful(r){
       return this.$theia.env.conf.colorful[r]||""
    },
    tagColor(r){
       return this.$theia.env.conf.colorTag[r]||""
    },
     goto(url,param){
       if (!url) {
        return;
      }
      if (!url.indexOf("://")) {
        this.$router.push({ path: url, query: param });
        return;
      }
      let p = this.$theia.url.encode(param)
      window.location =`${url}?${p}`
    },
    {{template "opts.tmpl.js" $table}}
  },
};
</script>
<style>
.el-form-item {
  margin-right: 10px !important;
}
</style>