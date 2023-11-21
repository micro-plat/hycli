{-{- $st := .}-}
{-{- $table := $st.Table}-}
{-{- $ccols := $st.Cols}-}
    loadEnums(){
        {-{- range $i,$c := $ccols}-}
        {-{- if eq true $c.Enum.IsEnum }-}
        // {-{$c.Enum.AssctColumn}-}
        {-{- $pid := f_string_contact `"` $c.Enum.PID `"`}-}
        {-{- if ne "" $c.Enum.AssctColumn}-}
        {-{- if eq "" $st.FormUNQ}-}
        {-{- $pid = f_string_contact "this.form." $c.Enum.AssctColumn `+""`}-}
        {-{- else}-}
        {-{- $pid = f_string_contact "this.form_" $st.FormUNQ "." $c.Enum.AssctColumn `+""`}-}
        {-{- end}-}
       
        {-{- end}-}
        {-{- if eq true (f_string_start $c.Cmpnt.Type "tree|cascader")}-}
        {-{- $deep := f_num_get ($c.GetOpt "deep") 99}-}
        {-{- $group:= f_string_trim $c.Enum.Group "#"}-}
        this.{-{.Name}-}List = this.$theia.enum.getTree("{-{$c.Enum.EnumType}-}",{-{$pid}-},{-{if f_string_start $c.Enum.Group "#"}-}this.$theia.user.get("{-{$group}-}"){-{else}-}"{-{$c.Enum.Group}-}" {-{end}-},{-{$deep}-})
        {-{- else if eq true (f_string_start $c.Cmpnt.Type "ddmenu")}-}
        this.{-{.Name}-}List = this.$theia.enum.getTree("{-{$c.Enum.EnumType}-}","{-{$c.Enum.EnumType}-}","")
        {-{- else if or (eq true $c.Enum.IsEnum) (eq true (f_string_start $c.Cmpnt.Type "multi"))}-}
        {-{- $group:= f_string_trim $c.Enum.Group "#"}-}
        this.{-{.Name}-}List = this.$theia.enum.get("{-{$c.Enum.EnumType}-}",{-{$pid}-},{-{if f_string_start $c.Enum.Group "#"}-}this.$theia.user.get("{-{$group}-}"){-{else}-}"{-{$c.Enum.Group}-}" {-{end}-})
        {-{- end}-}
        {-{- if and (eq "tabs" $c.Cmpnt.Type) (ne "" $c.Cmpnt.Format)}-}
        this.{-{.Name}-}TabList = this.$theia.enum.get("{-{$c.Cmpnt.Format}-}")
        {-{- end}-}
        {-{- end}-}
        {-{- end}-}
    },
