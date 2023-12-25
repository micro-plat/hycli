package data

import (
	"strings"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/types"
)

type color struct {
	Name      string
	FontColor string
	BgColor   string
}

// 显示样式
type displayStyle struct {
	ListWidth    string //列表宽度
	Rows         int    //显示的行数，如textArea行数
	Position     string //位置 default,默认，换行newline
	HideOverflow bool   //超出隐藏
	// CLR          color  //颜色信息
	HasBgColor bool
	Min        int //最小值
	Max        int //最大值
}

// 解析列表样式
func createStyle(r *md.Row) displayStyle {
	min, max := md.GetRanges(r.Constraints...)
	// fc, bc, _ := md.GetConsByTagIgnorecase("color", r.Constraints...)
	hasBGcolor := md.HasConstraint(r.Constraints, "color")

	return displayStyle{
		ListWidth:    md.GetConsNameByTagIgnorecase("lw", r.Constraints...),
		Rows:         types.GetInt(md.GetConsNameByTagIgnorecase("rows", r.Constraints...)),
		Position:     types.GetString(md.GetConsNameByTagIgnorecase("ps", r.Constraints...)),
		HideOverflow: md.HasConstraint(r.Constraints, "hof", "HOF"),
		Min:          types.GetInt(min),
		Max:          types.GetInt(max),
		HasBgColor:   hasBGcolor,
	}
}
func (d *displayStyle) IsFull() bool {
	return strings.EqualFold(d.Position, "full")
}
