package data

import (
	"fmt"
	"testing"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/assert"
)

func TestV(t *testing.T) {
	k := "tp(cut,le,4)"
	c, p, v := md.GetConsByTag("tp", k)
	assert.Equal(t, "cut", c)
	assert.Equal(t, "le", p)
	assert.Equal(t, "4", v)
}
func TestV2(t *testing.T) {
	k := "tp(cut,le|4)"
	c, p, v := md.GetConsByTag("tp", k)
	assert.Equal(t, "cut", c)
	assert.Equal(t, "le|4", p)
	assert.Equal(t, "", v)
}
func TestColor(t *testing.T) {
	c, p, v := md.GetConsByTagIgnorecase("color", "color")
	assert.Equal(t, "color", c)
	assert.Equal(t, "", p)
	assert.Equal(t, "", v)
}
func TestIdex(t *testing.T) {
	fc, _, _ := md.GetConsByTagIgnorecase(fmt.Sprintf("%sidx", "l"), "lidx(1)")
	assert.Equal(t, "1", fc)

}
func TestTP(t *testing.T) {
	fc, f1, f2 := md.GetConsByTagIgnorecase("tp", "tp(tabs,order_status_tabs,x)")
	assert.Equal(t, "tabs", fc)
	assert.Equal(t, "order_status_tabs", f1)
	assert.Equal(t, "x", f2)
}
func TestFltrMySQLType(t *testing.T) {
	assert.Equal(t, "varchar(32)", mySQLType("varchar2", 32, 0))
	assert.Equal(t, "varchar(64)", mySQLType("varchar", 64, 0))
	assert.Equal(t, "decimal(10,5)", mySQLType("number", 10, 5))
	assert.Equal(t, "bigint", mySQLType("number", 20, 0))
	assert.Equal(t, "bigint(22)", mySQLType("number", 22, 0))
	assert.Equal(t, "tinyint", mySQLType("number", 1, 0))
	assert.Equal(t, "int", mySQLType("number", 4, 0))

}

func TestMYSQLType(t *testing.T) {
	plst := md.GetIdx("idx", "idx", "idx(idx_order_info,1)", "IDX(idx_order_info2,2)", "l", "lidx(idx_order_info,1)")
	assert.Equal(t, 2, len(plst))
	assert.Equal(t, "idx_order_info", plst[0][0])
	assert.Equal(t, "1", plst[0][1])
	assert.Equal(t, "idx_order_info2", plst[1][0])
	assert.Equal(t, "2", plst[1][1])
}
func TestExpr(t *testing.T) {
	expr := str2Expr("@age>20")[0]
	assert.Equal(t, "", expr.Name)
	assert.Equal(t, "age", expr.Field)
	assert.Equal(t, ">", expr.Symbol)
	assert.Equal(t, "20", expr.Value)

	expr = str2Expr("age>20")[0]
	assert.Equal(t, "age", expr.Name)
	assert.Equal(t, "", expr.Field)
	assert.Equal(t, ">", expr.Symbol)
	assert.Equal(t, "20", expr.Value)
}
