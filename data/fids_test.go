package data

import (
	"fmt"
	"strings"
	"testing"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/assert"
	"github.com/micro-plat/lib4go/types"
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
	lst := createOptrs("", "view(项目,tab:wp_project_info,x:y),view(计划,tab:wp_plan_info)", "view")
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

	row := `
		### 4. 角色表[sso_role_info]

	| 字段名      | 类型        |      默认值       | 为空  |  约束   | 描述                 |
	| ----------- | ----------- | :---------------: | :---: | :-----: | :------------------- |
	| role_id     | bigint(20)  |                   |  否   | PK,SEQ,DI | 角色id               |
	| name        | varchar(64) |                   |  否   |   UNQ,DN  | 角色名称             |
	| price      | decimal(10,5)  |         0         |  是   |   alias(a,b,c),@change(name:#id;pwd:#id)      | 价格 |
	| status      | tinyint(1)  |         0         |  否   |         | 状态 1: 禁用 0: 正常 |
	| create_time | datetime    | sysdate |  否   |         | 创建时间             |`

	//转换md文件
	otbls, err := md.Md2TbByContent(row)
	assert.Equal(t, err, nil)
	//过滤表格
	tbls := otbls.Filter("", true)
	tbs := NewTables(tbls)

	assert.Equal(t, 1, len(tbs))

	tab := tbs[0]
	columns := tab.Columns.GetColumnsBy("alias")
	assert.Equal(t, 1, len(columns))
	assert.Equal(t, "a", columns[0].GetOtherCmpntValue("alias"))
	a, b, c := columns[0].GetOtherCmpntValues("alias")
	assert.Equal(t, "a", a)
	assert.Equal(t, "b", b)
	assert.Equal(t, "c", c)

	assert.Equal(t, "name:#id;pwd:#id", md.GetParamByTag("change", "change(name:#id;pwd:#id)"))
	columnxs := tab.Columns.GetColumns("pricex")

	params := columnxs.GetParams("@change")
	assert.Equal(t, 2, len(params))
	assert.Equal(t, "#id", params["name"])
	assert.Equal(t, "#id", params["pwd"])

}
func TestSplitConst(t *testing.T) {
	xfm := "1-1,2-2,3-4,4-5,5-5"
	fileds := []string{"name", "sex", "dept", "birth", "addr"}
	cfield := "content"
	input := types.NewXMap()
	input.SetValue("dept", "研发")
	input.SetValue("content", `杨磊,男,1983-09-11,成都,华阳
	杨展硕,男,2016-08-06,北京,海淀`)
	xm := TransformBatchContent(fileds, cfield, xfm, input)
	assert.Equal(t, 2, len(xm))
}
func TestSplitConst2(t *testing.T) {
	xfm := "1-2,2-2,2-3,3-4,4-5"
	fileds := []string{"source", "name", "content", "pmaster", "sponsor"}
	cfield := "bcontent"
	input := types.NewXMap()
	input.SetValue("source", "研发")
	input.SetValue(cfield, `REQ090701	除法定价	舒迪	陈丽君
	REQ090702	审批改造（客户单位可以发起审批）	陈艳	王春禹
	REQ090703	供应商修改商品审核	黄一航、杨兢	王佳妮
	REQ090704	积分发放优化	运营伙伴	王春禹
	REQ090705	优惠券需求	杨xx	陈丽君
	REQ090706	淘金需求	李建明	王春禹
	REQ090707	选品优化	黄一航	陈丽君
	SYS090708	小程序用户登录频繁掉线优化	汪海军	无
	REQ090709	去惠购、蜀道活动页配置运营需求(前端调整)	范思颖	陈丽君
	REQ090710	增加供应商联系方式	陈艳	许姣
	REQ090711	供应商售后订单优化	刘媛	许姣`)
	xm := TransformBatchContent(fileds, cfield, xfm, input)
	assert.Equal(t, 2, len(xm))
}

// TransformBatchContent 转换指定格式数据为map数组
// fileds := []string{"name", "sex", "dept", "birth", "addr"}
// cfield := "content"
// xfm := "1-1,2-2,3-4,4-5,5-5"
// input := types.NewXMap()
// input.SetValue("dept", "研发")
// input.SetValue("content", `杨磊,男,1983-09-11,成都,华阳
// 杨展硕,男,2016-08-06,北京,海淀`)
func TransformBatchContent(fileds []string, cfield string, xfm string, input types.XMap) types.XMaps {
	content := input.GetString(cfield)
	lines := strings.Split(content, "\n")
	nmaps := types.NewXMaps()
	idxs := getArray(xfm)
	for _, line := range lines {
		if strings.Trim(line, "") == "" {
			continue
		}
		//设置字段值
		nmap := types.NewXMap()
		for _, f := range fileds {
			if input.Has(f) {
				nmap.SetValue(f, strings.Trim(input.GetString(f), "\t"))
			}
		}
		//处理多行文本数据
		vlines := splitLines(line)
		for _, idx := range idxs {
			field := types.GetStringByIndex(fileds, idx[1]-1)
			value := strings.Trim(types.GetStringByIndex(vlines, idx[0]-1), "\t")
			if nmap.Has(field) {
				value = nmap.GetString(field) + value
			}
			nmap.SetValue(field, value)
		}
		nmaps.Append(nmap)
	}
	return nmaps
}

// 转换为二维数组 1-1,2-1,3-2,4-3
func getArray(f string) [][]int {
	lines := strings.Split(f, ",")
	lst := make([][]int, 0, len(lines))
	for i, line := range lines {
		xidx := strings.Split(line, "-")
		if len(xidx) == 2 {
			lst = append(lst, []int{
				types.GetInt(types.GetStringByIndex(xidx, 0), i+1),
				types.GetInt(types.GetStringByIndex(xidx, 1), i+1),
			})
		}
	}
	return lst
}
func splitLines(line string) []string {
	n := strings.Trim(strings.Trim(line, "\t"), " ")
	n = strings.ReplaceAll(n, "\t", ",")
	n = strings.ReplaceAll(n, " ", ",")
	n = strings.ReplaceAll(n, ",,", ",")
	return strings.Split(n, ",")
}
