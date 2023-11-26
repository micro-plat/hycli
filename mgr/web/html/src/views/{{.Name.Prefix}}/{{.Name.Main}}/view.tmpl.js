{-{- $table :=. }-}
{-{- $vcols :=  $table.Columns.GetViewColumns}-}
{-{- $viewOpts :=$table.Optrs.ViewOpts}-}
 show(form) {
    this.conf.visible = true
    this.form = form
    {-{- range $i,$c:= $viewOpts}-}
      {-{- if $c.Equal "TAB"}-}
      //{-{ $c.Label}-}查询
      let nform_{-{$c.UNQ}-} = {}
      {-{- $ct:=  $c.GetAssociatedTable true}-}
      {-{- $os :=  $c.GetParamsByAtPrefix }-}
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
      {-{- if $c.Enum.IsEnum}-}
        item.{-{$c.Name}-}_label = that.$theia.enum.getNames("{-{$c.Enum.EnumType}-}",item.{-{$c.Name}-})
      {-{- end }-}
      {-{- if $c.Cmpnt.Equal "switch"}-}
        item.{-{$c.Name}-}_switch = item.{-{$c.Name}-} == 0
      {-{- end}-}
    {-{- end}-}
    
    {-{- range $i,$c := $vcols}-}
        {-{- if and ($c.Cmpnt.HasFormat) ($c.Field.IsDate)}-}
        item.{-{$c.Name}-} = that.$theia.str.dateFormat(item.{-{$c.Name}-},'{-{$c.Cmpnt.Format}-}')
      {-{- else if and ($c.Cmpnt.HasFormat) ($c.Field.IsNumber)}-}
        item.{-{$c.Name}-} = that.$theia.str.numberFormat(item.{-{$c.Name}-},'{-{$c.Cmpnt.Format}-}')
      {-{- else if $c.Cmpnt.Equal "mobile"}-}
        item.{-{$c.Name}-} = that.$theia.str.phoneFormat(item.{-{$c.Name}-})
      {-{- else if $c.Cmpnt.Equal "cut"}-}
        item.{-{$c.Name}-} = that.$theia.str.cut(item.{-{$c.Name}-},{-{$c.Cmpnt.Format}-})
      {-{- else if gt $c.Field.Len 32}-}  
        item.{-{$c.Name}-} = (item.{-{$c.Name}-}||"")
      {-{- end}-}
      {-{- end }-}
      {-{- range $i,$m :=  $viewOpts.GetByName "step"}-}
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
  {-{- range $x,$m:=$viewOpts.GetCompnents }-}
    show_view_{-{$m.UNQ}-}(){
      let that = this;
      let query={}
      {-{- range $i,$c:= $table.Columns.GetColumns $m.RwName}-} 
        query.{-{$c.Name}-} = that.view.{-{$c.Name}-}
      {-{- end}-}
      that.$refs.view_{-{$m.UNQ}-}.show(query)
    },
    {-{- end}-}

 {-{- $stepOpts:= $viewOpts.GetByName "step"}-}
 {-{- if gt $stepOpts.Len 0}-}
 {-{- range $i,$c:= $stepOpts}-}
 getStepActive_{-{$c.UNQ}-}(view){
  {-{- $steps := $table.Columns.GetColumns $c.RwName}-}
  {-{- range $j,$s:= $steps}-}
      {-{- if and ($s.Cmpnt.HasFormat) ($s.Field.IsDate)}-}
        if(!view["{-{$s.Name}-}"]|| new Date(view["{-{$s.Name}-}"]) > new Date()){
          return {-{$j}-}
        }
      {-{- else}-}
      if(!view["{-{$s.Name}-}"]){
         return {-{$j}-}
      }
      {-{- end}-}
  {-{- end }-}
      return {-{$steps.SLen}-}
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