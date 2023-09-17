{-{- $table := .}-}
{-{- $opts:=$table.ListOpts.Merge $table.BarOpts $table.ChartOpts $table.ExtOpts}-}
{-{- range $x,$m:=  $opts.GetByName "CMPNT"}-}
// {-{$m.Label}-}
import {-{$m.UNQ}-} from "{-{f_string_translate $m.URL (f_table_find_by_name $m.Table)}-}"
{-{- end}-}