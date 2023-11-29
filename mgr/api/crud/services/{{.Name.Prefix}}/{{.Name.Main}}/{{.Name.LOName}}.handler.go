//go:build ignore

package {-{.Name.Main}-}
{-{- $table := .}-}
{-{- $ccols :=  $table.Columns.GetFontCreateColumns.GetRequiedColumns}-}
{-{- $ucols :=  $table.Columns.GetFontUpdateColumns.GetRequiedColumns}-}
{-{- $switchs :=$table.Columns.GetAllListColumns.GetValidColumns.GetByCmpnt "switch"}-}
{-{- $batinserts :=  $table.Optrs.BarOpts.GetByCmptName  "batinsert"}-}
{-{- $hasStatic :=  gt $table.Columns.GetAll.GetStaticColumns.Len 0}-}

import (
	"net/http"
	{-{- if or (gt $table.Columns.GetAllCreateColumns.Len 0) (gt $table.Columns.GetAllUpdateColumns.Len 0)}-}
	"strings"
	{-{- end}-}
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	{-{- if gt $table.Columns.GetAllQueryColumns.Len 0}-}
	"github.com/micro-plat/lib4go/types"
	{-{- end}-}
	{-{- if eq true $hasStatic}-}
	"{-{$table.PkgName}-}/modules/biz/member"
	{-{- end}-}
	{-{- if eq true $table.Conf.WriteLog}-}
	"{-{$table.PkgName}-}/modules/biz/system"
	"fmt"
	{-{- end}-}
	{-{- if gt $batinserts.Len 0}-}
	"{-{$table.PkgName}-}/services/utils"
	{-{- end}-}
)


//AuditInfoHandler 获取{-{.Desc}-}处理服务
type {-{.Name.CName}-}Handler struct {
	insertRequiredFields []string
	updateRequiredFields []string
	pkRequiredFields []string
	switchRequiredFields []string
}

func New{-{.Name.CName}-}Handler() *{-{.Name.CName}-}Handler {
	return &{-{.Name.CName}-}Handler{
		insertRequiredFields:[]string{ {-{- range $i,$v := $ccols}-}"{-{$v.Name}-}"{-{if lt $i $ccols.SLen}-},{-{end}-}{-{end}-}},
		updateRequiredFields:[]string{ {-{- range $i,$v :=  $ucols}-}"{-{$v.Name}-}"{-{if lt $i $ucols.SLen}-},{-{end}-}{-{end}-}},
		pkRequiredFields:[]string{ {-{- range $i,$v :=  $table.Columns.GetPKColumns}-}"{-{$v.Name}-}"{-{if lt $i  $table.Columns.GetPKColumns.SLen}-},{-{end}-}{-{end}-}},
		switchRequiredFields:[]string{ {-{- range $i,$v :=  $switchs }-}"{-{$v.Name}-}"{-{if lt $i $switchs.SLen}-},{-{end}-}{-{end}-}},
	}
}
{-{- $leColumns:=  $table.Columns.GetAllListColumns }-}
{-{- if gt $leColumns.Len 0}-}

//GetHandle  查询{-{.Desc}-}数据
func (u *{-{.Name.CName}-}Handler) GetHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------查询{-{.Desc}-}数据--------")

	ctx.Log().Info("1.检查必须字段")
	if err := ctx.Request().Check(u.pkRequiredFields...); err != nil {
		return err
	}
	ctx.Log().Info("2.查询数据")
	items, err := hydra.C.DB().GetRegularDB().Query(get{-{.Name.CName}-}, ctx.Request().GetMap())
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据出错:%+v", err)
	}
	return items.Get(0)
}
{-{- end}-}

{-{- if gt $table.Columns.GetAllUpdateColumns.GetValidColumns.Len 0}-}

