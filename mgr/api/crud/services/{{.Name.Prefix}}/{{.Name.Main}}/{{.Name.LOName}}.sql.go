//go:build ignore

package {-{.Name.Main}-}

{-{- $table := .}-}
{-{- $QColumns:=   $table.Columns.GetAllQueryColumns.GetValidColumns}-}
{-{- if gt $QColumns.Len 0}-}

//get{-{.Name.CName}-}ListCount 获取{-{.Desc}-}列表条数
var get{-{.Name.CName}-}ListCount = `
select 
	count(1)
from {-{.Name.Raw}-} t
where
{-{- range $i,$v := $QColumns}-}
{-{- if or ($v.Field.LikeQuery) ($v.Cmpnt.Equal "input")}-}
	?t.{-{$v.Name}-}
{-{- else if $v.Cmpnt.Equal "daterange"}-}
	and ((@start_{-{$v.Name}-}='' and @end_{-{$v.Name}-}='') or (t.{-{$v.Name}-} >=  if(@start_{-{$v.Name}-}='', t.{-{$v.Name}-},@start_{-{$v.Name}-})
	and t.{-{$v.Name}-} <  date_add(if(@end_{-{$v.Name}-}='',t.{-{$v.Name}-},@end_{-{$v.Name}-}), interval 1 day)))
{-{- else if $v.Cmpnt.Equal "date"}-}
	and ((@start_{-{$v.Name}-}='' and @end_{-{$v.Name}-}='') or (t.{-{$v.Name}-} >=  if(@{-{$v.Name}-}='',t.{-{$v.Name}-},@{-{$v.Name}-})
	and t.{-{$v.Name}-} <  date_add(if(@{-{$v.Name}-}='',t.{-{$v.Name}-},@{-{$v.Name}-}),interval 1 day)))	
{-{- else if $v.Cmpnt.StartWith "multi"}-}
	and (if(@{-{$v.Name}-} = '',NULL,@{-{$v.Name}-}) is null or  FIND_IN_SET(t.{-{$v.Name}-},@{-{$v.Name}-}))
{-{- else}-}
	&t.{-{$v.Name}-}
{-{- end}-}
{-{- end}-}`

{-{- $LColumns:=  $table.Columns.GetAllListColumns.GetValidColumns}-}
{-{- $orders :=  $table.Columns.GetValidColumns.GetColumns "ord"}-}

//get{-{.Name.CName}-}List 查询{-{.Desc}-}列表数据
var get{-{.Name.CName}-}List =[]string{ `
select
	{-{- range $i,$v :=  $LColumns}-}
	t.{-{$v.Name}-}{-{if lt $i $LColumns.SLen}-},{-{end}-}
	{-{- end}-}
from {-{.Name.Raw}-} t
where
	{-{- range $i,$v := $QColumns}-}
	{-{- if or ($v.Field.LikeQuery) ($v.Cmpnt.Equal "input")}-}
	?t.{-{$v.Name}-}
	{-{- else if $v.Cmpnt.Equal "daterange"}-}
	and ((@start_{-{$v.Name}-}='' and @end_{-{$v.Name}-}='') or (t.{-{$v.Name}-} >=  if(@start_{-{$v.Name}-}='', t.{-{$v.Name}-},@start_{-{$v.Name}-})
	and t.{-{$v.Name}-} <  date_add(if(@end_{-{$v.Name}-}='',t.{-{$v.Name}-},@end_{-{$v.Name}-}), interval 1 day)))
	{-{- else if $v.Cmpnt.Equal "date"}-}
	and ((@start_{-{$v.Name}-}='' and @end_{-{$v.Name}-}='') or (t.{-{$v.Name}-} >=  if(@{-{$v.Name}-}='',t.{-{$v.Name}-},@{-{$v.Name}-})
	and t.{-{$v.Name}-} <  date_add(if(@{-{$v.Name}-}='',t.{-{$v.Name}-},@{-{$v.Name}-}),interval 1 day)))	
	{-{- else if $v.Cmpnt.StartWith "multi"}-}
	and (if(@{-{$v.Name}-} = '',NULL,@{-{$v.Name}-}) is null or  FIND_IN_SET(t.{-{$v.Name}-},@{-{$v.Name}-}))
	{-{- else if $v.Cmpnt.StartWith "ddmenu"}-}
	and (if(@{-{$v.Name}-} = '',NULL,@{-{$v.Name}-}) is null or  FIND_IN_SET(@{-{$v.Name}-},t.{-{$v.Name}-}))
	{-{- else}-}
	&t.{-{$v.Name}-}
	{-{- end}-}
	{-{- end}-}
order by {-{ range $i,$m := $orders}-}
	t.{-{$m.Name}-}{-{if eq "asc" $m.Cmpnt.Format}-} asc {-{else}-} desc {-{end}-},
{-{- end}-}
{-{- range $i,$v := $table.Columns.GetPKColumns.GetValidColumns}-}
	t.{-{$v.Name}-} desc
{-{- end}-}
limit @ps offset @offset`,
}
{-{- end}-}

