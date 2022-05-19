package data

import (
	"testing"

	"github.com/micro-plat/lib4go/assert"
)

func TestFids(t *testing.T) {
	fd := newFids(3)
	assert.Equal(t, "aaa", fd.Current())
	assert.Equal(t, "aab", fd.Next())
}
func TestFids2(t *testing.T) {
	fd := &fids{charIndex: []int{0, 0, 25}}
	assert.Equal(t, "aaz", fd.Current())
	assert.Equal(t, "aba", fd.Next())
}
func TestType(t *testing.T) {
	v := getShowTypeName("Q", "int")
	assert.Equal(t, "number", v)
}
func TestGetExtOpt(t *testing.T) { //>/right/info,x
	lst := newOperations("lst(权限,link:/right/info,x:y),lst(启用,dialog:/right/save,m),pnl(启用,dialog:/right/save,m)", "lst")
	assert.Equal(t, 2, len(lst))
	assert.Equal(t, "LINK", lst[0].Name)
	assert.Equal(t, "/right/info", lst[0].URL)
	assert.Equal(t, "x", lst[0].RwName)
	assert.Equal(t, "y", lst[0].FwName)
	assert.Equal(t, "权限", lst[0].Label)

	assert.Equal(t, "DIALOG", lst[1].Name)
	assert.Equal(t, "/right/save", lst[1].URL)
	assert.Equal(t, "m", lst[1].RwName)
	assert.Equal(t, "", lst[1].FwName)
	assert.Equal(t, "启用", lst[1].Label)
}
