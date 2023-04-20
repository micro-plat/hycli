{-{- $table := . }-}
{-{- $optRow:= mergeOptrs $table.ListOpts $table.BarOpts}-}

{-{- $cmpnts:= fltrOptrs $optRow "CMPNT"}-}
{-{- if gt (len $cmpnts) 0}-}
show_cmpnt(cmd,row){
  if(this.cmpnt_funcs[cmd]){
    this.cmpnt_funcs[cmd](row)
  }
},
{-{- end}-}
{-{- range $x,$m:= $cmpnts}-}
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

      //2023.4.18添加
    {-{- $c := fltrOptrsByStatic $m }-}
    {-{- range $k,$v := $c}-}
    {-{ if eq true (fltrStart $v "@")}-}
    query.{-{$k}-} = this.{-{fltrTrim $v "@"}-}
    {-{else if eq true (fltrStart $v "&")}-}
    query.{-{$k}-} = (this.{-{fltrTrim $v "&"}-}||[]).join(",")
    {-{- else}-}
    query.{-{$k}-} = "{-{$v}-}"
    {-{- end}-}
    {-{- end}-}

    {-{- if eq true $m.IsMux }-}
    this.$refs.cmpnt_{-{$m.UNQ}-}.show_{-{$m.UNQ}-}(query)
    {-{- else}-}
    this.$refs.cmpnt_{-{$m.UNQ}-}.show(query)
    {-{- end}-}
  },
{-{- end}-}
{-{- $linkcmpnts:= fltrOptrs $optRow "LINK"}-}
{-{- range $x,$m:=$linkcmpnts}-}
  goto_{-{$m.UNQ}-}(fm = {}){
    let query = {}
    {-{- $rows:= fltrColumns $table $m.RwName}-}
      {-{- range $i,$c:=$rows}-} 
    query.{-{$c.Name}-} = fm.{-{$c.Name}-}
      {-{- end}-}
    this.goto('{-{$m.URL}-}',query)
  },
{-{- end}-}