package log

import (
	"fyne.io/fyne/v2"
	"os"
	"runtime/debug"
)

type (
	InterfaceLog interface {
		LogHexDump(title string, msg interface{})      //hex buf
		LogHex(title string, msg interface{})          //hex value
		LogInfo(title string, msg ...interface{})      //info
		LogTrace(title string, msg ...interface{})     //跟踪
		LogWarning(title string, msg ...interface{})   //警告
		LogPbMessage(title string, msg ...interface{}) //pb解码json文本
		LogSuccess(title string, msg ...interface{})   //成功
		LogStruct(msg ...interface{})                  //结构体
		LogError(msg ...interface{})                   //错误，堆栈
		IsStopPrintWriteLogFile(is bool)               //安卓需求
		AndroidLogFilePath() string
		AndroidJsonFilePath() string
		SetAndroidLogFilePath(a fyne.App, fileName string)
		SetAndroidJsonFilePath(a fyne.App, fileName string)
	}
	ObjLog struct {
		isStopPrintWriteLogFile bool
		androidLogFilePath      string
		androidJsonFilePath     string
		ctx
	}
	ctx struct {
		tag       string
		title     string
		color     uint8
		titleInfo string
		msg       string
		info      string
	}
)

//
//{BlackString, "%s", nil, "\x1b[30m%s\x1b[0m"},
//{RedString, "%s", nil, "\x1b[31m%s\x1b[0m"},
//{GreenString, "%s", nil, "\x1b[32m%s\x1b[0m"},
//{YellowString, "%s", nil, "\x1b[33m%s\x1b[0m"},
//{BlueString, "%s", nil, "\x1b[34m%s\x1b[0m"},
//{MagentaString, "%s", nil, "\x1b[35m%s\x1b[0m"},
//{CyanString, "%s", nil, "\x1b[36m%s\x1b[0m"},
//{WhiteString, "%s", nil, "\x1b[37m%s\x1b[0m"},
//{HiBlackString, "%s", nil, "\x1b[90m%s\x1b[0m"},
//{HiRedString, "%s", nil, "\x1b[91m%s\x1b[0m"},
//{HiGreenString, "%s", nil, "\x1b[92m%s\x1b[0m"},
//{HiYellowString, "%s", nil, "\x1b[93m%s\x1b[0m"},
//{HiBlueString, "%s", nil, "\x1b[94m%s\x1b[0m"},
//{HiMagentaString, "%s", nil, "\x1b[95m%s\x1b[0m"},
//{HiCyanString, "%s", nil, "\x1b[96m%s\x1b[0m"},
//{HiWhiteString, "%s", nil, "\x1b[97m%s\x1b[0m"},

const (
	colorRed = uint8(iota + 31)
	colorGreen
	colorYellow
	colorBlue
	colorMagenta
)

const (
	tagHex       = `[HEXV]`
	tagHexDump   = `[DUMP]`
	tagPbMessage = `[PMSG]`
	tagStruct    = `[STRC]`
	tagInfo      = `[INFO]`
	tagTrace     = `[TRAC]`
	tagError     = `[ERRO]`
	tagWarning   = `[WARN]`
	tagSuccess   = `[SUCC]`
)

const (
	colorFormat = "\x1b[1m\x1b[%dm%s\x1b[0m\n"
)

///usr/lib/go/src/os/example_test.go
func (p *ObjLog) writeAppend(filePath, buf string) (int, error) {
	if f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		p.LogError("maybe dir is not created", err.Error())
		p.creatDirectory(filePath)
		return p.writeAppend(filePath, buf)
	} else {
		//defer W.CheckErr(f.Close())
		return f.WriteString(buf + "\r\n")
	}
}

func (p *ObjLog) creatDirectory(path string) bool {
	return p.checkErr(os.MkdirAll(path, os.ModeAppend))
}

func (p *ObjLog) checkErr2(arg interface{}, err error) bool {
	return p.checkErr(err)
}

func (p *ObjLog) checkErr(err error) bool {
	if err != nil {
		println(err.Error())
		debug.PrintStack()
		return false
	}
	return true
}
