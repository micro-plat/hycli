package md

import (
	"fmt"
	"regexp"
	"strings"
)

// Lines2Table 表数据行变为表
func Lines2Table(lines Lines) (tables Tables, err error) {
	tables = Tables{}
	for _, tline := range lines.Lines {
		//markdown表格的表名，标题，标题数据区分行，共三行
		if len(tline) <= 3 {
			continue
		}
		var tb *Table
		for i, line := range tline {
			if i == 0 {
				//获取表名，描述名称
				name, err := getTableName(line)
				if err != nil {
					return nil, err
				}
				tb = NewTable(name, getTableDesc(line), getTableExtInfo(line))
				continue
			}
			if i < 3 {
				continue
			}
			c, err := Line2TableRow(line)
			if err != nil {
				return nil, err
			}

			if err := tb.AddRow(c); err != nil {
				return nil, err
			}
		}
		if tb != nil {
			tables = append(tables, tb)
		}
	}
	return tables, nil
}

func getTableDesc(line *Line) string {
	reg := regexp.MustCompile(`[^\d\.|\s]?[^\x00-\xff]+[^\[]+`)
	names := reg.FindAllString(line.Text, -1)
	if len(names) == 0 {
		return ""
	}
	return strings.TrimSpace(names[0])
}

func getTableExtInfo(line *Line) string {
	reg := regexp.MustCompile(`<(.*?)>`)
	names := reg.FindStringSubmatch(line.Text)
	if len(names) == 0 {
		return ""
	}
	return strings.TrimSpace(names[1])
}

func getTableName(line *Line) (string, error) {
	if !strings.HasPrefix(line.Text, "###") {
		return "", fmt.Errorf("%d行表名称标注不正确，请以###开头:(%s)", line.LineID, line.Text)
	}

	reg := regexp.MustCompile(`\[[\^]?[\w]+[,]?[\p{Han}A-Za-z0-9_]+\]`)
	names := reg.FindAllString(line.Text, -1)
	if len(names) == 0 {
		return "", fmt.Errorf("未设置表名称或者格式不正确:%s(行:%d)，格式：### 描述[表名,菜单名]，菜单名可选", line.Text, line.LineID)
	}
	s := strings.Split(strings.TrimRight(strings.TrimLeft(names[0], "["), "]"), ",")
	return s[0], nil
}
