{{- $qcols := .QColums}}
    form_{{.UNQ}}: {
        pi: 1,
        ps: 10,
        {{- range $i,$c := $qcols}}
        {{$c.Name}}:"",
        {{- end}}
        },
    {{- range $i,$c := $qcols -}}
    {{- if or (eq true $c.Enum.IsEnum) (eq "multiselect" $c.QCMPT.Type)}}
    {{.Name}}List:this.$theia.enum.get("{{$c.Enum.EnumType}}"),
    {{.Name}}Exts:this.$theia.enum.getExts("{{$c.Enum.EnumType}}"),
    {{- end}}
    {{- end}}
    dataList_{{.UNQ}}:[],
    total_{{.UNQ}}:0,