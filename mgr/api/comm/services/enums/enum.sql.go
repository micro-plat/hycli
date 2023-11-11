//go:build ignore

package enums

{-{- $etable := .}-}
var enumSQL = map[string]string{
	{-{- range $i,$v:=$etable}-}
	{-{- if ne "" $v.Enum.EnumType}-}
	"{-{$v.Enum.EnumType}-}":"select {-{$v.Enum.Id}-} value,{-{if ne "" $v.Enum.PID}-} {-{$v.Enum.PID}-} pid, {-{end}-}{-{$v.Enum.Name}-} name,{-{- range $j,$v:=$v.Enum.DEColumns}-}{-{$v.Name}-} {-{$v.Desc}-},{-{end}-}{-{if ne "" $v.Enum.Status}-}{-{$v.Enum.Status}-} status,{-{end}-}'{-{$v.Enum.EnumType}-}' type  {-{if ne "" $v.Enum.Group}-} ,{-{$v.Enum.Group}-} `group` {-{end}-} from {-{$v.Name.Raw}-} where 1=1   {-{if ne "" $v.Enum.Expire}-} and {-{$v.Enum.Expire}-} >= DATE_FORMAT(now(),'%Y-%m-%d'){-{end}-}{-{if ne "" $v.Enum.SortName}-} order by {-{$v.Enum.SortName}-} {-{$v.Enum.SortType}-} {-{end}-}",
	{-{- end}-}
	{-{- end}-}
}

var unspecifiedEnum = []string{
	{-{- range $i,$v:=$etable}-}
	{-{- if and (ne "" $v.Enum.EnumType) (eq true $v.Enum.Multiple)}-}
	"select {-{$v.Enum.Id}-} value,{-{if ne "" $v.Enum.PID}-} {-{$v.Enum.PID}-} pid, {-{end}-}{-{$v.Enum.Name}-} name,{-{- range $j,$v:=$v.Enum.DEColumns}-}{-{$v.Name}-} {-{$v.Desc}-},{-{end}-}{-{if ne "" $v.Enum.Status}-}{-{$v.Enum.Status}-} status,{-{end}-} {-{$v.Enum.Type}-} type {-{if ne "" $v.Enum.Group}-} ,{-{$v.Enum.Group}-} `group` {-{end}-} from {-{$v.Name.Raw}-} where {-{$v.Enum.Type}-}=if(@type='',{-{$v.Enum.Type}-},@type)  {-{if ne "" $v.Enum.Expire}-} and {-{$v.Enum.Expire}-} >= DATE_FORMAT(now(),'%Y-%m-%d'){-{end}-}{-{if ne "" $v.Enum.SortName}-} order by {-{$v.Enum.SortName}-} {-{$v.Enum.SortType}-} {-{end}-}",
	{-{- end}-}
	{-{- end}-}
}
var sqls = make([]string, 0, 1)

func getSQLs(tp string) []string {
	//返回所有枚举
	if tp == "" || tp == "*" {
		return sqls
	}

	//根据枚举查询具体提供SQL
	vsqls, ok := enumSQL[tp]
	if ok {
		return []string{vsqls}
	}

	//指定未明确类型的枚举
	return unspecifiedEnum
}
