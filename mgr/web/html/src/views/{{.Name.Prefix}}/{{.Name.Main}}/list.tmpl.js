{-{- $xtable := .}-}
{-{- $table := $xtable.Current}-}
{-{- $mtable := $xtable.Main}-}
{-{- $qrow :=  $table.Columns.GetFontQueryColumns}-}
{-{- $lstRow :=  $table.Columns.GetOnlyListColumns}-}
{-{- $leRow :=  $table.Columns.GetOnlyListExtColumns}-}
{-{- $LLERows:=  $table.Columns.GetFontListColumns}-}
{-{- $MLLERows:=  $mtable.Columns.GetFontListColumns.GetEnumColumns}-}

    queryData_{-{$table.UNQ}-}(mform = {},nform={}){
    //构建查询参数
    let queryForm = Object.assign({},this.form_{-{$table.UNQ}-})
    queryForm = Object.assign(queryForm,nform||{})
    
    //处理多个日期选择
    queryForm[this.form_{-{$table.UNQ}-}.single_date_range_name] = this.form_{-{$table.UNQ}-}.single_date_range_value
    queryForm[this.form_{-{$table.UNQ}-}.single_text_name] = this.form_{-{$table.UNQ}-}.single_text_value
   
    //处理日期范围选择
    {-{- range $i,$c:= $qrow}-}
    {-{- if $c.Cmpnt.Equal "daterange"}-}
      queryForm.start_{-{$c.Name}-} = null
      queryForm.end_{-{$c.Name}-} = null
    if(queryForm.{-{- $c.Name}-} && queryForm.{-{$c.Name}-}.length > 1){
      queryForm.start_{-{$c.Name}-} = queryForm.{-{$c.Name}-}[0]
      queryForm.end_{-{$c.Name}-} = queryForm.{-{$c.Name}-}[1]
    }
    {-{- else if  $c.Cmpnt.StartWith "multi"}-}
      queryForm.{-{$c.Name}-} = (queryForm.{-{$c.Name}-}||[]).join(",")
    {-{- end}-}
    {-{- end}-}

    queryForm[this.form_{-{$table.UNQ}-}.single_date_range_name] = null
    queryForm.single_date_range_value = null
    queryForm.single_date_range_name = null

    queryForm.single_text_value = null
    queryForm.single_text_name = null
 
    //处理关联表{-{$table.Name}-} {-{$xtable.Main.Name}-} {-{$table.Enum.EnumType}-}
    {-{- $exit := false}-}
    {-{- if eq true $xtable.IsTmpl}-}
    {-{- range $x,$v := $xtable.Main.Optrs.ViewOpts}-}
    {-{- if eq $v.URL $table.Name.Raw}-}
    {-{- if and (ne "" $v.RwName) (ne "" $v.FwName)}-}
    {-{- $exit = true}-}
    queryForm.{-{$v.FwName}-} = mform.{-{$v.RwName}-}   
    {-{- end}-}
    {-{- end}-}
    {-{- end}-}
    {-{- end}-}

    {-{- if eq false $exit}-}
    {-{- range $i,$c := $MLLERows}-}
        {-{- if eq $table.Enum.EnumType $c.Enum.EnumType}-}
        queryForm.{-{$table.Enum.Id}-} = mform.{-{$c.Name}-}
        {-{- end}-}
    {-{- end}-}
    {-{- end}-}
   
    //发送查询请求
    let that = this
    that.conf.loading = true
    
    //构建统计查询
    {-{- range $i,$v:= $table.Optrs.LStatOpts}-}
      this.$theia.http.get("{-{$v.URL|lower}-}",queryForm).then(res=>{
        let item = res||{}
        {-{- $uxcols :=  $table.Columns.GetColumns $v.RwName}-}
        {-{- range $i,$c := $uxcols}-}
        {-{- if and ($c.Cmpnt.HasFormat) ($c.Field.IsNumber)}-}
          that.stat_{-{$v.UNQ}-}.{-{$c.Name}-} = that.$theia.str.numberFormat(item.{-{$c.Name}-},'{-{$c.Cmpnt.Format}-}')
       {-{- else}-}
         that.stat_{-{$v.UNQ}-}.{-{$c.Name}-} = item.{-{$c.Name}-}
        {-{- end}-}
        {-{- end}-}
      });
    {-{- end}-}
    
  {-{- if gt $lstRow.Len 0}-}
  //数据查询
  this.$theia.http.get("/{-{$table.Name.MainPath|lower}-}/query",queryForm).then(res=>{
      that.conf.loading = false
      that.dataList_{-{$table.UNQ}-} = res.items||[]
      that.total_{-{$table.UNQ}-} = res.count
      that.resetItemData_{-{$table.UNQ}-}(that,that.dataList_{-{$table.UNQ}-})
    });
  {-{- end}-}
  {-{- range $i, $c:=  $table.Optrs.ChartOpts }-}
  this.$refs.chart_{-{ $c.UNQ }-}.show(this.form_{-{ $table.UNQ }-})
  {-{- end }-}
  {-{- range $i, $c:=  $table.Optrs.ExtOpts }-}
    this.$refs.ext_{-{ $c.UNQ }-}.show(this.form_{-{ $table.UNQ }-})
  {-{- end }-}
  },

  resetItemData_{-{$table.UNQ}-}(that,lst){
    lst.forEach(item => {
      item.__raw = Object.assign({}, item)
    {-{- range $i,$c := $LLERows}-}
    {-{- if $c.Cmpnt.Equal "tree"}-}
    item.{-{$c.Name}-}_label = that.$theia.enum.getTreeNames("{-{$c.Enum.EnumType}-}",item.{-{$c.Name}-})
    {-{- else }-}
    item.{-{$c.Name}-}_label = that.$theia.enum.getNames("{-{$c.Enum.EnumType}-}",item.{-{$c.Name}-})
    {-{- end}-}
    
    {-{- if $c.Cmpnt.Equal "switch"}-}
    item.{-{$c.Name}-}_switch = item.{-{$c.Name}-} == 0
    {-{- else if $c.Cmpnt.Equal "progress"}-}
    {-{- $qrow :=  $table.Columns.GetColumns $c.Cmpnt.Format}-}
    {-{- if gt $qrow.Len 0}-}
    {-{- $frow := f_colum_idx $qrow 0 $c}-}
    {-{- $srow := f_colum_idx $qrow 1 $c}-}
    // 计算日期占比
    item.{-{$c.Name}-}_progress = that.$js.process.getDays(item.{-{$frow.Name}-},item.{-{$srow.Name}-})
    {-{- end}-}
    {-{- else}-}
    item.{-{$c.Name}-}_progress = item.{-{$c.Name}-}
      {-{- end}-}
    {-{- end}-}

    {-{- range $i,$c := $leRow}-}
    {-{- if and ($c.Cmpnt.HasFormat) ($c.Field.IsDate)}-}
    item.le_{-{$c.Name}-} = that.$theia.str.dateFormat(item.{-{$c.Name}-},'{-{$c.Cmpnt.Format}-}')
    {-{- else if and ($c.Cmpnt.HasFormat) ($c.Field.IsNumber)}-}
    item.le_{-{$c.Name}-} = that.$theia.str.numberFormat(item.{-{$c.Name}-},'{-{$c.Cmpnt.Format}-}')
    {-{- else if $c.Cmpnt.Equal "mobile"}-}
    item.le_{-{$c.Name}-} = that.$theia.str.phoneFormat(item.{-{$c.Name}-})
    {-{- else if $c.Cmpnt.Equal "cut"}-}
    item.le_{-{$c.Name}-} = that.$theia.str.cut(item.{-{$c.Name}-},{-{$c.Cmpnt.Format}-})
      {-{else}-}
    item.le_{-{$c.Name}-} = item.{-{$c.Name}-}
    {-{- end}-}
    {-{- end}-}

    {-{- range $i,$c := $lstRow}-}
        {-{- if and ($c.Cmpnt.HasFormat) ($c.Field.IsDate)}-}
    item.{-{$c.Name}-} = that.$theia.str.dateFormat(item.{-{$c.Name}-},'{-{$c.Cmpnt.Format}-}')
      {-{- else if and ($c.Cmpnt.HasFormat) ( $c.Field.IsNumber)}-}
    item.{-{$c.Name}-} = that.$theia.str.numberFormat(item.{-{$c.Name}-},'{-{$c.Cmpnt.Format}-}')
      {-{- else if $c.Cmpnt.Equal "mobile"}-}
    item.{-{$c.Name}-} = that.$theia.str.phoneFormat(item.{-{$c.Name}-})
      {-{- else if $c.Cmpnt.Equal "cut"}-}
    item.{-{$c.Name}-} = that.$theia.str.cut(item.{-{$c.Name}-},{-{$c.Cmpnt.Format}-})
      {-{- end}-}
      {-{- end}-}
    if(item.children){
      that.resetItemData_{-{$table.UNQ}-}(that,item.children)
    }
  });

  },

  {-{- range $i,$c:= $qrow}-}
  {-{- if $c.Cmpnt.Equal "ddl"}-}
  on{-{.Name}-}dropdownClick(f, x) {
    let mf = f || {}
    x.{-{.Name}-}_label = mf.name
    x.{-{.Name}-} = mf.value
    this.onQuery()
  },
{-{- end}-}
{-{- end}-}

