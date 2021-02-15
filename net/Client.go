package net

import (
	"net/http"
	"net/http/cookiejar"
	"time"
)

func (p *ObjNet) CreatClient(data ProxyInfo) (ok bool) {
	Jar, err := cookiejar.New(nil)
	if !W.CheckErr(err) {
		return
	}
	if !p.CreatTransport(data) {
		W.CheckErr(`p.Transport is nil`)
		return
	}
	p.Client = &http.Client{
		Jar:       Jar,
		Timeout:   60 * time.Second,
		Transport: p.Transport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	return true
}
