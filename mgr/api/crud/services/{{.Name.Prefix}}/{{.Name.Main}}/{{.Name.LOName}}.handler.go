//go:build ignore

package {-{.Name.Main}-}
{-{- $table := .}-}
{-{- $ccols := fltrNotNullCols (fltrColumns $table "c")}-}
{-{- $clen := (len $ccols)|minus}-}
{-{- $ucols := fltrNotNullCols  (fltrColumns $table "u")}-}
{-{- $ulen := (len $ucols)|minus}-}
{-{- $pklen := (len $table.PKColumns)|minus}-}
{-{- $switchs := fltrCmpnt $table "switch" "l"}-}
{-{- $slen := (len $switchs)|minus}-}
{-{- $updator := fltrOptrsByCmd $table.BarOpts "lstupdator"}-}
{-{- $hasStatic := fltrHasStaticColumn  $table "bc-bu-bq" "#"}-}
import (
	"net/http"
	{-{- if or (gt (len (fltrColumns $table "c-bc")) 0) (gt (len (fltrColumnsExcludeExt (fltrColumns $table "u"))) 0)}-}
	"strings"
	{-{- end}-}

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	{-{- if gt (len (fltrColumns $table "q-bq")) 0}-}
	"github.com/micro-plat/lib4go/types"
	{-{- end}-}
	{-{- if eq true $hasStatic}-}
	"{-{$table.PkgName}-}/modules/biz/member"
	{-{- end}-}
	{-{- if eq true $table.Conf.WriteLog}-}
	"{-{$table.PkgName}-}/modules/biz/system"
	"fmt"
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
		insertRequiredFields:[]string{ {-{- range $i,$v := $ccols}-}"{-{$v.Name}-}"{-{if lt $i $clen}-},{-{end}-}{-{end}-}},
		updateRequiredFields:[]string{ {-{- range $i,$v :=  $ucols}-}"{-{$v.Name}-}"{-{if lt $i $ulen}-},{-{end}-}{-{end}-}},
		pkRequiredFields:[]string{ {-{- range $i,$v :=  $table.PKColumns}-}"{-{$v.Name}-}"{-{if lt $i $pklen}-},{-{end}-}{-{end}-}},
		switchRequiredFields:[]string{ {-{- range $i,$v :=  $switchs }-}"{-{$v.Name}-}"{-{if lt $i $slen}-},{-{end}-}{-{end}-}},
	}
}
{-{- $leColumns:= fltrColumnsExcludeExt (fltrColumns $table "l-le") }-}
{-{- if gt (len $leColumns) 0}-}

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

{-{- $uColumns := fltrColumnsExcludeExt (fltrColumns $table "u-bu")}-}
{-{- if gt (len $uColumns) 0}-}

//PutHandle  修改{-{.Desc}-}数据
func (u *{-{.Name.CName}-}Handler) PutHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------保存{-{.Desc}-}数据--------")

	ctx.Log().Info("1.检查必须字段")
	if err := ctx.Request().Check(u.updateRequiredFields...); err != nil {
		return err
	}
	
	ctx.Log().Info("2.修改数据")
	{-{- $memberClus := fltrStaticColumn $table "bu" "#"}-}
	{-{- if gt (len $memberClus) 0}-}
	member, err := member.GetMemberState(ctx)
	if err != nil {
		return err
	}
	{-{- range $i,$v := $memberClus}-}
	ctx.Request().SetValue("{-{$v.Name}-}",member.{-{fltrTrim $i "#"}-})
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

