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

func TestGetExtOpt(t *testing.T) { //>/right/info,x
	lst := createOptrs("view(项目,tab:wp_project_info,x:y),view(计划,tab:wp_plan_info)", "view")
	assert.Equal(t, 2, len(lst))
	fmt.Println("lst:", lst)
	assert.Equal(t, "TAB", lst[0].Name)
	assert.Equal(t, "wp_project_info", lst[0].URL)
	assert.Equal(t, "x", lst[0].RwName)
	assert.Equal(t, "y", lst[0].FwName)
	assert.Equal(t, "项目", lst[0].Label)
}

func TestRelation(t *testing.T) {

	// row := `
	// 	### 3. 用户信息[mgr_user_info]

	// | 字段名           | 类型        | 默认值  | 为空  |                 约束                  | 描述                         |
	// | ---------------- | ----------- | :-----: | :---: | :-----------------------------------: | :--------------------------- |
	// | user_id          | bigint(10)  |         |  否   |              PK,SEQ,L,DI              | 用户编号                     |
	// | full_name        | varchar(32) |         |  否   |             L,Q,DN,C,U,D              | 姓名                         |
	// | user_name        | varchar(64) |         |  否   |              UNQ,L,C,U,               | 登录名                       |
	// | password         | varchar(32) |         |  否   |               C,tp(pwd)               | 密码                         |
	// | mobile           | varchar(12) |         |  否   |        L,C,U,LE,DE,tp(mobile)         | 电话号码                     |
	// | wx_openid        | varchar(64) |         |  是   |                                       | 微信openid                   |
	// | scorp_id         | number(10)  |         |  是   |             c,u                        | 服务公司编号                 |
	// | role_id          | bigint(10)  |         |  否   |              c,u,LE,sl(role_info,scorp_id)                      | 角色编号                     |

	// 	### 4. 角色表[sso_role_info]

	// 	| 字段名      | 类型        |      默认值       | 为空  |  约束   | 描述                 |
	// 	| ----------- | ----------- | :---------------: | :---: | :-----: | :------------------- |
	// 	| role_id     | bigint(20)  |                   |  否   | PK,SEQ,DI | 角色id               |
	// 	| name        | varchar(64) |                   |  否   |   UNQ,DN  | 角色名称             |
	// 	| price      | decimal(10,5)  |         0         |  是   |         | 价格 |
	// 	| status      | tinyint(1)  |         0         |  否   |         | 状态 1: 禁用 0: 正常 |
	// 	| create_time | datetime    | sysdate |  否   |         | 创建时间             |`

	// //转换md文件
	// otbls, err := md.Md2TbByContent(row)
	// assert.Equal(t, err, nil)
	// // 过滤表格
	// tbls := otbls.Filter("", true)

	// ntbs := FltrUITables(tbls)
	// Caches(ntbs)
	// ResetSelectRelation(ntbs)
	// assert.Equal(t, 1, 2)
	// fmt.Println(ntbs)
}

func TestRelation2(t *testing.T) {

	// row := `
	// 	### 4. 角色表[sso_role_info]

	// | 字段名      | 类型        |      默认值       | 为空  |  约束   | 描述                 |
	// | ----------- | ----------- | :---------------: | :---: | :-----: | :------------------- |
	// | role_id     | bigint(20)  |                   |  否   | PK,SEQ,DI | 角色id               |
	// | name        | varchar(64) |                   |  否   |   UNQ,DN  | 角色名称             |
	// | price      | decimal(10,5)  |         0         |  是   |         | 价格 |
	// | status      | tinyint(1)  |         0         |  否   |         | 状态 1: 禁用 0: 正常 |
	// | create_time | datetime    | sysdate |  否   |         | 创建时间             |`

	// //转换md文件
	// otbls, err := md.Md2TbByContent(row)
	// assert.Equal(t, err, nil)
	// //过滤表格
	// tbls := otbls.Filter("", true)

	// ntbs := FltrUITables(tbls)
	// Caches(ntbs)
	// ResetSelectRelation(ntbs)
	// assert.Equal(t, ntbs[0].Rows[0].AllowNull, false)
}
