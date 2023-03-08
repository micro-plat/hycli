package data

import "github.com/micro-plat/hycli/data/internal/md"

type Table struct {
	*md.Table
	Colums    []*Column `json:'cols'`    //所有列
	QColums   []*Column `json:'qcols'`   //查询列
	LColums   []*Column `json:'lcols'`   //列表列
	LEColums  []*Column `json:'lecols'`  //列表扩展列
	CColums   []*Column `json:'ccols'`   //创建列
	UColums   []*Column `json:'ucols'`   //更新列
	VColums   []*Column `json:'vcols'`   //预览列
	DColums   []*Column `json:'dcols'`   //删除列
	PKColums  []*Column `json:'pk_cols'` //主键列
	Optrs     []*Operation
	BatchOpts []*Operation
	QueryOpts []*Operation
	ViewOpts  []*Operation //预览
	UNQ       string       `json:'unq'`
}
