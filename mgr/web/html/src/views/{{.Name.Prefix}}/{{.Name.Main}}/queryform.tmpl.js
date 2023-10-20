{-{- $table := .}-}
{-{- $dateRangeCmpnts :=$table.GetColumsByCmpnt "daterange" "q" -}-}
{-{- $textcmpnts :=$table.GetColumsByCmpnt "text" "q" -}-}
{-{- $qcols :=  $table.GetColumnsByTPName "q"}-}
    form_{-{.UNQ}-}: {
        pi: 1,
        ps: 15,
        single_date_range_name:"",
        single_date_range_value:[],
        single_text_name:"",
        single_text_value:"",
        {-{- range $i,$c := $qcols}-}
        {-{- if eq true (f_string_start $c.Cmpnt.Type "multi")}-}
        {-{$c.Name}-}:[],
        {-{- else}-}
        {-{$c.Name}-}:"",
        {-{- end}-}
        {-{- end}-}
        },
    {-{- range $i,$c :=  $table.GetColumnsByTPName "q-l-le" }-}
    {-{- if or (eq true $c.Enum.IsEnum) (eq true (f_string_start $c.Cmpnt.Type "multi"))}-}
    {-{.Name}-}List:this.$theia.enum.get("{-{$c.Enum.EnumType}-}","{-{$c.Enum.PID}-}","{-{$c.Enum.Group}-}"),
    {-{.Name}-}Exts:this.$theia.enum.getExts("{-{$c.Enum.EnumType}-}"),
    {-{- if and (eq "tabs" $c.Cmpnt.Type) (ne "" $c.Cmpnt.Format)}-}
    {-{.Name}-}TabList:this.$theia.enum.get("{-{$c.Cmpnt.Format}-}"),
    {-{- end}-}
    {-{- end}-}
    {-{- end}-}
    dataList_{-{.UNQ}-}:[],
    multiQueryDateRange:[{-{- range $i,$c:= $dateRangeCmpnts}-}{label:"{-{$c.Label}-}",value:"{-{$c.Name}-}"},{-{end}-}],
    multiQueryText:[{-{- range $i,$c:= $textcmpnts}-}{label:"{-{$c.Label}-}",value:"{-{$c.Name}-}"},{-{end}-}],
    total_{-{.UNQ}-}:0,