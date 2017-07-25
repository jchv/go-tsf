package tsf

import "unsafe"

// ITfTextEditSink COM interface.
type ITfTextEditSink struct {
	IUnknown
}

// ITfTextEditSinkVtbl COM interface vtable.
type ITfTextEditSinkVtbl struct {
	IUnknownVtbl
}

// ITfTextEditSink returns the ITfTextEditSink vtable.
func (obj *ITfTextEditSink) ITfTextEditSink() *ITfTextEditSinkVtbl {
	return (*ITfTextEditSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
