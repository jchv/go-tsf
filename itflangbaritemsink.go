package tsf

import "unsafe"

// ITfLangBarItemSink COM interface.
type ITfLangBarItemSink struct {
	IUnknown
}

// ITfLangBarItemSinkVtbl COM interface vtable.
type ITfLangBarItemSinkVtbl struct {
	IUnknownVtbl
}

// ITfLangBarItemSink returns the ITfLangBarItemSink vtable.
func (obj *ITfLangBarItemSink) ITfLangBarItemSink() *ITfLangBarItemSinkVtbl {
	return (*ITfLangBarItemSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
