{-{- $xtable := . }-}
{-{- $table := $xtable.Current }-}
{-{- $mtable := $xtable.Main}-}
{-{- $tmpl := f_table_is_tmp $xtable}-}
{-{- $optRow:=  $table.ListOpts.Merge $table.BarOpts}-}
{-{- if eq true $tmpl}-}
{-{- $optRow = $mtable.ViewExtCmptOpts.GetLstOptrs  $table.Name.Raw}-}
{-{- end}-}
{-{- $cmpnts:=  $optRow.GetByName "CMPNT"}-}
{-{- if gt (len $cmpnts) 0}-}
show_cmpnt(cmd,row){
  if(this.cmpnt_funcs[cmd]){
    this.cmpnt_funcs[cmd](row)
  }
},
{-{- end}-}
{-{- range $x,$m:= $cmpnts}-}

  //{-{$m.Label}-}
  show_cmpnt_{-{$m.UNQ}-}(fm = {}){
    let query = {}
    {-{- if and (ne "" $m.RwName) (ne "" $m.FwName)}-}
    query.{-{$m.FwName}-} = fm.{-{$m.RwName}-}|| fm.le_{-{$m.RwName}-}
    {-{- else}-}
    {-{- $rows:=  $table.GetColumnsByTPName $m.RwName}-}
    {-{- range $i,$c:=$rows}-} 
    query.{-{$c.Name}-} = fm.{-{$c.Name}-}|| fm.le_{-{$c.Name}-}
      {-{- end}-}
    {-{- end}-}
   
     {-{- $rows:=  $table.GetColumnsByTPName ($m.GetParam "labelName" "")}-}
     {-{- range $i,$c:=$rows}-} 
     query.{-{$c.Name}-} = fm.{-{$c.Name}-}|| fm.le_{-{$c.Name}-}
     {-{- end}-}  

    {-{- $pkrows:=  $table.PKColumns}-}
      {-{- range $i,$c:=$pkrows}-} 
    query.{-{$c.Name}-} = fm.{-{$c.Name}-}
      {-{- end}-}

      //2023.4.18添加---
    {-{- $c :=  $m.GetParamsByAtPrefix }-}
    {-{- range $k,$v := $c}-}
    {-{- if eq true (f_string_start $v "@")}-}
    query.{-{$k}-} = this.{-{f_string_trim $v "@"}-}
    {-{- else if eq true (f_string_start $v "&")}-}
    query.{-{$k}-} = (this.{-{f_string_trim $v "&"}-}||[]).join(",")
    {-{- else if eq true (f_string_start $v "#")}-}
    query.{-{$k}-} = fm.{-{f_string_trim $v "#"}-}
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
{-{- $linkcmpnts:=  $optRow.GetByName "LINK"}-}
{-{- range $x,$m:=$linkcmpnts}-}
  goto_{-{$m.UNQ}-}(fm = {}){
    let query = {}
    {-{- $rows:=  $table.GetColumnsByTPName $m.RwName}-}
      {-{- range $i,$c:=$rows}-} 
    query.{-{$c.Name}-} = fm.{-{$c.Name}-}
      {-{- end}-}
    this.goto('{-{$m.URL}-}',query)
  },
{-{- end}-}