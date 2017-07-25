package tsf

import "unsafe"

// ITfUIElementMgr COM interface.
type ITfUIElementMgr struct {
	IUnknown
}

// ITfUIElementMgrVtbl COM interface vtable.
type ITfUIElementMgrVtbl struct {
	IUnknownVtbl
}

// ITfUIElementMgr returns the ITfUIElementMgr vtable.
func (obj *ITfUIElementMgr) ITfUIElementMgr() *ITfUIElementMgrVtbl {
	return (*ITfUIElementMgrVtbl)(unsafe.Pointer(obj.Vtbl))
}
