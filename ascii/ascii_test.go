package ascii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	p := new(ObjAscii)
	ascii := "CFID"
	packedBuf := []byte{0x8e, 0x6a, 0x64}
	assert.Equal(t, p.AsciiPack(ascii), packedBuf)
	assert.Equal(t, p.AsciiUnPack(packedBuf), ascii)
}
