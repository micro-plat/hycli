{{- $qrow := .QRows}}
    form_{{.UNQ}}: {
        pi: 1,
        ps: 10,
        {{- range $i,$c := $qrow}}
        {{$c.Name}}:"",
        {{- end}}
        },
    {{- range $i,$c := $qrow -}}
    {{- if or (eq true $c.RType.IsSelect) (eq "multiselect" $c.RType.Type)}}
    {{.Name}}List:this.$theia.enum.get("{{$c.Ext.SLType}}"),
    {{.Name}}Exts:this.$theia.enum.getExts("{{$c.Ext.SLType}}"),
    {{- end}}
    {{- end}}
    dataList_{{.UNQ}}:[],
    total_{{.UNQ}}:0,