package utils

import (
	"fmt"
	"strings"

	"github.com/micro-plat/lib4go/errs"
	"github.com/micro-plat/lib4go/types"
)

// TransformBatchContent 转换指定格式数据为map数组
// fileds := []string{"name", "sex", "dept", "birth", "addr"}
// cfield := "content"
// xfm := "1-1,2-2,3-4,4-5,5-5"
// input := types.NewXMap()
// input.SetValue("dept", "研发")
// input.SetValue("content", `杨磊,男,1983-09-11,成都,华阳
// 杨展硕,男,2016-08-06,北京,海淀`)
func TransformBatchContent(fileds []string, enumsFields types.XMap, cfield string, xfm string, input types.XMap) (types.XMaps, error) {
	content := input.GetString(cfield)
	lines := strings.Split(content, "\n")
	nmaps := types.NewXMaps()
	idxs := getArray(xfm, fileds)
	for _, line := range lines {
		if strings.Trim(line, "") == "" {
			continue
		}
		//设置字段值
		nmap := types.NewXMap()
		for _, f := range fileds {
			if input.Has(f) {
				nmap.SetValue(f, strings.Trim(input.GetString(f), "\t"))
			}
		}
		//处理多行文本数据
		vlines := splitLines(line)
		for _, idx := range idxs {
			field := types.GetStringByIndex(fileds, idx[1]-1)
			value := strings.Trim(types.GetStringByIndex(vlines, idx[0]-1), "\t")
			//处理枚举类型数据
			if enumsFields.Has(field) {
				value = GetValue(enumsFields.GetString(field), value)
				if value == "" {
					return nil, errs.NewError(511, fmt.Sprintf("未找到%s对应的数据", value))
				}
			}
			if nmap.Has(field) {
				value = nmap.GetString(field) + value
			}
			nmap.SetValue(field, value)
		}
		nmaps.Append(nmap)
	}
	return nmaps, nil
}

// 转换为二维数组 1-1,2-1,3-2,4-3
func getArray(f string, fields []string) [][]int {
	lines := strings.Split(f, ",")
	lst := make([][]int, 0, len(lines))
	for i, line := range lines {
		xidx := strings.Split(line, "-")

		if len(xidx) == 2 {
			firstIdxName := types.GetStringByIndex(xidx, 0)
			secondIdxName := types.GetStringByIndex(xidx, 1)
			secondIdx := types.GetInt(secondIdxName, -1)
			if secondIdx == -1 {
				secondIdx = getIdx(secondIdxName, fields) + 1
			}
			lst = append(lst, []int{
				types.GetInt(firstIdxName, i+1),
				types.GetInt(secondIdx, i+1),
			})
		}
	}
	return lst
}
func getIdx(f string, fields []string) int {
	for i, v := range fields {
		if strings.EqualFold(v, f) {
			return i
		}
	}
	return -1
}
func splitLines(line string) []string {
	n := strings.Trim(strings.Trim(line, "\t"), " ")
	n = strings.ReplaceAll(n, "\t", ",")
	n = strings.ReplaceAll(n, " ", ",")
	n = strings.ReplaceAll(n, ",,", ",")
	return strings.Split(n, ",")
}
