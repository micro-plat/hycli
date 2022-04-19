{{- $crow := .}}
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