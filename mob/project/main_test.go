package project

import (
	"testing"

	"github.com/micro-plat/lib4go/assert"
)

func TestFile(t *testing.T) {

	entitys, err := tmpls.ReadDir(tmplName)
	assert.Equal(t, nil, err)
	assert.Equal(t, 6, len(entitys))
	assert.Equal(t, "babel.config.js", entitys[0].Name())
	assert.Equal(t, false, entitys[0].IsDir())

	assert.Equal(t, "jsconfig.json", entitys[1].Name())
	assert.Equal(t, false, entitys[1].IsDir())

	assert.Equal(t, "package.json", entitys[2].Name())
	assert.Equal(t, false, entitys[2].IsDir())

	_, err = tmpls.Open(tmplName + "/babel.config.js")
	assert.Equal(t, nil, err)

	entitys, err = tmpls.ReadDir(tmplName + "/public")
	assert.Equal(t, nil, err)
	assert.Equal(t, "env.conf.pro.json", entitys[1].Name())
	assert.Equal(t, false, entitys[1].IsDir())
}
