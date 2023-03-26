{{- $ccols := .}}
    rules:{
      {{- range $i,$c := $ccols}}
      {{$c.Name}}:[{
          required:{{$c.Field.Required}},
          message:"请输入{{$c.Label}}",
          trigger: 'blur',
      }],
    {{- end}}
    },
    form:{
        {{- range $i,$c := $ccols}}
        {{if eq "switch" $c.Cmpnt.Type  -}}
        {{$c.Name}}_switch:false,
        {{- else -}}
        {{$c.Name}}:"",
        {{- end}}
        {{- end}}
        },
        {{- range $i,$c := $ccols -}}
        {{- if or (eq true $c.Enum.IsEnum) (eq "multiselect" $c.Cmpnt.Type)}}
        {{.Name}}List:this.$theia.enum.get("{{$c.Enum.EnumType}}","{{$c.Enum.PID}}","{{$c.Enum.Group}}"),
        {{- end}}
        {{- end}}
      }