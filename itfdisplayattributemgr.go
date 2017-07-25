package tsf

import "unsafe"

// ITfDisplayAttributeMgr COM interface.
type ITfDisplayAttributeMgr struct {
	IUnknown
}

// ITfDisplayAttributeMgrVtbl COM interface vtable.
type ITfDisplayAttributeMgrVtbl struct {
	IUnknownVtbl
}

// ITfDisplayAttributeMgr returns the ITfDisplayAttributeMgr vtable.
func (obj *ITfDisplayAttributeMgr) ITfDisplayAttributeMgr() *ITfDisplayAttributeMgrVtbl {
	return (*ITfDisplayAttributeMgrVtbl)(unsafe.Pointer(obj.Vtbl))
}
