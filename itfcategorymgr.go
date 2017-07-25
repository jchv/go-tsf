package tsf

import "unsafe"

// ITfCategoryMgr COM interface.
type ITfCategoryMgr struct {
	IUnknown
}

// ITfCategoryMgrVtbl COM interface vtable.
type ITfCategoryMgrVtbl struct {
	IUnknownVtbl
}

// ITfCategoryMgr returns the ITfCategoryMgr vtable.
func (obj *ITfCategoryMgr) ITfCategoryMgr() *ITfCategoryMgrVtbl {
	return (*ITfCategoryMgrVtbl)(unsafe.Pointer(obj.Vtbl))
}
