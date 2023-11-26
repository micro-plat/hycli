package data

import (
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

type displayCmpnts map[string]displayCmpnt

func (d displayCmpnts) getCmpnt(r *md.Row, tps ...string) displayCmpnt {
	cmpnt := displayCmpnt{Type: "text"}
	var c displayCmpnt
	var tp string
	var ok bool
	for _, tp = range tps {
		if c, ok = d[strings.ToUpper(tp)]; ok {
			cmpnt = c
			break
		}
	}

	if !ok {
		if c, ok := d["*"]; ok {
			cmpnt = c
		} else if md.HasConstraint(r.Constraints, "sl", "SL") {
			cmpnt = displayCmpnt{Type: "select"}
		} else if strings.EqualFold(r.Type.Name, "date") ||
			strings.EqualFold(r.Type.Name, "datetime") {
			cmpnt = displayCmpnt{Type: "date"}
		}
	}
	if cmpnt.Format != "" {
		return cmpnt
	}
	cmpnt.Format = md.GetFormat(types.GetStringByIndex(tps, 0, tp), r.Constraints...)
	if cmpnt.Format != "" {
		return cmpnt
	}
	cmpnt.Format = md.GetFormat("f", r.Constraints...)
	return cmpnt
}
