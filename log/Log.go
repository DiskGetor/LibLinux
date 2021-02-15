package log

import (
	"encoding/hex"
	"fmt"
	"fyne.io/fyne/v2"
	"golang.org/x/text/width"
	"log"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

func (p *ObjLog) LogHexDump(title string, msg interface{}) {
	p.ctx = ctx{
		tag:   tagHexDump,
		title: title,
		color: 93, //colorYellow
		msg:   fmt.Sprintf("%s", hex.Dump(msg.([]byte))),
	}
	p.do()
}

func (p *ObjLog) LogHex(title string, msg interface{}) {
	p.ctx = ctx{
		tag:   tagHex,
		title: title,
		color: 36, //CyanString
		msg:   fmt.Sprintf("%#v", msg),
	}
	p.do()
}

func (p *ObjLog) LogInfo(title string, msg ...interface{}) {
	p.ctx = ctx{
		tag:   tagInfo,
		title: title,
		color: 96, //colorBlue
		msg:   fmt.Sprint(msg...),
	}
	p.do()
}

func (p *ObjLog) LogTrace(title string, msg ...interface{}) {
	p.ctx = ctx{
		tag:   tagTrace,
		title: title,
		color: 94, //HiBlue
		msg:   fmt.Sprint(msg...),
	}
	p.do()
}

func (p *ObjLog) LogWarning(title string, msg ...interface{}) {
	p.ctx = ctx{
		tag:   tagWarning,
		title: title,
		color: colorMagenta,
		msg:   fmt.Sprint(msg...),
	}
	p.do()
}

func (p *ObjLog) LogPbMessage(title string, msg ...interface{}) {
	p.ctx = ctx{
		tag:   tagPbMessage,
		title: title,
		color: colorGreen,
		msg:   fmt.Sprint(msg...),
	}
	p.do()
}

func (p *ObjLog) LogSuccess(title string, msg ...interface{}) {
	p.ctx = ctx{
		tag:   tagSuccess,
		title: title,
		color: colorGreen,
		msg:   fmt.Sprint(msg...),
	}
	p.do()
}

func (p *ObjLog) LogStruct(msg ...interface{}) {
	p.ctx = ctx{
		tag:   tagStruct,
		title: "",
		color: 92, //HiGreen
		msg:   fmt.Sprintf("%#v", msg...),
	}
	p.do()
}

func (p *ObjLog) LogError(msg ...interface{}) {
	p.ctx = ctx{
		tag:   tagError,
		title: "",
		color: colorRed,
		msg:   fmt.Sprint(msg...),
	}
	p.do()
}

func (p *ObjLog) IsStopPrintWriteLogFile(is bool) {
	p.isStopPrintWriteLogFile = is
}

func (p *ObjLog) AndroidLogFilePath() string {
	return p.androidLogFilePath
}
func (p *ObjLog) AndroidJsonFilePath() string {
	return p.androidJsonFilePath
}
func (p *ObjLog) SetAndroidLogFilePath(a fyne.App, fileName string) {
	p.androidLogFilePath = storagePath(a, fileName)
}
func (p *ObjLog) SetAndroidJsonFilePath(a fyne.App, fileName string) {
	p.androidJsonFilePath = storagePath(a, fileName)
}

func (p *ObjLog) do() {
	p.setInfo()
	isWriteLogFile := false
	isPrint := false
	if !p.isStopPrintWriteLogFile { //windows默认开启打印和写日志文件
		isPrint = true
		isWriteLogFile = true
	}
	if !isPrint {
		return
	}
	//if runtime.gOOS == `android` {
	//	return
	//}
	if isWriteLogFile { //安卓需要
		p.writeLogFile()
	}
	log.Print(p.info)
	if p.tag == tagError {
		debug.PrintStack()
	}
}

func getGOOS() string {
	goos := runtime.GOOS
	for _, v := range gOOS {
		if strings.Compare(v, goos) == 0 {
			return v
		}
	}
	return ""
}

var gOOS = map[int]string{
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

func (p *ObjLog) writeLogFile() {
	fileName := ""
	switch runtime.GOOS {
	case "windows", `linux`:
		fileName = "log.txt"
	case "android":
		return
		if p.androidLogFilePath == "" {
			p.LogError("androidLogFilePath is nil")
			return
		}
		fileName = p.androidLogFilePath
	default:
		println("log存储路径错误")
		debug.PrintStack()
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05 ")
	noColorInfo := fmt.Sprintf(p.titleInfo + p.msg)
	if !p.checkErr2(p.writeAppend(fileName, now+noColorInfo)) {
		return
	}
	if p.tag == tagError {
		if !p.checkErr2(p.writeAppend(fileName, string(debug.Stack()))) {
			return
		}
	}
}

func (p *ObjLog) setInfo() {
	p.indentTitle()
	if p.title == "" {
		p.titleInfo = p.tag
	} else {
		p.titleInfo = p.tag + p.title
	}
	switch p.tag {
	case tagPbMessage, tagHexDump:
		p.titleInfo += "\n"
	}
	p.info = fmt.Sprintf(colorFormat, p.color, p.titleInfo+p.msg)
}

func (p *ObjLog) indentTitle() {
	//https://blog.csdn.net/raoxiaoya/article/details/108982887
	const hexDumpIndentLen = 28
	Separate := ` | `
	spaceLen := hexDumpIndentLen - Width(p.title)
	alignStr := ``
	if spaceLen > 0 {
		alignStr = strings.Repeat(" ", spaceLen)
	}
	p.title = alignStr + p.title + Separate
	//title = fmt.Sprintf("%28s | ", title) //"%-28s 加个负号就是左对齐,英文状态下28的长度刚好与hexdump对齐，试了utf8没用的
}

func Width(s string) (w int) {
	for _, r := range []rune(s) {
		switch width.LookupRune(r).Kind() {
		case width.EastAsianWide, width.EastAsianFullwidth:
			w += 2
		default:
			w++
		}
	}
	return
}
