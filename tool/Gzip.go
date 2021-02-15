package tool

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func (p *ObjTool) GzipDecode(in []byte) (ok bool) {
	p.gzipReader, p.err = gzip.NewReader(bytes.NewReader(in))
	if !(W.CheckErr(p.err)) {
		return
	}
	defer func() { W.CheckErr(p.gzipReader.Close()) }()
	p.out, p.err = ioutil.ReadAll(p.gzipReader)
	if !(W.CheckErr(p.err)) {
		return
	}
	return true
}
