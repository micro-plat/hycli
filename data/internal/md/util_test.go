package md

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/micro-plat/lib4go/assert"
)

func TestRange(t *testing.T) {
	min, max := GetRanges("range(1-100)")
	assert.Equal(t, "1", min)
	assert.Equal(t, "100", max)

}
func TestRange1(t *testing.T) {
	min, max := GetRanges("range(100)")
	assert.Equal(t, "", min)
	assert.Equal(t, "100", max)

}
func TestTP(t *testing.T) {
	tp := GetConsNameByTag("tp", "tp(bl)")
	assert.Equal(t, "bl", tp)

	tp = GetConsNameByTag("l", "l(bl)")
	assert.Equal(t, "bl", tp)

	tp = GetConsNameByTag("DI", "DI")
	assert.Equal(t, "", tp)

}
func TestSelect(t *testing.T) {
	s, f, group := GetSelectName("status", "sl")
	assert.Equal(t, "status", s)
	assert.Equal(t, "", f)
	assert.Equal(t, "", group)
}
func TestSelect2(t *testing.T) {
	s, f, group := GetSelectName("status", "sl(bool)")
	assert.Equal(t, "bool", s)
	assert.Equal(t, "", f)
	assert.Equal(t, "", group)
}
func TestSelect3(t *testing.T) {
	s, f, group := GetSelectName("status", "sl(bool,fabc)")
	assert.Equal(t, "bool", s)
	assert.Equal(t, "fabc", f)
	assert.Equal(t, "", group)
}
func TestSelect3x(t *testing.T) {
	s, f, group := GetSelectName("status", "sl(bool,#fabc)")
	assert.Equal(t, "bool", s)
	assert.Equal(t, "#fabc", f)
	assert.Equal(t, "", group)
}
func TestSelect4(t *testing.T) {
	s, f, group := GetSelectName("status", "sl(bool,fabc,*)")
	assert.Equal(t, "bool", s)
	assert.Equal(t, "fabc", f)
	assert.Equal(t, "*", group)
}
func TestHasConstraint(t *testing.T) {
	b := HasConstraint([]string{"f"}, "f")
	assert.Equal(t, true, b)

	b = HasConstraint([]string{"PK", "SEQ", "L", "DI"}, "di")
	assert.Equal(t, true, b)

	b = HasConstraint([]string{"f(date)"}, "f")
	assert.Equal(t, true, b)

	b = HasConstraint([]string{"f(yyyy-mm-dd hh:mi:ss)"}, "f")
	assert.Equal(t, true, b)

	b = HasConstraint([]string{"f(10-20)"}, "f")
	assert.Equal(t, true, b)

	b = HasConstraint([]string{"sl(project_info)", "L", "LE", "C", "U"}, "sl")
	assert.Equal(t, true, b)

}
func TestRangeReg(t *testing.T) {
	reg := regexp.MustCompile(`range\((\w+)[-]?([\w]*)\)`)
	names := reg.FindAllStringSubmatch("range(100)", -1)

	t.Log(names[0][1:])
	assert.Equal(t, 1, 2)
}
func TestCons(t *testing.T) {
	reg := regexp.MustCompile(`[^,]+`)
	names := reg.FindAllString("l,v,E,range(1-100)", -1)
	fmt.Println(names)
	assert.Equal(t, 1, 2)
}

func TestCName(t *testing.T) {
	v := ToCName("user_id")
	assert.Equal(t, "UserId", v)

	v = ToCName("user__id")
	assert.Equal(t, "UserId", v)

	v = ToCName("user__i")
	assert.Equal(t, "UserI", v)

	v = ToCName("_user__i")
	assert.Equal(t, "UserI", v)
}
func TestGetName(t *testing.T) {
	v := mergeConstraint(getNames("IDX(idx_sso_sysinfo_urlident),@cUve "))
	assert.Equal(t, 5, len(v))
	assert.Equal(t, "U", v[2])
	assert.Equal(t, "le", v[4])
}

func TestGetName2(t *testing.T) {
	p := "l,v,q,sl(status_id,id),sl(bool),@cUve"
	xn := getNames(p)
	fmt.Println("xxxy:", strings.Join(xn, "|"))
	assert.Equal(t, 6, len(xn))

}
func TestGetName3x(t *testing.T) {
	p := "	PK,SEQ,L,DI              "
	xn := getNames(p)
	fmt.Println("xxxy:", strings.Join(xn, "|"))
	assert.Equal(t, 4, len(xn))

}
func TestGetNameFor(t *testing.T) {
	p := "Q,l,LE,C,U,lw(120),sl(project_system,project_id),tp(blist)"
	xn := getNames(p)
	fmt.Println("xxxy:", strings.Join(xn, "|"))
	assert.Equal(t, 4, len(xn))

}

func TestGetExtOpt(t *testing.T) { //>/right/info,x
	lst := GetExtOpt("lst(权限,link:/right/info,x:y),lst(启用,dialog:/right/save,m),pnl(启用,dialog:/right/save,m)", "lst")
	assert.Equal(t, 2, len(lst))
	assert.Equal(t, []string{"权限", "link", "/right/info", "x", "y"}, lst[0])
	assert.Equal(t, []string{"启用", "dialog", "/right/save", "m", ""}, lst[1])
}
func TestGetExtOptx(t *testing.T) { //>/right/info,x
	lst := GetExtOpt("lst(账户加款,dialog:/beanpay/account/info/balance,x),lst(授信,dialog:/beanpay/account/info/credit,y)", "lst")
	assert.Equal(t, 2, len(lst))
	assert.Equal(t, []string{"账户加款", "dialog", "/beanpay/account/info/balance", "x", ""}, lst[0])
	assert.Equal(t, []string{"授信", "dialog", "/beanpay/account/info/credit", "y", ""}, lst[1])
}

func TestGetFormat(t *testing.T) {
	f := GetFormat("   le,f(MM/dd HH:mm:ss),v")
	assert.Equal(t, "MM/dd HH:mm:ss", f)
}

func TestMergeConstraint(t *testing.T) {
	f := getNames("   le,f(MM/dd HH:mm:ss),v")
	fmt.Println("f:", f)
	assert.Equal(t, "MM/dd HH:mm:ss", f[1])
}
func TestExpr(t *testing.T) {
	p := GetExpr("age<20")
	assert.Equal(t, 3, len(p))
	assert.Equal(t, "age", p[0])
	assert.Equal(t, "<", p[1])
	assert.Equal(t, "20", p[2])

	p = GetExpr("age=20")
	assert.Equal(t, 3, len(p))
	assert.Equal(t, "age", p[0])
	assert.Equal(t, "=", p[1])
	assert.Equal(t, "20", p[2])

	p = GetExpr("age<=20")
	assert.Equal(t, 3, len(p))
	assert.Equal(t, "age", p[0])
	assert.Equal(t, "<=", p[1])
	assert.Equal(t, "20", p[2])

	p = GetExpr("age>=20")
	assert.Equal(t, 3, len(p))
	assert.Equal(t, "age", p[0])
	assert.Equal(t, ">=", p[1])
	assert.Equal(t, "20", p[2])

	p = GetExpr("age!=20")
	assert.Equal(t, 3, len(p))
	assert.Equal(t, "age", p[0])
	assert.Equal(t, "!=", p[1])
	assert.Equal(t, "20", p[2])

	p = GetExpr("age")
	assert.Equal(t, 0, len(p))

}
