package struct2bytes

import (
	"bytes"
	"github.com/DiskGetor/liblinux/errcheck"
	"github.com/DiskGetor/liblinux/log"
)

func (p *ObjStructBytes) StructBytes() []byte {
	return p.Bytes()
}

func (p *ObjStructBytes) Struct2Bytes(obj interface{}) bool {
	return p.Write(obj)
}

func (p *ObjStructBytes) Bytes2Struct(StructBytes []byte, obj interface{}) bool {
	return p.Read(StructBytes, obj)
}

//type和value的耦合度太高了
type (
	InterfaceStructBytes interface {
		StructBytes() []byte
		Struct2Bytes(obj interface{}) bool
		Bytes2Struct(StructBytes []byte, obj interface{}) bool
		InterfaceGob
	}
	ObjStructBytes struct {
		*bytes.Buffer
		ObjGob
	}
)

var (
	_ InterfaceStructBytes = (*ObjStructBytes)(nil)
	W                      = new(struct {
		log.ObjLog
		errcheck.ObjErrCheck
	})
)
