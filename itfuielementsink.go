package tsf

import "unsafe"

// ITfUIElementSink COM interface.
type ITfUIElementSink struct {
	IUnknown
}

// ITfUIElementSinkVtbl COM interface vtable.
type ITfUIElementSinkVtbl struct {
	IUnknownVtbl
}

// ITfUIElementSink returns the ITfUIElementSink vtable.
func (obj *ITfUIElementSink) ITfUIElementSink() *ITfUIElementSinkVtbl {
	return (*ITfUIElementSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
