package md

import (
	"bufio"
	"strings"
	"testing"

	"github.com/micro-plat/lib4go/assert"
)

func TestTBConf(t *testing.T) {
	row := `
	### 4. 角色表[^sso_role_info]
	
	| 字段名      | 类型        |      默认值       | 为空  |  约束   | 描述                 |
	| ----------- | ----------- | :---------------: | :---: | :-----: | :------------------- |
	| role_id     | bigint(20)  |                   |  否   | PK,SEQ | 角色id(角色)               |
	| name        | varchar(64) |                   |  否   |   UNQ   | 角色名称             |`

	reader, err := readMarkdownByReader(bufio.NewReader(strings.NewReader(row)))
	assert.Equal(t, err, nil)
	tbs, err := Lines2Table(line2TbLines(reader))
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(tbs))
	assert.Equal(t, "sso_role_info", tbs[0].Name.Raw)
	assert.Equal(t, strings.HasPrefix(tbs[0].Name.OName, "^"), tbs[0].Exclude)
	assert.Equal(t, "角色表", tbs[0].Desc)
	assert.Equal(t, 2, len(tbs[0].Rows))

	assert.Equal(t, "role_id", tbs[0].Rows[0].Name)
	assert.Equal(t, "角色id(角色)", tbs[0].Rows[0].Desc.Raw)
	assert.Equal(t, 2, len(tbs[0].Rows[0].Constraints))
	assert.Equal(t, "PK", tbs[0].Rows[0].Constraints[0])
	assert.Equal(t, "SEQ", tbs[0].Rows[0].Constraints[1])

	conf := `
	### 4. 角色表[^sso_role_info]
	
	| 字段名      |  约束   | 描述                 |
	| ----------- | :-----: | :------------------- |
	| role_id     |  DI | 角色id               |
	| name        |    DN,Q   | 角色名称             |`

	confReader, err := readMarkdownByReader(bufio.NewReader(strings.NewReader(conf)))
	assert.Equal(t, err, nil)
	ctbs, err := Lines2Conf(line2ConfLines(confReader))
	assert.Equal(t, err, nil)

	//重置配置
	tbs = tbs.resetConf(ctbs)
	assert.Equal(t, 1, len(tbs[0].Rows[0].Constraints))
	assert.Equal(t, "DI", tbs[0].Rows[0].Constraints[0])

	assert.Equal(t, 2, len(tbs[0].Rows[1].Constraints))
	assert.Equal(t, "DN", tbs[0].Rows[1].Constraints[0])
	assert.Equal(t, "Q", tbs[0].Rows[1].Constraints[1])
}

func TestXxx(t *testing.T) {
	getFiles("/root/yanglei/*.md")
}