//PutHandle  修改{-{.Desc}-}数据
func (u *{-{.Name.CName}-}Handler) PutHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------保存{-{.Desc}-}数据--------")

	ctx.Log().Info("1.检查必须字段")
	if err := ctx.Request().Check(u.updateRequiredFields...); err != nil {
		return err
	}
	
	ctx.Log().Info("2.修改数据")
	{-{- $memberClus :=  $table.Columns.GetAllUpdateColumns.GetStaticColumns}-}
	{-{- if gt $memberClus.Len 0}-}
	member, err := member.GetMemberState(ctx)
	if err != nil {
		return err
	}
	{-{- range $i,$v := $memberClus}-}
	ctx.Request().SetValue("{-{$v.Name}-}",member.Get("{-{$v.Ext.Name}-}"))
	{-{- end}-}
	{-{- end}-}
	rx, err := hydra.C.DB().GetRegularDB().Execute(update{-{.Name.CName}-}, ctx.Request().GetMap())
	if err != nil{
		if strings.Contains(err.Error(),"Error 1062"){
			return errs.NewErrorf(909, "数据重复，请修改后提交")
		}
		return errs.NewErrorf(http.StatusNotExtended, "数据出错:%+v,row:%d", err,rx)
	}
	if rx <= 0 {
		return errs.NewResult(204, nil)
	}
	{-{- if eq true $table.Conf.WriteLog}-}
	ctx.Log().Info("3.保存用户日志")
	name := ctx.Request().GetString("{-{$table.Enum.Name}-}")
	system.SaveLog(ctx,fmt.Sprintf("修改{-{$table.Desc}-} %s",name),string(ctx.Request().Marshal()))
	{-{- end}-}
	return
}
{-{- end}-}

{-{- if gt $table.Columns.GetAllCreateColumns.GetValidColumns.Len 0}-}

//PostHandle  保存{-{.Desc}-}数据
func (u *{-{.Name.CName}-}Handler) PostHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------保存{-{.Desc}-}数据--------")

	ctx.Log().Info("1.检查必须字段")
	if err := ctx.Request().Check(u.insertRequiredFields...); err != nil {
		return err
	}
	
	ctx.Log().Info("2.添加新数据")
	{-{- $memberClus :=  $table.Columns.GetAllCreateColumns.GetStaticColumns}-}
	{-{- if gt $memberClus.Len 0}-}
	member, err := member.GetMemberState(ctx)
	if err != nil {
		return err
	}
	{-{- range $i,$v := $memberClus}-}
	{-{- if $v.HasCmpnt "fc"}-}
	ctx.Request().SetValue("{-{$v.Name}-}",types.GetString(ctx.Request().GetValue("{-{$v.Name}-}"),member.Get("{-{$v.Ext.Name}-}")))
	{-{- else}-}
	ctx.Request().SetValue("{-{$v.Name}-}",member.Get("{-{$v.Ext.Name}-}"))
	{-{- end}-}
	{-{- end}-}
	{-{- end}-}
	rx, err := hydra.C.DB().GetRegularDB().Execute(insert{-{.Name.CName}-}, ctx.Request().GetMap())
	if err != nil{
		if strings.Contains(err.Error(),"Error 1062"){
			return errs.NewErrorf(909, "数据重复，请修改后提交")
		}
		return errs.NewErrorf(http.StatusNotExtended, "数据出错:%+v,row:%d", err,rx)
	}
	if rx <= 0 {
		return errs.NewResult(204, nil)
	}
	{-{- if eq true $table.Conf.WriteLog}-}
	ctx.Log().Info("3. 保存用户日志")
	name := ctx.Request().GetString("{-{$table.Enum.Name}-}")
	system.SaveLog(ctx,fmt.Sprintf("添加{-{$table.Desc}-} %s",name),string(ctx.Request().Marshal()))
	{-{- end}-}
	return
}
{-{- end}-}

{-{- if gt $table.Columns.GetAllQueryColumns.Len 0}-}
{-{- $qbar:=$table.Optrs.QueryOptrs.First}-}
{-{- $pkName := $table.Columns.GetPKColumns.First}-}
{-{- $treeNode :=  $qbar.GetParam "treeNode" ""}-}
{-{- $treeId :=  $qbar.GetParam "treeId" $pkName.Name}-}
{-{- $treeDef :=  $qbar.GetParam "firstNodeValue" "0"}-}

