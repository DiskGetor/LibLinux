package tool

import (
	"bytes"
	"encoding/hex"
)

func (p *ObjTool) SwapHexStringToBytes(hexStr string) bool {
	p.tmp, p.err = hex.DecodeString(hexStr)
	if !(W.CheckErr(p.err)) {
		return false
	}

	p.out = make([]byte, len(p.tmp))
	for i, v := range p.tmp {
		p.out[len(p.tmp)-i-1] = v
	}
	return true
}

func (p *ObjTool) SwapString(src []byte) (dst string) {
	to := bytes.Buffer{}
	for k, v := range src {
		if k%2 == 1 {
			to.WriteByte(v)
			to.WriteByte(src[k-1])
		}
	}
	return to.String()
}

func (p *ObjTool) SwapArray(src []byte) (dst []byte) {
	Size := len(src)
	dst = make([]byte, Size)
	for i := 0; i < Size; i++ {
		dst[i] = src[Size-i-1]
	}
	return
}

func (p *ObjTool) SwapUint8FromUint16(v uint16) uint8 { //6613-->16
	tmp := uint16(int32(v) << uint64(int32(4)) >> uint64(int32(8)))
	a := uint8(int32(tmp) << uint64(int32(12)) >> uint64(int32(12)))
	b := uint8(int32(a) >> uint64(int32(4)))
	return uint8(int32(a)<<uint64(int32(4)) | int32(b))
}
