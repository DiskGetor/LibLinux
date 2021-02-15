package ascii

import (
	"encoding/binary"
	"github.com/DiskGetor/liblinux/log"
	"github.com/DiskGetor/liblinux/tool"
)

var W = new(struct {
	log.ObjLog
	tool.ObjTool
})

type (
	InterfaceAscii interface {
		AsciiPack(ascii string) []byte
		AsciiUnPack(packedBuf []byte) string
	}
	ObjAscii struct {
	}
	ctx struct {
		ascii string
		value uint32
		buf   []byte
	}
)

func (ObjAscii) AsciiPack(ascii string) []byte {
	return new(ctx).AsciiPack(ascii)
}
func (ObjAscii) AsciiUnPack(packedBuf []byte) string {
	return new(ctx).AsciiUnPack(packedBuf)
}

const maxAsciiLen = 4

func (p ctx) AsciiPack(ascii string) []byte {
	p = ctx{
		ascii: ascii,
		value: 0,
		buf:   make([]byte, maxAsciiLen),
	}
	for i := 0; i < len(p.ascii); i++ {
		p.value |= (32 | uint32(p.ascii[i])&31) << ((3 - i) * 6)
	}
	binary.BigEndian.PutUint32(p.buf, p.value)
	p.buf = p.buf[1:]
	return p.buf
}

func (p ctx) AsciiUnPack(packedBuf []byte) string {
	p = ctx{
		ascii: "",
		value: 0,
		buf:   packedBuf,
	}
	buf0 := make([]byte, 1)
	p.buf = W.JoinBytes(buf0, p.buf)
	p.value = binary.BigEndian.Uint32(p.buf)
	tmp := make([]byte, maxAsciiLen)
	for i := 0; i < maxAsciiLen; i++ {
		v := (p.value >> ((3 - i) * 6)) & 63
		tmp[i] = byte(0x40 | (v & 0x1F))
	}
	p.ascii = string(tmp)
	return p.ascii
}
