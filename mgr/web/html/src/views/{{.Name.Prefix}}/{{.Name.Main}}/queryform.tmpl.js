{-{- $table := .}-}
{-{- $cmpnts :=$table.GetColumsByCmpnt "daterange" "q" -}-}
{-{- $qcols :=  $table.GetColumnsByName "q"}-}
    form_{-{.UNQ}-}: {
        pi: 1,
        ps: 15,
        single_date_range_name:"",
        single_date_range_value:[],
        {-{- range $i,$c := $qcols}-}
        {-{- if eq true (f_string_start $c.Cmpnt.Type "multi")}-}
        {-{$c.Name}-}:[],
        {-{- else}-}
        {-{$c.Name}-}:"",
        {-{- end}-}
        {-{- end}-}
        },
    {-{- range $i,$c :=  $table.GetColumnsByName "q-l-le" }-}
    {-{- if or (eq true $c.Enum.IsEnum) (eq true (f_string_start $c.Cmpnt.Type "multi"))}-}
    {-{.Name}-}List:this.$theia.enum.get("{-{$c.Enum.EnumType}-}","{-{$c.Enum.PID}-}","{-{$c.Enum.Group}-}"),
    {-{.Name}-}Exts:this.$theia.enum.getExts("{-{$c.Enum.EnumType}-}"),
    {-{- if and (eq "tabs" $c.Cmpnt.Type) (ne "" $c.Cmpnt.Format)}-}
    {-{.Name}-}TabList:this.$theia.enum.get("{-{$c.Cmpnt.Format}-}"),
    {-{- end}-}
    {-{- end}-}
    {-{- end}-}
    dataList_{-{.UNQ}-}:[],
    multiQueryDateRange:[{-{- range $i,$c:= $cmpnts}-}{label:"{-{$c.Label}-}",value:"{-{$c.Name}-}"},{-{end}-}],
    total_{-{.UNQ}-}:0,