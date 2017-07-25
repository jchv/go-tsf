package tsf

import "unsafe"

// ITfMouseSink COM interface.
type ITfMouseSink struct {
	IUnknown
}

// ITfMouseSinkVtbl COM interface vtable.
type ITfMouseSinkVtbl struct {
	IUnknownVtbl
}

// ITfMouseSink returns the ITfMouseSink vtable.
func (obj *ITfMouseSink) ITfMouseSink() *ITfMouseSinkVtbl {
	return (*ITfMouseSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
