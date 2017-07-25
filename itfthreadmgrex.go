package tsf

import "unsafe"

// ITfThreadMgrEx COM interface.
type ITfThreadMgrEx struct {
	IUnknown
}

// ITfThreadMgrExVtbl COM interface vtable.
type ITfThreadMgrExVtbl struct {
	IUnknownVtbl
}

// ITfThreadMgrEx returns the ITfThreadMgrEx vtable.
func (obj *ITfThreadMgrEx) ITfThreadMgrEx() *ITfThreadMgrExVtbl {
	return (*ITfThreadMgrExVtbl)(unsafe.Pointer(obj.Vtbl))
}
