{{- $ucols := .}}
    rules:{
      {{- range $i,$c := $ucols}}
      {{$c.Name}}:[{
          required:{{$c.Field.Required}},
          message:"请输入{{$c.Label}}",
          trigger: 'blur',
      }],
    {{- end}}
    },
    form:{
        {{- range $i,$c := $ucols}}
        {{if eq "switch" $c.ExCMPT.Type  -}}
        {{$c.Name}}_switch:false,
        {{- else -}}
        {{$c.Name}}:"",
        {{- end}}
        {{- end}}
        },
        {{- range $i,$c := $ucols -}}
        {{- if or (eq true $c.Enum.IsEnum) (eq "multiselect" $c.ExCMPT.Type)}}
        {{.Name}}List:this.$theia.enum.get("{{$c.Enum.EnumType}}"),
        {{- end}}
        {{- end}}
      }