//QueryHandle  获取{-{.Desc}-}列表数据
func (u *{-{.Name.CName}-}Handler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取{-{.Desc}-}数据列表--------")

	ctx.Log().Info("1.查询数据条数")
	{-{- $memberClus :=  $table.Columns.GetAllQueryColumns.GetStaticColumns}-}
	{-{- if gt $memberClus.Len 0}-}
	member, err := member.GetMemberState(ctx)
	if err != nil {
		return err
	}
	{-{- range $i,$v := $memberClus}-}
	ctx.Request().SetValue("{-{$v.Name}-}",types.GetString(ctx.Request().GetValue("{-{$v.Name}-}"),member.Get("{-{$v.Ext.Name}-}")))
	{-{- end}-}
	{-{- end}-}
	{-{- if eq "" $treeNode}-}
	m := ctx.Request().GetMap()
	m["ps"] = ctx.Request().GetInt("ps", 10) 
	m["offset"] = (ctx.Request().GetInt("pi", 1) - 1) * ctx.Request().GetInt("ps", 10) 
	{-{- else}-}
	ps := ctx.Request().GetInt("ps", 10)
	pi := ctx.Request().GetInt("pi", 1) - 1
	offset := pi * ps
	m := ctx.Request().GetMap()
	m["ps"] = 10000
	m["offset"] = 0
	{-{end}-}
	
	count, err := hydra.C.DB().GetRegularDB().Scalar(get{-{.Name.CName}-}ListCount, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据数量出错:%+v", err)
	}

	rcount:=types.GetInt(count)
	if  rcount== 0 {
		return map[string]interface{}{
			"count":0,
		}
	}

	ctx.Log().Info("2.查询数据列表")
	items, err := hydra.C.DB().GetRegularDB().QueryBatch(get{-{.Name.CName}-}List, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据出错:%+v", err)
	}
	
	{-{- if ne "" $treeNode}-}
	ctx.Log().Info("3.返回结果")
	nodes := getTreeNodes(items,"{-{$treeNode}-}","{-{$treeId}-}")
	return map[string]interface{}{
		"items": getRange(nodes, offset, offset+ps),
		"count":len(nodes),
	}
	{-{- else}-}
	ctx.Log().Info("3.返回结果")
	return map[string]interface{}{
		"items": items,
		"count":rcount,
	}
	{-{- end}-}
}
{-{- if ne "" $treeNode}-}
func getRange(xm types.XMaps, min int, max int) types.XMaps {
	if min >= len(xm) || min < 0 || max <= min {
		return nil
	}
	if max >= len(xm) {
		max = len(xm)
	}
	return xm[min:max]
}
func getTreeNodes(items types.XMaps, pidName string, pkName string) types.XMaps {
	treeMap := map[string]types.XMaps{}
	for _, r := range items {
		pid := r.GetString(pidName)
		if _, ok := treeMap[pid]; !ok {
			treeMap[pid] = make(types.XMaps, 0, 1)
		}
		treeMap[pid] = append(treeMap[pid], r)
	}
	lstMap := getTopNodeMap(treeMap, pidName, pkName)
	for _, m := range lstMap {
		setChildren(m, pkName, treeMap)
	}
	return lstMap
}

// {type:"---",value:"status",name:"状态"}
// {type:"status",value:"0",name:"启用"}
// {type:"status",value:"1",name:"禁用"}

// {type:"---",value:"pmstatus",name:"用处类型"}
// {type:"pmstatus",value:"pmstatus-10",name:"用途"}
// {type:"pmstatus",value:"pmstatus-11",name:"停用"}
// {type:"pmstatus-10",value:"px-10",name:"日用"}
// {type:"pmstatus-10",value:"px-11",name:"办公"}

func getTopNodeMap(treeMap map[string]types.XMaps, tpName string, valueName string) types.XMaps {
	lstMap := types.XMaps{}
	for _, onemap := range treeMap {
		endLoop:
		for _, v := range onemap {
			tp := v.GetString(tpName)
			for _, p0 := range treeMap {
				for _, p := range p0 {
					value2 := p.GetString(valueName)
					if tp == value2 {
						break endLoop
					}
				}
			}
			lstMap = append(lstMap, v)
		}
	}
	return lstMap

}
func setChildren(current types.XMap, idName string, treeMap map[string]types.XMaps) {
	id := current.GetString(idName)
	children, ok := treeMap[id]
	if !ok {
		return
	}
	current.SetValue("children", children)
	for _, v := range children {
		setChildren(v, idName, treeMap)
	}
}
{-{- end}-}
{-{- end}-}
{-{- $dColumns :=  $table.Columns.GetDelColumns}-}
{-{- if gt $dColumns.Len 0}-}

