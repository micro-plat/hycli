<template>
  {{- $table := .|fltrUITable}}
  {{- $qrow := $table.QRows}}
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
    {{ range $x,$m:=$optRow -}}
    {{- if eq "CMPNT" $m.Name -}}
    <{{$m.UNQ}} ref="cmpnt_{{$m.UNQ}}"  @onsaved="onQuery"></{{$m.UNQ}}>
    {{- end -}}
    {{end}}
    <el-dialog draggable v-model="conf.confirmVisible" title="删除" width="30%">
      <span>确认删除{{range $i,$c:= $delRow}} {{ $c.Desc }}:[<span
          v-text="delForm.{{$c.Name}}"></span>{{ end }}] 吗?</span>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="conf.confirmVisible = false">取消</el-button>
          <el-button type="primary" @click="onDelete">确认</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>


<script>
import CAdd from "./{{.Name}}.add"
import CEdit from "./{{.Name}}.edit"
import CView from "./{{.Name}}.view"
import DLGOpts from "./{{.Name}}.opts.vue"
  {{ range $x,$m:=$optRow -}}
    {{- if eq "CMPNT" $m.Name -}}
   import {{$m.UNQ}} from "{{$m.URL}}"
    {{- end -}}
    {{end}}
export default {
   components: {
		CAdd,
    CEdit,
    CView,
    DLGOpts,
    {{ range $x,$m:=$optRow -}}
    {{- if eq "CMPNT" $m.Name -}}
    {{$m.UNQ}},
    {{- end -}}
    {{end}}
  },
  data() {
    return {
      conf:{
        confirmVisible:false,
        loading:false,
        progressColor: this.$theia.env.conf.progress||[]
      },
      delForm:{},
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
     delItem(q){ 
       this.conf.confirmVisible = true
       this.delForm = q
    },
    onDelete(){
        let that = this
        this.$theia.http.del("/{{.Name.CName|lower}}",this.delForm).then(res=>{
          that.conf.confirmVisible = false
          that.$notify.success({title: '成功',message: '{{.Desc}}删除成功',duration:5000})
          that.onQuery()
       }).catch(err=>{
          that.conf.confirmVisible = false
          let code = res.response.status
          let msg= `{{.Desc}}保存失败(${code})`
          that.$notify.error({title: '失败',message:msg,duration:5000})
       })
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
     {{ range $x,$m:=$optRow -}}
    {{- if eq "DIALOG" $m.Name -}}
    show_dialog_{{$m.UNQ}}(fm){
       this.$refs.dlgOpts.show_{{$m.UNQ}}(fm)
    },
    {{- else if eq "CMPNT"  $m.Name -}}
    show_cmpnt_{{$m.UNQ}}(fm){
      let query={}
      {{- $rows:= fltrUIRows $table $m.RwName -}}
        {{range $i,$c:=$rows}} 
      query.{{$c.Name}} = fm.{{$c.Name}}
        {{- end}}
      this.$refs.cmpnt_{{$m.UNQ}}.show(query)
    },
    {{- end -}}
    {{- end}}

    {{ range $x,$m:=$optRow -}}
    {{- if eq "LINK" $m.Name -}}
    goto_{{$m.UNQ}}(fm){
      let query={}
      {{- $rows:= fltrUIRows $table $m.RwName -}}
        {{range $i,$c:=$rows}} 
      query.{{$c.Name}} = fm.{{$c.Name}}
        {{- end}}
      this.goto('{{$m.URL}}',query)
    }
    {{- end -}}
    {{- end}}
  },
};
</script>
<style>
.el-form-item {
  margin-right: 10px !important;
}
</style>