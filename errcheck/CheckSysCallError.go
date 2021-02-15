package errcheck

import (
	"fmt"
	"syscall"
)

const errorCodeIoPending = 997

var errorIoPending error = syscall.Errno(errorCodeIoPending)

func (p ObjErrCheck) CheckSysCallError(r1 uintptr, lastErr syscall.Errno) (ok bool) { //c语言临时变量不能返回指针的锅晕，cpp的class试了下是可以的
	switch r1 {
	case 0:
		switch lastErr {
		case 0:
			return true
		case errorCodeIoPending:
			p.CheckErr(`errorCodeIoPending:` + fmt.Sprint(errorIoPending))
			return
		default:
			p.CheckErr(` EINVAL:` + fmt.Sprint(syscall.EINVAL))
			return
		}
	default:
		return true
	}
}