{-{- $CColumns:=  $table.Columns.GetAllCreateColumns.GetValidColumns}-}
{-{- if gt $CColumns.Len 0}-}

// insert{-{.Name.CName}-} 保存{-{.Desc}-}数据
var insert{-{.Name.CName}-} = `
insert into {-{.Name.Raw}-}
(
	{-{- range $i,$v := $CColumns}-}
	{-{$v.Name}-}{-{if lt $i $CColumns.SLen}-},{-{end}-}
	{-{- end}-}
)
values
(
	{-{- range $i,$v := $CColumns}-}
	{-{- if eq false $v.Field.Required}-}
	if(@{-{$v.Name}-}='',{-{if ne "" $v.Field.DefaultValue}-} {-{$v.Field.DefaultValue}-} {-{else}-}NULL{-{end}-},@{-{$v.Name}-}){-{if lt $i $CColumns.SLen}-},{-{end}-}
	{-{- else}-}
	@{-{$v.Name}-}{-{if lt $i $CColumns.SLen}-},{-{end}-}
	{-{- end}-}
	{-{- end}-}
)`
{-{- end}-}

{-{- $UColumns:=  $table.Columns.GetAllUpdateColumns.GetValidColumns}-}
{-{- if gt $UColumns.Len 0}-}

//update{-{.Name.CName}-} 修改{-{.Desc}-}数据
var update{-{.Name.CName}-} = `
update {-{.Name.Raw}-} t set 
{-{- range $i,$v := $UColumns}-}
	t.{-{$v.Name}-} = if(@{-{$v.Name}-}='',{-{if ne "" $v.Field.DefaultValue}-} {-{$v.Field.DefaultValue}-} {-{else}-}NULL{-{end}-},@{-{$v.Name}-}){-{if lt $i $UColumns.SLen}-},{-{end}-}
{-{- end}-}
where 
{-{- range $i,$v := $table.Columns.GetPKColumns.GetValidColumns}-}
	&{-{$v.Name}-}
{-{- end}-}`
{-{- end}-}

{-{- $LColumns:=  $table.Columns.GetAllListColumns.GetValidColumns}-}
{-{- if gt $LColumns.Len 0}-}

//get{-{.Name.CName}-} 查询单条{-{.Desc}-}数据
var get{-{.Name.CName}-} = `
select
{-{- range $i,$v := $LColumns}-}
	t.{-{$v.Name}-}{-{if lt $i $LColumns.SLen}-},{-{end}-}
{-{- end}-}
from {-{.Name.Raw}-} t
where 
{-{- range $i,$v := $table.Columns.GetPKColumns.GetValidColumns}-}
	&{-{$v.Name}-}
{-{- end}-}`
{-{- end}-}

{-{- if gt $table.Columns.GetDelColumns.Len 0}-}

//delete{-{.Name.CName}-} 删除单条{-{.Desc}-}数据
var delete{-{.Name.CName}-} = `
delete from {-{.Name.Raw}-}
where 
{-{- range $i,$v := $table.Columns.GetPKColumns.GetValidColumns}-}
	&{-{$v.Name}-}
{-{- end}-}`
{-{- end}-}

{-{- $switchs :=  $table.Columns.GetAllListColumns.GetValidColumns.GetByCmpnt "switch"}-}
{-{- if gt $switchs.Len 0}-}

