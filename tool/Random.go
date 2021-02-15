package tool

import (
	"math/rand"
	"time"
)

func (p *ObjTool) RandomNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
