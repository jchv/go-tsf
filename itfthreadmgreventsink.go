package tsf

import "unsafe"

// ITfThreadMgrEventSink COM interface.
type ITfThreadMgrEventSink struct {
	IUnknown
}

// ITfThreadMgrEventSinkVtbl COM interface vtable.
type ITfThreadMgrEventSinkVtbl struct {
	IUnknownVtbl
}

// ITfThreadMgrEventSink returns the ITfThreadMgrEventSink vtable.
func (obj *ITfThreadMgrEventSink) ITfThreadMgrEventSink() *ITfThreadMgrEventSinkVtbl {
	return (*ITfThreadMgrEventSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
