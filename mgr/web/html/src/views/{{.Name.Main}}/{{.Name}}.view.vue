
<template>
  <el-dialog
    v-model="conf.visible"
    title="{{.Desc}}详情"
    draggable
    width="76%"
    :before-close="hide"
  >
    {{- $table := .|fltrUITable -}}
    {{- $vrow := $table.VRows -}}
    <el-tabs v-model="conf.selected">
      <el-tab-pane label="详情" name="first">
        {{template "view.tmpl.html" $vrow}}
      </el-tab-pane>
    </el-tabs>
    <template #footer>
      <span style="height: 60px"> </span>
    </template>
  </el-dialog>
</template>
<script>
export default {
  data() {
    return {
        conf:{
        visible:false,
        selected:"first"
      },
      view: {
        {{- range $i,$c := $vrow}}
        {{$c.Name}}:"",
        {{- end}}
        },
      }
  },
  methods: {
    
    show(form) {
      this.conf.visible = true
      let that = this;
      this.$theia.http
        .get("/{{.Name.CName|lower}}",form)
        .then((res) => {
          Object.assign(that.view, res)
      {{- range $i,$c := $vrow -}}
        {{- if eq true $c.Ext.IsSelect}}
          that.view.{{$c.Name}}_label = that.$theia.enum.getName("{{$c.Ext.SLType}}",res.{{$c.Name}})
       {{- end -}}
        {{- end}}
        })
        .catch((res) => {
          let code = res.response.status;
          let msg = `{{.Desc}}查询失败(${code})`;
          that.$notify.error({ title: "失败", message: msg, duration: 5000 });
        });
    },
     colorful(r){
       return this.$theia.env.conf.colorful[r]
    },
  },
};
</script>