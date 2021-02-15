package TCP

import (
	"fmt"
	"net"
)

func (ObjTransportTCP) TransportTCP(DstIP string, DstPort int) {
	p := New(DstIP, DstPort)
	if !p.getSrcAddrConn() {
		return
	}
	defer func() {
		W.CheckErr(p.SrcTCPListener.Close())
		W.CheckErr(p.DstTCPListener.Close())
	}()
	for {
		p.BufSize, p.err = p.SrcConn.Read(p.Bytes())
		if !W.CheckErr(p.err) {
			return
		}
		SrcBufChan <- p.Bytes()[:p.Len()]
		go p.do()
		if !W.CheckErr2(p.SrcConn.Write(<-DstBufChan)) {
			return
		}
	}
}

func (p *objTCP) getSrcAddrConn() (ok bool) { //设置源地址和连接
	p.SrcAddr, p.err = net.ResolveTCPAddr(p.ProtoCool, "0.0.0.0"+":"+fmt.Sprint(p.DstPort))
	if !W.CheckErr(p.err) {
		return
	}
	p.SrcTCPListener, p.err = net.ListenTCP(p.ProtoCool, p.SrcAddr)
	if !W.CheckErr(p.err) {
		return
	}
	p.SrcConn, p.err = p.SrcTCPListener.AcceptTCP()
	return W.CheckErr(p.err)
}

func (p *objTCP) do() {
	p.DstConn, p.err = net.DialTCP(p.ProtoCool, p.DstAddr, nil)
	if !W.CheckErr(p.err) {
		return
	}
	for {
		p.BufSize, p.err = p.DstConn.Write(<-SrcBufChan)
		if !W.CheckErr(p.err) {
			return
		}
		p.Reset()
		p.BufSize, p.err = p.DstConn.Read(p.Bytes())
		if !W.CheckErr(p.err) {
			return
		}
		DstBufChan <- p.Bytes()[:p.Len()]
	}
}
