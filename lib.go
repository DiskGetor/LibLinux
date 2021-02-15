package liblinux

import (
	"github.com/DiskGetor/liblinux/ascii"
	"github.com/DiskGetor/liblinux/asm"
	"github.com/DiskGetor/liblinux/crypto"
	"github.com/DiskGetor/liblinux/database"
	"github.com/DiskGetor/liblinux/errcheck"
	"github.com/DiskGetor/liblinux/js"
	"github.com/DiskGetor/liblinux/log"
	"github.com/DiskGetor/liblinux/net"
	//"github.com/DiskGetor/liblinux/pb"
	"github.com/DiskGetor/liblinux/struct2bytes"
	"github.com/DiskGetor/liblinux/syncMap"
	"github.com/DiskGetor/liblinux/tag"
	"github.com/DiskGetor/liblinux/theme"
	"github.com/DiskGetor/liblinux/tool"
)

type (
	InterfaceLib interface {
		ascii.InterfaceAscii
		asm.InterfaceLibAsm
		crypto.InterfaceLibCrypto
		database.InterfaceDataBase
		errcheck.InterfaceErrCheck
		js.InterfaceJs
		log.InterfaceLog
		net.InterfaceNet
		//pb.InterfacePb2
		struct2bytes.InterfaceStructBytes
		syncMap.InterfaceSyncMap
		tag.InterfaceTag
		theme.InterfaceLibTheme
		tool.InterfaceTool
	}
	ObjLib struct {
		ascii.ObjAscii
		asm.ObjLibAsm
		crypto.ObjLibCrypto
		database.ObjDataBase
		errcheck.ObjErrCheck
		js.ObjJs
		log.ObjLog
		net.ObjNet
		//pb.ObjPb2
		struct2bytes.ObjStructBytes
		syncMap.ObjSyncMap
		tag.ObjTag
		theme.ObjLibTheme
		tool.ObjTool
	}
)

var W InterfaceLib = new(ObjLib)
