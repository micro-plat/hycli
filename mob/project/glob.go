package project

import "github.com/micro-plat/hycli/md"

type Project struct {
	Name    string
	PkgName string
	Dot     string
}

func NewProject(name string) *Project {
	return &Project{
		Name:    name,
		PkgName: md.GetPkgName(),
		Dot:     ".",
	}
}
