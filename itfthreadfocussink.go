package tsf

import "unsafe"

// ITfThreadFocusSink COM interface.
type ITfThreadFocusSink struct {
	IUnknown
}

// ITfThreadFocusSinkVtbl COM interface vtable.
type ITfThreadFocusSinkVtbl struct {
	IUnknownVtbl
}

// ITfThreadFocusSink returns the ITfThreadFocusSink vtable.
func (obj *ITfThreadFocusSink) ITfThreadFocusSink() *ITfThreadFocusSinkVtbl {
	return (*ITfThreadFocusSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
