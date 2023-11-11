//go:build ignore

package {-{.Name.Main}-}

{-{- $table := .}-}
{-{- $CColumns:=  $table.GetValidColumnsByName "c-bc"}-}
{-{- $UColumns:=  $table.GetValidColumnsByName "u"}-}
{-{- $switchs :=  $table.GetColumsByCmpnt "switch" "l"}-}
{-{- $clen := (len $CColumns)|minus}-}
{-{- $ulen := (len $UColumns)|minus}-}
{-{- $slen := (len $switchs)|minus}-}
{-{- $sigleQueryCols:=  $table.GetValidColumnsByName "l-le-bl"}-}
{-{- $totalQCols:=   $table.GetValidColumnsByName "q-bq"}-}
{-{- $vlen := (len $sigleQueryCols)|minus}-}
{-{- $orders :=  $table.GetValidColumnsByName "ord"}-}

{-{- if gt (len ( $table.GetColumnsByTPName "q-bq")) 0}-}

//get{-{.Name.CName}-}ListCount 获取{-{.Desc}-}列表条数
const get{-{.Name.CName}-}ListCount = `
select 
	count(1)
from {-{.Name.Raw}-} t
where
{-{- range $i,$v := $totalQCols}-}
{-{- if or (eq true $v.Field.LikeQuery) (eq "input" $v.Cmpnt.Type)}-}
	?t.{-{$v.Name}-}
{-{- else if eq "daterange" $v.Cmpnt.Type}-}
	and ((@start_{-{$v.Name}-}='' and @end_{-{$v.Name}-}='') or (t.{-{$v.Name}-} >=  if(@start_{-{$v.Name}-}='', t.{-{$v.Name}-},@start_{-{$v.Name}-})
	and t.{-{$v.Name}-} <  date_add(if(@end_{-{$v.Name}-}='',t.{-{$v.Name}-},@end_{-{$v.Name}-}), interval 1 day)))
{-{- else if eq "date" $v.Cmpnt.Type}-}
	and ((@start_{-{$v.Name}-}='' and @end_{-{$v.Name}-}='') or (t.{-{$v.Name}-} >=  if(@{-{$v.Name}-}='',t.{-{$v.Name}-},@{-{$v.Name}-})
	and t.{-{$v.Name}-} <  date_add(if(@{-{$v.Name}-}='',t.{-{$v.Name}-},@{-{$v.Name}-}),interval 1 day)))	
{-{- else if eq true (f_string_start $v.Cmpnt.Type "multi")}-}
	and (if(@{-{$v.Name}-} = '',NULL,@{-{$v.Name}-}) is null or  FIND_IN_SET(t.{-{$v.Name}-},@{-{$v.Name}-}))
{-{- else}-}
	&t.{-{$v.Name}-}
{-{- end}-}
{-{- end}-}`

//get{-{.Name.CName}-}List 查询{-{.Desc}-}列表数据
const get{-{.Name.CName}-}List = `
select
	{-{- range $i,$v :=  $sigleQueryCols}-}
	t.{-{$v.Name}-},
	{-{- end}-}
	1
from {-{.Name.Raw}-} t
where
	{-{- range $i,$v := $totalQCols}-}
	{-{- if or (eq true $v.Field.LikeQuery) (eq "input" $v.Cmpnt.Type)}-}
	?t.{-{$v.Name}-}
	{-{- else if eq "daterange" $v.Cmpnt.Type}-}
	and ((@start_{-{$v.Name}-}='' and @end_{-{$v.Name}-}='') or (t.{-{$v.Name}-} >=  if(@start_{-{$v.Name}-}='', t.{-{$v.Name}-},@start_{-{$v.Name}-})
	and t.{-{$v.Name}-} <  date_add(if(@end_{-{$v.Name}-}='',t.{-{$v.Name}-},@end_{-{$v.Name}-}), interval 1 day)))
	{-{- else if eq "date" $v.Cmpnt.Type}-}
	and ((@start_{-{$v.Name}-}='' and @end_{-{$v.Name}-}='') or (t.{-{$v.Name}-} >=  if(@{-{$v.Name}-}='',t.{-{$v.Name}-},@{-{$v.Name}-})
	and t.{-{$v.Name}-} <  date_add(if(@{-{$v.Name}-}='',t.{-{$v.Name}-},@{-{$v.Name}-}),interval 1 day)))	
	{-{- else if eq true (f_string_start $v.Cmpnt.Type "multi")}-}
	and (if(@{-{$v.Name}-} = '',NULL,@{-{$v.Name}-}) is null or  FIND_IN_SET(t.{-{$v.Name}-},@{-{$v.Name}-}))
	{-{- else}-}
	&t.{-{$v.Name}-}
	{-{- end}-}
	{-{- end}-}
order by {-{ range $i,$m := $orders}-}
	t.{-{$m.Name}-}{-{if eq "asc" $m.Cmpnt.Format}-} asc {-{else}-} desc {-{end}-},
{-{- end}-}
{-{- range $i,$v := $table.PKColumns.GetValidColumns}-}
	t.{-{$v.Name}-} desc
{-{- end}-}
limit @ps offset @offset`
{-{- end}-}

