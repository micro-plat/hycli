package services

{{$etable := .|fltrCodeTables}}
import (
	"github.com/micro-plat/hydra"
	{{range $i,$v:=$etable -}}
	"{{$v.PkgName}}/services/{{$v.Name.Main}}"
	{{end}}
)

func init() {
	{{range $i,$v:=$etable -}}
	hydra.S.Micro("/{{.Name.CName|lower}}", {{$v.Name.Main}}.New{{.Name.CName}}Handler)
	{{end}}


	
}
