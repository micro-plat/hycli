{-{- $xtable := .}-}
{-{- $table := $xtable.Current}-}
{-{- $mtable := $xtable.Main}-}
{-{- $qrow := fltrColumns $table "q"}-}
{-{- $lstRow := fltrColumns $table "l"}-}
{-{- $leRow := fltrColumns $table "le"}-}
{-{- $LLERows:= fltrColumns $table "l-le"}-}
{-{- $MLLERows:= fltrColumns $mtable "l-le"}-}

    queryData_{-{$table.UNQ}-}(mform = {},nform={}){
    //构建查询参数
    let queryForm = Object.assign({},this.form_{-{$table.UNQ}-})
    queryForm = Object.assign(queryForm,nform||{})
    {-{- range $i,$c:= $qrow}-}
    {-{- if eq "daterange" $c.Cmpnt.Type}-}
      queryForm.start_{-{$c.Name}-} = null
      queryForm.end_{-{$c.Name}-} = null
    if(queryForm.{-{- $c.Name}-} && queryForm.{-{$c.Name}-}.length > 1){
      queryForm.start_{-{$c.Name}-} = queryForm.{-{$c.Name}-}[0]
      queryForm.end_{-{$c.Name}-} = queryForm.{-{$c.Name}-}[1]
    }
    {-{- else if eq true (fltrStart $c.Cmpnt.Type "multi")}-}
      queryForm.{-{$c.Name}-} = (queryForm.{-{$c.Name}-}||[]).join(",")
    {-{- end}-}
    {-{- end}-}
 
    //处理关联表{-{$table.Name}-} {-{$xtable.Main.Name}-} {-{$table.Enum.EnumType}-}
    {-{- $exit := false}-}
    {-{- range $x,$v := $xtable.Main.ViewOpts}-}
    {-{- if eq $v.URL $table.Name.Raw}-}
    {-{- if and (ne "" $v.RwName) (ne "" $v.FwName)}-}
    {-{- $exit = true}-}
    queryForm.{-{$v.FwName}-} = mform.{-{$v.RwName}-}   
  {-{- end}-}
    {-{- end}-}
{-{- end}-}

    {-{- if eq false $exit}-}
    {-{- range $i,$c := $MLLERows}-}
      {-{- if eq true $c.Enum.IsEnum}-}
        {-{- if eq $table.Enum.EnumType $c.Enum.EnumType}-}
        queryForm.{-{$table.Enum.Id}-} = mform.{-{$c.Name}-}
        {-{- end}-}
    {-{- end}-}
    {-{- end}-}
    {-{- end}-}
   
    //发送查询请求
    let that = this
    that.conf.loading = true
    
    //构建统计查询
    {-{- range $i,$v:= $table.LStatOpts}-}
      this.$theia.http.get("{-{$v.URL|lower}-}",queryForm).then(res=>{
        let item = res||{}
        {-{- $uxcols := fltrColumns $table $v.RwName}-}
        {-{- range $i,$c := $uxcols}-}
        {-{- if and (ne "" $c.Cmpnt.Format) (eq true $c.Field.IsNumber)}-}
          that.stat_{-{$v.UNQ}-}.{-{$c.Name}-} = that.$theia.str.numberFormat(item.{-{$c.Name}-},'{-{$c.Cmpnt.Format}-}')
       {-{- else}-}
         that.stat_{-{$v.UNQ}-}.{-{$c.Name}-} = item.{-{$c.Name}-}
        {-{- end}-}
        {-{- end}-}
      });
    {-{- end}-}
    
    {-{- if gt (len $lstRow) 0}-}
    //数据查询
    this.$theia.http.get("/{-{$table.Name.MainPath|lower}-}/query",queryForm).then(res=>{
        that.conf.loading = false
        that.dataList_{-{$table.UNQ}-} = res.items||[]
        that.total_{-{$table.UNQ}-} = res.count
        that.resetItemData(that,that.dataList_{-{$table.UNQ}-})
      });
    {-{- end}-}
  },

  resetItemData(that,lst){
    lst.forEach(item => {
      {-{- range $i,$c := $LLERows}-}
    {-{- if eq true $c.Enum.IsEnum}-}
    item.{-{$c.Name}-}_label = that.$theia.enum.getNames("{-{$c.Enum.EnumType}-}",item.{-{$c.Name}-})
      {-{- end}-}
      {-{- if eq "switch" $c.Cmpnt.Type}-}
    item.{-{$c.Name}-}_switch = item.{-{$c.Name}-} == 0
    {-{- else if eq "progress" $c.Cmpnt.Type}-}
      {-{- $qrow := fltrColumns $table $c.Cmpnt.Format}-}
      {-{- if gt (len $qrow) 0}-}
      {-{- $frow := getColumn $qrow 0 $c}-}
      {-{- $srow := getColumn $qrow 1 $c}-}
      // 计算日期占比
      item.{-{$c.Name}-}_progress = that.getDaysProgress(item.{-{$frow.Name}-},item.{-{$srow.Name}-})
      {-{- end}-}
      {-{- else}-}
    item.{-{$c.Name}-}_progress = item.{-{$c.Name}-}
      {-{- end}-}
    {-{- end}-}

    {-{- range $i,$c := $leRow}-}
      {-{- if and (ne "" $c.Cmpnt.Format) (eq true $c.Field.IsDate)}-}
    item.le_{-{$c.Name}-} = that.$theia.str.dateFormat(item.{-{$c.Name}-},'{-{$c.Cmpnt.Format}-}')
    {-{- else if and (ne "" $c.Cmpnt.Format) (eq true $c.Field.IsNumber)}-}
    item.le_{-{$c.Name}-} = that.$theia.str.numberFormat(item.{-{$c.Name}-},'{-{$c.Cmpnt.Format}-}')
    {-{- else if eq "mobile" $c.Cmpnt.Type}-}
    item.le_{-{$c.Name}-} = that.$theia.str.phoneFormat(item.{-{$c.Name}-})
    {-{- else if eq "cut" $c.Cmpnt.Type}-}
    item.le_{-{$c.Name}-} = that.$theia.str.cut(item.{-{$c.Name}-},{-{$c.Cmpnt.Format}-})
      {-{else}-}
    item.le_{-{$c.Name}-} = item.{-{$c.Name}-}
    {-{- end}-}
    {-{- end}-}

    {-{- range $i,$c := $lstRow}-}
        {-{- if and (ne "" $c.Cmpnt.Format) (eq true $c.Field.IsDate)}-}
    item.{-{$c.Name}-} = that.$theia.str.dateFormat(item.{-{$c.Name}-},'{-{$c.Cmpnt.Format}-}')
      {-{- else if and (ne "" $c.Cmpnt.Format) (eq true $c.Field.IsNumber)}-}
    item.{-{$c.Name}-} = that.$theia.str.numberFormat(item.{-{$c.Name}-},'{-{$c.Cmpnt.Format}-}')
      {-{- else if eq "mobile" $c.Cmpnt.Type}-}
    item.{-{$c.Name}-} = that.$theia.str.phoneFormat(item.{-{$c.Name}-})
      {-{- else if eq "cut" $c.Cmpnt.Type}-}
    item.{-{$c.Name}-} = that.$theia.str.cut(item.{-{$c.Name}-},{-{$c.Cmpnt.Format}-})
      {-{- end}-}
      {-{- end}-}
    if(item.children){
      that.resetItemData(that,item.children)
    }
  });

  },

  {-{- range $i,$c:= $qrow}-}
  {-{- if eq "ddl" $c.Cmpnt.Type}-}
  on{-{.Name}-}dropdownClick(f, x) {
    let mf = f || {}
    x.{-{.Name}-}_label = mf.name
    x.{-{.Name}-} = mf.value
    this.onQuery()
  },
{-{- end}-}
{-{- end}-}

