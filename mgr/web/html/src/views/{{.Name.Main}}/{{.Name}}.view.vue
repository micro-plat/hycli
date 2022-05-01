{{- $table := .|fltrUITable -}}
{{- $vrow := $table.VRows -}}
<template>
  <el-dialog
    v-model="conf.visible"
    title="{{.Desc}}详情"
    draggable
    width="76%"
    :before-close="hide"
  >
  
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
    {{template "view.tmpl.js" $table}}
  },
};
</script>