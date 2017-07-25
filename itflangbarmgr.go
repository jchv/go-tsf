package tsf

import "unsafe"

// ITfLangBarMgr COM interface.
type ITfLangBarMgr struct {
	IUnknown
}

// ITfLangBarMgrVtbl COM interface vtable.
type ITfLangBarMgrVtbl struct {
	IUnknownVtbl
}

// ITfLangBarMgr returns the ITfLangBarMgr vtable.
func (obj *ITfLangBarMgr) ITfLangBarMgr() *ITfLangBarMgrVtbl {
	return (*ITfLangBarMgrVtbl)(unsafe.Pointer(obj.Vtbl))
}
