<template>
  {{- $table := .|fltrUITable}}
  {{- $qrow := $table.QRows}}
  {{- $lstRow := $table.LRows}}
  {{- $leRow := $table.LERows}}
  {{- $optRow :=$table.Optrs}}
  {{- $delRow :=$table.DRows}}
  <div>

     <!-- 查询条件 -->
    {{template "query.tmpl.html" $qrow}}


    <!-- 列表 -->
    {{template "list.tmpl.html" $table}}
    <CAdd ref="cadd" @onsaved="onQuery"></CAdd>
    <CEdit ref="cedit" @onsaved="onQuery"></CEdit>
    <CView ref="cview" @onsaved="onQuery"></CView>
    <DLGOpts ref="dlgOpts" @onsaved="onQuery"></DLGOpts>
  <DLGCnfrm ref="dlgCnfrm"  @onsaved="onQuery"></DLGCnfrm>

    {{ range $x,$m:=$optRow -}}
    {{- if eq "CMPNT" $m.Name -}}
    <{{$m.UNQ}} ref="cmpnt_{{$m.UNQ}}"  @onsaved="onQuery"></{{$m.UNQ}}>
    {{- end -}}
    {{end}}
   
  </div>
</template>


<script>
import CAdd from "./{{.Name}}.add"
import CEdit from "./{{.Name}}.edit"
import CView from "./{{.Name}}.view"
import DLGOpts from "./{{.Name}}.opts.vue"
import DLGCnfrm from "./{{.Name}}.cnfrm.vue"

  {{ range $x,$m:=$optRow -}}
    {{- if eq "CMPNT" $m.Name  -}}
   import {{$m.UNQ}} from "{{$m.URL}}"
    {{- end -}}
    {{end}}
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
      {{- template "queryform.tmpl.js" $qrow -}}
    };
  },
    created() {
      this.queryData()
    },
  methods:{
    queryData(){

      //构建查询参数
      let that = this
       {{ range $i,$c:= $qrow }}
        {{- if eq "daterange" $c.RType.Type  -}}
     this.form.start_{{$c.Name}} = null
      this.form.end_{{$c.Name}} = null
      if(this.form.{{- $c.Name -}} && this.form.{{$c.Name}}.length>1){
        this.form.start_{{$c.Name}}=this.form.{{$c.Name}}[0]
        this.form.end_{{$c.Name}}=this.form.{{$c.Name}}[1]
      }
        {{end -}}
        {{- end -}}
      //发送查询请求
      this.conf.loading = true
      this.$theia.http.get("/{{.Name.CName|lower}}/query",this.form).then(res=>{
        that.conf.loading = false
        that.dataList = res.items
        that.total = res.count
        that.dataList.forEach(item => {
          {{- range $i,$c := $lstRow -}}
          {{- if and (ne "" $c.RType.Format) (eq true $c.RType.IsDate)}}
          item.{{$c.Name}} = that.$theia.str.dateFormat(item.{{$c.Name}},'{{$c.RType.Format}}')
        {{- else if and (ne "" $c.RType.Format) (eq true $c.RType.IsNumber)}}
           item.{{$c.Name}} = that.$theia.str.numberFormat(item.{{$c.Name}},'{{$c.RType.Format}}')
        {{- end}}
        {{- end }}
        {{- range $i,$c := $qrow -}}
       {{- if eq true $c.Ext.IsSelect}}
          item.{{$c.Name}}_label = that.$theia.enum.getName("{{$c.Ext.SLType}}",item.{{$c.Name}})
        {{- end -}}
        {{- if eq "switch" $c.RType.Type}}
          item.{{$c.Name}}_switch = item.{{$c.Name}} == 0
        {{- end}}
        {{- end}}
        });
      })
    },
    handleSizeChange(ps){
      this.form.ps = ps
      this.queryData()
    },
    handleCurrentChange(pi){
      this.form.pi = pi
      this.queryData()
    },
    onQuery(){
      this.form.pi = 1
      this.queryData()
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
       return this.$theia.env.conf.colorful[r]
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