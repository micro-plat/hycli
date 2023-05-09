package md

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/micro-plat/lib4go/jsons"
	"github.com/micro-plat/lib4go/types"
)

type RType struct {
	Raw  string `json:"raw"`
	Name string `json:"name"`
	Len  int    `json:"len"`
	DLen int    `json:"dlen"`
}

func (r *RType) String() string {
	return r.Name
}

type RDesc struct {
	Raw  string `json:"raw"`
	Name string `json:"name"`
}

func (r *RDesc) String() string {
	return r.Name
}

// Row 行信息
type Row struct {
	Index       int      `json:"index"` //索引
	Raw         string   `json:"raw"`
	Name        string   `json:"name"`        //字段名
	CName       string   `json:"cname"`       //字段名
	Type        *RType   `json:"type"`        //类型
	DefValue    string   `json:"defValue"`    //默认值
	AllowNull   bool     `json:"allowNull"`   //为空
	Constraints []string `json:"constraints"` //约束
	Desc        *RDesc   `json:"desc"`        //描述
}

func (r *Row) String() string {
	buff, _ := jsons.Marshal(r)
	return string(buff)
}
func (s *Row) Equal(t *Row) bool {
	return s.Raw == t.Raw
}

func Line2TableRow(line *Line) (*Row, error) {
	if strings.Count(line.Text, "|") != 7 {
		return nil, fmt.Errorf("数据表结构有误(行:%d)", line.LineID)
	}
	Columns := strings.Split(strings.Trim(strings.TrimSpace(line.Text), "|"), "|")
	if Columns[0] == "" {
		return nil, fmt.Errorf("字段名称不能为空 %s(行:%d)", line.Text, line.LineID)
	}
	rtp, err := getRType(strings.TrimSpace(Columns[1]), line.LineID)
	if err != nil {
		return nil, err
	}
	desc := strings.TrimSpace(strings.Replace(Columns[5], "&#124;", "|", -1))
	c := &Row{
		Index:       line.LineID,
		Raw:         strings.TrimSpace(strings.Replace(Columns[0], "&#124;", "|", -1)),
		Name:        strings.Trim(strings.TrimSpace(strings.Replace(Columns[0], "&#124;", "|", -1)), "^"),
		CName:       ToCName(strings.Trim(strings.TrimSpace(strings.Replace(Columns[0], "&#124;", "|", -1)), "^")),
		Type:        rtp,
		DefValue:    strings.TrimSpace(strings.Replace(Columns[2], "&#124;", "|", -1)),
		AllowNull:   strings.TrimSpace(Columns[3]) != "否",
		Constraints: mergeConstraint(getNames(strings.TrimSpace(Columns[4]))),
		Desc:        &RDesc{Raw: desc, Name: getShortDesc(desc)},
	}

	return c, nil
}
func getRType(t string, index int) (*RType, error) {
	reg := regexp.MustCompile(`[\w]+`)
	names := reg.FindAllString(t, -1)
	if len(names) == 0 || len(names) > 4 {
		return nil, fmt.Errorf("未设置字段类型:%v(行:%d)", names, index)
	}
	if len(names) == 1 {
		return &RType{Raw: t, Name: names[0], Len: 0, DLen: 0}, nil
	}
	if len(names) == 2 {
		return &RType{Raw: t, Name: names[0], Len: types.GetInt(names[1]), DLen: 0}, nil
	}
	return &RType{Raw: t, Name: names[0], Len: types.GetInt(names[1]), DLen: types.GetInt(names[2])}, nil
}
func getNamesx(t string) []string {
	reg := regexp.MustCompile(`[^,]+`)
	names := reg.FindAllString(t, -1)
	lname := make([]string, 0, len(names))
	for _, v := range names {
		lname = append(lname, strings.Trim(v, " "))
	}
	return lname
}

func getNames(p string) []string {
	// reg := regexp.MustCompile(`([\w]+[\(][\w]+[,]?[\w]*\))|([@]?[\w]+)`)
	reg := regexp.MustCompile(`([@&\^\?#\$%\*]?[\w]+\([^(]*\))|([@&\^\?#\$%\*]?[\w]+)`)
	xn := reg.FindAllStringSubmatch(p, -1)
	if len(xn) == 0 {
		return nil
	}
	names := make([]string, 0, len(xn))
	for _, r := range xn {
		for i := 1; i < len(r); i++ {
			if r[i] == "" {
				continue
			}
			names = append(names, r[i])
		}
	}
	return names
}
func mergeConstraint(input []string) []string {
	output := make([]string, 0, 1)
	for _, v := range input {
		if strings.HasPrefix(v, "?") {
			lst := []rune(v[1:])
			for _, c := range lst {
				output = append(output, string(c))
			}
			continue
		}
		output = append(output, v)
	}
	routput := make([]string, 0, len(output))
	for _, v := range output {
		if x, ok := mtp[strings.ToLower(v)]; ok {
			routput = append(routput, x...)
			continue
		}
		routput = append(routput, v)
	}
	return routput
}
func getShortDesc(t string) string {
	reg := regexp.MustCompile(`[\w\p{Han}]+`)
	names := reg.FindAllString(t, -1)
	if len(names) == 0 {
		return ""
	}
	return strings.TrimSpace(names[0])
}

var mtp = map[string][]string{
	"e": {"le"},
	"i": {"DI"},
	"n": {"DN"},
	"s": {"sl"},
	"q": {"q"},
	"l": {"l"},
	"c": {"c"},
	"u": {"u"},
	"d": {"d"},
}
