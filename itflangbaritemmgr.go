package tsf

import "unsafe"

// ITfLangBarItemMgr COM interface.
type ITfLangBarItemMgr struct {
	IUnknown
}

// ITfLangBarItemMgrVtbl COM interface vtable.
type ITfLangBarItemMgrVtbl struct {
	IUnknownVtbl
}

// ITfLangBarItemMgr returns the ITfLangBarItemMgr vtable.
func (obj *ITfLangBarItemMgr) ITfLangBarItemMgr() *ITfLangBarItemMgrVtbl {
	return (*ITfLangBarItemMgrVtbl)(unsafe.Pointer(obj.Vtbl))
}
