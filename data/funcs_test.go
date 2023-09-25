package data

import (
	"testing"

	"github.com/micro-plat/lib4go/assert"
)

func TestF_string_equal(t *testing.T) {
	assert.Equal(t, true, f_string_equal("PIE", "BAR-PIE-LINE"))
}
