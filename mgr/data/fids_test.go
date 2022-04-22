package data

import (
	"testing"

	"github.com/micro-plat/lib4go/assert"
)

func TestFids(t *testing.T) {
	fd := newFids(3)
	assert.Equal(t, "aaa", fd.Current())
	assert.Equal(t, "aab", fd.Next())
}
func TestFids2(t *testing.T) {
	fd := &fids{charIndex: []int{0, 0, 25}}
	assert.Equal(t, "aaz", fd.Current())
	assert.Equal(t, "aba", fd.Next())
}
