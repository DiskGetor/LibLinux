package js

import "testing"

import (
	"github.com/stretchr/testify/assert"
)

func TestInterfaceJs(t *testing.T) { //空框架
	var W InterfaceJs = new(ObjJs)
	assert.True(t, W.JsRun(`
function add(){
	return 1+1+""
}
add()

`))
	assert.Equal(t, W.JsValue(), "2")
}
