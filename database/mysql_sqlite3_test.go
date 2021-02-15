package database

import "testing"

type (
	InterfaceEA interface {
		GetValue() (ok bool)
		SetValue(value string)
	}
	ObjEA struct {
		maddenMySqldb    InterfaceDataBase
		maddenSqlLite3db InterfaceDataBase

		value string
		err   error
	}
)

func (p *ObjEA) GetValue() (ok bool) {
	panic("implement me")
	return true
}

func (p *ObjEA) SetValue(value string) {
	p.value = value
}

func (p *ObjEA) Value() string {
	return p.value
}

var (
	_ InterfaceEA = (*ObjEA)(nil)
)

func TestInterfaceDataBase(t *testing.T) {

}
