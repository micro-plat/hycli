{-{- $table :=. }-}
{-{- $vcols := fltrColumns $table "v"}-}
{-{- $viewOpts :=$table.ViewOpts}-}
 show(form) {
    this.conf.visible = true
    this.form = form
    {-{- range $i,$c:= $viewOpts}-}
      {-{- if eq "TAB" $c.Name}-}
     
      //{-{ $c.Label}-}查询
      let nform_{-{$c.UNQ}-} = {}
      {-{- $ct:= fltrSearchUITableAndResetUNQ $c}-}
      {-{- $os := fltrFrontOptrsByStatic $c }-}
      {-{- range $k,$v := $os}-}
      nform_{-{$c.UNQ}-}.{-{$k}-} =this.$route.params["{-{$k}-}"]|| "{-{$v}-}"
      {-{- end}-}
      this.queryData_{-{$ct.UNQ}-}(form,nform_{-{$c.UNQ}-})
      {-{- end}-}
     {-{- end}-}
  
   let that = this;
    this.$theia.http
      .get("/{-{.Name.MainPath|lower}-}",form)
      .then((res) => {
        let item = Object.assign({}, res)
    {-{- range $i,$c := $vcols }-}
    {-{- if eq true $c.Enum.IsEnum}-}
        item.{-{$c.Name}-}_label = that.$theia.enum.getNames("{-{$c.Enum.EnumType}-}",item.{-{$c.Name}-})
      {-{- end }-}
      {-{- if eq "switch" $c.Cmpnt.Type}-}
        item.{-{$c.Name}-}_switch = item.{-{$c.Name}-} == 0
      {-{- end}-}
    {-{- end}-}
    
    {-{- range $i,$c := $vcols}-}
        {-{- if and (ne "" $c.Cmpnt.Format) (eq true $c.Field.IsDate)}-}
        item.{-{$c.Name}-} = that.$theia.str.dateFormat(item.{-{$c.Name}-},'{-{$c.Cmpnt.Format}-}')
      {-{- else if and (ne "" $c.Cmpnt.Format) (eq true $c.Field.IsNumber)}-}
        item.{-{$c.Name}-} = that.$theia.str.numberFormat(item.{-{$c.Name}-},'{-{$c.Cmpnt.Format}-}')
      {-{- else if eq "mobile" $c.Cmpnt.Type}-}
        item.{-{$c.Name}-} = that.$theia.str.phoneFormat(item.{-{$c.Name}-})
      {-{- else if eq "cut" $c.Cmpnt.Type}-}
        item.{-{$c.Name}-} = that.$theia.str.cut(item.{-{$c.Name}-},{-{$c.Cmpnt.Format}-})
      {-{- else if gt $c.Field.Len 32}-}  
        item.{-{$c.Name}-} = (item.{-{$c.Name}-}||"").replace(/\n/g,"<br/>")
      {-{- end}-}
      {-{- end }-}
      {-{- $steps := fltrOptrs $viewOpts "step"}-}
      {-{- range $i,$m := $steps}-}
        that.conf.stepActive_{-{$m.UNQ}-} = that.getStepActive_{-{$m.UNQ}-}(res)
      {-{- end}-}
        that.view = item
            
      })
      .catch((res) => {
        let code = res.response.status;
        let msg = `{-{.Desc}-}查询失败(${code})`;
        that.$notify.error({ title: "失败", message: msg, duration: 5000 });
      });
  },
  {-{- range $x,$m:=$viewOpts }-}
    {-{- if eq "CMPNT" $m.Name }-}
    show_view_{-{$m.UNQ}-}(){
      let that = this;
      let query={}
      {-{- $rows:= fltrColumns $table $m.RwName }-}
      {-{- range $i,$c:=$rows}-} 
        query.{-{$c.Name}-} = that.view.{-{$c.Name}-}
      {-{- end}-}
      that.$refs.view_{-{$m.UNQ}-}.show(query)
    },
    {-{- end}-}
 {-{- end}-}

 {-{- $stepOpts:=fltrOptrs $viewOpts "step"}-}
 {-{- if gt (len $stepOpts) 0}-}
 {-{- range $i,$c:= $stepOpts}-}
 getStepActive_{-{$c.UNQ}-}(view){
 
  {-{- $steps :=  fltrColumns $table $c.RwName}-}
  {-{- range $j,$s:= $steps}-}
      {-{- if and (ne "" $s.Cmpnt.Format) (eq true $s.Field.IsDate)}-}
        if(!view["{-{$s.Name}-}"]|| new Date(view["{-{$s.Name}-}"]) > new Date()){
          return {-{$j}-}
        }
      {-{- else}-}
      if(!view["{-{$s.Name}-}"]){
         return {-{$j}-}
      }
      {-{- end}-}
  {-{- end }-}
        return {-{(len $steps)|minus}-}
 },
 {-{- end}-}
 {-{- end}-}
   colorful(r){
     return this.$theia.env.conf.colorful[r]
  },
  tagColor(r,name){
    if(this.$theia.env.conf.colorTag[name]){
          return this.$theia.env.conf.colorTag[name][r]||""
      }
     return this.$theia.env.conf.colorTag[r]||""
  },