package tool

import (
	"fmt"
	"strconv"
	"strings"
)

type (
	InterfaceVerSion interface {
		VerSionToString(p *ObjVerSion) string
		StringToVerSion(version string) (ok bool)
		Major() uint64
		SetMajor(major uint64)
		Minor() uint64
		SetMinor(minor uint64)
		Patch() uint64
		SetPatch(patch uint64)
	}
	ObjVerSion struct {
		major uint64
		minor uint64
		patch uint64
		err   error
	}
)

func (*ObjVerSion) SetPatch(patch uint64) {
	ctxV.patch = patch
}

func (*ObjVerSion) SetMinor(minor uint64) {
	ctxV.minor = minor
}

func (s *ObjVerSion) SetMajor(major uint64) {
	ctxV.major = major
}

var (
	ctxV                  = new(ObjVerSion)
	_    InterfaceVerSion = (*ObjVerSion)(nil)
)

func (*ObjVerSion) Major() uint64 {
	return ctxV.major
}

func (*ObjVerSion) Minor() uint64 {
	return ctxV.minor
}

func (*ObjVerSion) Patch() uint64 {
	return ctxV.patch
}

func (*ObjVerSion) VerSionToString(p *ObjVerSion) string {
	ctxV = p
	var arr []string
	arr = append(arr, fmt.Sprint(p.major), fmt.Sprint(p.minor), fmt.Sprint(p.patch))
	return strings.Join(arr, `.`)
}

func (*ObjVerSion) StringToVerSion(version string) (ok bool) {
	v := strings.Split(version, `.`)
	ctxV.major, ctxV.err = strconv.ParseUint(v[0], 10, 32)
	if !W.CheckErr(ctxV.err) {
		return
	}
	ctxV.minor, ctxV.err = strconv.ParseUint(v[1], 10, 32)
	if !W.CheckErr(ctxV.err) {
		return
	}
	ctxV.patch, ctxV.err = strconv.ParseUint(v[2], 10, 32)
	if !W.CheckErr(ctxV.err) {
		return
	}
	return true
}
