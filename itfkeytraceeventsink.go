package tsf

import "unsafe"

// ITfKeyTraceEventSink COM interface.
type ITfKeyTraceEventSink struct {
	IUnknown
}

// ITfKeyTraceEventSinkVtbl COM interface vtable.
type ITfKeyTraceEventSinkVtbl struct {
	IUnknownVtbl
}

// ITfKeyTraceEventSink returns the ITfKeyTraceEventSink vtable.
func (obj *ITfKeyTraceEventSink) ITfKeyTraceEventSink() *ITfKeyTraceEventSinkVtbl {
	return (*ITfKeyTraceEventSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
