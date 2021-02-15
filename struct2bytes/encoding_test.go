package struct2bytes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	type ObjXX struct {
		Buf []byte
	}
	var p InterfaceStructBytes = new(ObjStructBytes)
	StructBytes := []byte{
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88,
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88,
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99,
	}
	obj := ObjXX{Buf: StructBytes}
	assert.True(t, p.Struct2Bytes(&obj))
	assert.Equal(t, p.StructBytes(), StructBytes)

	objNew := ObjXX{Buf: make([]byte, len(StructBytes))}
	assert.True(t, p.Bytes2Struct(StructBytes, objNew))
	assert.Equal(t, objNew, obj)
}
