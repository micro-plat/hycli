package md

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/micro-plat/lib4go/assert"
	"github.com/micro-plat/lib4go/net/http"
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
func TestHasConstraint2(t *testing.T) {
	b := HasConstraint([]string{"c", "u", "l", "le", "sl", "DT"}, "DT")
	assert.Equal(t, true, b)
}
func TestHasConstraint3(t *testing.T) {
	b := HasConstraint([]string{"c(#ef)"}, "c")
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
	p := "Q,l,LE,C,U,lw(120),sl(project_system,project_id),tp(blist),&uid,@name,^abc,^ex(100),_x,_ef(ef)"
	xn := getNames(p)
	fmt.Println("xxxy:", strings.Join(xn, "|"))
	assert.Equal(t, 8, len(xn))

}
func TestGetNameFor2(t *testing.T) {
	p := "*abc,#USERiD,#wf(23),$ww,$ww(bw),%ef,%ef(wf),?2342,?234(ef),&uid,@name,^abc,^ex(100),_x,_ef(ef)"
	xn := getNames(p)
	fmt.Println("xxxy:", strings.Join(xn, "|"))
	assert.Equal(t, 15, len(xn))

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
	px := GetExprs("age<20")
	assert.Equal(t, 1, len(px))
	p := px[0]
	assert.Equal(t, 3, len(p))
	assert.Equal(t, "age", p[0])
	assert.Equal(t, "<", p[1])
	assert.Equal(t, "20", p[2])

	px = GetExprs("@age<20")
	assert.Equal(t, 1, len(px))
	p = px[0]
	assert.Equal(t, 3, len(p))
	assert.Equal(t, "@age", p[0])
	assert.Equal(t, "<", p[1])
	assert.Equal(t, "20", p[2])

	px = GetExprs("age=20")
	assert.Equal(t, 1, len(px))
	p = px[0]
	assert.Equal(t, 3, len(p))
	assert.Equal(t, "age", p[0])
	assert.Equal(t, "=", p[1])
	assert.Equal(t, "20", p[2])

	px = GetExprs("age==20")
	assert.Equal(t, 1, len(px))
	p = px[0]
	assert.Equal(t, 3, len(p))
	assert.Equal(t, "age", p[0])
	assert.Equal(t, "==", p[1])
	assert.Equal(t, "20", p[2])

	px = GetExprs("age<=20")
	assert.Equal(t, 1, len(px))
	p = px[0]
	assert.Equal(t, 3, len(p))
	assert.Equal(t, "age", p[0])
	assert.Equal(t, "<=", p[1])
	assert.Equal(t, "20", p[2])

	px = GetExprs("age>=20")
	assert.Equal(t, 1, len(px))
	p = px[0]
	assert.Equal(t, 3, len(p))
	assert.Equal(t, "age", p[0])
	assert.Equal(t, ">=", p[1])
	assert.Equal(t, "20", p[2])

	px = GetExprs("age!=20")
	assert.Equal(t, 1, len(px))
	p = px[0]
	assert.Equal(t, 3, len(p))
	assert.Equal(t, "age", p[0])
	assert.Equal(t, "!=", p[1])
	assert.Equal(t, "20", p[2])

	px = GetExprs("age")
	assert.Equal(t, 0, len(px))

}
func TestExprs(t *testing.T) {
	px := GetExprs("age!=20&name==3")
	assert.Equal(t, 2, len(px))
	p := px[0]
	assert.Equal(t, 3, len(p))
	assert.Equal(t, "age", p[0])
	assert.Equal(t, "!=", p[1])
	assert.Equal(t, "20", p[2])

	p = px[1]
	assert.Equal(t, 3, len(p))
	assert.Equal(t, "name", p[0])
	assert.Equal(t, "==", p[1])
	assert.Equal(t, "3", p[2])

}
func TestKDNO(t *testing.T) {
	reg := regexp.MustCompile(`^(\w+)[，,]?(\w*)[，,]?(\w*)$`)
	names := reg.FindAllStringSubmatch("SR000123122,09777343333", -1)
	assert.Equal(t, 1, len(names))
	assert.Equal(t, "SR000123122", names[0][1:][0])
	assert.Equal(t, "09777343333", names[0][1:][1])
	fmt.Println(names)

	names = reg.FindAllStringSubmatch("SR000123122", -1)
	assert.Equal(t, 1, len(names))
	assert.Equal(t, "SR000123122", names[0][1:][0])
	fmt.Println(names)
}
func TestVx(t *testing.T) {
	http, err := http.NewHTTPClient()
	assert.Equal(t, nil, err)
	content, status, err := http.Get("http://www.csdn.net")
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, status)

	reg := regexp.MustCompile(`<title>(.+)</title>`)
	titles := reg.FindAllStringSubmatch(content, -1)
	assert.Equal(t, 1, len(titles))

	assert.Equal(t, "企业微信文档", titles[0][0])

}
func TestGetTableSettingInfo(t *testing.T) {
	v := getTableSettingInfo(&Line{Text: `6. 文档[ws_dev_docs]<lstbar(添加文档,add:ws_dev_docs,style:{})>`})
	assert.Equal(t, "", v)

	v = getTableSettingInfo(&Line{Text: `6. 文档[ws_dev_docs]<lstbar(添加文档,add:ws_dev_docs,style:{})>(a:b)`})
	assert.Equal(t, "a:b", v)

	v = getTableSettingInfo(&Line{Text: `6. 文档[ws_dev_docs](a:b)`})
	assert.Equal(t, "a:b", v)

	v = getTableSettingInfo(&Line{Text: `6. 文档[ws_dev_docs](a:{@c:&d})`})
	assert.Equal(t, "a:{@c:&d}", v)
}
