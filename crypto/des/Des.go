package des

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"github.com/DiskGetor/liblinux/crypto/ctx"
	"github.com/DiskGetor/liblinux/errcheck"
	"github.com/DiskGetor/liblinux/log"
	"github.com/DiskGetor/liblinux/tool"
)

type (
	InterfaceDes interface {
		DesBytes() []byte
		DesString() string
		DesErrInfo() string
		DesEncode(in, key interface{}) bool
		DesDecode(in interface{}, key interface{}) bool
	}
	ObjDes struct {
		ctx.IoInfo
		block cipher.Block
	}
)

var (
	_ InterfaceDes = (*ObjDes)(nil)
	W              = new(struct {
		log.ObjLog
		errcheck.ObjErrCheck
		tool.ObjTool
	})
)

func (p *ObjDes) DesBytes() []byte {
	return p.Out
}
func (p *ObjDes) DesString() string {
	return hex.EncodeToString(p.Out)
}
func (p *ObjDes) DesErrInfo() string {
	return p.ErrInfo
}

func (p *ObjDes) DesEncode(in, key interface{}) (ok bool) {
	return p.do(ctx.Encode, in, key)
}

func (p *ObjDes) DesDecode(in, key interface{}) (ok bool) {
	return p.do(ctx.Decode, in, key)
}

func (p *ObjDes) do(model int, in, key interface{}) (ok bool) {
	if !(p.resetObj(in, key)) {
		return
	}
	switch model {
	case ctx.Encode:
		p.block.Encrypt(p.Out, p.In)
	case ctx.Decode:
		p.block.Decrypt(p.Out, p.In)
	}
	return true
}

func (p *ObjDes) resetObj(in, key interface{}) (ok bool) {
	fnCheck := func() (ok bool) {
		switch in.(type) {
		case []byte:
			if len(in.([]byte)) != 8 {
				p.ErrInfo = "src长度不是8字节"
				return
			}
			p.In = in.([]byte)
		case string:
			if len(in.(string)) != 8*2 {
				p.ErrInfo = "src长度不是16个字符"
				return
			}
			p.In, p.Err = hex.DecodeString(in.(string))
			if !W.CheckErr(p.Err) {
				return
			}
		}

		switch key.(type) {
		case []byte:
			if len(key.([]byte)) != 8 {
				p.ErrInfo = "key长度不是8字节"
				return
			}
			p.Key = key.([]byte)
		case string:
			if len(key.(string)) != 8*2 {
				p.ErrInfo = "key长度不是16个字符"
				return
			}
			p.Key, p.Err = hex.DecodeString(in.(string))
			if !W.CheckErr(p.Err) {
				return
			}
		}
		return true
	}

	if !(fnCheck()) {
		return
	}
	p.block, p.Err = des.NewCipher(p.Key)
	if !(W.CheckErr(p.Err)) {
		return
	}
	p.Out = make([]byte, des.BlockSize)
	return true
}