{-{- $cColumns := fltrColumnsExcludeExt (fltrColumns $table "c-bc")}-}
{-{- if gt (len $cColumns) 0}-}
//PostHandle  保存{-{.Desc}-}数据
func (u *{-{.Name.CName}-}Handler) PostHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------保存{-{.Desc}-}数据--------")

	ctx.Log().Info("1.检查必须字段")
	if err := ctx.Request().Check(u.insertRequiredFields...); err != nil {
		return err
	}
	
	ctx.Log().Info("2.添加新数据")
	{-{- $memberClus := fltrStaticColumn $table "bc" "#"}-}
	{-{- if gt (len $memberClus) 0}-}
	member, err := member.GetMemberState(ctx)
	if err != nil {
		return err
	}
	{-{- range $i,$v := $memberClus}-}
	ctx.Request().SetValue("{-{$v.Name}-}",member.{-{fltrTrim $i "#"}-})
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

{-{- $qColumns := fltrColumnsExcludeExt (fltrColumns $table "q-bq")}-}
{-{- if gt (len $qColumns) 0}-}
{-{- $qbar:=$table.QueryOptrs|getFirstOptr}-}
{-{- $pkName := $table.PKColumns|getFirstColumns}-}
{-{- $treeNode := fltrOptrPara $qbar "treeNode" ""}-}
//QueryHandle  获取{-{.Desc}-}列表数据
func (u *{-{.Name.CName}-}Handler) QueryHandle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------获取{-{.Desc}-}数据列表--------")

	ctx.Log().Info("1.查询数据条数")
	{-{- $memberClus := fltrStaticColumn $table "bq" "#"}-}
	{-{- if gt (len $memberClus) 0}-}
	member, err := member.GetMemberState(ctx)
	if err != nil {
		return err
	}
	{-{- range $i,$v := $memberClus}-}
	ctx.Request().SetValue("{-{$v.Name}-}",member.{-{fltrTrim $i "#"}-})
	{-{- end}-}
	{-{- end}-}
	m := ctx.Request().GetMap()
	m["ps"] ={-{- if ne "" $treeNode}-} 10000 {-{else}-} ctx.Request().GetInt("ps", 10) {-{end}-}
	m["offset"] = (ctx.Request().GetInt("pi", 1) - 1) * {-{- if ne "" $treeNode}-} 10000 {-{else}-} ctx.Request().GetInt("ps", 10) {-{end}-}

	
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
	items, err := hydra.C.DB().GetRegularDB().Query(get{-{.Name.CName}-}List, m)
	if err != nil {
		return errs.NewErrorf(http.StatusNotExtended, "查询数据出错:%+v", err)
	}
	
	{-{- if ne "" $treeNode}-}
	ctx.Log().Info("3.返回结果")
	return map[string]interface{}{
		"items": getTreeNodes(items,"{-{$treeNode}-}","{-{$pkName.Name}-}"),
		"count":rcount,
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
func getTreeNodes(items types.XMaps, pidName string, pkName string) types.XMaps {
	treeMap := map[string]types.XMaps{}
	lstMap := types.XMaps{}
	for _, r := range items {
		pid := r.GetString(pidName)
		if _, ok := treeMap[pid]; !ok {
			treeMap[pid] = make(types.XMaps, 0, 1)
		}
		treeMap[pid] = append(treeMap[pid], r)
		if pid == "0" {
			lstMap = append(lstMap, r)
		}
	}
	for _, m := range lstMap {
		setChildren(m, pkName, treeMap)
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

{-{- $dColumns := fltrColumnsExcludeExt (fltrColumns $table "d")}-}
{-{- if gt  (len $dColumns) 0}-}
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

{-{- if gt (len $updator) 0}-}
{-{- range $i,$c := $updator}-}
//{-{$c.ReqURL|fltr2CName}-}Handle  {-{$c.Label}-}
func (u *{-{$table.Name.CName}-}Handler) {-{$c.ReqURL|fltr2CName}-}Handle(ctx hydra.IContext) (r interface{}) {

	ctx.Log().Info("--------{-{$c.Label}-}--------")

	ctx.Log().Info("1.检查必须字段")
	reqFields :=[]string{
		{-{flterJoinColumnNames $table $c.FwName `"` `",`}-}
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