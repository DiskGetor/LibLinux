package tool

import (
	"fmt"
	"strconv"
)

func (p *ObjTool) Float64ToString(f float64, cut int) string {
	return strconv.FormatFloat(f, 'f', cut, 64)
}

func (p *ObjTool) Float64Cut(value float64, bits int) bool {
	p.float64Value, p.err = strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(bits)+"f", value), 64)
	if !(W.CheckErr(p.err)) {
		return false
	}
	return true
}

func (p *ObjTool) Float64ValueTool() float64 {
	return p.float64Value
}
