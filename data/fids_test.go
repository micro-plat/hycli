package data

import (
	"fmt"
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
	lst := newOperations("", "view(项目,tab:wp_project_info),view(计划,tab:wp_plan_info)", "view")
	assert.Equal(t, 2, len(lst))
	fmt.Println("lst:", lst)
	assert.Equal(t, "TAB", lst[0].Name)
	assert.Equal(t, "/right/info", lst[0].URL)
	assert.Equal(t, "x", lst[0].RwName)
	assert.Equal(t, "y", lst[0].FwName)
	assert.Equal(t, "权限", lst[0].Label)
}
