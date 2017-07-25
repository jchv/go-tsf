package tsf

import "unsafe"

// ITfThreadMgr2 COM interface.
type ITfThreadMgr2 struct {
	IUnknown
}

// ITfThreadMgr2Vtbl COM interface vtable.
type ITfThreadMgr2Vtbl struct {
	IUnknownVtbl
}

// ITfThreadMgr2 returns the ITfThreadMgr2 vtable.
func (obj *ITfThreadMgr2) ITfThreadMgr2() *ITfThreadMgr2Vtbl {
	return (*ITfThreadMgr2Vtbl)(unsafe.Pointer(obj.Vtbl))
}
