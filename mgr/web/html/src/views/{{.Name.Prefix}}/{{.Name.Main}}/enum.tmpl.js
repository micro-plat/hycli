{-{- $st := .}-}
    loadEnums_{-{$st.UNQ}-}(){
        {-{- range $i,$c := $st.Columns.GetEnumColumns}-}
        {-{- $deep := $c.GetOtherCmpntValueInt "deep" 99}-}
        {-{- $group:= $c.Enum.GetGroupValue}-}
        {-{- $pid := f_string_contact `"` $c.Enum.PID `"`}-}
        {-{- if ne "" $c.Enum.AssctColumn}-}
        {-{- $pid = f_string_contact "((this.form_" $st.UNQ "||this.form)." $c.Enum.AssctColumn  "||" `""` `)+""`}-}
        {-{- end}-}
        {-{- $url := $c.GetOtherCmpntValue "enumurl"}-}
        //{-{$c.Label}-}
        {-{- if $c.Cmpnt.StartWith "tree|cascader|ddmenu"}-}
        this.{-{.Name}-}List = this.$theia.enum.getTree("{-{$c.Enum.EnumType}-}",{-{$pid}-},{-{if $c.Enum.GroupIsStatic}-}this.$theia.user.get("{-{$group}-}"){-{else}-}"{-{$c.Enum.Group}-}" {-{end}-},{-{$deep}-})
        {-{- else if and ( $c.Cmpnt.Equal "tabs" ) ($c.Cmpnt.HasFormat)}-}
        this.{-{.Name}-}TabList = this.$theia.enum.get("{-{$c.Cmpnt.Format}-}","","",false,"{-{$url}-}")
        {-{- else}-}
        this.{-{.Name}-}List = this.$theia.enum.get("{-{$c.Enum.EnumType}-}",{-{$pid}-},{-{if $c.Enum.GroupIsStatic}-}this.$theia.user.get("{-{$group}-}"){-{else}-}"{-{$c.Enum.Group}-}" {-{end}-},false,"{-{$url}-}")
        {-{- end}-}
        {-{- end}-}
    },
