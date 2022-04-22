{{- $qrow := .}}
    form: {
        pi: 1,
        ps: 10,
        {{- range $i,$c := $qrow}}
        {{$c.Name}}:"",
        {{- end}}
        },
    {{- range $i,$c := $qrow -}}
    {{- if or (eq true $c.Ext.IsSelect) (eq "multiselect" $c.RType.Type)}}
    {{.Name}}List:this.$theia.enum.get("{{$c.Ext.SLType}}"),
    {{- end}}
    {{- end}}
    dataList:[],
    total:0,