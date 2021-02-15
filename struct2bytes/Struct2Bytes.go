package struct2bytes

import (
	"bytes"
	"encoding/binary"
	"reflect"
)

func (p *ObjStructBytes) Write(obj interface{}) bool {
	p.Buffer = &bytes.Buffer{}
	return p.WriteValue(obj, 0)
}

func (p *ObjStructBytes) WriteValue(obj interface{}, depth int) (ok bool) {
	v := reflect.ValueOf(obj)
	switch v.Kind() {
	case reflect.Interface:
	case reflect.Ptr:
		if !p.WriteValue(v.Elem().Interface(), depth+1) {
			return
		}
	case reflect.Struct:
		l := v.NumField()
		for i := 0; i < l; i++ {
			if !p.WriteValue(v.Field(i).Interface(), depth+1) {
				return
			}
		}

	case reflect.Slice, reflect.Array:
		l := v.Len()
		for i := 0; i < l; i++ {
			if !p.WriteValue(v.Index(i).Interface(), depth+1) {
				return
			}
		}

	case reflect.Int:
		i := int32(obj.(int))
		if int(i) != obj.(int) {
			W.CheckErr("Int does not fit into int32")
			return
		}
		if !W.CheckErr(binary.Write(p.Buffer, binary.BigEndian, i)) {
			return
		}

	case reflect.Bool:
		b := uint8(0)
		if v.Bool() {
			b = 1
		}
		if !W.CheckErr(binary.Write(p.Buffer, binary.BigEndian, b)) {
			return
		}

	default:
		if !W.CheckErr(binary.Write(p.Buffer, binary.BigEndian, obj)) {
			return
		}

	}
	return true
}
