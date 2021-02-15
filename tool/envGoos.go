package tool

import (
	"os"
	"runtime"
)

func (p *ObjTool) GetEnv(key string) string {
	ret := os.Getenv(key)
	if ret != "" {
		return ret
	}
	return "<nil>"
}

func (p *ObjTool) GetGOOS() string {
	return runtime.GOOS
	//goos := runtime.GOOS
	//for _, v := range GOOS {
	//	if strings.Compare(v, goos) == 0 {
	//		return v
	//	}
	//}
	//return ""
}

var GOOS = map[int]string{
	1:  `aix`,
	2:  `android`,
	3:  `darwin`,
	4:  `dragonfly`,
	5:  `freebsd`,
	6:  `hurd`,
	7:  `JS`,
	8:  `linux`,
	9:  `nacl`,
	10: `netbsd`,
	11: `openbsd`,
	12: `plan9`,
	13: `solaris`,
	14: `windows`,
	15: `zos`,
}
