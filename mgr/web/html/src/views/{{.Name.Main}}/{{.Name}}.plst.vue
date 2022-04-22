<template>
  <div>
{{- $table := .|fltrUITable}}
{{- $optRow :=$table.PlstOpts}}
{{ range $x,$m:=$optRow -}}
    {{- if eq "DIALOG" $m.Name -}}
   {{end}}
{{end}}
  </div>
</template>

<script>
export default {

}
</script>

<style>

</style>