{-{- if gt (len ( $table.GetColumnsByTPName "c-bc")) 0}-}

// insert{-{.Name.CName}-} 保存{-{.Desc}-}数据
const insert{-{.Name.CName}-} = `
insert into {-{.Name.Raw}-}
(
	{-{- range $i,$v :=  $table.GetValidColumnsByName "c-bc"}-}
	{-{$v.Name}-}{-{if lt $i $clen}-},{-{end}-}
	{-{- end}-}
)
values
(
	{-{- range $i,$v :=   $table.GetValidColumnsByName "c-bc"}-}
	{-{- if eq false $v.Field.Required}-}
	if(@{-{$v.Name}-}='',{-{if ne "" $v.Field.DefaultValue}-} {-{$v.Field.DefaultValue}-} {-{else}-}NULL{-{end}-},@{-{$v.Name}-}){-{if lt $i $clen}-},{-{end}-}
	{-{- else}-}
	@{-{$v.Name}-}{-{if lt $i $clen}-},{-{end}-}
	{-{- end}-}
	{-{- end}-}
)`
{-{- end}-}

{-{- if gt (len ( $table.GetValidColumnsByName "u")) 0}-}

//update{-{.Name.CName}-} 修改{-{.Desc}-}数据
const update{-{.Name.CName}-} = `
update {-{.Name.Raw}-} t set 
{-{- range $i,$v := $table.GetValidColumnsByName "u"}-}
	t.{-{$v.Name}-} = if(@{-{$v.Name}-}='',{-{if ne "" $v.Field.DefaultValue}-} {-{$v.Field.DefaultValue}-} {-{else}-}NULL{-{end}-},@{-{$v.Name}-}){-{if lt $i $ulen}-},{-{end}-}
{-{- end}-}
where 
{-{- range $i,$v := $table.PKColumns.GetValidColumns}-}
	&{-{$v.Name}-}
{-{- end}-}`
{-{- end}-}

{-{- if gt (len  ( $table.GetValidColumnsByName "l-le")) 0}-}

//get{-{.Name.CName}-} 查询单条{-{.Desc}-}数据
const get{-{.Name.CName}-} = `
select
{-{- range $i,$v := $sigleQueryCols}-}
	t.{-{$v.Name}-}{-{if lt $i $vlen}-},{-{end}-}
{-{- end}-}
from {-{.Name.Raw}-} t
where 
{-{- range $i,$v := $table.PKColumns.GetValidColumns}-}
	&{-{$v.Name}-}
{-{- end}-}`
{-{- end}-}

{-{- if gt (len ( $table.GetColumnsByTPName "D")) 0}-}

//delete{-{.Name.CName}-} 删除单条{-{.Desc}-}数据
const delete{-{.Name.CName}-} = `
delete from {-{.Name.Raw}-}
where 
{-{- range $i,$v := $table.PKColumns.GetValidColumns}-}
	&{-{$v.Name}-}
{-{- end}-}`
{-{- end}-}

{-{- if gt (len $switchs) 0}-}

//switch{-{.Name.CName}-} 删除单条{-{.Desc}-}数据
const switch{-{.Name.CName}-} = `
update {-{.Name.Raw}-} t set 
{-{- range $i,$v := $switchs}-}
	t.{-{$v.Name}-} = @{-{$v.Name}-}{-{if lt $i $slen}-},{-{end}-}
{-{- end}-}
where 
{-{- range $i,$v := $table.PKColumns.GetValidColumns}-}
	&{-{$v.Name}-}
{-{- end}-}`
{-{- end}-}