//DeleteHandle  删除{-{.Desc}-}数据
func (u *{-{.Name.CName}-}Handler) DelHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------删除{-{.Desc}-}数据--------")

	ctx.Log().Info("1.检查必须字段")
	if err := ctx.Request().Check(u.pkRequiredFields...); err != nil {
		return err
	}
	
	ctx.Log().Info("2.删除数据")
	rx, err := hydra.C.DB().GetRegularDB().Execute(delete{-{.Name.CName}-}, ctx.Request().GetMap())
	if err == nil && rx == 0{
		return errs.NewResult(204, nil)
	}
	if err != nil||rx <= 0 {
		return errs.NewErrorf(http.StatusNotExtended, "删除数据出错:%+v", err)
	}
	{-{- if eq true $table.Conf.WriteLog}-}
	ctx.Log().Info("3. 保存用户日志")
	name := ctx.Request().GetString("{-{$table.Enum.Name}-}")
	system.SaveLog(ctx,fmt.Sprintf("删除{-{$table.Desc}-} %s",name),string(ctx.Request().Marshal()))
	{-{- end}-}
	return
}
{-{- end}-}

{-{- if gt (len $switchs) 0}-}

//SwitchHandle  切换状态{-{.Desc}-}数据
func (u *{-{.Name.CName}-}Handler) SwitchHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------切换状态{-{.Desc}-}数据--------")

	ctx.Log().Info("1.检查必须字段")
	reqFields := u.pkRequiredFields
	reqFields = append(reqFields,u.switchRequiredFields...)
	if err := ctx.Request().Check(reqFields...); err != nil {
		return err
	}
	
	ctx.Log().Info("2.切换状态新数据")
	rx, err := hydra.C.DB().GetRegularDB().Execute(switch{-{.Name.CName}-}, ctx.Request().GetMap())
	if err == nil && rx == 0{
		return errs.NewResult(204, nil)
	}
	if err != nil || rx == 0 {
		return errs.NewErrorf(http.StatusNotExtended, "数据出错:%+v,row:%d", err,rx)
	}
	{-{- if eq true $table.Conf.WriteLog}-}
	ctx.Log().Info("3. 保存用户日志")
	name := ctx.Request().GetString("{-{$table.Enum.Name}-}")
	system.SaveLog(ctx,fmt.Sprintf("切换状态{-{$table.Desc}-} %s",name),string(ctx.Request().Marshal()))
	{-{- end}-}
	return
}
{-{- end}-}

{-{- $updator :=  $table.Optrs.All.GetByCmptName "batupdator-lstupdator"}-}
{-{- if gt $updator.Len 0}-}
{-{- range $i,$c := $updator}-}

//{-{$c.ReqURL|f_string_2cname}-}Handle  {-{$c.Label}-}
func (u *{-{$table.Name.CName}-}Handler) {-{$c.ReqURL|f_string_2cname}-}Handle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------{-{$c.Label}-}--------")

	ctx.Log().Info("1.检查必须字段")
	reqFields :=[]string{
		{-{ $table.JoinNames $c.FwName true `"` `",`}-}
	}
	if len(reqFields) == 0 {
		return errs.NewError(601, "FwName字段未配置")
	}
	if err := ctx.Request().Check(reqFields...); err != nil {
		return err
	}
	
	ctx.Log().Info("2.修改数据")
	rx, err := hydra.C.DB().GetRegularDB().Execute(updator{-{$table.Name.CName}-}{-{$c.ReqURL}-}, ctx.Request().GetMap())
	if err == nil && rx == 0{
		return errs.NewResult(204, nil)
	}
	if err != nil || rx == 0 {
		return errs.NewErrorf(http.StatusNotExtended, "数据出错:%+v,row:%d", err,rx)
	}
	{-{- if eq true $table.Conf.WriteLog}-}
	ctx.Log().Info("3. 保存用户日志")
	name := ctx.Request().GetString("{-{$table.Enum.Name}-}")
	system.SaveLog(ctx,fmt.Sprintf("更新{-{$table.Desc}-} {-{$c.Label}-} %s",name),string(ctx.Request().Marshal()))
	{-{- end}-}
	return
}
{-{- end}-}
{-{- end}-}

{-{- $barinsert :=  $table.Optrs.BarOpts.GetByCmptName "barinsert"}-}
{-{- if gt $barinsert.Len 0}-}
{-{- range $i,$c := $barinsert}-}

