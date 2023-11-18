{-{- $ccols := .}-}
    rules:{
      {-{- range $i,$c := $ccols}-}
      {-{$c.Name}-}:[{
          required:{-{$c.Field.Required}-},
          message:"请输入{-{$c.Label}-}",
          trigger: 'blur',
      }],
    {-{- end}-}
    },
    form:{
        {-{- range $i,$c := $ccols}-}
        {-{- if eq "switch" $c.Cmpnt.Type}-}
        {-{$c.Name}-}_switch:false,
        {-{- else if eq true (f_string_start $c.Cmpnt.Type "multi")}-}
        {-{$c.Name}-}:[],
        {-{- else}-}
        {-{$c.Name}-}:"",
        {-{- end}-}
        {-{- end}-}
        },
        {-{- range $i,$c := $ccols}-}
        {-{- if eq true $c.Enum.IsEnum}-}
        {-{.Name}-}List:[],
        {-{- end}-}
        {-{- end}-}
      }