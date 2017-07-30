package tsf

import "unsafe"

// ITfThreadMgrEx COM interface.
type ITfThreadMgrEx struct {
	IUnknown
}

// ITfThreadMgrExVtbl COM interface vtable.
type ITfThreadMgrExVtbl struct {
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
	ActivateEx            uintptr
	GetActiveFlags        uintptr
}

// ITfThreadMgrEx returns the ITfThreadMgrEx vtable.
func (obj *ITfThreadMgrEx) ITfThreadMgrEx() *ITfThreadMgrExVtbl {
	return (*ITfThreadMgrExVtbl)(unsafe.Pointer(obj.Vtbl))
}

// ActivateEx activates TSF for the calling thread.
func (obj *ITfThreadMgrEx) ActivateEx(ptid *TFClientID, flags uint32) error {
	return result(call2(
		obj.ITfThreadMgrEx().ActivateEx,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(ptid)),
		uintptr(flags),
	))
}
