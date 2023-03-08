package data

import "github.com/micro-plat/hycli/data/internal/md"

func Mds2Tables(mdpath string) (md.Tables, error) {
	return md.Mds2Tables(mdpath)
}
func GetPkgName() string {
	return md.GetPkgName()
}
