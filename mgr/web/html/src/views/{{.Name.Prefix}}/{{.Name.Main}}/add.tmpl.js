{-{- $ccols := .}-}
    rules:{
      {-{- range $i,$c := $ccols}-}
      {-{$c.Name}-}:[{required:{-{$c.Field.Required}-},message:"请输入{-{$c.Label}-}",trigger: 'blur'}],
    {-{- end}-}
    },
    form:{
        {-{- range $i,$c := $ccols}-}
        {-{- if $c.Cmpnt.Equal "switch"}-}
        {-{$c.Name}-}_switch:false,
        {-{- else if $c.Cmpnt.StartWith "multi"}-}
        {-{$c.Name}-}:[],
        {-{- else}-}
        {-{$c.Name}-}:"",
        {-{- end}-}
        {-{- end}-}
    },
    {-{- range $i,$c := $ccols.GetEnumColumns}-}
    {-{.Name}-}List:[],
    {-{- end}-}
    }