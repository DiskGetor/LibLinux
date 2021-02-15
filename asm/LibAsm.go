package asm

import "github.com/DiskGetor/liblinux/asm/Disassembler"

type (
	ObjLibAsm struct {
		Disassembler.ObjSortAsmBuf
	}
	InterfaceLibAsm interface {
		Disassembler.InterfaceSortAsmBuf
	}
)
