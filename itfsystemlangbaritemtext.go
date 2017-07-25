package tsf

import "unsafe"

// ITfSystemLangBarItemText COM interface.
type ITfSystemLangBarItemText struct {
	IUnknown
}

// ITfSystemLangBarItemTextVtbl COM interface vtable.
type ITfSystemLangBarItemTextVtbl struct {
	IUnknownVtbl
}

// ITfSystemLangBarItemText returns the ITfSystemLangBarItemText vtable.
func (obj *ITfSystemLangBarItemText) ITfSystemLangBarItemText() *ITfSystemLangBarItemTextVtbl {
	return (*ITfSystemLangBarItemTextVtbl)(unsafe.Pointer(obj.Vtbl))
}
