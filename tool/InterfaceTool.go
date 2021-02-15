package tool

import (
	"compress/gzip"
	"github.com/DiskGetor/liblinux/errcheck"
	"github.com/DiskGetor/liblinux/log"
	"os"
	"reflect"
)

var (
	W = new(struct {
		log.ObjLog
		errcheck.ObjErrCheck
	})
)

func (p *ObjTool) BytesTools() []byte {
	return p.out
}
func (p *ObjTool) StringTools() string {
	return string(p.out)
}
func (p *ObjTool) FileInfoTools() os.FileInfo {
	return p.fileInfo
}

type (
	InterfaceTool interface {
		InterfaceArrayToSlice
		InterfaceVerSion
		Float64ValueTool() float64                              //Float64Cut 返回值
		FileInfoTools() os.FileInfo                             //FileInfoTools 返回值
		BytesTools() []byte                                     //tool接口的全部实现返回值类型是[]byte的接收函数
		StringTools() string                                    //tool接口的全部实现返回值类型是string的接收函数
		UTF82GBK(src string) bool                               //UTF8转GBK
		GBK2UTF8(src []byte) bool                               //GBK转UTF8
		CmdBuf2ChineseString(arg interface{}) (str string)      //CmdBuf转ChineseString
		CmdRunWithCheck(arg string)                             //输出cmd执行结果
		JoinBytes(Bytes ...[]byte) []byte                       //字节切片合并
		CompareBufFromFile(buf []byte, fileName string)         //从文件读取buf比较
		RemoveRepeatedElement(arr []string) (newArr []string)   //数组去重1
		RemoveRepeatedElementV2(arr []string) (newArr []string) //数组去重2
		InterfacePath
		GetEnv(key string) string                                          //获取环境变量
		GetGOOS() string                                                   //获取当前系统类型
		GetFileInfo(filePath string) bool                                  //获取文件信息
		WriteAppend(filePath, buf string) (int, error)                     //循环写入buf到文件
		WriteGoFine(FileName string, buf []byte) (ok bool)                 //写go文件
		WriteFile(FileName string, obj interface{}) (ok bool)              //写文件
		JsonMarshalIndent2File(FileName string, obj interface{}) (ok bool) //序列化buf为json格式并写文件
		Float64ToString(f float64, cut int) string                         //Float64转String
		Float64Cut(value float64, bits int) bool                           //Float64截取，bits为保留几位小数
		FormatGoFile(FileName string) bool                                 //格式化go文件
		GzipDecode(in []byte) (ok bool)                                    //解码 Gzip buf
		Kind2TypeStr(kind reflect.Kind) (Type string)                      //遍历反射类型返回类型名称
		TypeStr2Kind(Type string) (Kind reflect.Kind)                      //遍历类型名称返回反射类型
		RandomNum(min, max int) int                                        //生成指定长度的随机数
		SwapHexStringToBytes(in string) bool                               //解码16进制字符串为buf并倒序
		SwapString(src []byte) (dst string)                                //倒序字符串
		SwapArray(src []byte) (dst []byte)                                 //倒序buf
		SwapUint8FromUint16(v uint16) uint8                                //从int16截取int8并倒序，用于数据恢复软件算法
		GetTimeStamp13Bits() int64                                         //获取指定长度的时间戳
		GetTimeStamp() string                                              //获取时间戳
	}
	ObjTool struct {
		out          []byte
		tmp          []byte
		err          error
		fileInfo     os.FileInfo
		file         *os.File
		float64Value float64
		gzipReader   *gzip.Reader
		ObjVerSion
		LibArrayToSlice
		ObjPath
	}
)

var _ InterfaceTool = (*ObjTool)(nil)
