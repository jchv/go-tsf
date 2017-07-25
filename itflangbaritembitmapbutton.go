package tsf

import "unsafe"

// ITfLangBarItemBitmapButton COM interface.
type ITfLangBarItemBitmapButton struct {
	IUnknown
}

// ITfLangBarItemBitmapButtonVtbl COM interface vtable.
type ITfLangBarItemBitmapButtonVtbl struct {
	IUnknownVtbl
}

// ITfLangBarItemBitmapButton returns the ITfLangBarItemBitmapButton vtable.
func (obj *ITfLangBarItemBitmapButton) ITfLangBarItemBitmapButton() *ITfLangBarItemBitmapButtonVtbl {
	return (*ITfLangBarItemBitmapButtonVtbl)(unsafe.Pointer(obj.Vtbl))
}
