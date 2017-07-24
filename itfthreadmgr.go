package tsf

import (
	"unsafe"
)

// ITfThreadMgr represents the ITfThreadMgr COM interface.
type ITfThreadMgr struct {
	IUnknown
}

// ITfThreadMgrVtbl COM interface vtable.
type ITfThreadMgrVtbl struct {
	IUnknownVtbl
	Activate              uintptr
	Deactivate            uintptr
	CreateDocumentMgr     uintptr
	EnumDocumentMgrs      uintptr
	GetFocus              uintptr
	SetFocus              uintptr
	AssociateFocus        uintptr
	IsThreadFocus         uintptr
	GetFunctionProvider   uintptr
	EnumFunctionProviders uintptr
	GetGlobalCompartment  uintptr
}

// ITfThreadMgr returns the ITfThreadMgr vtable.
func (obj *ITfThreadMgr) ITfThreadMgr() *ITfThreadMgrVtbl {
	return (*ITfThreadMgrVtbl)(unsafe.Pointer(obj.Vtbl))
}

// Activate activates TSF for the calling thread.
func (obj *ITfThreadMgr) Activate(ptid *TFClientID) error {
	return result(call1(
		obj.ITfThreadMgr().Activate,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(ptid)),
	))
}
