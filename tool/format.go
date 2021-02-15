package tool

import "os/exec"

func (p *ObjTool) FormatGoFile(FileName string) bool {
	//outBytes, err := format.Source(buf.Bytes())
	if !W.CheckErr(exec.Command("gofmt", "-w", FileName).Run()) {
		return false
	}
	return true
}
