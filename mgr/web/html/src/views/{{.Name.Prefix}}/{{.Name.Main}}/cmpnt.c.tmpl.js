{-{- $table := .}-}
{-{- $opts:= $table.ListOpts.Merge $table.BarOpts $table.ChartOpts $table.ExtOpts}-}
components: {
    {-{- range $x,$m:= fltrOptrs $opts "CMPNT"}-}
    {-{$m.UNQ}-},
    {-{- end}-}
},