package UDP

import (
	"bytes"
	"github.com/DiskGetor/liblinux/errcheck"
	"github.com/DiskGetor/liblinux/log"
	"net"
)

var (
	SrcBufChan = make(chan []byte, 1)
	DstBufChan = make(chan []byte, 1)
	W          = new(struct {
		log.ObjLog
		errcheck.ObjErrCheck
	})
)

type (
	InterfaceTransportUDP interface {
		TransportUDP(DstIP string, DstPort int)
		TransportUDPNoChan(DstIP string, DstPort int)
	}
	ObjTransportUDP struct{}
	objUDP          struct {
		ProtoCool string
		*bytes.Buffer
		BufSize int
		srcTransportCtx
		dstTransportCtx
		err error
	}
	srcTransportCtx struct {
		SrcConn *net.UDPConn
		SrcAddr *net.UDPAddr
	}
	dstTransportCtx struct {
		DstIP   string
		DstPort int
		DstConn *net.UDPConn
		DstAddr *net.UDPAddr
	}
)

func test() {
	p := new(ObjTransportUDP)
	p.TransportUDPNoChan(``, 6001)
	p.TransportUDP(``, 6001)
}

func New(DstIP string, DstPort int) (p *objUDP) {
	p = new(objUDP)
	p.ProtoCool = `udp`
	p.DstIP = DstIP
	p.DstPort = DstPort
	p.Buffer = bytes.NewBuffer(nil)
	return
}

func (p *objUDP) GetSrcAddrConn() bool { //设置源地址和连接
	p.SrcConn, p.err = net.ListenUDP(p.ProtoCool, &net.UDPAddr{IP: net.IPv4zero, Port: p.DstPort})
	if !W.CheckErr(p.err) {
		return false
	}
	p.BufSize, p.SrcAddr, p.err = p.SrcConn.ReadFromUDP(p.Bytes())
	return W.CheckErr(p.err)
}

func (p *objUDP) SetDstAddrConn() bool { //设置目标地址和连接
	p.SetDstAddr()
	p.DstConn, p.err = net.DialUDP(p.ProtoCool, nil, p.DstAddr)
	return W.CheckErr(p.err)
}

func (p *objUDP) SetDstAddr() {
	if p.DstAddr != nil {
		return
	}
	if p.BufSize == 48 {
		p.DstAddr = &net.UDPAddr{IP: net.ParseIP(p.DstIP), Port: p.DstPort}
		return
	}
	p.DstAddr = &net.UDPAddr{IP: p.Bytes()[14:18], Port: p.DstPort}
}
