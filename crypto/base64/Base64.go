package base64

import (
	"encoding/base64"
	"github.com/DiskGetor/liblinux/crypto/ctx"
	"github.com/DiskGetor/liblinux/errcheck"
	"github.com/DiskGetor/liblinux/log"
)

var W = new(struct {
	log.ObjLog
	errcheck.ObjErrCheck
})

type (
	ObjBase64 struct {
		ctx.IoInfo
	}
	InterfaceBase64 interface {
		Base64Encode(src []byte) string
		Base64Decode(result string) bool
		Base64Bytes() []byte
	}
)

var (
	base64coder                 = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	_           InterfaceBase64 = (*ObjBase64)(nil)
)

func (p *ObjBase64) Base64Encode(src []byte) string {
	return base64coder.EncodeToString(src)
}

func (p *ObjBase64) Base64Bytes() []byte {
	return p.Out
}

func (p *ObjBase64) Base64Decode(result string) bool {
	p.Out, p.Err = base64coder.DecodeString(result)
	if !W.CheckErr(p.Err) {
		return false
	}
	return true
}
