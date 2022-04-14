package md

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/micro-plat/lib4go/types"
)

func ToCName(f string) string {
	l := strings.Split(f, "_")
	sb := strings.Builder{}
	for _, v := range l {
		if v == "" {
			continue
		}
		sb.WriteString(strings.ToUpper(string(v[0])))
		if len(v) > 1 {
			sb.WriteString(strings.ToLower(strings.ToLower(v[1:])))
		}

	}
	return sb.String()
}

func GetFormat(cns ...string) string {
	lst := GetConstraintByReg(cns, `f\(([^\(^\)]+)\)`)
	return types.GetStringByIndex(lst, 0)
}
func GetDefValue(cns ...string) string {
	lst := GetConstraintByReg(cns, `def\((\w+)\)`)
	return types.GetStringByIndex(lst, 0)
}
func GetVdlValue(cns ...string) string {
	lst := GetConstraintByReg(cns, `valid\((\w+)\)`)
	return types.GetStringByIndex(lst, 0)
}
func GetRangeName(cns ...string) string {
	if r := types.GetStringByIndex(GetConstraintByReg(cns, `range\((\w+)\)`), 0); r != "" {
		return "range"
	}
	return ""
}
func GetTPName(t string, cns ...string) string {
	if r := types.GetStringByIndex(GetConstraintByReg(cns,
		fmt.Sprintf(`%s\((\w+)[-]?([\w]*)\)`, t)), 0); r != "" {
		return r
	}
	return ""
}

func GetSelectName(fname string, cns ...string) string {
	if r := types.GetStringByIndex(GetConstraintByReg(cns, `sl\((\w+)\)`), 0); r != "" {
		return r
	}
	if HasConstraint(cns, "sl") {
		return fname
	}
	return ""
}
func GetRanges(cns ...string) (string, string) {
	lst := GetConstraintByReg(cns, `range\((\w+)[-]?([\w]*)\)`)
	if len(lst) > 1 && lst[1] != "" {
		return lst[0], lst[1]
	}
	if len(lst) > 0 {
		return "", lst[0]
	}
	return "", ""
}

func HasConstraint(cns []string, f string) bool {
	for _, c := range cns {
		if strings.EqualFold(c, f) {
			return true
		}
		reg := regexp.MustCompile(fmt.Sprintf(`%s\([\w-\s:]+\)`, f))
		if len(reg.FindAllString(c, -1)) > 0 {
			return true
		}
	}
	return false
}
func GetConstraintByReg(cns []string, f string) []string {
	reg := regexp.MustCompile(f)
	for _, c := range cns {
		names := reg.FindAllStringSubmatch(c, -1)
		if len(names) > 0 {
			return names[0][1:]
		}
	}
	return nil
}
