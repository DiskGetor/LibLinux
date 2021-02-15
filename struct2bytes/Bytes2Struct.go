package struct2bytes

import (
	"bytes"
	"encoding/binary"
	"reflect"
)

func (p *ObjStructBytes) Read(StructBytes []byte, obj interface{}) bool {
	p.Buffer = bytes.NewBuffer(StructBytes)
	return p.ReadValue(reflect.ValueOf(obj), 0)
}

func (p *ObjStructBytes) ReadValue(v reflect.Value, depth int) (ok bool) {
	switch v.Kind() {
	case reflect.Interface:
		if v.IsNil() {
			v.Set(reflect.ValueOf(v.Type()))
		}
		fallthrough
	case reflect.Ptr:
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		if !p.ReadValue(v.Elem(), depth+1) {
			return
		}

	case reflect.Struct:
		l := v.NumField()
		for i := 0; i < l; i++ {
			if !p.ReadValue(v.Field(i), depth+1) {
				return
			}
		}

	case reflect.Slice:
		if v.IsNil() {
			W.CheckErr("切片必须在解码之前初始化为正确的长度")
			return
		}
		fallthrough
	case reflect.Array:
		l := v.Len()
		for i := 0; i < l; i++ {
			if !p.ReadValue(v.Index(i), depth+1) {
				return
			}
		}

	case reflect.Int:
		var i int32
		if !W.CheckErr(binary.Read(p.Buffer, binary.BigEndian, &i)) {
			return
		}
		v.SetInt(int64(i))

	case reflect.Bool:
		boolValue := uint8(0)
		if !W.CheckErr(binary.Read(p.Buffer, binary.BigEndian, &boolValue)) {
			return
		}
		v.SetBool(boolValue != 0)

	default:
		if !W.CheckErr(binary.Read(p.Buffer, binary.BigEndian, v.Addr().Interface())) {
			return
		}
	}
	return true
}
