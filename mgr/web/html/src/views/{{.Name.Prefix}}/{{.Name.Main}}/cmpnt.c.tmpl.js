{-{- $table := .}-}
{-{- $opts:=mergeOptrs $table.ListOpts $table.BarOpts $table.ChartOpts $table.ExtOpts}-}
components: {
    {-{- range $x,$m:= fltrOptrs $opts "CMPNT"}-}
    {-{$m.UNQ}-},
    {-{- end}-}
},