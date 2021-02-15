package tool

import (
	"encoding/json"
	"fmt"
	"os"
)

func (p *ObjTool) GetFileInfo(filePath string) bool {
	p.fileInfo, p.err = os.Stat(filePath)
	if !(W.CheckErr(p.err)) {
		return false
	}
	return true
}

//const (
//	// 单字符是被String方法用于格式化的属性缩写。
//	ModeDir        FileMode = 1 << (32 - 1 - iota) // d: 目录
//	ModeAppend                                     // a: 只能写入，且只能写入到末尾
//	ModeExclusive                                  // l: 用于执行
//	ModeTemporary                                  // T: 临时文件（非备份文件）
//	ModeSymlink                                    // L: 符号链接（不是快捷方式文件）
//	ModeDevice                                     // D: 设备
//	ModeNamedPipe                                  // p: 命名管道（FIFO）
//	ModeSocket                                     // S: Unix域socket
//	ModeSetuid                                     // u: 表示文件具有其创建者用户id权限
//	ModeSetgid                                     // g: 表示文件具有其创建者组id的权限
//	ModeCharDevice                                 // c: 字符设备，需已设置ModeDevice
//	ModeSticky                                     // t: 只有root/创建者能删除/移动文件
//	// 覆盖所有类型位（用于通过&获取类型位），对普通文件，所有这些位都不应被设置
//	ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
//	ModePerm FileMode = 0777 // 覆盖所有Unix权限位（用于通过&获取类型位）
//)
//打开的权限是可读可写，权限是644
func (p *ObjTool) WriteAppend(filePath, buf string) (int, error) {
	if f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		W.LogError("maybe dir is not created", err.Error())
		new(ObjPath).CreatDirectory(filePath)
		return p.WriteAppend(filePath, buf)
	} else {
		//defer W.CheckErr(f.Close())
		return f.WriteString(buf + "\r\n")
	}
}

func (p *ObjTool) WriteGoFine(FileName string, buf []byte) (ok bool) {
	f, err := os.Create(FileName)
	if !W.CheckErr(err) {
		return
	}
	W.CheckErr2(f.WriteString(fmt.Sprintf("%s\n\n", `package main`)))
	W.CheckErr2(f.WriteString(fmt.Sprintf("%s\n", `	var buf = []byte{`)))
	for _, v := range buf {
		W.CheckErr2(f.WriteString(fmt.Sprintf("0x%02X,", v)))
	}
	if !W.CheckErr2(f.WriteString(fmt.Sprintf("\n%s", `	}`))) {
		return
	}
	if !p.FormatGoFile(FileName) {
		return
	}
	return true
}

func (p *ObjTool) WriteFile(FileName string, obj interface{}) (ok bool) {
	p.file, p.err = os.Create(FileName)
	if !(W.CheckErr(p.err)) {
		return false
	}
	switch obj.(type) {
	case []byte:
		if !W.CheckErr2(p.file.Write(obj.([]byte))) {
			return
		}
	case string:
		if !W.CheckErr2(p.file.WriteString(obj.(string))) {
			return
		}
	}
	return true
}

func (p *ObjTool) JsonMarshalIndent2File(FileName string, obj interface{}) (ok bool) {
	p.file, p.err = os.OpenFile(FileName, os.O_CREATE|os.O_APPEND, 0644)
	if p.err != nil {
		if !W.CheckErr(os.Truncate(FileName, 0)) {
			if !W.CheckErr(p.file.Close()) {
				return
			}
			return
		}
	}
	//if W.CheckErr2(ioutil.ReadFile(FileName)) {//别的进程正在使用，改一下逻辑;清空
	//	if !W.CheckErr(os.Remove(FileName)) {//| remove env.json: The process cannot access the file because it is being used by another process.
	//		return
	//	}
	//}
	p.out, p.err = json.MarshalIndent(obj, " ", " ")
	if !W.CheckErr(p.err) {
		return
	}
	if !p.WriteFile(FileName, string(p.out)) {
		return
	}
	ok = true
	return
}
