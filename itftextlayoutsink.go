package tsf

import "unsafe"

// ITfTextLayoutSink COM interface.
type ITfTextLayoutSink struct {
	IUnknown
}

// ITfTextLayoutSinkVtbl COM interface vtable.
type ITfTextLayoutSinkVtbl struct {
	IUnknownVtbl
}

// ITfTextLayoutSink returns the ITfTextLayoutSink vtable.
func (obj *ITfTextLayoutSink) ITfTextLayoutSink() *ITfTextLayoutSinkVtbl {
	return (*ITfTextLayoutSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
