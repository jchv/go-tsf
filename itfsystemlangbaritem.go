package tsf

import "unsafe"

// ITfSystemLangBarItem COM interface.
type ITfSystemLangBarItem struct {
	IUnknown
}

// ITfSystemLangBarItemVtbl COM interface vtable.
type ITfSystemLangBarItemVtbl struct {
	IUnknownVtbl
}

// ITfSystemLangBarItem returns the ITfSystemLangBarItem vtable.
func (obj *ITfSystemLangBarItem) ITfSystemLangBarItem() *ITfSystemLangBarItemVtbl {
	return (*ITfSystemLangBarItemVtbl)(unsafe.Pointer(obj.Vtbl))
}