{-{- range $i,$c:= $lstRow}-}
{-{- if $c.Cmpnt.Equal "switch" }-}
  on{-{$c.Name}-}SwitchChange(xfrom,v) {
      let form = {}
      form.{-{$c.Name}-} = v == true? 0 : 1;
      {-{- range $i,$v :=  $table.Columns.GetPKColumns}-}
      form.{-{$v.Name}-} = xfrom.{-{$v.Name}-}
      {-{- end}-}
      let that = this
      this.$theia.http.post("/{-{$table.Name.MainPath|lower}-}/switch",form).then(res=>{
        that.$notify.success({title: '成功',message: '修改{-{$c.Label}-}成功',duration:5000})
      }).catch(res=>{
        v = !v
        let code = ((res||{}).response||{}).status||0
        let msg = `修改{-{$c.Label}-}失败(${code})`
        that.$notify.error({title: '失败',message: msg,duration:5000})
      });
  },
{-{- end}-}
{-{- end}-}
{-{- range $i,$c:= $qrow }-}
    {-{- $aColumns :=  $qrow.GetColumnsByTriggerChangeEvent $c.Name }-}
    {-{- if gt $aColumns.Len 0}-} 
    onChange_{-{$c.Name}-}(val){
      {-{- range $j,$x:= $aColumns  }-}
      this.{-{$x.Name}-}List = this.$theia.enum.get("{-{$x.Enum.EnumType}-}",val)
      this.form_{-{$table.UNQ}-}.{-{$x.Name}-} = null
      {-{- end}-}
    },
    {-{- end}-}
    {-{- end}-}
  
    {-{- if eq $table.NeedBatchCheck true}-}
    handleSelectionChange(lst=[]){
      this.bcheck = []
      lst.forEach(f=>{
        {-{- range $i,$v :=  $table.Columns.GetPKColumns}-}
        this.bcheck.push(f.{-{$v.Name}-})
        {-{- end}-}
      })
    },
     {-{- end}-}
     columnfilterHandler(value,row,column){
        const property = column['property']
        return row[property] === value
     },
 {-{- $st := $table.NewSceneByList $table.Columns.GetAll}-}
 {-{- template "enum.tmpl.js" $st }-}