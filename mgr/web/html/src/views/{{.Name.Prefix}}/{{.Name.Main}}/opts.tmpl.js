{-{- $table := . }-}
{-{- $optRow:= mergeOptrs $table.ListOpts $table.BarOpts}-}
{-{- range $x,$m:= $optRow}-}
{-{- if eq "CMPNT" $m.Name}-}
  show_cmpnt_{-{$m.UNQ}-}(fm = {}){
    let query = {}
    {-{- $rows:= fltrColumns $table $m.RwName}-}
      {-{- range $i,$c:=$rows}-} 
    query.{-{$c.Name}-} = fm.{-{$c.Name}-}|| fm.le_{-{$c.Name}-}
      {-{- end}-}
    {-{- $pkrows:=  $table.PKColumns}-}
      {-{- range $i,$c:=$pkrows}-} 
    query.{-{$c.Name}-} = fm.{-{$c.Name}-}
      {-{- end}-}
    {-{- if eq true $m.IsMux }-}
    this.$refs.cmpnt_{-{$m.UNQ}-}.show_{-{$m.UNQ}-}(query)
    {-{- else}-}
    this.$refs.cmpnt_{-{$m.UNQ}-}.show(query)
    {-{- end}-}
  },
{-{- end}-}
{-{- end}-}
{-{- range $x,$m:=$optRow}-}
{-{- if eq "LINK" $m.Name}-}
  goto_{-{$m.UNQ}-}(fm = {}){
    let query = {}
    {-{- $rows:= fltrColumns $table $m.RwName}-}
      {-{- range $i,$c:=$rows}-} 
    query.{-{$c.Name}-} = fm.{-{$c.Name}-}
      {-{- end}-}
    this.goto('{-{$m.URL}-}',query)
  }
{-{- end}-}
{-{- end}-}