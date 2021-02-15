package TCP

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

func test() {
	p := new(ObjTransportTCP)
	p.TransportTCP(``, 6001)
}

type (
	InterfaceTransportTCP interface {
		TransportTCP(DstIP string, DstPort int)
	}
	ObjTransportTCP struct{}
	objTCP          struct {
		ProtoCool string
		*bytes.Buffer
		BufSize int
		srcTransportCtx
		dstTransportCtx
		err error
	}
	srcTransportCtx struct {
		SrcConn        *net.TCPConn
		SrcAddr        *net.TCPAddr
		SrcTCPListener *net.TCPListener
	}
	dstTransportCtx struct {
		DstIP          string
		DstPort        int
		DstConn        *net.TCPConn
		DstAddr        *net.TCPAddr
		DstTCPListener *net.TCPListener
	}
)

func New(DstIP string, DstPort int) (p *objTCP) {
	p = new(objTCP)
	p.ProtoCool = `tcp`
	p.DstIP = DstIP
	p.DstPort = DstPort
	p.Buffer = bytes.NewBuffer(nil)
	return
}
