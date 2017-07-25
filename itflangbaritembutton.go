package tsf

import "unsafe"

// ITfLangBarItemButton COM interface.
type ITfLangBarItemButton struct {
	IUnknown
}

// ITfLangBarItemButtonVtbl COM interface vtable.
type ITfLangBarItemButtonVtbl struct {
	IUnknownVtbl
}

// ITfLangBarItemButton returns the ITfLangBarItemButton vtable.
func (obj *ITfLangBarItemButton) ITfLangBarItemButton() *ITfLangBarItemButtonVtbl {
	return (*ITfLangBarItemButtonVtbl)(unsafe.Pointer(obj.Vtbl))
}