//{-{$c.ReqURL|f_string_2cname}-}Handle  保存{-{.Desc}-}数据
func (u *{-{$table.Name.CName}-}Handler) {-{$c.ReqURL|f_string_2cname}-}Handle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------保存{-{.Desc}-}数据--------")

	ctx.Log().Info("1.检查必须字段")
	reqFields :=[]string{
		{-{ $table.JoinNames $c.RwName true `"` `",`}-}
	}
	if len(reqFields) == 0 {
		return errs.NewError(601, "RwName字段未配置")
	}
	if err := ctx.Request().Check(reqFields...); err != nil {
		return err
	}
	
	ctx.Log().Info("2.添加新数据")
	{-{- $memberClus :=  $table.Columns.GetAllCreateColumns.GetStaticColumns}-}
	{-{- if gt $memberClus.Len 0}-}
	member, err := member.GetMemberState(ctx)
	if err != nil {
		return err
	}
	{-{- range $i,$v := $memberClus}-}
	ctx.Request().SetValue("{-{$v.Name}-}",member.Get("{-{$v.Ext.Name}-}"))
	{-{- end}-}
	{-{- end}-}

	rx, err := hydra.C.DB().GetRegularDB().Execute(barinsert{-{$c.ReqURL|f_string_2cname}-}{-{$table.Name.CName}-}, ctx.Request().GetMap())
	if err != nil{
		if strings.Contains(err.Error(),"Error 1062"){
			return errs.NewErrorf(909, "数据重复，请修改后提交")
		}
		return errs.NewErrorf(http.StatusNotExtended, "数据出错:%+v,row:%d", err,rx)
	}
	if rx <= 0 {
		return errs.NewResult(204, nil)
	}
	{-{- if eq true $table.Conf.WriteLog}-}
	ctx.Log().Info("3. 保存用户日志")
	name := ctx.Request().GetString("{-{$table.Enum.Name}-}")
	system.SaveLog(ctx,fmt.Sprintf("添加{-{$table.Desc}-} %s",name),string(ctx.Request().Marshal()))
	{-{- end}-}
	return
}
{-{- end}-}
{-{- end}-}

{-{- $batinserts :=  $table.Optrs.BarOpts.GetByCmptName  "batinsert"}-}
{-{- if gt $batinserts.Len 0}-}
{-{- range $i,$c := $batinserts}-}

//{-{$c.ReqURL|f_string_2cname}-}Handle  {-{$c.Label}-}
func (u *{-{$table.Name.CName}-}Handler) {-{$c.ReqURL|f_string_2cname}-}Handle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------{-{$c.Label}-}--------")

	ctx.Log().Info("1.检查必须字段")
	reqFields :=[]string{
		{-{ $table.JoinNames $c.RwName true `"` `",`}-}
	}
	if len(reqFields) == 0 {
		return errs.NewError(601, "RwName字段未配置")
	}
	if err := ctx.Request().Check(reqFields...); err != nil {
		return err
	}
	fileds := []string{
		{-{ $table.JoinNames $c.FwName true `"` `",`}-}
	}
	enumsMap := map[string]interface{}{
		{-{- range $i,$c := $table.Columns.GetEnumColumns}-}
		"{-{$c.Name}-}":"{-{$c.Enum.EnumType}-}",
		{-{- end}-}
	}
	{-{ $binsertColumns := $table.Columns.GetColumns "batisert" -}-}
	{-{- $binsertColum := $binsertColumns.First -}-}
	cfield := "{-{- $binsertColum.Name}-}"
	xfm :="{-{- f_string_join ($table.Columns.GetCmpntValues "batisert") ","}-}"
	xms,err := utils.TransformBatchContent(fileds,enumsMap, cfield, xfm, ctx.Request().GetMap())
	if err != nil{
		return err
	}
	if xms.Len() == 0 {
		return errs.NewResult(512, "导入数据格式错误")
	}

	ctx.Log().Info("2.修改数据")
	db, err := hydra.C.DB().GetRegularDB().Begin()
	if err != nil {
		return err
	}
	for _, xm := range xms {
		rx, err := db.Execute(batInsert{-{$table.Name.CName}-}{-{$c.ReqURL}-}, xm)
		if err != nil || rx == 0 {
			db.Rollback()
			return errs.NewErrorf(http.StatusNotExtended, "数据出错:%+v,row:%d", err, rx)
		}
	}
	db.Commit()
	{-{- if eq true $table.Conf.WriteLog}-}
	ctx.Log().Info("3. 保存用户日志")
	name := ctx.Request().GetString("{-{$table.Enum.Name}-}")
	system.SaveLog(ctx,fmt.Sprintf("批量添加{-{$table.Desc}-} {-{$c.Label}-} %s",name),string(ctx.Request().Marshal()))
	{-{- end}-}
	return
}
{-{- end}-}
{-{- end}-}