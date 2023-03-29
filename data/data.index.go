package data

import (
	"sort"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

type dbIdxs []*dbIdx
type dbIdx struct {
	Name    string
	Columns Columns
}

func createNormalIdx(t *Table) dbIdxs {
	return getDBIdxs("idx", t.Columns)
}
func createUNQIdx(t *Table) dbIdxs {
	return getDBIdxs("unq", t.Columns)
}
func getDBIdxs(tp string, colus Columns) dbIdxs {
	nameMap := map[string]*dbIdx{}
	for i, c := range colus {
		idxs := md.GetIdx(tp, c.RawConsts...)
		for _, idxArr := range idxs {
			name := types.GetStringByIndex(idxArr, 0)
			idx := types.GetInt(types.GetStringByIndex(idxArr, 1), i)
			if name == "" {
				continue
			}
			if v, ok := nameMap[name]; ok {
				v.Columns = append(v.Columns, c.NIndex(idx))
			} else {
				nameMap[name] = &dbIdx{
					Name:    name,
					Columns: Columns{c.NIndex(idx)},
				}
			}
		}
	}
	//对字段顺序进行排序
	lst := make(dbIdxs, 0, 1)
	for _, v := range nameMap {
		sort.Sort(v.Columns)
		lst = append(lst, v)
	}
	return lst
}
