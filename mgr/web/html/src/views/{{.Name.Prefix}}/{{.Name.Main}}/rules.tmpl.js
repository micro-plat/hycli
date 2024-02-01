{-{- $st:= .}-}
rules_{-{$st.UNQ}-}:{
    {-{- range $i,$c := $st.Columns}-}
    {-{$c.Name}-}:[{required:{-{$c.Field.Required}-},message:"请输入{-{$c.Label}-}",trigger: 'blur'}],
    {-{- end}-}
  },