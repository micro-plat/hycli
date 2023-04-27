package menu

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/types"
)

// 获取用户菜单s
func GetMenusByUserID(userID int64, ident string) (types.XMaps, error) {
	db := hydra.C.DB().GetRegularDB()
	data, err := db.Query(getMenuByUserId, map[string]interface{}{
		"user_id": userID,
		"ident":   ident,
	})
	if err != nil {
		return nil, err
	}
	return getTreeNodes(data, "parent", "id"), nil
}
func getTreeNodes(items types.XMaps, pidName string, pkName string) types.XMaps {
	treeMap := map[string]types.XMaps{}
	lstMap := types.XMaps{}
	for _, r := range items {
		pid := r.GetString(pidName)
		if _, ok := treeMap[pid]; !ok {
			treeMap[pid] = make(types.XMaps, 0, 1)
		}
		treeMap[pid] = append(treeMap[pid], r)
		if pid == "0" {
			lstMap = append(lstMap, r)
		}
	}
	for _, m := range lstMap {
		setChildren(m, pkName, treeMap)
	}
	return lstMap
}
func setChildren(current types.XMap, idName string, treeMap map[string]types.XMaps) {
	id := current.GetString(idName)
	children, ok := treeMap[id]
	if !ok {
		return
	}
	current.SetValue("children", children)
	for _, v := range children {
		setChildren(v, idName, treeMap)
	}
}
