{-{- $table := .}-}
{-{- $dateRangeCmpnts := $table.Columns.GetFontQueryColumns.GetByCmpnt "daterange"}-}
{-{- $textcmpnts := $table.Columns.GetFontQueryColumns.GetByCmpnt "text" }-}
        form_{-{.UNQ}-}: {
            pi: 1,
            ps: 15,
            disabled:false,
            single_date_range_name:"",
            single_date_range_value:[],
            single_text_name:"",
            single_text_value:"",
            {-{- range $i,$c := $table.Columns.GetFontQueryColumns}-}
            {-{- if  $c.Cmpnt.StartWith "multi"}-}
            {-{$c.Name}-}:[],
            {-{- else}-}
            {-{$c.Name}-}:"",
            {-{- end}-}
            {-{- end}-}
            },
    {-{- range $i,$c :=  $table.Columns.GetAllListColumns.GetEnumColumns}-}
            {-{- if $c.Cmpnt.StartWith "multi"}-}
            {-{.Name}-}Exts:[],
            {-{- end}-}
            {-{- if $c.Cmpnt.StartWith "ddmenu"}-}
            {-{.Name}-}List:[],
            {-{- else}-}
            {-{.Name}-}List:[],
            {-{- end}-}
            {-{- if and ($c.Cmpnt.Equal "tabs") ($c.Cmpnt.HasFormat)}-}
            {-{.Name}-}TabList:[],
            {-{- end}-}
            {-{- end}-}
            dataList_{-{.UNQ}-}:[],
            multiQueryDateRange:[{-{- range $i,$c:= $dateRangeCmpnts}-}{label:"{-{$c.Label}-}",value:"{-{$c.Name}-}"},{-{end}-}],
            multiQueryText:[{-{- range $i,$c:= $textcmpnts}-}{label:"{-{$c.Label}-}",value:"{-{$c.Name}-}"},{-{end}-}],
            total_{-{.UNQ}-}:0,