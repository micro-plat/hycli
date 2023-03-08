package project

import "github.com/micro-plat/hycli/data"

type Project struct {
	Name    string
	PkgName string
	Dot     string
}

func NewProject(name string) *Project {
	return &Project{
		Name:    name,
		PkgName: data.GetPkgName(),
		Dot:     ".",
	}
}
