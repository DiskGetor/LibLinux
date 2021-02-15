package tool

import (
	"reflect"
	"unsafe"
)

type (
	InterfaceArrayToSlice interface {
		ArrayToSlice(ArrayPtr unsafe.Pointer, Len int) []byte
	}
	LibArrayToSlice struct {
	}
)

func (p *LibArrayToSlice) ArrayToSlice(ArrayPtr unsafe.Pointer, Len int) (Slice []byte) {
	SliceHeader := reflect.SliceHeader{
		Data: uintptr(ArrayPtr),
		Len:  Len,
		//Cap:  Len,
	}
	//核心思想：用反射构造一个切片指针，然后传入数组指针append到反射构造的切片指针，append之前填充一下反射构造的切片Len,Cap自动根据len管理，不要操作
	Slice = make([]byte, 0)                                            //不初始化不行，Cap自动填充len
	Slice = append(Slice, *(*[]byte)(unsafe.Pointer(&SliceHeader))...) //这个是重点,unsafe.Pointer 是*int类型
	return
}
