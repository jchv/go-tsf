package tsf

import "unsafe"

// ITfStatusSink COM interface.
type ITfStatusSink struct {
	IUnknown
}

// ITfStatusSinkVtbl COM interface vtable.
type ITfStatusSinkVtbl struct {
	IUnknownVtbl
}

// ITfStatusSink returns the ITfStatusSink vtable.
func (obj *ITfStatusSink) ITfStatusSink() *ITfStatusSinkVtbl {
	return (*ITfStatusSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
