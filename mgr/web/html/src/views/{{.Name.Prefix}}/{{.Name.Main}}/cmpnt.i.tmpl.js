{-{- $table := .}-}
{-{- $opts:=mergeOptrs $table.ListOpts $table.BarOpts $table.ChartOpts $table.ExtOpts}-}
{-{- range $x,$m:= $opts}-}
// {-{$m.Label}-}
 {-{- if eq "CMPNT" $m.Name}-}
import {-{$m.UNQ}-} from "{-{fltrTranslate $m.URL (fltrFindTable $m.Table)}-}"
{-{- end}-}
{-{- end}-}