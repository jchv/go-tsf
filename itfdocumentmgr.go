package tsf

import (
	"syscall"
	"unsafe"
)

// TFDocumentMgr represents the ITfDocumentMgr COM interface.
type TFDocumentMgr struct {
	Vtbl *itfDocumentMgrVtbl
}

type itfDocumentMgrVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	CreateContext uintptr
	EnumContexts  uintptr
	GetBase       uintptr
	GetTop        uintptr
	Pop           uintptr
	Push          uintptr
}

func (obj *TFDocumentMgr) AddRef() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.Vtbl.AddRef,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

func (obj *TFDocumentMgr) Release() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.Vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}
