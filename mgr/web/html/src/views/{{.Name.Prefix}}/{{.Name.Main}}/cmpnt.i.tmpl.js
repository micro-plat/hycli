{-{- $table := .}-}

{-{- range $x,$m:=  $table.Optrs.All.GetCompnents}-}
// {-{$m.Label}-}
import {-{$m.UNQ}-} from "{-{f_string_translate $m.URL (f_table_find_by_name $m.Table)}-}"
{-{- end}-}

{-{- range $i,$c:= ($table.Optrs.BarOpts.GetByCmptName "LIST").GetParamsOptrs "gantt" }-}
// {-{$c.Label}-}
import gantt from "@/views/cmpnts/gantt.vue"
{-{- end}-}

{-{- range $i,$c:= $table.Columns.GetFontQueryColumns.GetEnumColumns.GetByCmpnt "ddmenu"}-}
// {-{$c.Label}-}
import ddmenu from "@/views/cmpnts/ddMenu.vue"
{-{- end  }-}


