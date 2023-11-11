{-{- $table := .}-}
{-{- $opts:=$table.ListOpts.Merge $table.BarOpts $table.ChartOpts $table.ExtOpts}-}
{-{- range $x,$m:=  $opts.GetByName "CMPNT"}-}
// {-{$m.Label}-}
import {-{$m.UNQ}-} from "{-{f_string_translate $m.URL (f_table_find_by_name $m.Table)}-}"
{-{- end}-}

{-{- range $i,$c:= $table.BarOpts.GetByCmptName "LIST" }-}
{-{- $gantt := $c.GetParams "gantt" -}-}
{-{- if ne $gantt ""}-}
import gantt from "@/views/cmpnts/gantt.vue"
{-{- end}-}
{-{- end}-}

{-{- $qcols :=  $table.GetColumnsByTPName "q"}-}
{-{- range $i,$c:= $qcols}-}
{-{- if and (eq "ddmenu" $c.Cmpnt.Type) (eq true $c.Enum.IsEnum)}-}
import ddmenu from "@/views/cmpnts/ddMenu.vue"
{-{- end  }-}
{-{- end  }-}


