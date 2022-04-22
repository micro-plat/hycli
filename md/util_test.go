package md

import (
	"fmt"
	"regexp"
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
	tp := GetTPName("tp", "tp(bl)")
	assert.Equal(t, "bl", tp)

	tp = GetTPName("l", "l(bl)")
	assert.Equal(t, "bl", tp)

	tp = GetTPName("DI", "DI")
	assert.Equal(t, "", tp)

}
func TestSelect(t *testing.T) {
	s := GetSelectName("status", "sl")
	assert.Equal(t, "status", s)
}
func TestSelect2(t *testing.T) {
	s := GetSelectName("status", "sl(bool)")
	assert.Equal(t, "bool", s)
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
	v := getNames("IDX(idx_sso_sysinfo_urlident),c,U ,V ")
	assert.Equal(t, 4, len(v))
	assert.Equal(t, "U", v[2])

}
func TestGetExtOpt(t *testing.T) { //>/right/info,x
	lst := GetExtOpt("lst(权限,link>/right/info,x),lst(启用,dialog>/right/save,m),pnl(启用,dialog>/right/save,m)", "lst")
	assert.Equal(t, "", lst)
}