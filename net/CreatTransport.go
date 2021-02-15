package net

import (
	"crypto/tls"
	"github.com/DiskGetor/liblinux/net/Transport/TCP"
	"github.com/DiskGetor/liblinux/net/Transport/UDP"
	"net"
	"net/http"
	"net/url"
	"time"
)

type (
	ObjTransport struct {
		Transport *http.Transport
		TCP.ObjTransportTCP
		UDP.ObjTransportUDP
		dialFunc     func(network, addr string) (net.Conn, error)
		proxyURLFunc func(*http.Request) (*url.URL, error)
	}
	ProxyInfo struct {
		ProxyProtocol string
		ProxyIpPort   string
	}
)

func (p *ObjNet) CreatTransport(data ProxyInfo) (ok bool) {
	if !p.CheckProtocol(data.ProxyProtocol, data.ProxyIpPort) {
		return
	}
	switch data.ProxyProtocol {
	case Socks4, Socks5:
		p.dialFunc = SDial(data.ProxyProtocol + "://" + data.ProxyIpPort + "?timeout=90s")
	case Http, Https:
		URL, err := url.Parse(Http + "://" + data.ProxyIpPort)
		if !W.CheckErr(err) {
			return
		}
		p.proxyURLFunc = http.ProxyURL(URL)
		p.dialFunc = (&net.Dialer{Timeout: 60 * time.Second}).Dial
	}
	p.Transport = &http.Transport{
		Dial:              p.dialFunc,
		Proxy:             p.proxyURLFunc,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: true, //禁用端口转发连接池
	}
	return true
}
