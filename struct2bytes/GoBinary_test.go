package struct2bytes

type User struct {
	Name string
	Age  int
	Buf  []byte
}

type Out struct {
	Age  int
	Name string
	Buf  []byte
}

//
//func TestInterfaceGob(t *testing.T) {
//	var p InterfaceGob = new(ObjGob)
//	u := New()
//	assert.True(t, p.GoBinaryEncode(u))
//	W.LogHexDump("GoBinaryEncode", p.GoBinaryBytes())
//	var out Out
//	assert.Equal(t, p.GoBinaryDecode(p.GoBinaryBytes(), &out), u)
//	W.LogStruct(out)
//}
//
//func New() *User {
//	return &User{
//		Name: "xxxxx",
//		Age:  99,
//		Buf: []byte{
//			0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88,
//			0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88,
//			0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99,
//		},
//	}
//}
