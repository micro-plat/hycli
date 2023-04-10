{-{- $table := .}-}
{-{- $opts:=mergeOptrs $table.ListOpts $table.BarOpts $table.ChartOpts}-}
{-{- range $x,$m:= $opts}-}
// {-{$m.Label}-}
 {-{- if eq "CMPNT" $m.Name}-}
import {-{$m.UNQ}-} from "{-{fltrTranslate $m.URL $table}-}"
{-{- end}-}
{-{- end}-}