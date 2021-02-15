package errcheck

import (
	"github.com/DiskGetor/liblinux/log"
	"html/template"
)

var W = new(log.ObjLog)

type (
	InterfaceErrCheck interface {
		CheckErr(err interface{}) bool
		CheckBool2(retCtx interface{}, ok bool) bool
		CheckErr2(retCtx interface{}, err error) (ok bool)
	}
	ObjErrCheck struct{}
)

func (ObjErrCheck) CheckErr(err interface{}) bool {
	if err == nil {
		return true
	}
	W.LogError(err)
	return false
}

func (ObjErrCheck) CheckBool2(retCtx interface{}, ok bool) bool {
	if !ok {
		return false
	}
	return true
}

func (ObjErrCheck) CheckErr2(retCtx interface{}, err error) (ok bool) {
	if err == nil {
		switch retCtx.(type) { //这里可以封装为一个可变参数的函数，然后range一波
		case string:
			if retCtx == "" {
				W.LogError("nil string")
				return
			}
			if retCtx == "undefined" {
				W.LogError("JsRun return undefined")
				return
			}
			if retCtx == "{}" {
				W.LogError("json Structure member names must be uppercase")
				return
			}
		case int: //加入其它的整形？再观察下是否需要断言或者更多的判断？
			if retCtx == 0 { //windows api似乎是返回0就是错误
				W.LogError("Write 0 bytes to file")
				return
			}
		case []byte: //这些应该堆栈回溯判断下函数名？
			if retCtx.([]byte) == nil {
				W.LogError("The network request did not return content")
				return //个别爬虫网站确实需要忽略某个包，那么这里要返回true,也可以修改业务代码配合这个错误检查逻辑使之通用
			}
		case *template.Template:
			if retCtx.(*template.Template) == nil {
				W.LogError("The html template file returns a null pointer, please check the content of the html file")
				return
			}
		default:
			//检查别的类型
		}
		return true
	}
	W.LogError(err.Error()) //干掉层层返回，一出错就能卡到代码了
	return
}
