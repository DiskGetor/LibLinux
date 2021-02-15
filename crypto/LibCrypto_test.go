package crypto

import "testing"

func TestLibCrypto(t *testing.T) {
	var e InterfaceLibCrypto = new(ObjLibCrypto)
	e.DesErrInfo()
}
