package tool

import (
	"bufio"
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
)

type charset string

const (
	UTF8    = charset("UTF-8")
	GB18030 = charset("GB18030")
)

func convertByte2String(byte []byte, charset charset) string {
	var str string
	switch charset {
	case GB18030:
		decodeBytes, err := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		if !W.CheckErr(err) {
			return "err GB18030.NewDecoder()"
		}
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

func checkCmdResult(arg string, cmd *exec.Cmd) {
	outReader, err := cmd.StdoutPipe()
	if !W.CheckErr(err) {
		return
	}
	errReader, err := cmd.StderrPipe()
	if !W.CheckErr(err) {
		return
	}
	cmdReader := io.MultiReader(outReader, errReader)

	if !W.CheckErr(cmd.Start()) {
		return
	}
	Stdin := bufio.NewScanner(cmdReader)
	for Stdin.Scan() {
		cmdRe := convertByte2String(Stdin.Bytes(), GB18030)
		cmdRe = strings.Replace(cmdRe, "\r\n", "", -1)
		W.LogInfo("cmd命令 "+arg+" 返回:", cmdRe)
	}
	if !W.CheckErr(cmd.Wait()) {
		return
	}
}

func (p *ObjTool) CmdRunWithCheck(arg string) {
	checkCmdResult(arg, exec.Command("C:\\Windows\\SysWOW64\\cmd.exe", "/C", arg))
}

func (p *ObjTool) CmdBuf2ChineseString(arg interface{}) (str string) {
	switch arg.(type) {
	case string:
		str = convertByte2String([]byte(arg.(string)), GB18030)
	case []byte:
		str = convertByte2String(arg.([]byte), GB18030)
	}
	return
}

// UTF82GBK : transform UTF8 rune into GBK byte array
func (p *ObjTool) UTF82GBK(src string) bool {
	GB18030 := simplifiedchinese.All[0]
	p.out, p.err = ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), GB18030.NewEncoder()))
	if !(W.CheckErr(p.err)) {
		return false
	}
	return true
}

// GBK2UTF8 : transform  GBK byte array into UTF8 string
func (p *ObjTool) GBK2UTF8(src []byte) bool {
	GB18030 := simplifiedchinese.All[0]
	p.out, p.err = ioutil.ReadAll(transform.NewReader(bytes.NewReader(src), GB18030.NewDecoder()))
	if !(W.CheckErr(p.err)) {
		return false
	}
	return true
}
