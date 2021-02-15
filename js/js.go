package js

import (
	"github.com/DiskGetor/liblinux/errcheck"
	"github.com/DiskGetor/liblinux/log"
	"github.com/dop251/goja"
)

var (
	W = new(struct {
		log.ObjLog
		errcheck.ObjErrCheck
	})
)

type (
	InterfaceJs interface {
		JsValue() string
		JsRun(src string) bool
	}
	ObjJs struct {
		str   string
		value goja.Value
		err   error
	}
)

func (p *ObjJs) JsValue() string {
	return p.str
}

func (p *ObjJs) JsRun(src string) bool {
	p.value, p.err = goja.New().RunString(src)
	if !(W.CheckErr(p.err)) {
		return false
	}
	p.str = p.value.Export().(string)
	return true
}
