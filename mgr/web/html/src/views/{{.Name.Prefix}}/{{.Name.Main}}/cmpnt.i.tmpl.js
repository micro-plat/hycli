{-{- $table := .}-}
{-{- $opts:=mergeOptrs $table.ListOpts $table.BarOpts $table.ChartOpts $table.ExtOpts}-}
{-{- range $x,$m:= fltrOptrs $opts "CMPNT"}-}
// {-{$m.Label}-}
import {-{$m.UNQ}-} from "{-{fltrTranslate $m.URL (fltrFindTable $m.Table)}-}"
{-{- end}-}