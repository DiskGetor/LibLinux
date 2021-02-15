package des

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterfaceDes(t *testing.T) {
	type desDta struct {
		key []byte
		in  []byte
		out []byte
	}

	var data = desDta{
		key: []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88},
		in:  []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88},
		out: []byte{0xcd, 0x9, 0xbc, 0x48, 0x76, 0xac, 0xf, 0x2b},
	}
	var d InterfaceDes = new(ObjDes)

	assert.True(t, d.DesEncode(data.in, data.key))
	assert.Equal(t, d.DesBytes(), data.out)

	assert.True(t, d.DesDecode(data.out, data.key))
	assert.Equal(t, d.DesBytes(), data.in)
}
