package crypto

import (
	"github.com/DiskGetor/liblinux/crypto/base64"
	"github.com/DiskGetor/liblinux/crypto/des"
)

type (
	ObjLibCrypto struct {
		base64.ObjBase64
		des.ObjDes
	}
	InterfaceLibCrypto interface {
		des.InterfaceDes
		base64.InterfaceBase64
	}
)

var _ InterfaceLibCrypto = (*ObjLibCrypto)(nil)
