package struct2bytes

import (
	"bytes"
	"encoding/gob"
)

type (
	InterfaceGob interface {
		GoBinaryBytes() []byte
		GoBinaryEncode(obj interface{}) (ok bool)
		GoBinaryDecode(buf []byte, obj interface{}) (ok bool)
	}
	ObjGob struct {
		bytes.Buffer
		err error
	}
)

var (
	_ InterfaceGob = (*ObjGob)(nil)
)

func (p *ObjGob) GoBinaryBytes() []byte {
	return p.Bytes()
}

func (p *ObjGob) GoBinaryEncode(obj interface{}) (ok bool) {
	enc := gob.NewEncoder(&p.Buffer)
	if !W.CheckErr(enc.Encode(obj)) {
		return
	}
	return true
}

func (p *ObjGob) GoBinaryDecode(buf []byte, obj interface{}) (ok bool) {
	if !W.CheckErr2(p.Write(buf)) {
		return
	}
	dec := gob.NewDecoder(&p.Buffer)
	return W.CheckErr(dec.Decode(obj))
}
