package UDP

func (ObjTransportUDP) TransportUDP(DstIP string, DstPort int) {
	p := New(DstIP, DstPort)
	if !p.GetSrcAddrConn() {
		return
	}
	defer func() {
		W.CheckErr(p.SrcConn.Close())
		W.CheckErr(p.DstConn.Close())
	}()
	for {
		SrcBufChan <- p.Bytes()[:p.BufSize]
		if !p.SetDstAddrConn() {
			return
		}
		go p.readDstBuf()
		if !W.CheckErr2(p.SrcConn.WriteToUDP(<-DstBufChan, p.SrcAddr)) { //这句提到协程内即可不用信道
			return
		}
	}
}

func (p *objUDP) readDstBuf() { //读目标buf
	select {
	case b := <-SrcBufChan:
		if !W.CheckErr2(p.DstConn.Write(b)) {
			return
		}
		p.Reset()
		p.BufSize, p.err = p.DstConn.Read(p.Bytes())
		if !W.CheckErr(p.err) {
			return
		}
		DstBufChan <- p.Bytes()[:p.BufSize]
	}
}