{-{- $barlst := $table.BarOpts.Merge $table.ListOpts }-}
{-{- $updator :=  $barlst.GetByCmptName "batupdator-lstupdator"}-}
{-{- if gt (len $updator) 0}-}
{-{- range $i,$c := $updator}-}

//update{-{$table.Name.CName}-} 修改{-{$table.Desc}-}数据
const updator{-{$table.Name.CName}-}{-{$c.ReqURL}-} = `
update {-{$table.Name.Raw}-} t set 
{-{- $fds := f_string_contact $c.RwName "-" (f_string_contact "b" $c.RwName)}-}
{-{- $fields :=  $table.GetColumnsByTPName $fds}-}
{-{- $uplen := minus (len $fields)}-}
{-{- range $i,$v := $fields}-}
	t.{-{$v.Name}-} = if(@{-{$v.Name}-}='',{-{if ne "" $v.Field.DefaultValue}-} {-{$v.Field.DefaultValue}-} {-{else}-}t.{-{$v.Name}-}{-{end}-},@{-{$v.Name}-}){-{if lt $i $uplen}-},{-{end}-}
{-{- end}-}
where 1 = 1
{-{- range $i,$v :=  $table.GetColumnsByTPName $c.FwName -}-}
	{-{- if  $table.IsMutilValue $c $v.Name}-}
	and FIND_IN_SET(t.{-{$v.Name}-},@{-{$v.Name}-})
	{-{- else}-}
	&{-{$v.Name}-}
	{-{- end}-}
{-{- end}-}`
{-{- end}-}
{-{- end}-}

{-{- $barinsert :=  $table.BarOpts.GetByCmptName  "barinsert"}-}
{-{- if gt (len $barinsert) 0}-}
{-{- range $i,$c := $barinsert}-}
{-{- $fds := f_string_contact $c.RwName "-" (f_string_contact "b" $c.RwName)}-}
{-{- $colums:=$table.GetValidColumnsByName $fds}-}
{-{- $clen :=minus (len $colums)}-}
// barinsert{-{$table.Name.CName}-} 保存{-{.Desc}-}数据
const barinsert{-{$c.ReqURL|f_string_2cname}-}{-{$table.Name.CName}-} = `
insert into {-{$table.Name.Raw}-}
(
	{-{- range $i,$v :=  $colums}-}
	{-{$v.Name}-}{-{if lt $i $clen}-},{-{end}-}
	{-{- end}-}
)
values
(
	{-{- range $i,$v := $colums}-}
	{-{- if eq false $v.Field.Required}-}
	if(@{-{$v.Name}-}='',{-{if ne "" $v.Field.DefaultValue}-} {-{$v.Field.DefaultValue}-} {-{else}-}NULL{-{end}-},@{-{$v.Name}-}){-{if lt $i $clen}-},{-{end}-}
	{-{- else}-}
	@{-{$v.Name}-}{-{if lt $i $clen}-},{-{end}-}
	{-{- end}-}
	{-{- end}-}
)`
{-{- end}-}
{-{- end}-}

//批量添加数据
{-{- $batinserts :=  $table.BarOpts.GetByCmptName  "batinsert"}-}
{-{- $uplen := minus (len $batinserts)}-}
{-{- if gt (len $batinserts) 0}-}
{-{- range $i,$c := $batinserts}-}

//batInsert{-{$table.Name.CName}-}{-{$c.ReqURL}-} 修改{-{$table.Desc}-}数据
const batInsert{-{$table.Name.CName}-}{-{$c.ReqURL}-} = `
insert {-{$table.Name.Raw}-} (
	{-{- $fields :=  $table.GetColumnsByTPName $c.FwName}-}
	{-{- $flen := minus (len $fields)}-}
{-{- range $i,$v := $fields}-}
	{-{$v.Name}-}{-{if lt $i $flen}-},{-{end}-}
{-{- end}-})values(
	{-{- range $i,$v := $fields}-}
	@{-{$v.Name}-}{-{if lt $i $flen}-},{-{end}-}
{-{- end}-}
)`
{-{- end}-}
{-{- end}-}