//switch{-{.Name.CName}-} 删除单条{-{.Desc}-}数据
var switch{-{.Name.CName}-} = `
update {-{.Name.Raw}-} t set 
{-{- range $i,$v := $switchs}-}
	t.{-{$v.Name}-} = @{-{$v.Name}-}{-{if lt $i $switchs.SLen}-},{-{end}-}
{-{- end}-}
where 
{-{- range $i,$v := $table.Columns.GetPKColumns.GetValidColumns}-}
	&{-{$v.Name}-}
{-{- end}-}`
{-{- end}-}

{-{- $updator :=  $table.Optrs.All.GetByCmptName "batupdator-lstupdator"}-}
{-{- if gt $updator.Len 0}-}
{-{- range $i,$c := $updator}-}
{-{- $fds := f_string_contact $c.RwName "-" (f_string_contact "b" $c.RwName)}-}
{-{- $fields :=  $table.Columns.GetColumns $fds}-}

//update{-{$table.Name.CName}-} 修改{-{$table.Desc}-}数据
var updator{-{$table.Name.CName}-}{-{$c.ReqURL}-} = `
update {-{$table.Name.Raw}-} t set 
{-{- range $i,$v := $fields}-}
	t.{-{$v.Name}-} = if(@{-{$v.Name}-}='',{-{if ne "" $v.Field.DefaultValue}-} {-{$v.Field.DefaultValue}-} {-{else}-}t.{-{$v.Name}-}{-{end}-},@{-{$v.Name}-}){-{if lt $i $fields.SLen}-},{-{end}-}
{-{- end}-}
where
{-{- range $i,$v :=  $table.Columns.GetColumns $c.FwName -}-}
	{-{- if  $c.IsMutilValue $v.Name $table.Columns}-}
	and FIND_IN_SET(t.{-{$v.Name}-},@{-{$v.Name}-})
	{-{- else}-}
	&{-{$v.Name}-}
	{-{- end}-}
{-{- end}-}`
{-{- end}-}
{-{- end}-}

{-{- $barinsert :=  $table.Optrs.All.GetByCmptName  "barinsert"}-}
{-{- if gt $barinsert.Len 0}-}
{-{- range $i,$c := $barinsert}-}
{-{- $fds := f_string_contact $c.RwName "-" (f_string_contact "b" $c.RwName)}-}
{-{- $fds = f_string_contact $fds "-" (f_string_contact "f" $c.RwName)}-}
{-{- $Columns:=$table.Columns.GetValidColumns.GetColumns $fds}-}

// barinsert{-{$table.Name.CName}-} 保存{-{.Desc}-}数据
var barinsert{-{$c.ReqURL|f_string_2cname}-}{-{$table.Name.CName}-} = `
insert into {-{$table.Name.Raw}-}
(
	{-{- range $i,$v :=  $Columns}-}
	{-{$v.Name}-}{-{if lt $i $Columns.SLen}-},{-{end}-}
	{-{- end}-}
)
values
(
	{-{- range $i,$v := $Columns}-}
	{-{- if eq false $v.Field.Required}-}
	if(@{-{$v.Name}-}='',{-{if ne "" $v.Field.DefaultValue}-} {-{$v.Field.DefaultValue}-} {-{else}-}NULL{-{end}-},@{-{$v.Name}-}){-{if lt $i $Columns.SLen}-},{-{end}-}
	{-{- else}-}
	@{-{$v.Name}-}{-{if lt $i $Columns.SLen}-},{-{end}-}
	{-{- end}-}
	{-{- end}-}
)`
{-{- end}-}
{-{- end}-}


{-{- $batinserts :=  $table.Optrs.All.GetByCmptName  "batinsert"}-}
{-{- if gt $batinserts.Len 0}-}
{-{- range $i,$c := $batinserts}-}

//batInsert{-{$table.Name.CName}-}{-{$c.ReqURL}-} 批量添加{-{$table.Desc}-}数据
var batInsert{-{$table.Name.CName}-}{-{$c.ReqURL}-} = `
insert {-{$table.Name.Raw}-} (
	{-{- $fields :=  $table.Columns.GetColumns $c.FwName}-}
{-{- range $i,$v := $fields}-}
	{-{$v.Name}-}{-{if lt $i $fields.SLen}-},{-{end}-}
{-{- end}-})values(
	{-{- range $i,$v := $fields}-}
	@{-{$v.Name}-}{-{if lt $i $fields.SLen}-},{-{end}-}
{-{- end}-}
)`
{-{- end}-}
{-{- end}-}