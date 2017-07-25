package tsf

import "unsafe"

// ITfKeyEventSink COM interface.
type ITfKeyEventSink struct {
	IUnknown
}

// ITfKeyEventSinkVtbl COM interface vtable.
type ITfKeyEventSinkVtbl struct {
	IUnknownVtbl
}

// ITfKeyEventSink returns the ITfKeyEventSink vtable.
func (obj *ITfKeyEventSink) ITfKeyEventSink() *ITfKeyEventSinkVtbl {
	return (*ITfKeyEventSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
