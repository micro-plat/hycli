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
        {-{- if and (eq true $c.Enum.IsEnum) (eq true (f_string_start $c.Cmpnt.Type "tree"))}-}
        {-{.Name}-}List:this.$theia.enum.getTree("{-{$c.Enum.EnumType}-}","{-{$c.Enum.PID}-}","{-{$c.Enum.Group}-}"),
        {-{- else if or (eq true $c.Enum.IsEnum) (eq true (f_string_start $c.Cmpnt.Type "multi"))}-}
        {-{.Name}-}List:this.$theia.enum.get("{-{$c.Enum.EnumType}-}","{-{$c.Enum.PID}-}","{-{$c.Enum.Group}-}"),
        {-{- end}-}
        {-{- end}-}
      }