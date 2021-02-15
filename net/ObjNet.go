package net

import (
	"github.com/DiskGetor/liblinux/errcheck"
	"github.com/DiskGetor/liblinux/log"
	"github.com/DiskGetor/liblinux/net/Transport/TCP"
	"github.com/DiskGetor/liblinux/net/Transport/UDP"
	"github.com/DiskGetor/liblinux/tool"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type (
	InterfaceNet interface {
		CreatClient(data ProxyInfo) (ok bool)
		IsCookieInJar(jar *cookiejar.Jar, cookieName, Host string) (ok bool)
		CreatTransport(data ProxyInfo) (ok bool)
		InterToIP(ip int64) string
		CheckProtocol(Protocol, IpPort string) (ok bool)
		IsKnownProtocol(Protocol string) bool
		RegexpWebBodyBlocks(tagName string) string
		HttpRequest(data *HttpRequestCtx) (Response *http.Response, err error)
		HttpReadRequestBody(data *HttpRequestCtx) (buf []byte, err error)
		UDP.InterfaceTransportUDP
		TCP.InterfaceTransportTCP
		ClientGet() *http.Client
		ClientSet(c *http.Client)
		ClientInit()
	}
	ObjNet struct {
		ObjTransport
		Client        *http.Client
		Jar           *cookiejar.Jar
		RequestReader io.Reader
		Request       *http.Request
		Response      *http.Response
	}
)

func (p *ObjNet) ClientGet() *http.Client {
	return p.Client
}

func (p *ObjNet) ClientSet(c *http.Client) {
	p.Client = c
}

func (p *ObjNet) ClientInit() {
	p.Client = &http.Client{
		Timeout: 60 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}

var W = new(struct {
	log.ObjLog
	errcheck.ObjErrCheck
	tool.ObjTool
})
