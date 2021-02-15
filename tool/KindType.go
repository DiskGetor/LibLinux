package tool

import "reflect"

func (p *ObjTool) Kind2TypeStr(kind reflect.Kind) (Type string) {
	switch kind {
	case reflect.Bool:
		Type = `bool`
	case reflect.Int:
		Type = `int`
	case reflect.Int8:
		Type = `int8`
	case reflect.Int16:
		Type = `int16`
	case reflect.Int32:
		Type = `int32`
	case reflect.Int64:
		Type = `int64`
	case reflect.Uint:
		Type = `uint`
	case reflect.Uint8:
		Type = `uint8`
	case reflect.Uint16:
		Type = `uint16`
	case reflect.Uint32:
		Type = `uint32`
	case reflect.Uint64:
		Type = `uint64`
	case reflect.Uintptr:
		Type = `uintptr`
	case reflect.Float32:
		Type = `float32`
	case reflect.Float64:
		Type = `float64`
	case reflect.Complex64:
		Type = `complex64`
	case reflect.Complex128:
		Type = `complex128`
	case reflect.Array:
		Type = `array`
	case reflect.Map:
		Type = `Map`
	case reflect.Ptr:
		Type = `ptr`
	case reflect.Slice:
		Type = `slice`
	case reflect.String:
		Type = `string`
	case reflect.Struct:
		Type = `struct`
	case reflect.UnsafePointer:
		Type = `unsafepointer`
	default:
		W.CheckErr(`bad Kind`)
		return
	}
	return
}

func (p *ObjTool) TypeStr2Kind(Type string) (Kind reflect.Kind) {
	switch Type {
	case `bool`:
		Kind = reflect.Bool
	case `int`:
		Kind = reflect.Int
	case `int8`:
		Kind = reflect.Int8
	case `int16`:
		Kind = reflect.Int16
	case `int32`:
		Kind = reflect.Int32
	case `int64`:
		Kind = reflect.Int64
	case `uint`:
		Kind = reflect.Uint
	case `uint8`:
		Kind = reflect.Uint8
	case `uint16`:
		Kind = reflect.Uint16
	case `uint32`:
		Kind = reflect.Uint32
	case `uint64`:
		Kind = reflect.Uint64
	case `uintptr`:
		Kind = reflect.Uintptr
	case `float32`:
		Kind = reflect.Float32
	case `float64`:
		Kind = reflect.Float64
	case `complex64`:
		Kind = reflect.Complex64
	case `complex128`:
		Kind = reflect.Complex128
	case `array`:
		Kind = reflect.Array
	case `Map`:
		Kind = reflect.Map
	case `ptr`:
		Kind = reflect.Ptr
	case `slice`:
		Kind = reflect.Slice
	case `string`:
		Kind = reflect.String
	case `struct`:
		Kind = reflect.Struct
	case `unsafepointer`:
		Kind = reflect.UnsafePointer
	default:
		Kind = reflect.Struct
		//panic(`bad Type`)
	}
	return
}
