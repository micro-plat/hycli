{-{- $table := .}-}
{-{- $opts:= $table.Optrs.All.GetCompnents}-}
components: {
    {-{- range $x,$m:= $opts}-}
    {-{$m.UNQ}-},
    {-{- end}-}
    {-{- range $i,$c:= $table.Optrs.BarOpts.GetByCmptName "list" }-}
    {-{- $gantt := $c.GetParams "gantt" -}-}
    {-{- if ne $gantt ""}-}
    gantt,
    {-{- end}-}
    {-{- end}-}

    {-{- $qcols :=  $table.GetColumnsByTPName "q"}-}
    {-{- range $i,$c:= $qcols}-}
    {-{- if and (eq "ddmenu" $c.Cmpnt.Type) (eq true $c.Enum.IsEnum)}-}
     ddmenu,
    {-{- end  }-}
    {-{- end  }-}
},