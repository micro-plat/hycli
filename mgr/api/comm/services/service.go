//go:build ignore

package services

{{- $etable := .|fltrCodeTables -}}
{{- $mtable :=$etable|flterMainTable -}}
{{$ft := getFirstCodeTable $etable}}
import (
	"github.com/micro-plat/hydra"
	{{range $i,$v:=$mtable -}}
	"{{$v.PkgName}}/services/{{$v.Name.Main}}"
	{{end -}}
	"{{$ft.PkgName}}/services/enums"
)

func init() {
	{{range $i,$v:=$etable -}}
	hydra.S.Micro("/{{.Name.CName|lower}}", {{$v.Name.Main}}.New{{.Name.CName}}Handler)
	{{end -}}
	hydra.S.Micro("/enums",enums.NewEnumsHandler)
}
