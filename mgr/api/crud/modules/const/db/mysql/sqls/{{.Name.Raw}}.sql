{-{- $table := .}-}
{-{- $rows := $table.Columns.GetValidColumns}-}
{-{- $count :=$rows.Len}-}
-- {-{$table.Name}-} {-{$table.Desc}-}
CREATE TABLE IF NOT EXISTS `{-{.Name.Raw}-}` (
	{-{- range $i,$c:=$rows}-}
	`{-{$c.Name}-}` {-{$c.Field.TMYSQL}-} {-{if eq false $c.AllowNull}-} NOT NULL {-{end}-}{-{if $c.Field.IsSEQ}-} AUTO_INCREMENT {-{else}-} {-{if ne "" $c.Field.VMYSQL}-} DEFAULT {-{$c.Field.VMYSQL}-}{-{end}-}{-{end}-} comment '{-{$c.Desc}-}',
	{-{- end}-}
	{-{- if gt $table.Columns.GetPKColumns.Len 0}-}
	PRIMARY KEY (
	{-{- range $i,$m := $table.Columns.GetPKColumns -}-}
	`{-{$m.Name}-}`{-{if lt $i $table.Columns.GetPKColumns.SLen}-},{-{end}-}
	{-{- end}-}) USING BTREE {-{if or (gt $table.DBIdx.UNQ.Len 0) (gt $table.DBIdx.Normal.Len 0)}-},{-{end}-}
	{-{- end}-}
	{-{- if gt $table.DBIdx.UNQ.Len 0}-}
	{-{- range $i,$unq := $table.DBIdx.UNQ}-}
	UNIQUE KEY `{-{$unq.Name}-}`(
	{-{- range $i,$m := $unq.Columns -}-}
	`{-{$m.Name}-}`{-{if lt $i $unq.Columns.SLen}-},{-{end}-}
	{-{- end}-}) USING BTREE {-{if lt $i $table.DBIdx.UNQ.SLen}-},{-{end}-}
	{-{- end}-}
	{-{- end}-}
	{-{- if gt $table.DBIdx.Normal.Len 0}-}
	{-{- range $i,$nr := $table.DBIdx.Normal}-}
	KEY `{-{$nr.Name}-}`(
	{-{- range $i,$m := $nr.Columns -}-}
	`{-{$m.Name}-}`{-{if lt $i $nr.Columns.SLen}-},{-{end}-}
	{-{- end}-}) USING BTREE {-{if lt $i $table.DBIdx.Normal.SLen}-},{-{end}-}
	{-{- end}-}
	{-{- end}-}
) ENGINE=InnoDB {-{- range $i,$c:=$rows}-} {-{- if and (eq true $c.Field.IsSEQ) (ne "" $c.DefValue)}-} AUTO_INCREMENT = {-{$c.DefValue|f_mysql_get_def_value}-}{-{- end}-}{-{- end}-} DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='{-{$table.Desc}-}'