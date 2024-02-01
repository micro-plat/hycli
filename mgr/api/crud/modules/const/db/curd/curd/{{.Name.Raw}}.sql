{-{- $table := .}-}
{-{- $rows := $table.Columns.GetValidColumns}-}
SELECT {-{ range $i,$c:=$rows}-}t.{-{$c.Name}-}{-{if lt $i $rows.SLen}-},{-{end}-} {-{end}-} 
FROM  {-{$table.Name.Raw}-} t
where {-{range $i,$m := $table.Columns.GetPKColumns}-}t.{-{$m.Name}-} = @{-{$m.Name}-}{-{if lt $i $table.Columns.GetPKColumns.SLen}-},{-{end}-}
{-{- end}-}