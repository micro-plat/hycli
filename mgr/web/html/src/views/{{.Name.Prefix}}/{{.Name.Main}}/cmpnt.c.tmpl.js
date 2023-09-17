{-{- $table := .}-}
{-{- $opts:= $table.ListOpts.Merge $table.BarOpts $table.ChartOpts $table.ExtOpts}-}
components: {
    {-{- range $x,$m:= $opts.GetByName "CMPNT"}-}
    {-{$m.UNQ}-},
    {-{- end}-}
},