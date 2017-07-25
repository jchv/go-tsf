package tsf

import "unsafe"

// ITfKeystrokeMgr COM interface.
type ITfKeystrokeMgr struct {
	IUnknown
}

// ITfKeystrokeMgrVtbl COM interface vtable.
type ITfKeystrokeMgrVtbl struct {
	IUnknownVtbl
}

// ITfKeystrokeMgr returns the ITfKeystrokeMgr vtable.
func (obj *ITfKeystrokeMgr) ITfKeystrokeMgr() *ITfKeystrokeMgrVtbl {
	return (*ITfKeystrokeMgrVtbl)(unsafe.Pointer(obj.Vtbl))
}
