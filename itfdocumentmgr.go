package tsf

import "unsafe"

// ITfDocumentMgr COM interface.
type ITfDocumentMgr struct {
	IUnknown
}

// ITfDocumentMgrVtbl COM interface vtable.
type ITfDocumentMgrVtbl struct {
	IUnknownVtbl
}

// ITfDocumentMgr returns the ITfDocumentMgr vtable.
func (obj *ITfDocumentMgr) ITfDocumentMgr() *ITfDocumentMgrVtbl {
	return (*ITfDocumentMgrVtbl)(unsafe.Pointer(obj.Vtbl))
}
