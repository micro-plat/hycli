package md

import (
	"fmt"
	"os"
	"path/filepath"
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

func GetFormat(tp string, cns ...string) string {
	lst := GetConstraintByReg(cns, fmt.Sprintf(`^%s\(([^\(^\)]+)\)$`, tp))
	return types.GetStringByIndex(lst, 0)
}
func GetValue(name string, cns ...string) string {
	lst := GetConstraintByReg(cns, fmt.Sprintf(`%s\((\w+)\)`, name))
	return types.GetStringByIndex(lst, 0)
}
func GetDefValue(cns ...string) string {
	return GetValue("def", cns...)
	// lst := GetConstraintByReg(cns, `def\((\w+)\)`)
	// return types.GetStringByIndex(lst, 0)
}
func GetVdlValue(cns ...string) string {
	return GetValue("valid", cns...)
	// lst := GetConstraintByReg(cns, `valid\((\w+)\)`)
	// return types.GetStringByIndex(lst, 0)
}
func GetRangeName(cns ...string) string {
	if r := types.GetStringByIndex(GetConstraintByReg(cns, `range\((\w+)\)`), 0); r != "" {
		return "range"
	}
	return ""
}
func GetConsByTagIgnorecase(t string, cns ...string) (string, string) {
	k, v := GetConsByTag(strings.ToLower(t), cns...)
	if k == "" {
		k, v = GetConsByTag(strings.ToUpper(t), cns...)
	}
	return k, v
}
func GetConsNameByTagIgnorecase(t string, cns ...string) string {
	k := GetConsNameByTag(strings.ToLower(t), cns...)
	if k == "" {
		k = GetConsNameByTag(strings.ToUpper(t), cns...)
	}
	return k
}

func GetConsNameByTag(t string, cns ...string) string {
	if r := types.GetStringByIndex(GetConstraintByReg(cns,
		fmt.Sprintf(`%s\((\w+)[,]?([\w]*)\)`, t)), 0); r != "" {
		return r
	}
	return ""
}

func GetConsByTag(t string, cns ...string) (string, string) {
	cons := GetConstraintByReg(cns, fmt.Sprintf(`%s\((\w+)[,]?([\w]*)\)`, t))
	if len(cons) == 2 {
		return cons[0], cons[1]
	}
	if len(cons) == 1 {
		return cons[0], ""
	}
	return "", ""
}

func GetSelectName(fname string, cns ...string) (string, string) {

	lst := GetConstraintByReg(cns, `sl\((\w+)[,]?([\w]*)\)`)
	if len(lst) > 1 && lst[1] != "" {
		return lst[0], lst[1]
	}
	if len(lst) > 0 {
		return lst[0], ""
	}
	if HasConstraint(cns, "sl") {
		return fname, ""
	}
	return "", ""
}
func GetRanges(cns ...string) (string, string) {
	lst := GetConstraintByReg(cns, `range\((\w+)[,]?([\w]*)\)`)
	if len(lst) > 1 && lst[1] != "" {
		return lst[0], lst[1]
	}
	if len(lst) > 0 {
		return "", lst[0]
	}
	return "", ""
}

func HasConstraint(cns []string, fx ...string) bool {
	for _, f := range fx {
		for _, c := range cns {
			if strings.EqualFold(c, f) {
				return true
			}
			reg := regexp.MustCompile(fmt.Sprintf(`^%s\([\w-\s:]+\)$`, f))
			if len(reg.FindAllString(c, -1)) > 0 {
				return true
			}
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
func GetExtOpt(t string, tag string) [][]string {
	reg := regexp.MustCompile(fmt.Sprintf(`%s\([^(^)]+\)`, tag))
	lst := reg.FindAllString(t, -1)
	rlst := make([][]string, 0, 1)
	for _, l := range lst {
		n := regexp.MustCompile(`\(([\w\p{Han}]+),([\w]+):([/\w\.]+)[,]?([/\w\.]*)[:]?([\w]*)\)`)
		xn := n.FindAllStringSubmatch(l, -1)
		if len(xn) > 0 && len(xn[0]) > 1 {
			rlst = append(rlst, xn[0][1:])
		}
	}
	return rlst
}
func GetPkgName() string {
	gopath, _ := os.LookupEnv("GOPATH")
	if gopath == "" {
		return ""
	}
	currentPath, _ := os.Getwd()
	root := filepath.Join(gopath, "src")
	if strings.HasPrefix(strings.ToLower(currentPath), strings.ToLower(root)) {
		currentPath = currentPath[len(root)+1:]
	}
	return strings.Trim(strings.Replace(currentPath, "\\", "/", -1), "/")

}
