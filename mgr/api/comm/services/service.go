//go:build ignore

package services

{{- $etable := .}}
{{- $mtable :=$etable|flterMainTable}}
{{- $ft := getFirstTable $etable}}
import (
	"github.com/micro-plat/hydra"
	{{- range $i,$v:=$mtable}}
	{{- if eq false $v.Exclude}}
	"{{$v.PkgName}}/services/{{$v.Name.Prefix}}/{{$v.Name.Main}}"
	{{- end}}
	{{- end}}
	"{{$ft.PkgName}}/services/enums"
	_ "{{$ft.PkgName}}/modules/const/db/mysql"
)

func init() {
	{{- range $i,$v:=$etable}}
	{{- if eq false $v.Exclude}}
	hydra.S.Micro("/{{.Name.MainPath|lower}}", {{$v.Name.Main}}.New{{.Name.CName}}Handler)
	{{- end}}
	{{- end}}
	hydra.S.Micro("/enums",enums.NewEnumsHandler)
}
