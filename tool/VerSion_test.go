package tool

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterfaceVerSion(t *testing.T) {
	var W InterfaceVerSion = new(ObjVerSion)
	v := `56196.439.0`
	assert.True(t, W.StringToVerSion(v))
	f := &ObjVerSion{
		major: 0xdb84,
		minor: 0x1b7,
		patch: 0,
	}
	assert.Equal(t, f.VerSionToString(f), v)
}
