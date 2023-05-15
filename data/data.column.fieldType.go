package data

import "github.com/micro-plat/hycli/data/internal/md"

func (c *fieldType) LikeQuery() bool {
	return md.HasConstraint(c.row.Constraints, "#like")
}
