package tsf

import "unsafe"

// ITfLangBarItemBitmap COM interface.
type ITfLangBarItemBitmap struct {
	IUnknown
}

// ITfLangBarItemBitmapVtbl COM interface vtable.
type ITfLangBarItemBitmapVtbl struct {
	IUnknownVtbl
}

// ITfLangBarItemBitmap returns the ITfLangBarItemBitmap vtable.
func (obj *ITfLangBarItemBitmap) ITfLangBarItemBitmap() *ITfLangBarItemBitmapVtbl {
	return (*ITfLangBarItemBitmapVtbl)(unsafe.Pointer(obj.Vtbl))
}
