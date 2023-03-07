{{- $table := . -}}
{{- $vrow := $table.VRows -}}
{{- $viewOptRow :=$table.ViewOpts}}
<template>
  <el-dialog
    v-model="conf.visible"
    title="{{.Desc}}详情"
    draggable
    width="76%"
    :close-on-click-modal="false"
    :before-close="hide"
  >
    <el-tabs v-model="conf.selected">
      <el-tab-pane label="详情" name="first">
        {{template "view.tmpl.html" $vrow}}
      </el-tab-pane>
      {{range $i,$c:= $viewOptRow}}
      {{if eq "CMPNT" $c.Name}}
      <el-tab-pane label="{{$c.Label}}" name="{{$c.Label}}"  @tab-click="show_view_{{$c.UNQ}}">
         <{{$c.UNQ}} ref="view_{{$c.UNQ}}"></{{$c.UNQ}}>
      </el-tab-pane>
       {{else if eq "TAB" $c.Name}}
        <el-tab-pane label="{{$c.Label}}" name="{{$c.Label}}"  @tab-click="show_view_{{$c.UNQ}}">
         {{$ct:= fltrSearchUITable $c}}
        {{- template "list.tmpl.html" $ct}}
         </el-tab-pane>
      {{end}}
      {{ end }}
    </el-tabs>
    <template #footer>
      <span style="height: 60px"> </span>
    </template>
  </el-dialog>
</template>
<script>
 {{ range $x,$m:=$viewOptRow -}}
    {{- if eq "CMPNT" $m.Name  -}}
   import {{$m.UNQ}} from "{{$m.URL}}"
    {{- end -}}
    {{end}}
export default {
   components: {
    {{ range $x,$m:=$viewOptRow -}}
     {{- if eq "CMPNT" $m.Name -}}
    {{$m.UNQ}},
    {{- end -}}
    {{end}}
  },
  data() {
    return {
        conf:{
        visible:false,
        selected:"first"
      },
       {{- range $i,$c:= $viewOptRow -}}
       {{ if eq "TAB" $c.Name}}
          {{$ct:= fltrSearchUITable  $c}}
        {{- template "queryform.tmpl.js" $ct}}
      {{end}}
      {{ end }}
      view: {
        {{- range $i,$c := $vrow}}
        {{$c.Name}}:"",
        {{- end}}
        },
      }
  },
  methods: {
    {{template "view.tmpl.js" $table}}
     {{- range $i,$c:= $viewOptRow -}}
       {{- if eq "TAB" $c.Name -}}
          {{- $ct:= fltrSearchUITable  $c -}}
          {{- $tbs := contactTBS  $ct $table -}}
        {{- template "list.tmpl.js" $tbs}}
      {{- end -}}
      {{- end -}}
  },
};
</script>