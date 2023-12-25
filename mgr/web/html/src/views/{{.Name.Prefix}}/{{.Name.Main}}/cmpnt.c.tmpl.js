{-{- $table := .}-}
components: {
    {-{- range $x,$m:= $table.Optrs.All.GetCompnents}-}
    {-{$m.UNQ}-},
    {-{- end}-}
    {-{- range $i,$c:= ($table.Optrs.BarOpts.GetByCmptName "LIST").GetParamsOptrs "gantt"}-}
    gantt,
    {-{- end}-}
    {-{- range $i,$c:= $table.Columns.GetFontQueryColumns.GetEnumColumns.GetByCmpnt "ddmenu"}-}
     ddmenu,
    {-{- end  }-}
},