package tsf

import "unsafe"

// ITfLangBarEventSink COM interface.
type ITfLangBarEventSink struct {
	IUnknown
}

// ITfLangBarEventSinkVtbl COM interface vtable.
type ITfLangBarEventSinkVtbl struct {
	IUnknownVtbl
}

// ITfLangBarEventSink returns the ITfLangBarEventSink vtable.
func (obj *ITfLangBarEventSink) ITfLangBarEventSink() *ITfLangBarEventSinkVtbl {
	return (*ITfLangBarEventSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
