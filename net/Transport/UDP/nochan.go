package UDP

func (ObjTransportUDP) TransportUDPNoChan(DstIP string, DstPort int) {
	p := New(DstIP, DstPort)
	if !p.GetSrcAddrConn() {
		return
	}
	defer func() {
		W.CheckErr(p.SrcConn.Close())
		W.CheckErr(p.DstConn.Close())
	}()
	for {
		if !p.SetDstAddrConn() {
			return
		}
		go p.srcWriteDstBuf()
		if !W.CheckErr2(p.DstConn.Write(p.Bytes()[:p.Len()])) {
			return
		}
	}
}

func (p *objUDP) srcWriteDstBuf() { //源连接写入目标buf
	p.Reset()
	p.BufSize, p.err = p.DstConn.Read(p.Bytes())
	if !W.CheckErr(p.err) {
		return
	}
	if !W.CheckErr2(p.SrcConn.WriteToUDP(p.Bytes()[:p.BufSize], p.SrcAddr)) {
		return
	}
}
