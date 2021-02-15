package tool

import (
	"strconv"
	"time"
)

func (p *ObjTool) GetTimeStamp13Bits() int64 {
	return time.Now().UnixNano() / 1000000
}

func (p *ObjTool) GetTimeStamp() string {
	return strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
}
