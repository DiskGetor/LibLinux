package base64

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterfaceBase64(t *testing.T) {
	var Base64Obj InterfaceBase64 = new(ObjBase64)
	src := []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88}
	result := `ESIzRFVmd4g=`
	assert.Equal(t, Base64Obj.Base64Encode(src), result)
	assert.True(t, Base64Obj.Base64Decode(result))
	assert.Equal(t, Base64Obj.Base64Bytes(), src)
}
