package md

import (
	"bufio"
	"strings"
	"testing"

	"github.com/micro-plat/lib4go/assert"
)

func TestTB(t *testing.T) {
	row := `
	### 4. 角色表[^sso_role_info]
	
	| 字段名      | 类型        |      默认值       | 为空  |  约束   | 描述                 |
	| ----------- | ----------- | :---------------: | :---: | :-----: | :------------------- |
	| role_id     | bigint(20)  |                   |  否   | PK,SEQ | 角色id               |
	| name        | varchar(64) |                   |  否   |   UNQ   | 角色名称             |
	| price      | decimal(10,5)  |         0         |  是   |         | 价格 |
	| status      | tinyint(1)  |         0         |  否   |         | 状态 1: 禁用 0: 正常 |
	| create_time | datetime    | sysdate |  否   |         | 创建时间             |`

	reader, err := readMarkdownByReader(bufio.NewReader(strings.NewReader(row)))
	assert.Equal(t, err, nil)
	tbs, err := Lines2Table(line2Lines(reader))
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(tbs))
	assert.Equal(t, "sso_role_info", tbs[0].Name.Raw)
	assert.Equal(t, strings.HasPrefix(tbs[0].Name.OName, "^"), tbs[0].Exclude)
	assert.Equal(t, "角色表", tbs[0].Desc)
	assert.Equal(t, 5, len(tbs[0].Rows))

	assert.Equal(t, "role_id", tbs[0].Rows[0].Name)
	assert.Equal(t, "bigint(20)", tbs[0].Rows[0].Type.Raw)
	assert.Equal(t, "bigint", tbs[0].Rows[0].Type.Name)
	assert.Equal(t, 20, tbs[0].Rows[0].Type.Len)
	assert.Equal(t, 0, tbs[0].Rows[0].Type.DLen)
	assert.Equal(t, "", tbs[0].Rows[0].DefValue)
	assert.Equal(t, false, tbs[0].Rows[0].AllowNull)
	assert.Equal(t, 2, len(tbs[0].Rows[0].Constraints))
	assert.Equal(t, "PK", tbs[0].Rows[0].Constraints[0])
	assert.Equal(t, "SEQ", tbs[0].Rows[0].Constraints[1])
	assert.Equal(t, "角色id", tbs[0].Rows[0].Desc.Raw)
	assert.Equal(t, "角色id", tbs[0].Rows[0].Desc.Name)

	assert.Equal(t, "name", tbs[0].Rows[1].Name)
	assert.Equal(t, "varchar(64)", tbs[0].Rows[1].Type.Raw)
	assert.Equal(t, "varchar", tbs[0].Rows[1].Type.Name)
	assert.Equal(t, 64, tbs[0].Rows[1].Type.Len)
	assert.Equal(t, 0, tbs[0].Rows[1].Type.DLen)
	assert.Equal(t, "", tbs[0].Rows[1].DefValue)
	assert.Equal(t, false, tbs[0].Rows[1].AllowNull)
	assert.Equal(t, 1, len(tbs[0].Rows[1].Constraints))
	assert.Equal(t, "UNQ", tbs[0].Rows[1].Constraints[0])
	assert.Equal(t, "角色名称", tbs[0].Rows[1].Desc.Raw)
	assert.Equal(t, "角色名称", tbs[0].Rows[1].Desc.Name)

	assert.Equal(t, "price", tbs[0].Rows[2].Name)
	assert.Equal(t, "decimal(10,5)", tbs[0].Rows[2].Type.Raw)
	assert.Equal(t, "decimal", tbs[0].Rows[2].Type.Name)
	assert.Equal(t, 10, tbs[0].Rows[2].Type.Len)
	assert.Equal(t, 5, tbs[0].Rows[2].Type.DLen)
	assert.Equal(t, "0", tbs[0].Rows[2].DefValue)
	assert.Equal(t, true, tbs[0].Rows[2].AllowNull)
	assert.Equal(t, 0, len(tbs[0].Rows[2].Constraints))
	assert.Equal(t, "价格", tbs[0].Rows[2].Desc.Raw)
	assert.Equal(t, "价格", tbs[0].Rows[2].Desc.Name)

	assert.Equal(t, "status", tbs[0].Rows[3].Name)
	assert.Equal(t, "tinyint(1)", tbs[0].Rows[3].Type.Raw)
	assert.Equal(t, "tinyint", tbs[0].Rows[3].Type.Name)
	assert.Equal(t, 1, tbs[0].Rows[3].Type.Len)
	assert.Equal(t, 0, tbs[0].Rows[3].Type.DLen)
	assert.Equal(t, "0", tbs[0].Rows[3].DefValue)
	assert.Equal(t, false, tbs[0].Rows[3].AllowNull)
	assert.Equal(t, 0, len(tbs[0].Rows[3].Constraints))
	assert.Equal(t, "状态 1: 禁用 0: 正常", tbs[0].Rows[3].Desc.Raw)
	assert.Equal(t, "状态", tbs[0].Rows[3].Desc.Name)

	assert.Equal(t, "create_time", tbs[0].Rows[4].Name)
	assert.Equal(t, "datetime", tbs[0].Rows[4].Type.Raw)
	assert.Equal(t, "datetime", tbs[0].Rows[4].Type.Name)
	assert.Equal(t, 0, tbs[0].Rows[4].Type.Len)
	assert.Equal(t, 0, tbs[0].Rows[4].Type.DLen)
	assert.Equal(t, "sysdate", tbs[0].Rows[4].DefValue)
	assert.Equal(t, false, tbs[0].Rows[4].AllowNull)
	assert.Equal(t, 0, len(tbs[0].Rows[4].Constraints))
	assert.Equal(t, "创建时间", tbs[0].Rows[4].Desc.Raw)
	assert.Equal(t, "创建时间", tbs[0].Rows[4].Desc.Name)
}
func TestV(t *testing.T) {
	ln := &Line{Text: "系统信息[sso_system_info]<lst(权限,link>/right/info,x),lst(启用,dialog>/right/save,m)>"}
	v := getTableExtInfo(ln)
	assert.Equal(t, "lst(权限,link>/right/info,x),lst(启用,dialog>/right/save,m)", v)

	ln = &Line{Text: "商户信息[ots_merchant_info]<view(货架,tab>ots_merchant_shelf)>"}
	v = getTableExtInfo(ln)
	assert.Equal(t, "view(货架,tab>ots_merchant_shelf)", v)
}
