{-{- $table := .}-}
{-{- $rows :=fltrColumnsExcludeExt $table.Columns}-}
{-{- $count :=len $rows}-}
-- {-{$table.Name}-} {-{$table.Desc}-}
CREATE TABLE IF NOT EXISTS `{-{.Name.Raw}-}` (
	{-{- range $i,$c:=$rows}-}
	`{-{$c.Name}-}` {-{$c.Field.TMYSQL}-} {-{if eq false $c.AllowNull}-} NOT NULL {-{end}-}{-{if $c.Field.IsSEQ}-} AUTO_INCREMENT {-{else}-} {-{if ne "" $c.Field.VMYSQL}-} DEFAULT {-{$c.Field.VMYSQL}-}{-{end}-}{-{end}-} comment '{-{$c.Desc}-}',
	{-{- end}-}
	{-{- if gt (len $table.PKColumns) 0}-}
	PRIMARY KEY (
	{-{- range $i,$m := $table.PKColumns -}-}
	`{-{$m.Name}-}`{-{if lt $i ((len $table.PKColumns)|minus)}-},{-{end}-}
	{-{- end}-}) USING BTREE {-{if or (gt (len $table.UNQIndex) 0) (gt (len $table.NormalIdx) 0)}-},{-{end}-}
	{-{- end}-}
	{-{- if gt (len $table.UNQIndex) 0}-}
	{-{- range $i,$unq := $table.UNQIndex}-}
	UNIQUE KEY `{-{$unq.Name}-}`(
	{-{- range $i,$m := $unq.Columns -}-}
	`{-{$m.Name}-}`{-{if lt $i ((len $unq.Columns)|minus)}-},{-{end}-}
	{-{- end}-}) USING BTREE {-{if or (gt (len $table.NormalIdx) 0)  (lt $i ((len $table.UNQIndex)|minus))}-},{-{end}-}
	{-{- end}-}
	{-{- end}-}
	{-{- if gt (len $table.NormalIdx) 0}-}
	{-{- range $i,$nr := $table.NormalIdx}-}
	KEY `{-{$nr.Name}-}`(
	{-{- range $i,$m := $nr.Columns -}-}
	`{-{$m.Name}-}`{-{if lt $i ((len $nr.Columns)|minus)}-},{-{end}-}
	{-{- end}-}) USING BTREE {-{if lt $i ((len $table.NormalIdx)|minus)}-},{-{end}-}
	{-{- end}-}
	{-{- end}-}
) ENGINE=InnoDB {-{- range $i,$c:=$rows}-} {-{- if and (eq true $c.Field.IsSEQ) (ne "" $c.DefValue)}-} AUTO_INCREMENT = {-{$c.DefValue|fltrMYSQLDef}-}{-{- end}-}{-{- end}-} DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='{-{$table.Desc}-}'