{-{- range $i,$c:= $lstRow}-}
{-{- if eq "switch" $c.Cmpnt.Type}-}
  on{-{$c.Name}-}SwitchChange(xfrom) {
    let form = {}
    form.{-{$c.Name}-} = xfrom.{-{$c.Name}-} == true?0:1;
    {-{- range $i,$v :=  $table.PKColumns}-}
    form.{-{$v.Name}-} = xfrom.{-{$v.Name}-}
    {-{- end}-}
    let that = this
    this.$theia.http.post("/{-{$table.Name.MainPath|lower}-}/switch",form).then(res=>{
      that.$notify.success({title: '成功',message: '修改{-{$c.Label}-}成功',duration:5000})
    }).catch(res=>{
      let code = ((res||{}).response||{}).status||0
      let msg = `修改{-{$c.Label}-}失败(${code})`
      that.$notify.error({title: '失败',message: msg,duration:5000})
    });
  },
{-{- end}-}
{-{- end}-}
{-{- range $i,$c:= $qrow }-}
    {-{- $aColumns := fltrAssctColumns $qrow $c.Name }-}
    {-{- if gt (len $aColumns) 0}-} 
    onChange_{-{$c.Name}-}(val){
      {-{- range $j,$x:= $aColumns  }-}
      this.{-{$x.Name}-}List = this.$theia.enum.get("{-{$x.Enum.EnumType}-}",val)
      this.form_{-{$table.UNQ}-}.{-{$x.Name}-} = null
      {-{- end}-}
    },
    {-{- end}-}
    {-{- end}-}
    {-{- $progress := fltrCmpnt $table "progress" "l"}-}
    {-{- if gt (len $progress) 0}-}
    getDaysProgress(s, e) {
      let start = s ? new Date(s.replace(/-/g, "/")) : null;
      let end = e ? new Date(e.replace(/-/g, "/")) : null;
      let now = new Date()

      //结束时间小于当前时间
      if (end && end < now) {
        return 100
      }

      //开始时间为空，或大于当前时间
      if (!start || start > now) {
        return 0
      }

      //结束时间为空，开始时间不为空，默认为50%
      if (!end) {
        return 50
      }

      //开始 当前时间 结束时间
      let total = this.getWeekHours(start, end)
      let pass = this.getWeekHours(start, now)
      let pv = (pass / total) * 100
      let persent = pv >100?"100":pv  + ""
      let lst = persent.split(".")
      return lst.length == 0 ? persent : lst[0]
    },
    getWeekHours(beginDate, endDate) {
      //日期差值,即包含周六日、以天为单位的工时，86400000=1000*60*60*24.  
      var workDayVal = (endDate - beginDate) / 86400000 + 1;
      //工时的余数  
      var remainder = workDayVal % 7;
      //工时向下取整的除数  
      var divisor = Math.floor(workDayVal / 7);
      var weekendDay = 2 * divisor;
      //起始日期的星期，星期取值有（1,2,3,4,5,6,0）  
      var nextDay = beginDate.getDay();
      //从起始日期的星期开始 遍历remainder天  
      for (var tempDay = remainder; tempDay >= 1; tempDay--) {
        //第一天不用加1  
        if (tempDay == remainder) {
          nextDay = nextDay + 0;
        } else if (tempDay != remainder) {
          nextDay = nextDay + 1;
        }
        //周日，变更为0  
        if (nextDay == 7) {
          nextDay = 0;
        }
        //周六日  
        if (nextDay == 0 || nextDay == 6) {
          weekendDay = weekendDay + 1;
        }
      }
      //实际工时（天） = 起止日期差 - 周六日数目。  
      let rd = workDayVal - weekendDay;
      return rd < 0 ? 0 : rd * 8; //换算为每天8小时
    },
    {-{- end }-}

    {-{- if eq $table.NeedBatchCheck true}-}
    handleSelectionChange(lst=[]){
      this.bcheck = []
      lst.forEach(f=>{
        {-{- range $i,$v :=  $table.PKColumns}-}
        this.bcheck.push(f.{-{$v.Name}-})
        {-{- end}-}
      })
    },
     {-{- end}-}