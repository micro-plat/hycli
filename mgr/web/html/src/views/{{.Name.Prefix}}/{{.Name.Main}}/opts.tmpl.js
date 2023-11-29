{-{- $xtable := . }-}
{-{- $table := $xtable.Current }-}
{-{- $mtable := $xtable.Main}-}
{-{- $optRow:=  $table.Optrs.ListOpts.Merge $table.Optrs.BarOpts}-}
{-{- if $xtable.IsTmpl}-}
{-{- $optRow = $mtable.Optrs.ViewExtCmptOpts.GetLstOptrs  $table.Name.Raw}-}
{-{- end}-}
{-{- $cmpnts:=  $optRow.GetCompnents}-}
{-{- if gt $cmpnts.Len 0}-}
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
    {-{- $rows:=  $table.Columns.GetColumns $m.RwName}-}
    {-{- range $i,$c:=$rows}-} 
    query.{-{$c.Name}-} = fm.{-{$c.Name}-}|| fm.le_{-{$c.Name}-}
      {-{- end}-}
    {-{- end}-}
   
     {-{- $rows:=  $table.Columns.GetColumns ($m.GetParam "labelName" "")}-}
     {-{- range $i,$c:=$rows}-} 
     query.{-{$c.Name}-} = fm.{-{$c.Name}-}|| fm.le_{-{$c.Name}-}
     {-{- end}-}  

    {-{- $pkrows:=  $table.Columns.GetPKColumns}-}
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

    query.$title="{-{$m.Label}-}"
    {-{- if $m.IsMux }-}
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
    {-{- $rows:=  $table.Columns.GetColumns $m.RwName}-}
      {-{- range $i,$c:=$rows}-} 
    query.{-{$c.Name}-} = fm.{-{$c.Name}-}
      {-{- end}-}
    this.$js.page.goto('{-{$m.URL}-}',query)
  },
{-{- end}-}