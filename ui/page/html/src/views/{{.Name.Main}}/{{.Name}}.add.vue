<template>
  <el-dialog
    v-model="conf.visible"
    title="添加 {{.Desc}}"
    width="30%"
    :before-close="hide"
  >
    {{ $table := .|fltrUITable}}
  {{ $crow := $table.CRows}}
    <el-form :model="form" size="small" ref="addForm" :rules="rules">
      {{ range $i,$c:= $crow }}
      {{if eq "password" $c.RType.Type  -}}
      <el-form-item
        prop="{{$c.Name}}"
        label="{{$c.Desc}}"
        label-width="100px"
        label-position="right"
      >
        <el-input
          clearable
          type="password"
          show-password
          style="width: 100%"
          v-model="form.{{$c.Name}}"
          maxlength="{{$c.RType.Len}}"
          placeholder="请输入{{$c.Desc}}"
        />
      </el-form-item>
      {{- end -}}
      {{if eq "switch" $c.RType.Type  -}}
      <el-form-item
        prop="{{$c.Name}}"
        label="{{$c.Desc}}"
        label-width="100px"
        label-position="right"
      >
        <el-switch v-model="form.{{$c.Name}}_switch" />
      </el-form-item>
      {{- end -}}
      {{if or (eq "date" $c.RType.Type) (eq "daterange" $c.RType.Type)  -}}
      <el-form-item
        prop="{{$c.Name}}"
        label-width="100px"
        label-position="right"
      >
        <el-date-picker
          style="width: 100%"
          v-model="form.{{$c.Name}}"
          clearable
          type="date"
          placeholder="请选择{{.Desc}}"
          format="YYYY-MM-DD"
          value-format="YYYY-MM-DD"
        />
      </el-form-item>
      {{- end -}}
      {{if eq "select" $c.RType.Type  -}}
      <el-form-item
        prop="{{$c.Name}}"
        label="{{$c.Desc}}"
        label-width="100px"
        label-position="right"
      >
        <el-select
          v-model="form.{{.Name}}"
          style="width: 100%"
          clearable
          placeholder="请选择{{.Desc}}"
        >
          <el-option
            v-for="item in {{.Name}}List"
            :key="item.value"
            :label="item.name"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      {{- end -}}
      {{if eq "multiselect" $c.RType.Type  -}}
      <el-form-item
        prop="{{$c.Name}}"
        label="{{$c.Desc}}"
        label-width="100px"
        label-position="right"
      >
        <el-select
          style="width: 100%"
          v-model="form.{{.Name}}"
          clearable
          multiple
          collapse-tags
          placeholder="请选择{{.Desc}}"
        >
          <el-option
            v-for="item in {{.Name}}List"
            :key="item.value"
            :label="item.name"
            :value="item.value"
          />
        </el-select>
      </el-form-item>
      {{- end -}}
      {{if eq "number" $c.RType.Type  -}}
      <el-form-item
        prop="{{$c.Name}}"
        label="{{$c.Desc}}"
        label-width="100px"
        label-position="right"
      >
        <el-input-number style="width: 100%" v-model="form.{{ $c.Name }}"
        {{ if ne "" $c.RType.Min}} :min="{{ $c.RType.Min }}" {{ end }}
        {{ if ne "" $c.RType.Max}} :max="{{ $c.RType.Max }}" {{ end }}
        />
      </el-form-item>
      {{- end -}}
      {{if eq "input" $c.RType.Type  -}}
      <el-form-item
        prop="{{$c.Name}}"
        label="{{$c.Desc}}"
        label-width="100px"
        label-position="right"
      >
        <el-input
          clearable
          style="width: 100%"
          v-model="form.{{$c.Name}}"
          maxlength="{{$c.RType.Len}}"
          placeholder="请输入{{$c.Desc}}"
        />
      </el-form-item>
      {{- end -}} {{ end }}
    </el-form>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="hide">取消</el-button>
        <el-button type="primary" @click="save">提交</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script>
export default {
  data() {
    return {
      conf: {
        visible: false,
      },
      rules:{
           {{- range $i,$c := $crow}}
            {{$c.Name}}:[{
                required:{{$c.Required}},
                message:"请输入{{$c.Desc}}",
                trigger: 'blur',
            }],
        {{- end}}
      },
      form:{
        {{- range $i,$c := $crow}}
        {{if eq "switch" $c.RType.Type  -}}
        {{$c.Name}}_switch:false,
        {{- else -}}
        {{$c.Name}}:"",
        {{- end}}
         {{- end}}
        },
        {{- range $i,$c := $crow -}}
        {{- if or (eq "select" $c.RType.Type) (eq "multiselect" $c.RType.Type)}}
        {{.Name}}List:this.$theia.enum.get("{{$c.Ext.SLType}}"),
        {{- end}}
        {{- end}}
      }
  },
  methods: {
    show() {
      this.conf.visible = true;
    },
    save(){
       {{range $i,$c:= $crow -}}
         {{- if eq "switch" $c.RType.Type  -}}
        this.form.{{$c.Name}} = this.form.{{$c.Name}}_switch?0:1;
         {{- end -}}
          {{end}}
        this.$refs.addForm.validate((v=>{
            if(v){
                this.onSave()
               
            }
        }))
    },
    onSave(){
        let that = this
        this.$theia.http.post("/{{.Name.CName|lower}}",this.form).then(res=>{
            that.$notify.success({title: '成功',message: '{{.Desc}}保存成功',duration:5000})
            that.$emit("onsaved")
            that.hide()
        }).catch(res=>{
          that.$notify.error({title: '失败',message: '{{.Desc}}保存失败',duration:5000})
        })
    },
    hide() {
      this.conf.visible = false;
      this.$refs.addForm.resetFields();
    },
  },
};
</script>

<style>
</style>