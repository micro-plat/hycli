package data

import (
	"testing"

	"github.com/micro-plat/hycli/data/internal/md"
	"github.com/micro-plat/lib4go/assert"
)

func TestV(t *testing.T) {
	k := "tp(cut,le,4)"
	c, p, v := md.GetConsByTag("tp", k)
	assert.Equal(t, "cut", c)
	assert.Equal(t, "le", p)
	assert.Equal(t, "4", v)
}
func TestV2(t *testing.T) {
	k := "tp(cut,le|4)"
	c, p, v := md.GetConsByTag("tp", k)
	assert.Equal(t, "cut", c)
	assert.Equal(t, "le|4", p)
	assert.Equal(t, "", v)
}
func TestColor(t *testing.T) {
	c, p, v := md.GetConsByTagIgnorecase("color", "color")
	assert.Equal(t, "color", c)
	assert.Equal(t, "", p)
	assert.Equal(t, "", v)
}
