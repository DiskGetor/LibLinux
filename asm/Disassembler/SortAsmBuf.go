package Disassembler

type (
	InterfaceSortAsmBuf interface {
		DisassemblerFormatLine(src []byte) string
	}
	ObjSortAsmBuf struct {
		Title string
		Buf   string
	}
)

const (
	hexTable  = "0123456789abcdef"
	IndentLen = 56
)

func encodedLen(n int) int { return n * 3 }

func (p *ObjSortAsmBuf) DisassemblerFormatLine(src []byte) string { //可能是反汇编显示需要
	dst := make([]byte, encodedLen(len(src)))
	Encode(dst, src)
	return string(dst)
}

func Encode(dst, src []byte) int {
	j := 0
	for _, v := range src {
		dst[j] = hexTable[v>>4]
		dst[j+1] = hexTable[v&0x0f]
		dst[j+2] = ' '
		j += 3
	}
	return len(src) * 3
}

func (p *ObjSortAsmBuf) indentTitle() {
	if len(p.Title) < IndentLen {
		for {
			if len(p.Title) == IndentLen {
				p.Title += " | "
				return
			}
			p.Title += " "
		}
	}
}
