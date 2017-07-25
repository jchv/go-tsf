package tsf

import "unsafe"

// ITfSystemLangBarItemSink COM interface.
type ITfSystemLangBarItemSink struct {
	IUnknown
}

// ITfSystemLangBarItemSinkVtbl COM interface vtable.
type ITfSystemLangBarItemSinkVtbl struct {
	IUnknownVtbl
}

// ITfSystemLangBarItemSink returns the ITfSystemLangBarItemSink vtable.
func (obj *ITfSystemLangBarItemSink) ITfSystemLangBarItemSink() *ITfSystemLangBarItemSinkVtbl {
	return (*ITfSystemLangBarItemSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
