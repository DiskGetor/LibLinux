package net

func (p *ObjNet) CheckProtocol(Protocol, IpPort string) (ok bool) {
	if IpPort == "" {
		W.CheckErr(`bad IpPort ` + IpPort)
		return
	}
	if Protocol == "" {
		W.CheckErr(`bad Protocol ` + Protocol)
		return
	}
	if !p.IsKnownProtocol(Protocol) {
		W.CheckErr(`unknown Protocol ` + Protocol + ` called CheckProtocol`)
		return
	}
	return true
}

func (*ObjNet) IsKnownProtocol(Protocol string) bool {
	return Protocol == Socks5 || Protocol == Socks4 || Protocol == Http || Protocol == Https
}
