package tsf

import (
	"syscall"
	"unsafe"
)

// TFContext represents the ITfContext COM interface.
type TFContext struct {
	Vtbl *itfContextVtbl
}

type itfContextVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	CreateRangeBackup  uintptr
	EnumProperties     uintptr
	EnumViews          uintptr
	GetActiveView      uintptr
	GetAppProperty     uintptr
	GetDocumentMgr     uintptr
	GetEnd             uintptr
	GetProperty        uintptr
	GetSelection       uintptr
	GetStart           uintptr
	GetStatus          uintptr
	InWriteSession     uintptr
	RequestEditSession uintptr
	SetSelection       uintptr
	TrackProperties    uintptr
}

func (obj *TFContext) AddRef() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.Vtbl.AddRef,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

func (obj *TFContext) Release() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.Vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}
