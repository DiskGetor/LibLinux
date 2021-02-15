package tool

import (
	"os"
	"path/filepath"
	"strings"
)

type (
	InterfacePath interface {
		CreatDirectory(directoryName string) bool  //创建目录
		CurrentDirectory() string                  //返回当前目录
		GetCurrentDirectory() (ok bool)            //获取当前目录
		GetCurrentRunningPath() (ok bool)          //获取当前运行路径
		CurrentRunningPath() string                //返回当前运行路径
		GetFilepath(fileName string) (ok bool)     //获取文件路径
		Filepath() string                          //返回文件路径
		GetFileDirectory(Filepath string) string   //返回文件目录
		GetFilename(Filepath string) string        //返回文件名
		GetExtensionPath(Filepath string) string   //返回文件扩展名
		GetAbsolutePath(Filepath string) (ok bool) //获取绝对路径
		AbsolutePath() string                      //返回绝对路径
		GetUNCPath(fileName string) (ok bool)      //获取UNC路径
		UNCPath() string                           //返回UNC路径
	}
	ObjPath struct {
		currentDirectory   string
		absolutePath       string
		currentRunningPath string
		filepath           string
		uncPath            string
		err                error
	}
)

var (
	_ InterfacePath = (*ObjPath)(nil)
)

func (p *ObjPath) UNCPath() string {
	return p.uncPath
}

func (p *ObjPath) GetUNCPath(fileName string) (ok bool) {
	if !p.GetCurrentRunningPath() {
		return
	}
	if !p.GetFilepath(fileName) {
		return
	}
	p.uncPath = strings.Replace(p.filepath, `\`, `\\`, -1)
	return true
}

func (p *ObjPath) CreatDirectory(directoryName string) bool {
	return W.CheckErr(os.MkdirAll(directoryName, os.ModePerm))
}

func (p *ObjPath) CurrentDirectory() string {
	return p.currentDirectory
}

func (p *ObjPath) AbsolutePath() string {
	return p.absolutePath
}

func (p *ObjPath) CurrentRunningPath() string {
	return p.currentRunningPath
}

//----------------------------------------------------

func (p *ObjPath) GetCurrentRunningPath() (ok bool) {
	dir, err := os.Executable()
	if !W.CheckErr2(dir, err) {
		return
	}
	p.currentRunningPath = filepath.Dir(dir)
	return true
}

func (p *ObjPath) GetCurrentDirectory() (ok bool) {
	p.currentDirectory, p.err = os.Getwd()
	return W.CheckErr(p.err)
}

func (p *ObjPath) Filepath() string {
	return p.filepath
}
func (p *ObjPath) GetFilepath(fileName string) (ok bool) {
	if !p.GetCurrentDirectory() {
		return
	}
	p.filepath = filepath.Join(p.currentDirectory, fileName)
	return true
}
func (p *ObjPath) GetFileDirectory(Filepath string) string {
	return filepath.Dir(Filepath)
}
func (p *ObjPath) GetFilename(Filepath string) string {
	return filepath.Base(Filepath)
}
func (p *ObjPath) GetAbsolutePath(Filepath string) (ok bool) {
	p.absolutePath, p.err = filepath.Abs(Filepath)
	return W.CheckErr(p.err)
}
func (p *ObjPath) GetExtensionPath(Filepath string) string {
	return filepath.Ext(Filepath)
}
