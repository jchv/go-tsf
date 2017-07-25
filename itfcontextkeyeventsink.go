package tsf

import "unsafe"

// ITfContextKeyEventSink COM interface.
type ITfContextKeyEventSink struct {
	IUnknown
}

// ITfContextKeyEventSinkVtbl COM interface vtable.
type ITfContextKeyEventSinkVtbl struct {
	IUnknownVtbl
}

// ITfContextKeyEventSink returns the ITfContextKeyEventSink vtable.
func (obj *ITfContextKeyEventSink) ITfContextKeyEventSink() *ITfContextKeyEventSinkVtbl {
	return (*ITfContextKeyEventSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
