<template>
  <div>
    <div class="query">
      <el-form :model="form" label-width="120px">
        {{ range.Rows }}
        {{ $R := .|to_ui_row}}
        <el-form-item label="{{$R.Desc.Name}}">
          {{- if eq "input" $R.T.InputType -}}
          <el-input v-model="form.{{$R.Name}}" />
          {{- else if eq "date" $R.T.InputType  -}}
          <el-date-picker
            v-model="form.{{$R.Name}}"
            type="date"
            placeholder="{{$R.Desc.Raw}}"
            style="width: 100%"
          />
          {{- else if eq "number" $R.T.InputType  -}}
          <el-input-number v-model="form.{{$R.Name}}" />
          {{- else if eq "select" $R.T.InputType  -}}
          <el-select
            v-model="form.{{$R.Name}}"
            clearable
            placeholder="{{.Desc}}"
          >
            <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
          {{- end -}}</el-form-item
        >{{ end }}
      </el-form>
    </div>
    <div class="list"></div>
  </div>
</template>


<script lang="ts" setup>
import { reactive } from 'vue'

// do not use same name with ref
const form = reactive({
   {{ range.Rows }} {{.Name}}:'',
   {{end}}
})

const onSubmit = () => {
}
</script>


<style>
</style>