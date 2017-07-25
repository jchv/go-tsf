package tsf

import "unsafe"

// ITfLangBarItem COM interface.
type ITfLangBarItem struct {
	IUnknown
}

// ITfLangBarItemVtbl COM interface vtable.
type ITfLangBarItemVtbl struct {
	IUnknownVtbl
}

// ITfLangBarItem returns the ITfLangBarItem vtable.
func (obj *ITfLangBarItem) ITfLangBarItem() *ITfLangBarItemVtbl {
	return (*ITfLangBarItemVtbl)(unsafe.Pointer(obj.Vtbl))
}
